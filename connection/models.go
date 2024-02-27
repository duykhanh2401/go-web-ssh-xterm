package connection

import "golang.org/x/crypto/ssh"

type ptyRequestMsg struct {
	Term     string
	Columns  uint32
	Rows     uint32
	Width    uint32
	Height   uint32
	Modelist string
}

type ptyWindowChangeMsg struct {
	Columns uint32
	Rows    uint32
	Width   uint32
	Height  uint32
}

type Terminal struct {
	Columns uint32 `json:"cols"`
	Rows    uint32 `json:"rows"`
}

type WindowSize struct {
	High  int `json:"high"`
	Width int `json:"width"`
}

type SSHClient struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	IpAddress string `json:"ipaddress"`
	Port      int    `json:"port"`
	Session   *ssh.Session
	Client    *ssh.Client
	channel   ssh.Channel
}

func NewSSHClient() SSHClient {
	client := SSHClient{}
	client.Username = ""
	client.Port = 2252
	client.Password = ""
	client.IpAddress = ""
	return client
}
