# Dockerfile for the `go-meeting` service
FROM golang
MAINTAINER Navneet

# Add the cwd to its corresponding location in the docker container
ADD . /go/src/github.com/Navneet923/go-meeting

# Install deps
RUN go get github.com/gorilla/mux

# Build the application
RUN go install github.com/Navneet923/go-meeting/go-meeting

# Setup the entry-point
WORKDIR /go/src/github.com/Navneet923/go-meeting
ENTRYPOINT /go/bin/go-meeting

# Expose any needed ports
EXPOSE 80

