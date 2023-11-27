package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var userConnections = make(map[string]*websocket.Conn)
var broadcast = make(chan Message)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil) // upgrade http request to websocket
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	userConnections[conn] = true // add connection to map

	for {
		var msg Message = Message{}
		err := conn.ReadJSON(&msg) // read message from websocket
		if err != nil {
			fmt.Printf("got error: %v", err)
			delete(userConnections, conn)
			return
		}
	}
}

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/ws", handleConnections)

	fmt.Println("starting server...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
