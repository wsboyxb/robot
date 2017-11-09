package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"

	"reflect"

	log "github.com/sirupsen/logrus"
	"github.com/wsboyxb/robot/conf"
	"github.com/wsboyxb/robot/login"
	"github.com/wsboyxb/robot/network"
)

type resultPair struct {
	idx int
	err error
}

var (
	loginAction             network.LoginAction
	interfaceAction         network.InterfaceAction
	activityAction          network.ActivityAction
	activityGlobalPVPAction network.ActivityGlobalPVPAction
	arenaAction             network.ArenaAction
	autoUpLevelAction       network.AutoUpLevelAction
	dailyActivityAction     network.DailyActivityAction
	deviceFactoryAction     network.DeviceFactoryAction
	drawAction              network.DrawAction
	guildAction             network.GuildAction
	payAction               network.PayAction
	playerAction            network.PlayerAction

	actions = []interface{}{
		interfaceAction,
		activityAction,
		activityGlobalPVPAction,
		arenaAction,
		autoUpLevelAction,
		dailyActivityAction,
		deviceFactoryAction,
		drawAction,
		guildAction,
		payAction,
		playerAction,
	}

	begin = flag.Int("b", 0, "begin user id")
)

func reg() {
	cnt := 2000
	r := make(chan resultPair, cnt)
	for i := 0; i < cnt; i++ {
		go func(i int) {
			id := fmt.Sprintf("%s%d", conf.AccountPrefix, i)
			r <- resultPair{
				i,
				login.Register(id),
			}
		}(i)
		time.Sleep(100 * time.Microsecond)
	}

	for i := 0; i < cnt; i++ {
		v := <-r
		fmt.Printf("%d,%v\n", v.idx, v.err)
	}
}
func lo11g() {
	cnt := 50
	r := make(chan resultPair, cnt)
	for i := 0; i < cnt; i++ {
		go func(i int) {
			id := fmt.Sprintf("%s%d", conf.AccountPrefix, i)
			_, err := login.Login(id)
			r <- resultPair{
				i,
				err,
			}
			//fmt.Print(sk)
		}(i)
		time.Sleep(10 * time.Microsecond)
	}

	for i := 0; i < cnt; i++ {
		v := <-r
		fmt.Printf("%d,%v\n", v.idx, v.err)
	}
}
func gologin(sk chan<- string) {
	cnt := 500
	for i := 0; i < cnt; i++ {
		id := fmt.Sprintf("%s%d", conf.AccountPrefix, i+*begin)
		s, err := login.Login(id)
		if err == nil {
			sk <- s
		} else {
			log.Errorln("login account err", err)
		}
		time.Sleep(10 * time.Microsecond)
	}
	close(sk)
}

func process(sk string) {
	//login http server gen session keys

	client := network.NewTcpClient(conf.GameServerIP, conf.GameServerPort)
	defer client.Close()
	go client.Run()

	u := &network.User{
		TcpClient:  client,
		SessionKey: sk,
	}
	client.SetUser(u)
	for client.GetUser().PlayerID == 0 {
		loginAction.InitLogin(client, sk)
		time.Sleep(time.Second * 2)
	}
	args := []reflect.Value{reflect.ValueOf(client)}
	for {
		time.Sleep(3 * time.Second)
		for _, act := range actions {

			t := reflect.ValueOf(act)
			for i := 0; i < t.NumMethod(); i++ {
				time.Sleep(1 * time.Second)
				go t.Method(i).Call(args)
			}
		}
	}
}
func main() {
	flag.Parse()
	//reg()
	//var pa network.PlayerAction
	sks := make(chan string, 100)
	go gologin(sks)
	var i int
	for {
		time.Sleep(500 * time.Microsecond)
		sk, ok := <-sks
		if ok {
			i++
			go process(sk)
		}
		if i >= 500 {
			break
		}
		log.Infoln(i)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	log.Infof("Leaf closing down (signal: %v)", sig)
}
func ac() {
	//reg()
	//log()
	//fmt.Printf("%v", LoginAction.InitLogin)
	//t := reflect.TypeOf(LoginAction.InitLogin)
	sk, err := login.Login("Q4Oi")
	fmt.Println(sk, err)
	client := network.NewTcpClient(conf.GameServerIP, conf.GameServerPort)
	//c.Send(nil)
	var a network.LoginAction
	a.InitLogin(client, "")
	//a.LoadPlayer(client)
	//for i := 0; i < 100; i++ {
	//	go a.InitLogin(client)
	//}
	time.Sleep(time.Second)
	go client.Run()
	//go a.InitLogin(client)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	fmt.Printf("Leaf closing down (signal: %v)", sig)
}
