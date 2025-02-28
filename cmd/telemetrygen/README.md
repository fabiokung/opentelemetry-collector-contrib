# Telemetry generator for OpenTelemetry

<!-- status autogenerated section -->
| Status        |           |
| ------------- |-----------|
| Stability     | [alpha]: traces   |
|               | [development]: metrics, logs   |
| Issues        | ![Open issues](https://img.shields.io/github/issues-search/open-telemetry/opentelemetry-collector-contrib?query=is%3Aissue%20is%3Aopen%20label%3Acmd%2Ftelemetrygen%20&label=open&color=orange&logo=opentelemetry) ![Closed issues](https://img.shields.io/github/issues-search/open-telemetry/opentelemetry-collector-contrib?query=is%3Aissue%20is%3Aclosed%20label%3Acmd%2Ftelemetrygen%20&label=closed&color=blue&logo=opentelemetry) |

[alpha]: https://github.com/open-telemetry/opentelemetry-collector#alpha
[development]: https://github.com/open-telemetry/opentelemetry-collector#development
<!-- end autogenerated section -->

This utility simulates a client generating **traces**, **metrics**, and **logs**. It is useful for testing and demonstration purposes.

## Installing

To install the latest version run the following command:

```console
go install github.com/open-telemetry/opentelemetry-collector-contrib/cmd/telemetrygen@latest
```

Check the [`go install` reference](https://go.dev/ref/mod#go-install) to install specific versions.

## Running

First, you'll need an OpenTelemetry Collector to receive the telemetry data. Follow the project's instructions for a detailed setting up guide. The following configuration file should be sufficient:

```yaml
receivers:
  otlp:
    protocols:
      grpc:
        endpoint: localhost:4317

processors:

exporters:
  logging:

service:
  pipelines:
    traces:
      receivers:
      - otlp
      processors: []
      exporters:
      - logging
```

Once the OpenTelemetry Collector instance is up and running, run `telemetrygen` for your desired telemetry:

### Traces

```console
telemetrygen traces --otlp-insecure --duration 5s
```

Or, to generate a specific number of traces:

```console
telemetrygen traces --otlp-insecure --traces 1
```

Check `telemetrygen traces --help` for all the options.
