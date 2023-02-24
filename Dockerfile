############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && \
    apk add git && \
    apk add build-base upx

WORKDIR /src/app
COPY . .
# Fetch dependencies.
# Using go get.
# Build the binary.
RUN go build  -o /go/bin/app cmd/main.go
RUN upx /go/bin/app

############################
# STEP 2 build a small image
############################
FROM alpine
# Copy our static executable.
RUN apk update && apk add --no-cache  vips-dev
COPY --from=builder /go/bin/app /go/bin/app
COPY --from=builder /src/app/cmd/config.yaml /go/bin

# Run the hello binary.
ENTRYPOINT ["/go/bin/app", "-configpath=/go/bin/config.yaml"]
