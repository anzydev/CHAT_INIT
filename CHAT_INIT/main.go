package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/charmbracelet/lipgloss"
)

// structure of the json data base

type user struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
	Password string `json:"passowrd"`
}

// json file read

func jsonreade() ([]user, error) {

	var jsondata []user
	bytes, err := os.ReadFile("database.json")
	if err != nil {

		return nil, err

	}

	err = json.Unmarshal(bytes, &jsondata)
	if err != nil {

		return nil, err
	}

	return jsondata, nil

}

// the login funciton

func login(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/login" {

		fmt.Printf(" login page faild ! : %s ", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "\n Login page is active \n")

	var incamingdata struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	json.NewDecoder(r.Body).Decode(&incamingdata)

	// loading the data base

	jsondata, _ := jsonreade()

	userfound := false
	passwordfund := false

	// checking the user name is in data base or not

	for _, user := range jsondata {

		if user.UserName == incamingdata.Username {

			userfound = true
			if user.Password == incamingdata.Password {

				passwordfund = true
				break

			}

		}

	}

	if !userfound {
		fmt.Fprintf(w, " user not found ")
	} else if !passwordfund {
		fmt.Fprintf(w, " passsowrd not found ")
	} else {
		fmt.Fprintf(w, " logig success ")
	}

	fmt.Printf(" %+v ", incamingdata)

}

// this is register function

func register(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/signup" {

		fmt.Printf(" signup page faild : %s ", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, " \n Signup page is active <3 \n")

	var incamingdats struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	json.NewDecoder(r.Body).Decode(&incamingdats)

	fmt.Printf(" %+v ", incamingdats)

}

// this is forget password email capture

func forgetpass(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/forgetpass" {
		fmt.Printf(" faild show forget passpage : %+v", http.StatusNotFound)
		return
	} else {

		fmt.Fprintf(w, " \n Forgetpassrod gape is active :D \n")
	}

	var imcamingdata struct {
		Email string `json:"email"`
	}

	json.NewDecoder(r.Body).Decode(&imcamingdata)

	fmt.Printf(" \n we got the email that forget passwrod : %+v ", imcamingdata)

}

func main() {

	// jsondata, err := jsonreade()
	// if err != nil {
	// 	fmt.Printf(" Faild to Load json file : " , err)
	// 	return
	// }

	// fmt.Printf("\n %+v \n", jsondata )

	// this are only style of text colors in terminal

	var text = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff0000")).Italic(true)
	var oktext = lipgloss.NewStyle().Foreground(lipgloss.Color("#3cff00")).Bold(true)

	// this are the main functin that canect logical funcion to the url quaris

	http.HandleFunc("/login", login)
	http.HandleFunc("/signup", register)
	http.HandleFunc("/forgetpass", forgetpass)

	// http.HandleFunc("/chat-init" , long)

	// server startig indicatin
	fmt.Printf("%s", text.Render("\n\n The server is starting... \n"))

	// this the last stage and the biginig of the server the code will be stay here in a loop and the server will be start
	fmt.Printf("%s", oktext.Render("\n ( CTRL + C ) TO STOP THE SRVER <3 \n\n"))

	if err := http.ListenAndServe(":4040", nil); err != nil {

		fmt.Println(" local Fort 4040 is unusebal right this moment ! ", err)

	}

}
