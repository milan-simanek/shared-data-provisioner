# Dynamic Provisioning of specific directory locations using Kubernetes HostPath

This is a dynamic provisioner for Kubernetes. It can connect to your
application PVC to a predefined directory with a data.

It is based on [Rimusz's hostPath provisioner](https://github.com/rimusz/hostpath-provisioner) project.

The provisioner is intended to use on a single node Kubernetes.

# Use cases

- **Situation**: 
  
  A pod (or multiple pods) run an imgage (e.g. php) which requires some
  library (tcpdf) to be mounted on a specific location.

  **Solution**:

  1. Make a subdirectory of the baseDir: /var/shared-data/tcpdf.
  2. Copy the library into that directory.
  3. Create a PVC with spec.selector.matchLabels.component=tcpdf
  4. Shared-data provisioner will create a PV to match this PVC using
     hostPath /var/shared-data/tcpdf.

- **Situation**:

  A couple of pods need a shared directory for a comminucation.

  **Solution**:

  1. Create a subdirectory of the baseDir: /var/shared-data/communication
  2. Create a PVC with spec.selector.matchLabels.component=communication for each pod
  3. Shared-data provisioner will create a PVs to match this PVCs using
     hostPath /var/shared-data/communication

# Installation

There are two methods how to install this provisioner: using shell script or
using helm chart. Helm chart method is prefered.

## Helm chart install method

```bash
helm repo add shared-data https://milan-simanek.github.io/shared-data-provisioner
helm repo update
helm search repo shared-data
helm install --namespace shared-data-provisioner --create-namespace shared-data shared-data/shared-data-provisioner
```

Alternatively using oneliner:

```bash
helm install --namespace shared-data-provisioner --create-namespace shared-data https://milan-simanek.github.io/shared-data-provisioner/shared-data-provisioner-1.0.0/shared-data-provisioner-1.0.0.tgz
```

## Shell script install method

```bash
# git clone https://github.com/milan-simanek/shared-data-provisioner
# cd shared-data-provisioner/manifests
# ./INSTALL
```

