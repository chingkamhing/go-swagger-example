project=example

# install necessary program
.PHONY: install
install:
	go install github.com/go-swagger/go-swagger/cmd/swagger

# build the source to native OS and platform
.PHONY: build
build:
	go build -ldflags '-extldflags "-static"' -o ${project} main.go

# build the source to Linux amd64 binary
.PHONY: release
release:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-extldflags "-static"' -o ${project} main.go

# - generate swagger source codes
# - build the sources to binar
.PHONY: all
all: generate build

# go update libraries
.PHONY: update
update:
	go get -u ./...
	go mod tidy

# validate if swagger.yml is valid
.PHONY: validate
validate:
	swagger validate ./swagger/swagger.yml

# generate server source code base on input of swagger.yml
.PHONY: generate
generate: validate
	swagger generate server --name Example --spec ./swagger/swagger.yml --principal interface{} --target gen --exclude-main --flag-strategy=flag

# generate go models structure base on input of swagger.yml
.PHONY: model
model:
	swagger generate model --spec ./swagger/swagger.yml --target gen

# remove all previously generated codes and re-generate all server source code base on input of swagger.yml again; this can eliminate some redundant files
.PHONY: regenerate
regenerate: clean generate

# clean all the binary and the generated code
# note: leave restapi/configure_telesales_admin.go untouch as it is supposed to be customized by user
.PHONY: clean
clean:
	# remove all the generated sources that can be re-generated with the swagger.yml file
	rm -rf gen/models gen/restapi
	# remove all the compiled binaries
	rm -f ${project}
