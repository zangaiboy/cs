package asr

import (
	"flag"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func ServerRegister() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", echo)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
	// 解决跨域问题
	return true
}}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		if mt == websocket.TextMessage {
			handText(mt, c, message)
		} else if mt == websocket.BinaryMessage {
			handleBinary(mt, c, message)
		}
	}
}

func handText(mt int, c *websocket.Conn, message []byte) {
	log.Printf("asr收到用户的文本信息：%s", message)
	err := c.WriteMessage(mt, message)
	if err != nil {
		log.Println("write:", err)
	}
}

func handleBinary(mt int, c *websocket.Conn, message []byte) {
	log.Println("收到语音文件")
	err := c.WriteMessage(mt, []byte(message))
	if err != nil {
		log.Println("write:", err)
	}
}
