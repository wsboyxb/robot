package conf

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var Server struct {
	TCPAddr string
}

func finit() {
	data, err := ioutil.ReadFile("conf/server.json")
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &Server)
	if err != nil {
		log.Fatal("%v", err)
	}
}
