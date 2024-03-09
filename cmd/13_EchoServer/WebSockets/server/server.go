package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} // use default options

func echo(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)

		return
	}
	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			go log.Println("read:", err)

			break
		}

		go log.Printf("recived message: %s", message)

		errWriteBack := conn.WriteMessage(messageType, message)
		if errWriteBack != nil {
			go log.Println("write:", errWriteBack)

			break
		}
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, "ws://"+r.Host+"/echo")
}

func main() {
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/", home)

	port := ":7000"

	fmt.Printf("listening on http://localhost%s\n", port)

	log.Fatal(http.ListenAndServe(":7000", nil))
}
