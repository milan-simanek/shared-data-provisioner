# This is the Shared Data Provisioner for Kubernetes.

Let us have a directory ```/var/data``` (called base directory) containing 
several subdirectories called components. 
These components contain various data (like libraries) which
need to be mounted inside a container. To do this, one can create a
PersistentVolumeClaim with a ```spec.selector.matchLabels.component=abc``` where ``abc`` is the
name of the component to be mounted. The provisioner will automatically
provision a PersistentVolume accessing the requested component of the base
directory.

The provisioner was designed for a single node clusters, but can work also
on multiple nodes, if you run more instances and have a copy of the data on
multiple nodes.

Using this approach, multiple different pod can mount the same component as
a shared directory. The mount can be read-only or not.
The component can provide a library data (like tools, libraries, commands,
constant data files) or can be used for pod-to-pod communication (shared
filesystem).

See more on https://github.com/milan-simanek/shared-data-provisioner

## Namespace

The helm chart deploys the provisioner into the namespace specified as an
argument to helm (``--namespace``) or as a value. The namespace has to be
created before ``helm install`` and it should follow ``privileged`` Pod
Security Standard level. 

## Values

- baseDir
  
  Host directory where components are located
  (defaults to ``/var/shared-data``)

- image
  The OCI image with provisioner binary 
  (defaults to `` milansimanek/shared-data-provisioner:v1.0.0``)

- namespace
  The target namespace for the installation
  (defaults to the helm release namespace)

### other values - usually you need not to change them
    ClusterRole and ClusterRoleBinding allowing to create PersistentVolumes and modify PersistentVolumeClaim
```
clusterRoleName: shared-data-provisioner
clusterRoleBindingName: shared-data-provisioner
serviceAccountName: shared-data-provisioner
leaderLockingRoleName: shared-data-provisioner-leader-locking
```
