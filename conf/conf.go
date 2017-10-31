package conf

import "time"

var (
	//LoginURL       = "http://119.29.254.55:8095" //正式服
	LoginURL = "http://192.168.1.243:8095" //正式服
	//GameServerIP   = "139.199.157.127"
	GameServerIP   = "192.168.1.42"
	GameServerPort = 9097

	AccountPrefix = "debug1030-"
	UserPWD       = "123456"
	ChannelID     = 114

	// Duration for every cmd
	Duration = 2 * time.Second

	// RcvCnt for receive msg from server loop
	RcvCnt = 10000
)
