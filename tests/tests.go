// backend selector for testings
package tests

import (
	"time"
	"github.com/calce/sgo"
	"github.com/calce/sgo/mocks"
)

func mockBackend() *sgo.Backend {
	sm := mocks.New()
	backend, _ := sgo.NewBackend("yay", "http://connect.squareup.com/v1", "me", time.Second * 30, sm.Client)
	return backend
}

func liveBackend() *sgo.Backend {
	backend, _ := sgo.NewBackend("2H0GuQYJzJw97VYs_T1BUw", "https://connect.squareup.com/v1", "me", time.Second * 30, nil)
	return backend	
}

func GetTestBackend() *sgo.Backend {
	return mockBackend()
}

func GetMockBackend() *sgo.Backend {
	return mockBackend()
}