package main

import (
	"bufio"
	"bytes"
	"strings"

	// "time"

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

	url = fmt.Sprintf("%s/login", string(url))

	data := map[string]string{

		"username": username,
		"password": password,
	}

	jsondata, _ := json.Marshal(data)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsondata))

	if err != nil {

		fmt.Printf("\n Faild to sent post req : %v ", err)
		return
	} else {

		fmt.Printf(" \n sucessfully sented : %v ", resp.Status)
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

	urle := fmt.Sprintf("%s/signup", url)

	data := map[string]string{

		"email":    email,
		"username": username,
	}

	jsondata, err := json.Marshal(data)
	if err != nil {

		fmt.Printf(" \n faild to marshel json data : %v ", err)
		return
	}
	resp, err := http.Post(urle, "application/json-data", bytes.NewBuffer(jsondata))

	if err != nil {
		fmt.Printf(" \n failt to post register info : %v ", err)
		return
	} else {

		fmt.Printf(" \n successfuly posted register data : %v ", err)

	}

	defer resp.Body.Close()

	newbytes, _ := io.ReadAll(resp.Body)

	massage := string(newbytes)

	if massage == "success" {
		register(url, email, username, password)
	}

}

func register(url string, email string, username string, password string) {

	url = fmt.Sprintf("%s/confarmregister", url)

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

		fmt.Printf(" \n faild to marshel json data : %v ", err)
		return
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsondata))

	if err != nil {
		fmt.Printf(" \n failt to post register info : %v ", err)
		return
	} else {

		fmt.Printf(" \n successfuly posted register data : %v ", err)

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

func forgetpass(baseURL string, email string) {
	// 1. REQUEST THE OTP
	url := fmt.Sprintf("%s/forgetpass", baseURL)
	data := map[string]string{"email": email}
	jsondata, _ := json.Marshal(data)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsondata))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	parts := strings.Split(string(body), ":")

	if parts[0] == "done" {
		fmt.Printf("\n %s", parts[1])

		var otpInput string
		var newPassInput string
		fmt.Printf("\n Enter OTP: ")
		fmt.Scan(&otpInput)
		fmt.Printf("\n Enter New Password: ")
		fmt.Scan(&newPassInput)

		// 2. SEND OTP AND PASSWORD BACK
		// We use the same URL but add ?otp=...&user=...&new=...
		resetURL := fmt.Sprintf("%s/forgetpass?otp=%s&user=%s&new=%s", baseURL, otpInput, email, newPassInput)

		// We send nil for the body because the data is in the URL now
		resp2, err := http.Post(resetURL, "application/json", nil)
		if err != nil {
			return
		}
		defer resp2.Body.Close()

		finalBody, _ := io.ReadAll(resp2.Body)
		fmt.Printf("\n Final Result: %s", string(finalBody))
	}
}

func savecradenshial(username string, tokeen string) {

	contents := fmt.Sprintf("user=%s\ntoken=%s", username, tokeen)

	err := os.WriteFile(".env", []byte(contents), 0644)

	if err != nil {
		fmt.Printf("\n faild to save the credential : %v", err)
	}
	fmt.Printf("\n successflyy cradentional saved \n")

}

func chate(tusr string, token string, user string) {

	url := fmt.Sprintf("ws://localhost:4040/chat?user=%s&token=%s", user, token)

	conn, _, err := websocket.DefaultDialer.Dial(url, nil)

	if err != nil {
		fmt.Printf("\n handshake faild \n : %v", err)
		return
	}

	defer conn.Close()

	fmt.Printf("\n We are in Cannected to the server \n")

	go func() {

		for {

			_, p, err := conn.ReadMessage()

			if err != nil {
				fmt.Printf("\n Faild to resive message : %v ", err)
				return
			}

			fmt.Println(string(p))

		}

	}()

	scanner := bufio.NewScanner(os.Stdin)

	for {

		if scanner.Scan() {
			fmt.Printf(" TO --> %s \n > ", tusr)
			text := scanner.Text()

			if text == "" {
				continue
			}

			msg := fmt.Sprintf("tusr:%s:user:%s", tusr, text)
			err := conn.WriteMessage(websocket.TextMessage, []byte(msg))

			if err != nil {
				fmt.Printf("\n %v ", err)
				break
			}

		}
	}

}

// ACtions SRQ for sentfriend req RFQ for refect friend req AFQ for accept firend req
// DLF for delate form friend
func todo(url, token string, user string, action string, targetuser string) {

	var act string

	switch action {

	case "SRQ":
		act = "sentfreq"
	case "RFQ":
		act = "rejectfreq"
	case "AFQ":
		act = "acceptfreq"
	case "DLF":
		act = "delatfre"

	}

	Url := fmt.Sprintf("%s/do?user=%s&token=%s&act=%s&tar=%s", url, user, token, act, targetuser)

	resp, err := http.Post(Url, "application/json", nil)

	if err != nil {
		fmt.Printf("\n Failed to send friend request: %v", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Printf("\n Successfully action sented to %s!", targetuser)
	} else {
		fmt.Printf("\n Server returned error: %s", resp.Status)
	}

	rebytes, _ := io.ReadAll(resp.Body)

	message := string(rebytes)

	fmt.Printf(" \n %v  \n", message)
}

// func rejectfreindreq(url, token string, user string) {

// 	act := "rejectfreq"
// 	Url := fmt.Sprintf("%s/do?user=%s&token=%s&act=%s", url, user, token, act)

// 	resp, err := http.Post(Url, "application/json", nil)

// 	if err != nil {
// 		fmt.Printf("\n Failed to send friend request: %v", err)
// 		return
// 	}

// 	defer resp.Body.Close()

// 	if resp.StatusCode == http.StatusOK {
// 		fmt.Printf("\n Friend request sent to %s!", user)
// 	} else {
// 		fmt.Printf("\n Server returned error: %s", resp.Status)
// 	}
// }

// func acceptfreindreq(url, token string, user string) {

// 	act := "acceptfreq"
// 	Url := fmt.Sprintf("%s/do?user=%s&token=%s&act=%s", url, user, token, act)

// 	resp, err := http.Post(Url, "application/json", nil)

// 	if err != nil {
// 		fmt.Printf("\n Failed to send friend request: %v", err)
// 		return
// 	}

// 	defer resp.Body.Close()

// 	if resp.StatusCode == http.StatusOK {
// 		fmt.Printf("\n Friend request sent to %s!", user)
// 	} else {
// 		fmt.Printf("\n Server returned error: %s", resp.Status)
// 	}
// }

// func delatefreind(url, token string, user string) {

// 	act := "delatfre"
// 	Url := fmt.Sprintf("%s/do?user=%s&token=%s&act=%s", url, user, token, act)

// 	resp, err := http.Post(Url, "application/json", nil)

// 	if err != nil {
// 		fmt.Printf("\n Failed to send friend request: %v", err)
// 		return
// 	}

// 	defer resp.Body.Close()

// 	if resp.StatusCode == http.StatusOK {
// 		fmt.Printf("\n Friend request sent to %s!", user)
// 	} else {
// 		fmt.Printf("\n Server returned error: %s", resp.Status)
// 	}
// }

func main() {

	url := "http://localhost"
	port := ":4040"

	url = fmt.Sprintf("%v%v", url, port)

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("%v ", err)
	}

	mytoken = os.Getenv("token")
	myuser = os.Getenv("user")

	// login("http://localhost:4040/login", "dra34ken", "paf32453")
	// login(url, "mikey2", "paf32453")

	// emailcheck(url, "mda891526@gmail.com", "mikey2", "paf32453")

	todo(url, mytoken, myuser, "SRQ", "dra34ken")
	// tosend := "dra34ken"

	// // chate(tosend, mytoken, myuser)

	// forgetpass(url, "mda891526@gmail.com")

}
