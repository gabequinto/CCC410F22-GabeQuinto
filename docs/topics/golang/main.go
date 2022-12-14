package main

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
)

var username = "USERNAME"
var password = "PASSWORD"
var host = "RemoteSystemIP"
var command = "COMMAND"

func main() {
	//authenitcation
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	//connect to remote system
	client, err := ssh.Dial("tcp", host, config)
	if err != nil {
		panic("Error dialing server. " + err.Error())
	}
	//create session
	session, err := client.NewSession()
	if err != nil {
		panic("Failed creating session" + err.Error())
	}
	defer session.Close()

	var b bytes.Buffer
	//session output
	session.Stdout = &b
	//run command on remote system

	if err := session.Run(command); err != nil {
		panic("Command Failed. " + err.Error())
	}
	fmt.Println(b.String())
}
