package main

import (
	"fmt"
	"net/http"

	"github.com/charmbracelet/lipgloss"
)

// the login funciton

func login(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/login" {

		fmt.Printf(" login page faild ! : %s ", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		fmt.Printf(" login page faild ! methode ! : %s ", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "\n Login page is active \n")

	// Inside your login handler on the server:
	email := r.FormValue("email")
	pass := r.FormValue("password")

}

func main() {

	// this are only style of text colors in terminal

	var text = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff0000")).Italic(true)
	var oktext = lipgloss.NewStyle().Foreground(lipgloss.Color("#3cff00")).Bold(true)

	// this are the main functin that canect logical funcion to the url quaris

	http.HandleFunc("/login", login)
	// http.HandleFunc("/register" , long)
	// http.HandleFunc("/forgetpass" , long)
	// http.HandleFunc("/chat-init" , long)

	// server startig indicatin
	fmt.Printf("%s", text.Render("\n\n The server is starting... \n"))

	// this the last stage and the biginig of the server the code will be stay here in a loop and the server will be start
	fmt.Printf("%s", oktext.Render("\n ( CTRL + C ) TO STOP THE SRVER <3 \n\n"))

	if err := http.ListenAndServe(":4040", nil); err != nil {

		fmt.Println(" local Fort 8080 is unusebal right this moment ! ", err)

	}

}
