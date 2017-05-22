FROM golang:latest

# Copy the local package files to the container’s workspace.
ADD . /go/src/github.com/hathbanger/butterfli-api

# Install our dependencies
RUN go get github.com/hathbanger/butterfli-api/...


# Install api binary globally within container 
RUN go install github.com/hathbanger/butterfli-api

# Set binary as entrypoint
ENTRYPOINT /go/bin/butterfli-api

# Expose default port (3000)
EXPOSE 3000 
