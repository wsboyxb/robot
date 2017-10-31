package network

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/wsboyxb/robot/login"
	"github.com/wsboyxb/robot/utils"
	"github.com/zemirco/uid"
)

type LoginAction struct {
}

//step 0
//State 0: no player; 1:have player
//if InitLogin return state = 0, then run func CreatePlayer to create player
//else Go to func LoadPlayer
func (la LoginAction) InitLogin(c *TcpClient, sk string) error {
	initLoginMsg := login.SK{sk}
	data, _ := json.Marshal(initLoginMsg)
	a := reflect.TypeOf(la).Name()
	m := "initLogin"
	cb := fmt.Sprintf("%s.%s", a, m)
	data, err := utils.ActionToData(a, m, cb, data)
	if err != nil {
		return err
	}
	return c.GetMsgParser().Write(c, data)
}

//step 1.0
func (la LoginAction) CreatePlayer(c *TcpClient, sk string) error {
	d := struct {
		SK   string `json:"sessionKey"`
		Name string `json:"name"`
		Icon int    `json:"icon"`
	}{sk, uid.New(10), 1}

	data, _ := json.Marshal(d)
	a := reflect.TypeOf(la).Name()
	m := "createPlayer"
	cb := fmt.Sprintf("%s.%s", a, m)
	data, err := utils.ActionToData(a, m, cb, data)
	if err != nil {
		return err
	}
	return c.GetMsgParser().Write(c, data)
}

//step 1.1
//state=1 is ok
func (la LoginAction) LoadPlayer(c *TcpClient, p int) error {
	d := struct {
		PlayerId int `json:"playerId"`
	}{p}
	data, _ := json.Marshal(d)
	a := reflect.TypeOf(la).Name()
	m := "loadPlayer"
	cb := fmt.Sprintf("%s.%s", a, m)
	data, err := utils.ActionToData(a, m, cb, data)
	if err != nil {
		return err
	}
	return c.GetMsgParser().Write(c, data)
}

func (la LoginAction) InitLoginCB(c *TcpClient, data []byte) {
	var vo initLogin
	err := json.Unmarshal(data, &vo)
	if err != nil {
		fmt.Println(err)
		return
	}

	if vo.PlayerID == 0 {
		la.CreatePlayer(c, c.GetUser().SessionKey)
		return
	}
	c.GetUser().PlayerID = vo.PlayerID
	la.LoadPlayer(c, c.GetUser().PlayerID)
}

func (la LoginAction) CreatePlayerCB(c *TcpClient, data []byte) {
	var vo initLogin
	err := json.Unmarshal(data, &vo)
	if err != nil {
		fmt.Println(err)
		return
	}
	if vo.State == 1 {
		c.GetUser().PlayerID = vo.PlayerID
		la.LoadPlayer(c, c.GetUser().PlayerID)
	}
}

type InterfaceAction struct {
}

func (ifa InterfaceAction) LoadInterface(c *TcpClient) error {
	d := struct{}{}
	data, _ := json.Marshal(d)
	a := reflect.TypeOf(ifa).Name()
	m := "loadInterface"
	cb := fmt.Sprintf("%s.%s", a, m)
	data, err := utils.ActionToData(a, m, cb, data)
	if err != nil {
		return err
	}
	return c.GetMsgParser().Write(c, data)
}
