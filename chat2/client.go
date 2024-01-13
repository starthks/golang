package main

import "github.com/gorilla/websocket"

// client는 한 명의 채팅 사용자를 나타낸다.
type client struct {
	socket *websocket.Conn //socket은 이 클라이언트의 웹 소캣이다
	send   chan []byte     //send는 메시지가 전송되는 채널이다
	room   *room           //room은 클라이언트가 채팅하는 방이다.
}

func (c *client) read() {
	defer c.socket.Close()
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		c.room.forward <- msg
	}
}

func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}
