my_d = $(shell pwd -P)
installer_d = src/github.com/openshift/installer
GOPATH = $(my_d)

.PHONY: init
init:
	env GOPATH=$(GOPATH) \
	    go get bou.ke/monkey && \
	    git clone https://github.com/openshift/installer $(installer_d)

.PHONY: clean
clean:
	@rm -rf pkg src osi-monkey
.PHONY: build
build:
	cp -avu monkey.go $(installer_d)/pkg/rhcos
	cd $(installer_d); \
	    env GOPATH=$(GOPATH) \
	        hack/build.sh && cp bin/openshift-install $(my_d)/osi-monkey
