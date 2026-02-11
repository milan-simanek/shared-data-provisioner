/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
	"syscall"

	"sigs.k8s.io/sig-storage-lib-external-provisioner/v7/controller"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	klog "k8s.io/klog/v2"
)

// Fetch provisioner name from environment variable SHARED_DATA_PROVISIONER_NAME
// if not set uses default shared-data name
func GetProvisionerName() string {
	provisionerName := os.Getenv("SHARED_DATA_PROVISIONER_NAME")
	if provisionerName == "" {
		provisionerName = "shared-data"
	}
	return provisionerName
}


type sharedDataProvisioner struct {
	// The base directory where the sahred data is located
	baseDir string

	// Identity of this sharedDataProvisioner, set to node's name. Used to identify
	// "this" provisioner's PVs.
	identity string
}

// NewSharedDataProvisioner creates a new shared-data provisioner
func NewSharedDataProvisioner() controller.Provisioner {
	nodeName := os.Getenv("NODE_NAME")
	if nodeName == "" {
		klog.Fatal("env variable NODE_NAME must be set so that this provisioner can identify itself")
	}
	nodeBaseDir := os.Getenv("NODE_BASE_DIR")
	if nodeBaseDir == "" {
		nodeBaseDir = "/var/shared-data"
	}
	return &sharedDataProvisioner{
		baseDir:  nodeBaseDir,
		identity: nodeName,
	}
}

var _ controller.Provisioner = &sharedDataProvisioner{}

// Provision verifies that the data directory exists and returns a PV object representing it.
func (p *sharedDataProvisioner) Provision(ctx context.Context, options controller.ProvisionOptions) (*v1.PersistentVolume, controller.ProvisioningState, error) {
        component := options.PVC.GetLabels()["component"]
        
        if strings.Contains(component, "/") {
        	return nil, controller.ProvisioningFinished, fmt.Errorf("PVC label 'component' contains invalid character '/'")
        }
        
        if len(component) == 0 {
        	return nil, controller.ProvisioningFinished, fmt.Errorf("PVC label 'component' is empty or not defined")
	}

	path := path.Join(p.baseDir, component)
	info, err := os.Stat(path)
	if err != nil {
		return nil, controller.ProvisioningFinished, fmt.Errorf("failed to stat path: %v", err)
	}
	if !info.IsDir() {
		return nil, controller.ProvisioningFinished, fmt.Errorf("path is not a directory: %s", path)
	}
	
	fmt.Printf("provisioning directory '%s'\n", path)
	
	pv := &v1.PersistentVolume{
		ObjectMeta: metav1.ObjectMeta{
			Name: options.PVName,
			Annotations: map[string]string{
				"sharedDataProvisionerIdentity": p.identity,
			},
		},
		Spec: v1.PersistentVolumeSpec{
			PersistentVolumeReclaimPolicy: *options.StorageClass.ReclaimPolicy,
			AccessModes:                   []v1.PersistentVolumeAccessMode{v1.ReadOnlyMany},
			Capacity: v1.ResourceList{
				v1.ResourceName(v1.ResourceStorage): options.PVC.Spec.Resources.Requests[v1.ResourceName(v1.ResourceStorage)],
			},
			PersistentVolumeSource: v1.PersistentVolumeSource{
				HostPath: &v1.HostPathVolumeSource{
					Path: path,
				},
			},
		},
	}

	return pv, controller.ProvisioningFinished, nil
}

// Delete does not remove the data (the data is shared)
func (p *sharedDataProvisioner) Delete(ctx context.Context, volume *v1.PersistentVolume) error {
	ann, ok := volume.Annotations["sharedDataProvisionerIdentity"]
	if !ok {
		return errors.New("identity annotation not found on PV")
	}
	if ann != p.identity {
		return &controller.IgnoredError{Reason: "identity annotation on PV does not match ours"}
	}

	return nil
}

func main() {
	syscall.Umask(0)

	flag.Parse()
	flag.Set("logtostderr", "true")

	// Create an InClusterConfig and use it to create a client for the controller
	// to use to communicate with Kubernetes
	config, err := rest.InClusterConfig()
	if err != nil {
		klog.Fatalf("Failed to create config: %v", err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		klog.Fatalf("Failed to create client: %v", err)
	}

	// Create the provisioner: it implements the Provisioner interface expected by
	// the controller
	sharedDataProvisioner := NewSharedDataProvisioner()

	// Start the provision controller which will dynamically provision hostPath
	// PVs
	pc := controller.NewProvisionController(clientset, GetProvisionerName(), sharedDataProvisioner)

	// Never stops.
	pc.Run(context.Background())
}
