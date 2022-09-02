package ssh

import "fmt"

func SSHConnection() {
	fmt.Println("connected")
}

// package main

// import (
// 	"fmt"

// 	"golang.org/x/crypto/ssh"
// 	"golang.org/x/crypto/ssh/knownhosts"
// )
// func main() {
// //collecting necessary info from user

// // var hostname string
// 	// var usrnm string
// 	// var pswd string
// 	// fmt.Println("thpe the hostname and port -> <hostname>:<port>")
// 	// fmt.Scanln(&hostname)
// 	// fmt.Println("the the username")
// 	// fmt.Scanln(&usrnm)
// 	// fmt.Println("thpe the hostname and port -> <hostname>:<port>")
// 	// fmt.Scanln(&pswd)
// 	host := "<hostname>:<port>"
// 	user := "<username>"
// 	pwd := "<password>"
// 	pKey := []byte("<privateKey>")

// 	var err error
// 	var signer ssh.Signer

// 	signer, err = ssh.ParsePrivateKey(pKey)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	var hostKeyCallback ssh.hostKeyCallback

// 	hostKeyCallback, err = knownhosts.new("C:\\cygwin64\\home\\username\\. ssh\\known_hosts")

// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	conf := &ssh.ClientConfig{
// 		User:            user,
// 		HostKeyCallback: hostKeyCallback,
// 		Auth: []ssh.AuthMethod{
// 		ssh.Password(pwd),
// 		ssh.PublicKeys(signer),
// 		},
// 	}

// 	var conn *ssh.Client
// 	conn, err = ssh.Dial("tcp", host, conf)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	defer conn.Close()
// 	var session *ssh.Session
// 	session, err = conn.NewSession()

// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	defer session.Close()

//}
/*								yarıda kaldı moruq								*/
