package integration

import (
	"os"
	"runtime"
	"testing"

	"github.com/opencontainers/runc/libcontainer"
	_ "github.com/opencontainers/runc/libcontainer/nsenter"

	"github.com/sirupsen/logrus"
)

// init runs the libcontainer initialization code because of the busybox style needs
// to work around the go runtime and the issues with forking
func init() {
	if len(os.Args) < 2 || os.Args[1] != "init" {
		return
	}
	runtime.GOMAXPROCS(1)
	runtime.LockOSThread()
	if err := libcontainer.StartInitialization(); err != nil {
		logrus.Fatal(err)
	}
}

func TestMain(m *testing.M) {
	logrus.SetOutput(os.Stderr)
	logrus.SetLevel(logrus.InfoLevel)

	ret := m.Run()
	os.Exit(ret)
}
