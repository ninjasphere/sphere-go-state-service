FROM       golang:1.4
MAINTAINER Ninja Blocks <developers@ninjablocks.com>

RUN        apt-get -qy update && apt-get -qy install vim-common gcc mercurial

WORKDIR    /go/src/github.com/ninjablocks/sphere-go-state-service
ADD        . /go/src/github.com/ninjablocks/sphere-go-state-service

RUN        go get -v

RUN  go build -ldflags " \
       -X main.buildVersion  $(grep "const Version " version.go | sed -E 's/.*"(.+)"$/\1/' ) \
       -X main.buildRevision $(git rev-parse --short HEAD) \
       -X main.buildBranch   $(git rev-parse --abbrev-ref HEAD) \
       -X main.buildDate     $(date +%Y%m%d-%H:%M:%S) \
       -X main.goVersion     $GOLANG_VERSION \
     "

EXPOSE     6100
CMD        ./sphere-go-state-service