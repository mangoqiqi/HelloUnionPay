package server

import (
	"github.com/vinkdong/gox/log"
	"testing"
)

func TestStartServers(t *testing.T) {
	err := StartServers(8080)
	if err != nil {
		log.Info(err)
	}
}

func TestNetworkCheckHandler(t *testing.T)  {
}

