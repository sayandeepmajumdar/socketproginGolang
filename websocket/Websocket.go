// websocket.go
package main

import (
	"fmt"
	"net/http"
	"net"
	"log"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
}

func main() {
	 //Read msg from browser and sent to server as well as to browser again

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

		for {
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			// Print the message to the console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			// Write message back to browser
			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}

		
	})

  //Checking which ports are online and unreachable or closed
   http.HandleFunc("/portcheck", func(w http.ResponseWriter, r *http.Request) {
	    var status string
	        conn2, err := net.Dial("tcp", "127.0.0.1:80")
	        if err != nil {
	                log.Println("Connection error:", err)
	                status = "Unreachable"
	                fmt.Fprintf(w, "Hello, Port is: %s\n",status)
	        } else {
	                status = "Online"
	                fmt.Fprintf(w, "Hello, Port is: %s\n", status)
	                defer conn2.Close()
	        }
			log.Println(status)
	})

//Particuler html where you can see output
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websocket.html")
	})



	http.ListenAndServe(":8080", nil)
}