# This is the Shared Data Provisioner for Kubernetes.

Let us have a directory ```/var/data``` (called base directory) containing 
several subdirectories called components. 
These components contain various data (like libraries) which
need to be mounted inside a container. To do this, one can create a
PersistentVolumeClaim with a label ```component=abc``` where ``abc`` is the
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

## Values

- namespace
  
  Kubernetes namespace where the provisioner should run
  (defaults to ``shared-data-provisioner``)

- baseDir
  
  Host directory where components are located
  (defaults to ``/var/shared-data``)

- image
  The OCI image with provisioner binary 
  (defaults to `` docker.io/milansimanek/shared-data-provisioner:v0.1.3-label``)
