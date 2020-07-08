# golang image where workspace (GOPATH) configured at /go.
FROM golang:1.7

# Install dependencies
RUN go get gopkg.in/mgo.v2
RUN go get github.com/gorilla/mux

# copy the local package files to the container workspace
ADD . /go/src/github.com/gkganesh126/nokia-interview

# Setting up working directory
WORKDIR /go/src/github.com/gkganesh126/nokia-interview

# Build command inside the container.
RUN go install github.com/gkganesh126/nokia-interview

# Run the microservice when the container starts.
ENTRYPOINT /go/bin/nokia-interview

# Service listens on port 8083.
EXPOSE 8084
EXPOSE 27017
