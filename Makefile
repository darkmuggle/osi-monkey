my_d = $(shell pwd -P)
installer_d = src/github.com/openshift/installer
GOPATH = $(my_d)
export GOPATH
export PATH=$(shell echo "$(GOPATH)/bin:$$PATH")
export CGO=1

.PHONY: init
init:
	@echo "** THIS WILL TAKE A BIT..."
	@mkdir -p bin
	go get bou.ke/monkey
	git clone https://github.com/openshift/installer $(installer_d)
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
	@cd $(installer_d) && dep ensure

.PHONY: clean
clean:
	@rm -rf pkg src osi-monkey bin
.PHONY: build
build:
	cp -avu monkey.go $(installer_d)/pkg/rhcos
	@cd $(installer_d) && \
	        TAGS=libvirt hack/build.sh && cp bin/openshift-install $(my_d)/osi-monkey
