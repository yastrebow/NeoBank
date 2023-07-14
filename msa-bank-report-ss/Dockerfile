FROM golang:1.19-alpine as builder
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o /main cmd/msa-bank-report-ss/main.go
FROM alpine:3
USER root
COPY --from=builder main /bin/main
COPY config.yml /
ENTRYPOINT ["/bin/main"]