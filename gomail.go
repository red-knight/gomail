/* -----------
To a great extent, the following was borrowed/inspired by Nathan LaClaire's work at:
http://nathanleclaire.com/blog/2013/12/17/sending-email-from-gmail-using-golang/

I've found it very helpful in writing this quick tool, so check out some of his
other work while you're over there.
===
POC THISTA:  Red Thomas, red(dot)thomas(at)redknightllc(dot)com
----------- */

package main

import (
	"fmt"
	"io/ioutil"
	"net/smtp"
	"strconv"
)

var confile = "./red.conf"

type EmailUser struct {
	Username    string
	Password    string
	EmailServer string
	Port        int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	emailUser := EmailUser{"red.thomas@redknightllc.com", "A!A!S@S@D#D#a1a1s2s2d3d3", "smtp.gmail.com", 587}
	auth := smtp.PlainAuth("", emailUser.Username, emailUser.Password, emailUser.EmailServer)

	dat, err := ioutil.ReadFile(confile)
	check(err)
	fmt.Print(string(dat))

	doc := []byte("test email")

	err = smtp.SendMail(emailUser.EmailServer+":"+strconv.Itoa(emailUser.Port), // in our case, "smtp.google.com:587"
		auth,
		emailUser.Username,
		[]string{"red.thomas@redknightllc.com"},
		doc)
	if err != nil {
		fmt.Print("ERROR: attempting to send a mail ", err)
	}
}
