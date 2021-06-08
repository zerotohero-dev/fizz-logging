package main

import "github.com/zerotohero-dev/fizz-logging/pkg/log"

const appName = "demo-svc"

func main() {
	log.Init(appName)

	log.Info("Smoke test: '%d'", 42)
}
