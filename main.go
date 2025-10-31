// Package main
package main

import (
	"github.com/andreykaipov/goobs"
	"github.com/andreykaipov/goobs/api/requests/general"
	"log"
	"os"
)

func main() {
	pass, ok := os.LookupEnv("OBS_WEBSOCKET_PASS")
	if !ok {
		log.Fatalf("Please set WebSocket Password to OBS_WEBSOCKET_PASS")
	}
	port, ok := os.LookupEnv("OBS_WEBSOCKET_PORT")
	if !ok {
		port = "4455"
	}
	client, err := goobs.New("localhost:"+port, goobs.WithPassword(pass))
	if err != nil {
		log.Fatalf("Faild to connect OBS WebSocket Server: %v", err)
	}
	defer client.Disconnect()
	_, err = client.Outputs.StopReplayBuffer()
	if err != nil {
		log.Fatalf("Faild to stop Replay Buffer: %v", err)
	}
	params := general.NewCallVendorRequestParams().
		WithVendorName("obs-shutdown-plugin").
		WithRequestType("shutdown").
		WithRequestData(map[string]any{
			"reason":      "cleaning up",
			"support_url": "",
			"force":       true,
		})
	_, err = client.General.CallVendorRequest(params)
	if err != nil {
		log.Fatalf("Faild to shutdown OBS: %v", err)
	}
}
