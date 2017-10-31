package login

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/google/go-querystring/query"
	"github.com/wsboyxb/robot/conf"
)

// sk hold sessionKey
type SK struct {
	SessionKey string `json:"sessionKey"`
}

//httpResp hold login server response
type resp struct {
	Code   int  `json:"code"`
	Status bool `json:"status"`
	Data   SK   `json:"data"`
}

// base hold the common parts of login server message
type req struct {
	Action    string `url:"action"`
	M         string `url:"m"`
	Account   string `url:"account"`
	Password  string `url:"password"`
	ChannelId int    `url:"channelId"`
}

func Register(account string) (err error) {
	m := req{
		"login.do",
		"register",
		account,
		conf.UserPWD,
		conf.ChannelID,
	}
	qs, err := query.Values(m)
	if err != nil {
		return
	}

	res, err := doPost(qs.Encode())
	if err != nil {
		return
	}

	var r resp
	err = json.Unmarshal(res, &r)
	if err != nil {
		return
	}

	// success
	if r.Status && r.Code == 0 {
		return
	}

	return fmt.Errorf("%t\t%d", r.Status, r.Code)
}

// Login for get session key
func Login(account string) (sessionKey string, err error) {
	m := req{
		"login.do",
		"login",
		account,
		conf.UserPWD,
		conf.ChannelID,
	}
	qs, err := query.Values(m)
	if err != nil {
		return
	}

	res, err := doPost(qs.Encode())
	if err != nil {
		return
	}

	var r resp
	err = json.Unmarshal(res, &r)
	if err != nil {
		return
	}

	// success
	if r.Status && r.Code == 0 {
		sessionKey = r.Data.SessionKey
		return
	}

	err = fmt.Errorf("%t\t%d", r.Status, r.Code)
	return
}

func doPost(message string) (resp []byte, err error) {
	req, err := http.NewRequest("POST", conf.LoginURL, strings.NewReader(message))
	if err != nil {
		return
	}
	req.Close = true

	c := &http.Client{}
	r, err := c.Do(req)
	if err != nil {
		return
	}
	defer r.Body.Close()

	resp, err = ioutil.ReadAll(r.Body)
	return
}
