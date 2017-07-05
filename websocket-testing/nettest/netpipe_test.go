package main

import (
	"net"
	"golang.org/x/net/nettest"
	"testing"
)

func TestNetPipe(t *testing.T) {
	nettest.TestConn(t,
		func() (c1, c2 net.Conn, stop func(), err error) {
			c1, c2 = net.Pipe()
			stop = func() { c1.Close(); c2.Close() }
			return
		},
	)
}

func main() {
	testing.RunTests(
		func(_, _ string) (bool, error) { return true, nil },
		[]testing.InternalTest{
			{Name: "TestNetPipe", F: TestNetPipe},
		},
	)
}
