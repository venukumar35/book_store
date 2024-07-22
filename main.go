package main

import "time"

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"strings"

// 	"github.com/gorilla/websocket"
// )

// var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
// 	return true
// }}

// var todoList []string

// func getCmd(input string) string {
// 	inputArr := strings.Split(input, " ")
// 	return inputArr[0]
// }

// func getMessage(input string) string {
// 	inputArr := strings.Split(input, " ")
// 	var result string
// 	for i := 1; i < len(inputArr); i++ {
// 		result += inputArr[i]
// 	}
// 	return result
// }

// func updateTodoList(input string) {
// 	tmpList := todoList
// 	todoList = []string{}
// 	for _, val := range tmpList {
// 		if val == input {
// 			continue
// 		}
// 		todoList = append(todoList, val)
// 	}
// }

// func main() {

// 	http.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
// 		// Upgrade upgrades the HTTP server connection to the WebSocket protocol.
// 		conn, err := upgrader.Upgrade(w, r, nil)
// 		if err != nil {
// 			log.Print("upgrade failed: ", err)
// 			return
// 		}
// 		defer conn.Close()

// 		// Continuosly read and write message
// 		// Inside your main function
// 		for {
// 			mt, message, err := conn.ReadMessage()
// 			if err != nil {
// 				//	if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
// 				log.Printf("error: %v", err)
// 				//	}
// 				break
// 			}
// 			fmt.Print(string(message), mt)
// 			input := string(message)
// 			cmd := getCmd(input)
// 			msg := getMessage(input)

// 			if cmd == "add" {
// 				todoList = append(todoList, msg)
// 			} else if cmd == "done" {
// 				updateTodoList(msg)
// 			}
// 			fmt.Println("todoList", todoList)
// 			output := "Current Todos: \n"
// 			for _, todo := range todoList {
// 				output += "\n - " + todo + "\n"
// 			}
// 			output += "\n----------------------------------------"

// 			message = []byte(output)

// 			// Check WebSocket state before writing
// 			if err := conn.WriteMessage(mt, message); err != nil {
// 				log.Println("write failed:", err)
// 				break
// 			}
// 		}

// 	})

// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		http.ServeFile(w, r, "websockets.html")
// 	})
// 	fmt.Printf("eneter 1")

// 	http.ListenAndServe(":8081", nil)
// }
/*******************web socket two*******************************/
// var upgrader = websocket.Upgrader{
// 	CheckOrigin: func(r *http.Request) bool {
// 		return true
// 	},
// }

// func handleWebSocket(w http.ResponseWriter, r *http.Request) {
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		fmt.Println("Error upgrading connection:", err)
// 		return
// 	}
// 	defer conn.Close()

// 	for {
// 		messageType, message, err := conn.ReadMessage()
// 		if err != nil {
// 			fmt.Println("Error reading message:", err)
// 			break
// 		}
// 		fmt.Printf("Received message from client: %s \n", message, messageType)

// 		err = conn.WriteMessage(messageType, message)
// 		if err != nil {
// 			fmt.Println("Error writing message:", err)
// 			break
// 		}
// 	}
// }

func main() {
	// go numbersPrint(100)

	go numbersPrint(1000)

	// go numbersPrint([]int{1, 2, 3, 4, 5})

	// go numbersPrint([]int{6, 7, 8, 9, 0})

	// go numbersPrint([]int{16, 17, 18, 19, 10})

	time.Sleep(time.Minute * 2)

	// http.HandleFunc("/websocket", handleWebSocket)
	// fmt.Println("WebSocket server listening on port 8080...")
	// http.ListenAndServe(":8080", nil)
}
