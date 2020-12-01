FROM alpine:latest as certs
RUN apk --update add ca-certificates
ARG version=0.0.1
FROM scratch
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY ./datadog-apm ./datadog-apm
LABEL Name=oem-audi Version=$version
CMD [ "./datadog-apm" ]