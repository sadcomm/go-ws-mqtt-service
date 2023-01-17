package main

import (
	"github.com/sadcomm/ws"
	"github.com/sadcomm/mqtt"
)

func main() {
	ws.SetupRoutes();
	mqtt.SetupMqttClient();
}