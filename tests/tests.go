// backend selector for testings
package tests

import (	
	"log"
	"time"
	"io/ioutil"
	"path/filepath"	
	"github.com/calce/sgo"
	"github.com/calce/sgo/mocks"
	"github.com/BurntSushi/toml"
)

type Config struct {
	Token string
	Timeout int
}

func mockBackend() *sgo.Backend {
	sm := mocks.New("../mocks")
	backend := sgo.NewBackend("yay", "http://connect.squareup.com/v1", "me", time.Second * 30, sm.Client)
	return backend
}

func GetliveBackend(token string, timeout int) *sgo.Backend {
	backend := sgo.NewBackend(token, "https://connect.squareup.com/v1", "me", time.Second * time.Duration(timeout), nil)
	return backend	
}

func GetTestBackend() *sgo.Backend {
	return mockBackend()
	var config Config
	
	path, err := filepath.Abs("../config.toml")
	if err != nil { panic("Config file not found") }
	
	log.Println("Loading config file", path)
	
	tomlData, err := ioutil.ReadFile(path)
	if err != nil { panic("Unable to read configuration file") }
	
	if _, err = toml.Decode(string(tomlData[:]), &config); err != nil {
		log.Println(err)
	  panic("Unable to load configurations")
	}	
	backend := sgo.NewBackend(config.Token, "https://connect.squareup.com/v1", "me", time.Second * time.Duration(config.Timeout), nil)
	return backend	
}

func GetMockBackend() *sgo.Backend {
	return mockBackend()
}