package sql

import (
	"github.com/vinkdong/gox/log"
	"testing"
)

func TestNewEngine(t *testing.T) {
	engine, err := NewEngine("test", "root","root","192.168.56.11","3306")
	if err != nil {
		log.Error(err)
	}
	err = engine.Ping()
	if err != nil {
		log.Error(err)
	}
}
