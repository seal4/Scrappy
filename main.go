package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/5elenay/revoltgo"
)

const help = "Type !meme for a random meme"

func Shitposting() string {
	cmd, err := exec.Command("/bin/sh", "./wget.sh").Output()
	if err != nil {
		fmt.Printf("error %s", err)
	}
	output := string(cmd)
	return output
}

func main() {
	// Init a new client.
	client := revoltgo.Client{
		Token: "qy1X_Z9f7ENEhVcXMkBjMgwcR3P_gGjsAbNbbiLfl0s5UgXcJ394VIJLVrTfELJm",
	}
	content, err := ioutil.ReadFile("output.txt") // the file is inside the local directory
	if err != nil {
		fmt.Println("Err")
	}
	// Listen a on message event.
	client.OnMessage(func(m *revoltgo.Message) {
		if m.Content == "!meme" {
			Shitposting()
			sendMsg := &revoltgo.SendMessage{}
			sendMsg.SetContent(string(content))
			m.Reply(true, sendMsg)
		}
		if m.Content == "!help" {
			sendMSG := &revoltgo.SendMessage{}
			sendMSG.SetContent(help)
			m.Reply(true, sendMSG)
		}
	})

	log.Println(string(content))
	//Start the client.
	client.Start()

	// Wait for close.
	sc := make(chan os.Signal, 1)

	signal.Notify(
		sc,
		syscall.SIGINT,
		syscall.SIGTERM,
		os.Interrupt,
	)
	<-sc

	// Destroy client.
	client.Destroy()

}
