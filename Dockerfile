# Baseimage
FROM golang:1.16-alpine AS builder

WORKDIR $GOPATH/src/github.com/Zferg/simple-http

# Copying codebase to container
COPY . .

ENV CGO_ENABLED=0

# Building the binary
RUN go build -ldflags="-s -w" -v -o /bin/simple-syrup ./cmd/simple-syrup

##### Runtime image

FROM scratch

COPY --from=builder /bin/simple-syrup /bin/simple-syrup

# Setting env
ENV PORT=8080

EXPOSE 8080

# Run command for binary
ENTRYPOINT ["/bin/simple-syrup"]