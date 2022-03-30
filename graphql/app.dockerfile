FROM golang:1.18.0-alpine3.15 AS build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /go/src/wpg_pretest/graphql
COPY vendor ../vendor
COPY account ../account
COPY catalog ../catalog
COPY order ../order
COPY graphql ./
ENV GO111MODULE=auto
RUN go build -o /go/bin/app_gql

FROM alpine:3.15
WORKDIR /usr/bin
COPY --from=build /go/bin .
EXPOSE 8080
CMD ["app_gql"]
