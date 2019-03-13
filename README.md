openshift-install moneky
===

This is a helper to use the upstream [Openshift Install CLI](https://github.com/openshift/installer) for RHCOS developers.

The resulting binary will allow you to use environment variables to override the hard-coded defaults.

Building
===

This assumes you have a GoLang build environment, [have completed the Libvirt setup](https://github.com/openshift/installer/blob/master/docs/dev/libvirt-howto.md).
```
make init
make build
```

The resulting file is `osi-monkey`, as in `openshift-installer-monkey`.

envVars
===

Three environment variables are exposed:
* `RHCOS_CHANNEL` defaults to Ootpa, but allows you change it
* `RHCOS_URL` which allows you to follow a different RHCOS stream
* `RHCOS_NAME` to override the build name
* `RHCOS_QCOW` to override Libvirt QCOW2 images to _local_ file.

For example:
```
export RHCOS_URL="https://horcruxes.me/storage/releases/"
osi-monkey create cluster
```
