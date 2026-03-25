package main

import (
	"bufio"
	"bytes"
	"strings"

	// "encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	// "golang.org/x/tools/go/analysis/checker"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	// "net/url"
	// "net/http"
)

var (
	mytoken string
	myuser  string
)

// this is login function it do post login info in a url in json form package

func login(url string, username string, password string) {

	data := map[string]string{

		"username": username,
		"password": password,
	}

	jsondata, _ := json.Marshal(data)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsondata))

	if err != nil {

		fmt.Printf("\n Faild to sent post req : ", err)
		return
	} else {

		fmt.Printf(" \n sucessfully sented : ", resp.Status)
	}

	defer resp.Body.Close()

	bodybytes, _ := io.ReadAll(resp.Body)
	massage := string(bodybytes)

	parts := strings.Split(massage, ":")

	if len(parts) > 0 && parts[0] == "success" {

		savecradenshial(username, parts[2])
		mytoken = parts[2]

		fmt.Printf(" \n successfully login \n")
	}

}

// this is register function this do post register info to the srever in json form

func emailcheck(url string, email string, username string, password string) {

	data := map[string]string{

		"email":    email,
		"username": username,
	}

	jsondata, err := json.Marshal(data)
	if err != nil {

		fmt.Printf(" \n faild to marshel json data : ", err)
		return
	}
	resp, err := http.Post(url, "application/json-data", bytes.NewBuffer(jsondata))

	if err != nil {
		fmt.Printf(" \n failt to post register info : ", err)
		return
	} else {

		fmt.Printf(" \n successfuly posted register data : ", err)

	}

	defer resp.Body.Close()

	newbytes, _ := io.ReadAll(resp.Body)

	massage := string(newbytes)

	if massage == "success" {
		register("http://localhost:4040/confarmregister", email, username, password)
	}

}

func register(url string, email string, username string, password string) {

	var userinput string
	fmt.Printf(" \n Enter the otp : ")
	fmt.Scan(&userinput)

	data := map[string]string{

		"email":    email,
		"username": username,
		"password": password,
		"otp":      userinput,
	}

	jsondata, err := json.Marshal(data)
	if err != nil {

		fmt.Printf(" \n faild to marshel json data : ", err)
		return
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsondata))

	if err != nil {
		fmt.Printf(" \n failt to post register info : ", err)
		return
	} else {

		fmt.Printf(" \n successfuly posted register data : ", err)

	}

	defer resp.Body.Close()

	newbystes, _ := io.ReadAll(resp.Body)
	servermessage := string(newbystes)

	partsmessge := strings.Split(servermessage, ":")

	if len(partsmessge) > 0 && partsmessge[0] == "success" {
		fmt.Printf("\n login successfully \n")
		savecradenshial(username, partsmessge[2])
		mytoken = partsmessge[2]
	}

}

// forget password

func forgetpass(url string, email string) {

	data := map[string]string{

		"email": email,
	}

	jsondata, err := json.Marshal(data)

	if err != nil {

		fmt.Printf(" \n faild to marshal json data : %+s ", err)
		return
	}

	resp, err := http.Post(url, "/application/paasforget", bytes.NewBuffer(jsondata))

	if err != nil {

		fmt.Printf(" \n failt to sent the email to server : %+s ", err)
		return
	} else {
		fmt.Printf(" \n success fully sented forget password email : %+s", resp.Status)

	}

	defer resp.Body.Close()

}

func savecradenshial(username string, tokeen string) {

	contents := fmt.Sprintf("user=%s\ntoken=%s", username, tokeen)

	err := os.WriteFile(".env", []byte(contents), 0644)

	if err != nil {
		fmt.Printf("\n faild to save the credential : %s", err)
	}
	fmt.Printf("\n successflyy cradentional saved \n")

}

func chate(tusr string, token string, user string) {

	url := fmt.Sprintf("ws://localhost:4040/chat?user=%s&token=%s", user, token)

	conn, _, err := websocket.DefaultDialer.Dial(url, nil)

	if err != nil {
		fmt.Printf("\n handshake faild \n : %s", &err)
		return
	}

	defer conn.Close()

	fmt.Printf("\n We are in Cannected to the server \n")

	go func() {

		for {

			_, p, err := conn.ReadMessage()

			if err != nil {
				fmt.Printf("\n Faild to resive message : %s ", &err)
				return
			}

			fmt.Println(string(p))

		}

	}()

	scanner := bufio.NewScanner(os.Stdin)

	for {

		if scanner.Scan() {
			fmt.Printf("\n > ")
			text := scanner.Text()

			if text == "" {
				continue
			}

			msg := fmt.Sprintf("tusr:%s:user:%s", tusr, text)
			err := conn.WriteMessage(websocket.TextMessage, []byte(msg))

			if err != nil {
				fmt.Printf("\n %s ", &err)
				break
			}

		}
	}

}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("%s", err)
	}

	mytoken = os.Getenv("token")
	myuser = os.Getenv("user")

	// login("http://localhost:4040/login", "dra34ken", "paf32453")
	login("http://localhost:4040/login", "mikey2", "paf32453")

	// emailcheck("http://localhost:4040/signup", "mda891526@gmail.com", "mikey2", "paf32453")

	tosend := "dra34ken"

	chate(tosend, mytoken, myuser)

	// forgetpass("http://localhost:4040/forgetpass", "mda35345345@gmail.com")

}
