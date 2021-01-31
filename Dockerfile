FROM golang:1.14 as builder

WORKDIR /src
COPY ./go.mod ./go.sum ./

ENV GOPROXY=https://proxy.golang.org
ENV GOSUMDB=sum.golang.org
ENV GO111MODULE=on
RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 go build -installsuffix 'static'

FROM alpine:latest as certs
RUN apk --update add ca-certificates
ARG version=0.0.1
FROM scratch
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY --from=builder /src ./
CMD [ "./datadog-apm" ]