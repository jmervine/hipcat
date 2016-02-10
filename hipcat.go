package main

import (
	"github.com/jmervine/hipcat/cli"
	"github.com/jmervine/hipcat/client"
	"github.com/jmervine/hipcat/config"

	"fmt"
)

func main() {
	cli.Run(func(cfg *config.Config) error {

		h, err := client.NewClient(cfg)
		if err != nil {
			return err
		}

		err = client.Notify(h, cfg)
		if err != nil {
			return err
		}

		fmt.Printf("%+v\n", string(cfg.Message))
		return nil

	})
}
