// a88a_prototype_a_min.go
// Prototype a minimalist automation script notifier

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Configuration holds the notification settings
type Configuration struct {
	NotificationInterval time.Duration
	NotificationTimeout  time.Duration
	NotificationEndpoint string
}

// Notifier is responsible for sending notifications
type Notifier struct {
	Config Configuration
}

// NewNotifier returns a new Notifier instance
func NewNotifier(config Configuration) *Notifier {
	return &Notifier{Config: config}
}

// SendNotification sends a notification to the specified endpoint
func (n *Notifier) SendNotification(message string) error {
	client := &http.Client{}
	req, err := http.NewRequest("POST", n.Config.NotificationEndpoint, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func main() {
	config := Configuration{
		NotificationInterval: 5 * time.Minute,
		NotificationTimeout:  30 * time.Second,
		NotificationEndpoint: "https://example.com/notify",
	}

	notifier := NewNotifier(config)

	for {
		select {
		case <-time.After(config.NotificationInterval):
			log.Println("Sending notification...")
			err := notifier.SendNotification("Automation script executed successfully!")
			if err != nil {
				log.Println(err)
			}
		}
	}
}