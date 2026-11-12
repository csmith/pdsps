# Changelog

## 1.1.0 - 2026-11-12

### Changed

- Log with structured logging, configurable with the new `log.level` and `log.format` flags
- Update to Go 1.27.4

### Fixed

- Time out connections that are slow to send request headers or sit idle, so stalled clients can't exhaust the server
- Forward requests with unclean paths (such as `//xrpc/foo`) to the PDS instead of redirecting them

## 1.0.2 - 2026-11-01

### Fixed

- Allow the `x-atproto-device-id` and `x-atproto-session-id` headers in age assurance responses, which current Bluesky clients send

## 1.0.1 - 2026-10-23

### Changed

- Build with Go 1.27.3 and update the base image

### Fixed

- Stop the proxy forwarding queries with more parameters than the configured limit (GO-2026-4976)
- Fix HTTP/1 connection desynchronisation after CONNECT requests, in both the server and the proxy (GO-2026-6613, GO-2026-6605)
- Enforce the memory limit when parsing request headers (GO-2026-6608)
- Fix denial of service and excessive work in TLS and certificate verification when talking to an HTTPS upstream (GO-2026-4870, GO-2026-4946, GO-2026-4947, GO-2026-5037, GO-2026-6090)
- Fix the HTTP/2 transport hanging or accepting malformed frames from an upstream (GO-2026-4918, GO-2026-6610)
- Distrust compromised intermediate certificates when verifying an HTTPS upstream (GO-2026-6809)

## 1.0.0 - 2026-04-02

_Initial release._
