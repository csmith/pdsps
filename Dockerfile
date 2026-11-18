FROM golang:1.27.4 AS build
WORKDIR /go/src/app
COPY . .

RUN --mount=type=cache,target=/go/pkg/mod \
    set -eux; \
    CGO_ENABLED=0 GO111MODULE=on go install .; \
    mkdir -p /notices; \
    go tool gocredits -skip-missing . > /notices/CREDITS;

FROM ghcr.io/greboid/dockerbase/nonroot:1.20261020.1
COPY --from=build /go/bin/pdsps /pdsps
COPY --from=build /notices /notices
ENTRYPOINT ["/pdsps"]
