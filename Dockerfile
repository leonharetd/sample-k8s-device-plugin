FROM golang:1.16 as builder

WORKDIR /workspace
COPY ./ ./
RUN go mod download
RUN go build -mod=readonly ./cmd/sample/main.go

FROM centos:7

COPY --from=builder /workspace/main /usr/bin/sample-device-plugin

ENTRYPOINT ["/usr/bin/sample-device-plugin"]