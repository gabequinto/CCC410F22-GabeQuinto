package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"strings"

	"golang.org/x/crypto/ssh"
)

type Connection struct {
	*ssh.Client
	password string
}
//connecting to remote machine without host keys
func Connect(addr, user, password string) (*Connection, error) {
	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", addr, sshConfig)
	if err != nil {
		return nil, err
	}

	return &Connection{conn, password}, nil

}

// run commands as sudo user
func (conn *Connection) SendCommands(cmds string) ([]byte, error) {
	session, err := conn.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	stdoutB := new(bytes.Buffer)
	session.Stdout = stdoutB
	in, _ := session.StdinPipe()

	go func(in io.Writer, output *bytes.Buffer) {
		for {
			if strings.Contains(string(output.Bytes()), "[sudo] password for ") {
				_, err = in.Write([]byte(conn.password + "\n"))
				if err != nil {
					break
				}
				fmt.Println("put the password ---  end .")
				break
			}
		}
	}(in, stdoutB)

	err = session.Run(cmds)
	if err != nil {
		return []byte{}, err
	}
	return stdoutB.Bytes(), nil
}

//define computer, user, and password
func main() {

	conn, err := Connect("192.168.229.135:22", "gabe", "password123")
	if err != nil {
		log.Fatal(err)
	}
	// command to execute
	output, err := conn.SendCommands("")
	if err != nil {
		log.Fatal(err)
	}
	// print output of remote command
	fmt.Println(string(output))
	
