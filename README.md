# logslice

Fast structured log filter and formatter for JSON log streams from Kubernetes pods.

## Installation

```bash
go install github.com/yourname/logslice@latest
```

## Usage

Pipe Kubernetes pod logs directly into `logslice` to filter and format structured JSON output:

```bash
kubectl logs -f my-pod | logslice
```

Filter by log level and format output:

```bash
kubectl logs -f my-pod | logslice --level error --format pretty
```

Filter by a specific field value:

```bash
kubectl logs -f my-pod | logslice --field service=auth --level warn
```

### Flags

| Flag | Description | Default |
|------|-------------|---------|
| `--level` | Minimum log level to display (`debug`, `info`, `warn`, `error`) | `info` |
| `--format` | Output format: `pretty`, `json`, `compact` | `pretty` |
| `--field` | Filter by field key=value pair | — |
| `--timestamp` | Show timestamp in output | `true` |

### Example Output

```
2024-01-15 12:34:56  ERROR  auth-service  failed to validate token  {"user_id": "u_123", "reason": "expired"}
2024-01-15 12:34:57  WARN   auth-service  retry attempt              {"attempt": 2, "max": 3}
```

## Requirements

- Go 1.21+
- Kubernetes cluster with `kubectl` configured

## License

MIT © 2024 yourname