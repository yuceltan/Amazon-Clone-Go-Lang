package mails

import (
	"log"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
)

type MailData struct{
	To string
	From string
	Subject string
	Content string
}

func ListenForMail(mainChain chan MailData){

	go func(){
		for {
			msg := <- mainChain
			sendMsg(msg)
		}
	}()

}

func sendMsg(m MailData){
	server := mail.NewSMTPClient()
	server.Host = "localhost"
	server.Port = 1025 
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second


	client, err := server.Connect()
	if err != nil {
		log.Println(err)
		return
	}

	email := mail.NewMSG()
	email.SetFrom(m.From).AddTo(m.To).SetSubject(m.Subject)
	email.SetBody(mail.TextHTML, m.Content)

	err = email.Send(client)
	if err != nil {
		log.Println("failed to send mail", err)
	} else {
		log.Println("Email send")
	}
}