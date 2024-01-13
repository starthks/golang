package main

import (
	"chat2/trace"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type room struct {
	forward chan []byte      //forward는 수신 메시지를 보관하는 채널
	join    chan *client     //방에 들어오려는 클라이언트를 위한 채널
	leave   chan *client     //방을 나가길 원하는 클라이언트를 위한 채널
	clients map[*client]bool //현재 채팅방에 있는 모든 클라이언트를 보유
	tracer  trace.Tracer
}

// 복잡성 제거를 위한 헬퍼 함수 사용
// newRoom은 새로운 방을 만든다.
func newRoom() *room {
	return &room{
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
		tracer: trace.Off(),
	}
}

// 채팅방 안의 join, leave, forward를 계속 모니터링
func (r *room) run() {
	for {
		select { //한번에 한 케이스 코드만 실행한다는 점이 중요
		case client := <-r.join:
			r.clients[client] = true
			r.tracer.Trace("New client joined")
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.send)
			r.tracer.Trace("Client left")
			//forward채널에서 메시지를 받으면 모든 클라이언트를 반복하고
			// 각 클라이언트의 send채널에 메시지를 추가한다.
			// 그런 다음 write매소드가 이를 받아들여 소켓에서 브라우저로 보낸다.
		case msg := <-r.forward:
			for client := range r.clients {
				client.send <- msg
				r.tracer.Trace("--sent to client")
			}
		}
	}
}

// 채팅방을 HTTP 핸들러로 변환
// 웹 소캣을 사용하려면 websocket.Upgrader 타입을 사용해 HTTP연결을 업그레이드
const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

// websocket.Upgrader 타입을 한개만 만들고 그런 다음 ServeHTTP메소드를
var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize,
	WriteBufferSize: socketBufferSize}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 통해 요청이 들어오면 upgrader.Upgrade 메소드를 호출해 소켓을 가져온다.
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}
	client := &client{
		socket: socket,
		send:   make(chan []byte, messageBufferSize),
		room:   r,
	}
	// 문제가 없다면 client를 생성해 join채널에 전달
	r.join <- client
	defer func() { r.leave <- client }()
	go client.write()
	client.read()
}
