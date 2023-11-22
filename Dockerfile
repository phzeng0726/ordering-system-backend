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
RUN go get -u ./... && go mod tidy

# Build the binary.
# -mod=readonly ensures immutable go.mod and go.sum in container builds.
# -mod=mod automatically update go.mod and go.sum
RUN cd app && CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o /docker-gs-ping

# Use the official Alpine image for a lean production container.
# https://hub.docker.com/_/alpine
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine:latest

WORKDIR /

RUN apk add --no-cache ca-certificates

# Copy the binary to the production image from the builder stage.
COPY --from=builder /docker-gs-ping /docker-gs-ping
COPY --from=builder /osbe/assets /assets

# Run the web service on container startup.
CMD ["/docker-gs-ping", "-c"]
