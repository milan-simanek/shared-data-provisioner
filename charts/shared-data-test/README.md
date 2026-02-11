# Shared Data Provisioner TEST

This chart creates a new component and a sample data file inside the base
directory. Then it starts a pod mounting this directory using PVC and tries
to access it.

Uninstalling of this chart also removes the sample data and its associated
directory component.

