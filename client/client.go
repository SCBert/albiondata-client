package client

import (
	"github.com/Sirupsen/logrus"
	"github.com/regner/albiondata-client/log"
)

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (client *Client) Run() {
	log.SetFormatter(&logrus.TextFormatter{FullTimestamp: true, DisableSorting: true})

	log.Info("Starting the Client...")
	log.Info("This client was built off of the Albiondata Project client")

	createDispatcher()

	if ConfigGlobal.Offline {
		processOfflinePcap(ConfigGlobal.OfflinePath)
	} else {
		pw := newProcessWatcher()
		pw.run()
	}
}
