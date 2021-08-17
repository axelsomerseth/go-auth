# Accept the Go version for the image to be set as a build argument.
# Default to Go 1.13.
ARG GO_VERSION=1.13

# First stage: build the executable.
FROM golang:${GO_VERSION}-alpine AS builder

# Add Maintainer info.
LABEL maintainer="Axel Somerseth Cordova <axelsomerseth@gmail.com>"

# Set proxy environment variable.
RUN go env -w GOPROXY=direct

# Git is required for fetching the dependencies.
RUN apk update 
RUN apk add --no-cache git
RUN apk add ca-certificates

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /src

# Fetch dependencies first; they are less susceptible to change on every build
# and will therefore be cached for speeding up the next build.
COPY ./go.mod ./go.sum ./
RUN go mod download

# Import the code from the context.
COPY ./ ./

# Build the executable to `/app`. Mark the build as statically linked.
RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /app .

# Final stage: the running container.
FROM scratch AS final

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Import the compiled executable from the first stage.
COPY --from=builder /app /app

# Expose the port for the server.
EXPOSE 8080

# Run the compiled binary.
ENTRYPOINT ["/app"]