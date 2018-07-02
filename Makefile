
#export GOPATH:=$(shell pwd)

BUILDTAGS=debug

#GOARCH="amd64"
#GOOS="linux"
#GOHOSTARCH="amd64"
#GOHOSTOS="linux"

all: gohome
gohome:
	GOARCH=amd64 GOOS=linux go build -tags '$(BUILDTAGS)'  gohome.go 
release: BUILDTAGS=release
release: gohome
clean:
	rm gohome
