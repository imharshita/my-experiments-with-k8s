FROM golang:1.8-alpine AS builder
WORKDIR /go/src/fileserver
ADD file-server.go /go/src/fileserver
RUN go build fileserver

FROM alpine:latest
COPY --from=builder /go/src/fileserver .
COPY index.html /
COPY serve.html /serve/data/
ENV PORT 8080
CMD ["./fileserver"]