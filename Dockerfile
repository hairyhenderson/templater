FROM golang:1.11.5-alpine@sha256:8dea7186cf96e6072c23bcbac842d140fe0186758bcc215acb1745f584984857 AS build

RUN apk add --no-cache \
    make \
    git \
    upx=3.94-r0

RUN mkdir -p /go/src/github.com/hairyhenderson/gomplate
WORKDIR /go/src/github.com/hairyhenderson/gomplate
COPY . /go/src/github.com/hairyhenderson/gomplate

RUN make build-x compress-all

FROM scratch AS artifacts

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /go/src/github.com/hairyhenderson/gomplate/bin/* /bin/

CMD [ "/bin/gomplate_linux-amd64" ]

FROM scratch AS gomplate

ARG BUILD_DATE
ARG VCS_REF
ARG OS=linux
ARG ARCH=amd64

LABEL org.opencontainers.image.created=$BUILD_DATE \
      org.opencontainers.image.revision=$VCS_REF \
      org.opencontainers.image.source="https://github.com/hairyhenderson/gomplate"

COPY --from=artifacts /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=artifacts /bin/gomplate_${OS}-${ARCH} /gomplate

ENTRYPOINT [ "/gomplate" ]

FROM alpine:3.9@sha256:b3dbf31b77fd99d9c08f780ce6f5282aba076d70a513a8be859d8d3a4d0c92b8 AS gomplate-alpine

ARG BUILD_DATE
ARG VCS_REF
ARG OS=linux
ARG ARCH=amd64

LABEL org.opencontainers.image.created=$BUILD_DATE \
      org.opencontainers.image.revision=$VCS_REF \
      org.opencontainers.image.source="https://github.com/hairyhenderson/gomplate"

RUN apk add --no-cache ca-certificates
COPY --from=artifacts /bin/gomplate_${OS}-${ARCH}-slim /bin/gomplate

ENTRYPOINT [ "/bin/gomplate" ]

FROM scratch AS gomplate-slim

ARG BUILD_DATE
ARG VCS_REF
ARG OS=linux
ARG ARCH=amd64

LABEL org.opencontainers.image.created=$BUILD_DATE \
      org.opencontainers.image.revision=$VCS_REF \
      org.opencontainers.image.source="https://github.com/hairyhenderson/gomplate"

COPY --from=artifacts /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=artifacts /bin/gomplate_${OS}-${ARCH}-slim /gomplate

ENTRYPOINT [ "/gomplate" ]
