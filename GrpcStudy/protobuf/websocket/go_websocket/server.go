package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func main() {
	// http 升级
	http.HandleFunc("/ws", Handler(
		&websocket.Upgrader{
			//解决同源问题
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	))

	http.ListenAndServe(":9090", nil)
}

func Handler(u *websocket.Upgrader) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := u.Upgrade(w, r, nil)
		defer conn.Close()
		if err != nil {
			log.Printf("upgrade err %v", err.Error())
			return
		}

		// other srv
		// 其他的业务逻辑,如接入rmq,mqtt处理消息 or other
		msgs := make(chan []byte)

		// 根据业务去区分 the srv is recv msg or push msg
		// 这里我对msg同时进行recv && push
		// 引入chan模式采用一收一发
		//对标实际业务为物联网,通过mqtt发送消息give me,我进行业务处理后返回一定的数据或其他处理
		done := make(chan struct{})
		go recvMsg(conn, msgs, done)

		go sendMsg(conn, msgs, done)

		<-done
		fmt.Printf("sk end...")
	}
}

func sendMsg(conn *websocket.Conn, msgs chan []byte, done chan struct{}) {
	for {
		select {
		case msg := <-msgs:
			//err := conn.WriteJSON(msg)
			err := conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Printf("cannot write JSON:%s\n", err.Error())
			}
		case <-done:
			return
		}
	}
}

func recvMsg(conn *websocket.Conn, msgs chan []byte, done chan struct{}) {
	for {
		// 这里写入recv 到的消息进行处理等
		_, data, err := conn.ReadMessage()
		if err != nil {
			if !websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure, websocket.CloseNoStatusReceived) {
				log.Printf("unexploggerected read error:%v\n", err.Error())
			}
			break
		}

		// 模拟处理
		//marshal, err := json.Marshal(data)
		//if err != nil {
		//	log.Printf("json marshal error:%v\n", err.Error())
		//	done <- struct{}{}
		//	break
		//}
		// chan通信交由sendMsg
		msgs <- data
	}
}
