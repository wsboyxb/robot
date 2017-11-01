package network

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/wsboyxb/robot/login"
	"github.com/wsboyxb/robot/utils"
	"github.com/zemirco/uid"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

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

func (ifa InterfaceAction) ChangeOutskirtsInterface(c *TcpClient) error {
	a := reflect.TypeOf(ifa).Name()
	m := "changeOutskirtsInterface"
	return innerSendNoData(c, a, m)
}

func (ifa InterfaceAction) GetGametips(c *TcpClient) error {
	a := reflect.TypeOf(ifa).Name()
	m := "getGametips"
	return innerSendNoData(c, a, m)
}

func (ifa InterfaceAction) GetRifferenceInfo(c *TcpClient) error {
	a := reflect.TypeOf(ifa).Name()
	m := "getRifferenceInfo"
	return innerSendNoData(c, a, m)
}

func (ifa InterfaceAction) GetShortcutInfo(c *TcpClient) error {
	a := reflect.TypeOf(ifa).Name()
	m := "getShortcutInfo"
	return innerSendNoData(c, a, m)
}

func (ifa InterfaceAction) LoadInterface(c *TcpClient) error {
	a := reflect.TypeOf(ifa).Name()
	m := "loadInterface"
	return innerSendNoData(c, a, m)
}

type ActivityAction struct {
}

func (act ActivityAction) OpenActivityList(c *TcpClient) error {
	a := reflect.TypeOf(act).Name()
	m := "openActivityList"
	return innerSendNoData(c, a, m)
}

type ActivityGlobalPVPAction struct {
}

func (act ActivityGlobalPVPAction) OpenActivityGlobalPVP(c *TcpClient) error {
	a := reflect.TypeOf(act).Name()
	m := "openActivityGlobalPVP"
	return innerSendNoData(c, a, m)
}

type ArenaAction struct {
}

func (act ArenaAction) OpenArena(c *TcpClient) error {
	a := reflect.TypeOf(act).Name()
	m := "openArena"
	return innerSendNoData(c, a, m)
}
func (act ArenaAction) GetRank(c *TcpClient) error {
	a := reflect.TypeOf(act).Name()
	m := "getRank"
	return innerSendNoData(c, a, m)
}

type AutoUpLevelAction struct {
}

func (act AutoUpLevelAction) RequestBuildAutoUp(c *TcpClient) error {
	a := reflect.TypeOf(act).Name()
	m := "requestBuildAutoUp"
	return innerSendNoData(c, a, m)
}

type DailyActivityAction struct {
}

func (act DailyActivityAction) GetFreeEnergy(c *TcpClient) error {
	a := reflect.TypeOf(act).Name()
	m := "getFreeEnergy"
	t := rand.Intn(2)
	d := struct {
		Type int `json:"type"`
	}{101 + t}
	return innerSend(c, a, m, d)
}

func (act DailyActivityAction) OpenDailyActivity(c *TcpClient) error {
	a := reflect.TypeOf(act).Name()
	m := "openDailyActivity"
	return innerSendNoData(c, a, m)
}

type DeviceFactoryAction struct {
}

func (act DeviceFactoryAction) GetItemsNum(c *TcpClient) error {
	a := reflect.TypeOf(act).Name()
	m := "getItemsNum"
	return innerSendNoData(c, a, m)
}

type DrawAction struct {
}

func (act DrawAction) GetCostItemNum(c *TcpClient) error {
	a := reflect.TypeOf(act).Name()
	m := "getCostItemNum"
	return innerSendNoData(c, a, m)
}

type GuildAction struct {
}

func (act GuildAction) ShowGuildTips(c *TcpClient) error {
	a := reflect.TypeOf(act).Name()
	m := "showGuildTips"
	return innerSendNoData(c, a, m)
}

func (act GuildAction) SearchGuildPlayer(c *TcpClient) error {
	a := reflect.TypeOf(act).Name()
	m := "searchGuildPlayer"
	return innerSendNoData(c, a, m)
}

func (act GuildAction) GetGuildPlayerInfo(c *TcpClient) error {
	a := reflect.TypeOf(act).Name()
	m := "getGuildPlayerInfo"
	return innerSendNoData(c, a, m)
}

func (act GuildAction) GetContrRank(c *TcpClient) error {
	a := reflect.TypeOf(act).Name()
	m := "getContrRank"
	return innerSendNoData(c, a, m)
}

func (act GuildAction) GetGuildApplyList(c *TcpClient) error {
	a := reflect.TypeOf(act).Name()
	m := "getGuildApplyList"
	return innerSendNoData(c, a, m)
}

func (act GuildAction) OpenGuildTech(c *TcpClient) error {
	a := reflect.TypeOf(act).Name()
	m := "openGuildTech"
	return innerSendNoData(c, a, m)
}

func (act GuildAction) OpenGuildShop(c *TcpClient) error {
	a := reflect.TypeOf(act).Name()
	m := "openGuildShop"
	return innerSendNoData(c, a, m)
}

func (act GuildAction) GetJoinGuildGift(c *TcpClient) error {
	a := reflect.TypeOf(act).Name()
	m := "getJoinGuildGift"
	return innerSendNoData(c, a, m)
}

func (act GuildAction) OpenGuildActive(c *TcpClient) error {
	a := reflect.TypeOf(act).Name()
	m := "openGuildActive"
	return innerSendNoData(c, a, m)
}

func (act GuildAction) ShowGuildActiveRank(c *TcpClient) error {
	a := reflect.TypeOf(act).Name()
	m := "showGuildActiveRank"
	return innerSendNoData(c, a, m)
}

func (act GuildAction) GetWelfareRes(c *TcpClient) error {
	a := reflect.TypeOf(act).Name()
	m := "getWelfareRes"
	return innerSendNoData(c, a, m)
}

func (act GuildAction) ShowGuildWelfare(c *TcpClient) error {
	a := reflect.TypeOf(act).Name()
	m := "showGuildWelfare"
	return innerSendNoData(c, a, m)
}

func (act GuildAction) OpenGuildFb(c *TcpClient) error {
	a := reflect.TypeOf(act).Name()
	m := "openGuildFb"
	return innerSendNoData(c, a, m)
}

func (act GuildAction) OpenGuild(c *TcpClient) error {
	a := reflect.TypeOf(act).Name()
	m := "openGuild"
	return innerSendNoData(c, a, m)
}

func (act GuildAction) OpenGuildInfo(c *TcpClient) error {
	a := reflect.TypeOf(act).Name()
	m := "openGuildInfo"
	d := struct {
		Page int `json:"nowPage"`
		Type int `json:"type"`
	}{0, 1 + rand.Intn(2)}
	return innerSend(c, a, m, d)
}

func (act GuildAction) ChooseCanJoinGuild(c *TcpClient) error {
	a := reflect.TypeOf(act).Name()
	m := "chooseCanJoinGuild"
	d := struct {
		Page int `json:"nowPage"`
	}{0}
	return innerSend(c, a, m, d)
}

func (act GuildAction) SearchGuild(c *TcpClient) error {
	a := reflect.TypeOf(act).Name()
	m := "searchGuild"
	d := struct {
		Name string `json:"guildName"`
		Page int    `json:"nowPage"`
	}{uid.New(10), 0}
	return innerSend(c, a, m, d)
}

func (act GuildAction) NoOpenGuildList(c *TcpClient) error {
	a := reflect.TypeOf(act).Name()
	m := "noOpenGuildList"
	d := struct {
		Page int `json:"nowPage"`
	}{0}
	return innerSend(c, a, m, d)
}

type PayAction struct {
}

func (act PayAction) GetMonthCardGold(c *TcpClient) error {
	a := reflect.TypeOf(act).Name()
	m := "getMonthCardGold"
	return innerSendNoData(c, a, m)
}

func (act PayAction) IsSpecialMonthCard(c *TcpClient) error {
	a := reflect.TypeOf(act).Name()
	m := "isSpecialMonthCard"
	return innerSendNoData(c, a, m)
}

func (act PayAction) GetMonthCardInfo(c *TcpClient) error {
	a := reflect.TypeOf(act).Name()
	m := "getMonthCardInfo"
	return innerSendNoData(c, a, m)
}

type PlayerAction struct {
}

func (act PlayerAction) OpenWarShop(c *TcpClient) error {
	a := reflect.TypeOf(act).Name()
	m := "openWarShop"
	return innerSendNoData(c, a, m)
}

func (act PlayerAction) Heart(c *TcpClient) error {
	a := reflect.TypeOf(act).Name()
	m := "heart"
	return innerSendNoData(c, a, m)
}

func (act PlayerAction) GetCombatInfo(c *TcpClient) error {
	a := reflect.TypeOf(act).Name()
	m := "getCombatInfo"
	return innerSendNoData(c, a, m)
}

func (act PlayerAction) GetNotice(c *TcpClient) error {
	a := reflect.TypeOf(act).Name()
	m := "getNotice"
	return innerSendNoData(c, a, m)
}

func (act PlayerAction) GetPlayerInfo(c *TcpClient) error {
	a := reflect.TypeOf(act).Name()
	m := "getPlayerInfo"
	return innerSendNoData(c, a, m)
}

func (act PlayerAction) GetProtectVolume(c *TcpClient) error {
	a := reflect.TypeOf(act).Name()
	m := "getProtectVolume"
	return innerSendNoData(c, a, m)
}

func (act PlayerAction) Rename(c *TcpClient) error {
	a := reflect.TypeOf(act).Name()
	m := "rename"
	d := struct {
		Name string `json:"name"`
	}{uid.New(10)}
	return innerSend(c, a, m, d)
}

func (act PlayerAction) SetUseGoldConfirm(c *TcpClient) error {
	a := reflect.TypeOf(act).Name()
	m := "setUseGoldConfirm"
	d := struct {
		UseGoldConfirm int `json:"useGoldConfirm"`
	}{1}
	return innerSend(c, a, m, d)
}

func (act PlayerAction) SetAutoDef(c *TcpClient) error {
	a := reflect.TypeOf(act).Name()
	m := "setAutoDef"
	d := struct {
		AutoDef int `json:"autoDef"`
	}{1}
	return innerSend(c, a, m, d)
}

func (act PlayerAction) modifyPlayerName(c *TcpClient) error {
	a := reflect.TypeOf(act).Name()
	m := "modifyPlayerName"
	d := struct {
		Name  string `json:"name"`
		Icon  int    `json:"icon"`
		Guide int    `json:"guide"`
	}{uid.New(10), 1, 1}
	return innerSend(c, a, m, d)
}

func innerSendNoData(c *TcpClient, a, m string) error {
	d := struct{}{}
	return innerSend(c, a, m, d)
}

func innerSend(c *TcpClient, a, m string, d interface{}) error {
	data, _ := json.Marshal(d)
	cb := fmt.Sprintf("%s.%s", a, m)
	data, err := utils.ActionToData(a, m, cb, data)
	if err != nil {
		return err
	}
	return c.GetMsgParser().Write(c, data)
}
