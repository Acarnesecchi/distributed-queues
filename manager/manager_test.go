package manager_test

import (
	"testing"

	"github.com/Acarnesecchi/distributed-queue/manager"
)

func TestNewConfig(t *testing.T) {
	conf := manager.NewConfig().WithListenAddr(":2000")
	if conf.ListenAddr != ":2000" {
		t.Errorf("conf.ListenAddr = %s; want :2000", conf.ListenAddr)
	}
}
