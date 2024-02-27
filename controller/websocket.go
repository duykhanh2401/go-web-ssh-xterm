package controller

import (
	"go-web-ssh/connection"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WSHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var err error
		msg := ctx.DefaultQuery("msg", "")
		cols := ctx.DefaultQuery("cols", "159")
		rows := ctx.DefaultQuery("rows", "24")
		col, _ := strconv.Atoi(cols)
		row, _ := strconv.Atoi(rows)
		terminal := connection.Terminal{
			Columns: uint32(col),
			Rows:    uint32(row),
		}

		log.Println(msg)

		sshClient, err := connection.DecodedMsgToSSHClient(msg)
		if err != nil {
			ctx.Error(err)
			return
		}

		conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			ctx.Error(err)
			return
		}

		err = sshClient.GenerateClient()
		if err != nil {
			conn.WriteMessage(1, []byte(err.Error()))
			conn.Close()
			return
		}

		sshClient.RequestTerminal(&terminal)
		sshClient.Connect(conn)
	}
}
