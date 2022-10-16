### Installing GO
`sudo apt update`    
`sudo apt upgrade`    
`sudo apt-get install golang-go`

###### Set Go Paths
Add the following lines to the file `~/.bash_profile`:  
`export GOPATH=$HOME/go`    
`export PATH=$PATH:/usr/localgo/bin:$GOPATH/bin`  

###### Running and Building a GO program 
To run the program 'hello.go' use the command `go run hello.go`. To build and then run the program use the commands `go build hello.go` than `./hello`.

Source: https://www.cyberciti.biz/faq/how-to-install-gol-ang-on-ubuntu-linux/
### Creating SSH Session
For my first task I needed to create an SSH session with a remote machine. The following code is how that was completed. 
![image](https://user-images.githubusercontent.com/78443183/196046868-f0a9857d-0eb9-4cd6-b614-4d3cff8c3cca.png)

Please note that in this example I am connecting to the remote system without hostkeys but it is recommended that you use hostkeys when possible. 

### Execute Remote Commands through SSH session
My second task was to create execute commands on the remote machine through the SSH session that was created in the first task. The following code is how that was completed. 
![image](https://user-images.githubusercontent.com/78443183/196046907-f5756094-5552-4f34-9de1-d1d104f62da9.png)

Note this is a very small snippet of code and relies on having the already set up ssh connection. Another thing to note is that if you want to run mulitple commands over an ssh
session you will need to create a loop as only one command can be run per ssh session. 
