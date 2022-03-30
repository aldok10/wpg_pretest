FROM golang:1.18.0-alpine3.15 AS build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /go/src/wpg_pretest/catalog
COPY vendor ../vendor
COPY catalog ./
ENV GO111MODULE=auto
RUN go build -o /go/bin/app_catalog ./cmd/catalog/main.go

FROM alpine:3.15
WORKDIR /usr/bin
COPY --from=build /go/bin .
EXPOSE 8080
CMD ["app_catalog"]
