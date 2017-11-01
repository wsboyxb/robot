package network

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
	"github.com/wsboyxb/robot/msg"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	f := &log.TextFormatter{}
	f.DisableTimestamp = true
	f.ForceColors = true
	log.SetFormatter(f)

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

type TcpClient struct {
	addr *net.TCPAddr
	conn *net.TCPConn
	sync.RWMutex
	msgParser *MsgParser
	user      *User
}

func NewTcpClient(ip string, port int) *TcpClient {
	addr := &net.TCPAddr{
		IP:   net.ParseIP(ip),
		Port: port,
	}

	conn, err := net.DialTCP("tcp", nil, addr)
	if err == nil {
		client := &TcpClient{
			addr:      addr,
			conn:      conn,
			msgParser: NewMsgParser(),
		}
		return client
	} else {
		panic(err)
	}
}

func (this *TcpClient) SetUser(u *User) {
	this.user = u
}

func (this *TcpClient) Run() {
	for {
		data, err := this.msgParser.Read(this)
		if err != nil {
			//fmt.Println(err)
			log.Errorln(err)
			break
		}

		var resp msg.HMResponse
		err = proto.Unmarshal(data, &resp)
		if err != nil {
			break
		}

		if resp.GetIsCompress() {
			gz, err := gzip.NewReader(bytes.NewReader(resp.GetData()))
			if err != nil {
				break
			}

			var ret bytes.Buffer
			_, err = ret.ReadFrom(gz)
			if err != nil {
				break
			}
			resp.Data = ret.Bytes()
		}

		var rs map[string]interface{}
		err = json.Unmarshal(resp.GetData(), &rs)
		var str string
		if err != nil {
			log.Errorln("json unmarshal err", string(resp.GetData()), err)
		} else {
			for k, v := range rs {
				str += fmt.Sprintf("%s = %v ", k, v)
			}
			//log.Print(str)
		}

		//log.Println(len(rs))
		sublen := len(str)
		if sublen > 120 {
			sublen = 120
		}
		str = str[:sublen]
		log.Printf("%d,%d,%s,[%s]", this.user.PlayerID, resp.GetCode(), resp.GetCallback(), str)

		if resp.GetCallback() == "LoginAction.initLogin" {
			var la LoginAction
			la.InitLoginCB(this, resp.GetData())
		} else if resp.GetCallback() == "LoginAction.createPlayer" {
			var la LoginAction
			la.CreatePlayerCB(this, resp.GetData())
		}
	}
}

func (this *TcpClient) ReConnection() bool {
	for i := 0; i < 10; i++ {
		conn, err := net.DialTCP("tcp", nil, this.addr)
		if err == nil {
			this.conn = conn
			return true
		} else {
			time.Sleep(3 * time.Second)
		}
	}
	return false
}

func (this *TcpClient) Send(data []byte) error {
	this.Lock()
	defer this.Unlock()

	if _, err := this.conn.Write(data); err != nil {
		//fmt.Printf("rpc client send data error.reason: %s", err)
		log.Errorf("rpc client send data error.reason: %s", err)
		return err
	}
	return nil
}

func (this *TcpClient) Close() {
	this.Lock()
	this.conn.Close()
	this.Unlock()
}

func (this *TcpClient) GetConnection() *net.TCPConn {
	return this.conn
}

func (this *TcpClient) GetMsgParser() *MsgParser {
	return this.msgParser
}

func (this *TcpClient) GetUser() *User {
	return this.user
}

func tool(conn net.Conn, act, mtd string) {
	data, _ := json.Marshal(struct{}{})
	req := &msg.HMRequest{
		Action:   proto.String(act),
		Method:   proto.String(mtd),
		Data:     data,
		Callback: proto.String("fuck"),
		//ChannelId: proto.Int32(int32(conf.ChannelID)),
	}

	sz := proto.Size(req)
	buff := make([]byte, sz+4)
	binary.BigEndian.PutUint32(buff, uint32(sz))

	msgByte, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}

	copy(buff[4:], msgByte)

	conn.Write(buff)
}
