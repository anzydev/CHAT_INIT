package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

func main() {

	userrname := bufio.NewReader(os.Stdin)
	fmt.Printf("\n Enter ur user name : \n ")
	user, _ := userrname.ReadString('\n')

	user = strings.TrimSpace(user)

	url := fmt.Sprintf("ws://localhost:6060/chat?user=%s", user)

	conn, _, err := websocket.DefaultDialer.Dial(url, nil)

	if err != nil {
		log.Fatal(" faild to cannect server ! ; ", err)
	}

	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Printf("\n > ")

		if !scanner.Scan() {
			break
		}

		text := scanner.Text()

		if text == "" {
			continue
		}
		err := conn.WriteMessage(websocket.TextMessage, []byte(text))

		if err != nil {
			fmt.Printf(" Write error : %v ", err)
		}

		// _, replay, err := conn.ReadMessage()

		// if err != nil {
		// 	fmt.Printf(" \n faild to read server reaplya : %v ", err)
		// 	break
		// }
		// fmt.Printf("\n Server sented : %v", string(replay))

	}

	fmt.Printf(" \n hello  \n")

}
