package main

import "github.com/lightstep/otel-launcher-go/launcher"

func main() {
	// note: you only need the access token for connecting to Lightstep
	otel := launcher.ConfigureOpentelemetry(
		launcher.WithServiceName("service-123"),
	)
	defer otel.Shutdown()
}
