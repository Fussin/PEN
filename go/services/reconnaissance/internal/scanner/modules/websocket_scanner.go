package modules

import (
	"context"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func (x *XSSScanner) ScanWebSocket(ctx context.Context, url string, payload string) {
	// Create a new WebSocket connection
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Printf("Error dialing WebSocket: %v", err)
		return
	}
	defer c.Close()

	// Send the payload
	err = c.WriteMessage(websocket.TextMessage, []byte(payload))
	if err != nil {
		log.Printf("Error writing to WebSocket: %v", err)
		return
	}

	// Read the response
	_, message, err := c.ReadMessage()
	if err != nil {
		log.Printf("Error reading from WebSocket: %v", err)
		return
	}

	// Check for reflection
	if x.validator.Validate(payload, string(message)) {
		log.Printf("XSS found in WebSocket message: %s", message)
	}
}
