FROM golang:1.11.5-alpine3.9 as builder

ENV SRC_DIR="/go/src/github.com/satriajidam/hello-country"

# Install tools & packages required to build the project.
# We will need to run `docker build --no-cache .` to update those dependencies.
RUN apk add --no-cache git
RUN go get github.com/golang/dep/cmd/dep

# Gopkg.toml and Gopkg.lock lists source dependencies.
# These layers will only be re-built when Gopkg files are updated.
COPY Gopkg.lock Gopkg.toml $SRC_DIR/

WORKDIR $SRC_DIR

# Install library dependencies.
RUN dep ensure --vendor-only

# Copy all source and build it.
# This layer will be rebuilt whenever a file has changed in the source directory.
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -i -a -v -installsuffix nocgo -o /bin/hello-country .

# Build final image.
FROM gcr.io/distroless/static
WORKDIR /app
COPY --from=builder /bin/hello-country hello-country
ENTRYPOINT ["./hello-country"]
