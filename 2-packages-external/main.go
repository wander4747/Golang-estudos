package main

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {
	mail()
	domain()
	hostUser()
}

func mail() {
	err := checkmail.ValidateFormat("wander.douglas14@gmail.com")
	if err != nil {
		fmt.Println(err)
	}
}

func domain() {
	err := checkmail.ValidateHost("email@x-unkown-domain.com")
	if err != nil {
		fmt.Println(err)
	}
}

func hostUser() {
	var (
		serverHostName    = "smtp.gmail.com"             // set your SMTP server here
		serverMailAddress = "wander.douglas14@gmail.com" // set your valid mail address here
	)
	err := checkmail.ValidateHostAndUser(serverHostName, serverMailAddress, "wander.douglas14_1231Vcs@gmail.com")
	if smtpErr, ok := err.(checkmail.SmtpError); ok && err != nil {
		fmt.Printf("Code: %s, Msg: %s", smtpErr.Code(), smtpErr)
	}
}
