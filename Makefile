default:
	@export GOPATH=$$GOPATH:$$(pwd)  && go install btrz-packer-awss3
	@cp bin/btrz-packer-awss3 ~/.packer.d/plugins/packer-provisioner-s3loader
run: default
	@bin/btrz-packer-awss3
	@echo ""
push:
	git add --all && git commit -am 'update' && git push origin master
clean:
	@rm -rf bin
	@rm -rf pkg
edit:
	@export GOPATH=$$GOPATH:$$(pwd)  && atom .
edit2:
	@export GOPATH=$$GOPATH:$$(pwd)  && code .
test:
	@export GOPATH=$$GOPATH:$$(pwd) && go test ./...
setup:
	go get github.com/hashicorp/packer
	go get github.com/hashicorp/packer/packer/plugin
	go get -u github.com/aws/aws-sdk-go/...
	go get github.com/mitchellh/mapstructure
