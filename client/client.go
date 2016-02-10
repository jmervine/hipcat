package client

import (
	"github.com/jmervine/hipcat/config"

	"net/url"

	"github.com/jmervine/hipcat/Godeps/_workspace/src/github.com/tbruyelle/hipchat-go/hipchat"
)

func NewClient(cfg *config.Config) (*hipchat.Client, error) {
	var err error

	client := hipchat.NewClient(cfg.Token)

	if cfg.Host != "" {
		baseURL, err := url.Parse(cfg.Host)

		if err != nil {
			return client, err
		}

		client.BaseURL = baseURL
	}

	return client, err
}

// some simple wrappers for common actions
func Notify(client *hipchat.Client, cfg *config.Config) error {
	req := &hipchat.NotificationRequest{
		Message: cfg.FormattedMessage(),
		Notify:  true,
	}

	_, err := client.Room.Notification(cfg.Room, req)

	return err
}
