package network

type User struct {
	TcpClient  *TcpClient
	SessionKey string
	PlayerID   int
}
