# pdsps

A reverse proxy that sits in front of a Bluesky PDS and returns a fixed "assured" response for `getAgeAssuranceState`, proxying all other requests to the upstream.

## Usage

```
UPSTREAM=https://your-pds.example.com ./pdsps
```

### Flags

| Flag | Env var | Default | Description |
|------|---------|---------|-------------|
| `-upstream` | `UPSTREAM` | (required) | URL of the upstream PDS server |
| `-port` | `PORT` | `8080` | Port to listen on |

### Docker Compose

```yaml
services:
  pdsps:
    image: pdsps
    environment:
      - UPSTREAM=http://pds:2583
      - PORT=2583
    ports:
      - "2583:2583"

  pds:
    image: your-pds-image
```

## Provenance

This project was primarily created with Claude Code, but with a strong guiding
hand. It's not "vibe coded", but an LLM was still the primary author of most
lines of code. I believe it meets the same sort of standards I'd aim for with
hand-crafted code, but some slop may slip through. I understand if you
prefer not to use LLM-created software, and welcome human-authored alternatives
(I just don't personally have the time/motivation to do so).
