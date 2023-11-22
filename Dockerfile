# Use the offical golang image to create a binary.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.20.5 as builder

# Create and change to the app directory.
WORKDIR /osbe

# Retrieve application dependencies using go modules.
# Allows container builds to reuse downloaded dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./
RUN rm -rf tools && go get -u ./... && go mod tidy

# Build the binary.
# -mod=readonly ensures immutable go.mod and go.sum in container builds.
# -mod=mod automatically update go.mod and go.sum
RUN cd app && CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o /go/bin/app

# Use the official Alpine image for a lean production container.
# https://hub.docker.com/_/alpine
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine:latest
RUN apk add --no-cache ca-certificates

# Copy the binary to the production image from the builder stage.
COPY --from=builder /go/bin/app /app

# Run the web service on container startup.
CMD ["/app", "-c"]