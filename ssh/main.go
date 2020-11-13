package main

import (
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

func main() {
	var cp ssh.Config
	cp = ssh.Config{
		Ciphers: []string{"aes128-ctr", "aes192-ctr", "aes256-ctr", "aes128-gcm@openssh.com", "arcfour256", "arcfour128", "aes128-cbc", "3des-cbc", "aes192-cbc", "aes256-cbc"},
	}
	// var hostKey ssh.PublicKey
	config := &ssh.ClientConfig{
		User: "tjx",
		Auth: []ssh.AuthMethod{
			ssh.Password("tanjianxiong"),
		},
		Config:          cp,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", "192.168.0.192:22", config)
	if err != nil {
		log.Fatal("Failed to dial: ", err.Error())
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer session.Close()

	// 单词执行，无交互
	// out, errOutput := session.Output("ls\npwd") // 封装了Run在里面,如果使用output session的stdout不能外部初始化
	// if errOutput != nil {
	// 	log.Fatal("Failed to run: " + err.Error())
	// }
	// fmt.Println(string(out))

	// // 交互
	// session.Stdout = os.Stdout
	// session.Stderr = os.Stderr
	// session.Stdin = os.Stdin

	// modes := ssh.TerminalModes{
	// 	ssh.ECHO:          0,     // disable echoing
	// 	ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
	// 	ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	// }

	// // Request pseudo terminal
	// if err := session.RequestPty("xterm", 40, 80, modes); err != nil {
	// 	log.Fatal("request for pseudo terminal failed: ", err)
	// }

	// if err := session.Shell(); err != nil {
	// 	log.Fatal("failed to start shell: ", err)
	// }
	// session.Wait()

	session.Stderr = os.Stderr
	session.Stdout = os.Stdout
	session.Stdin = os.Stdin
	// stdin, err := session.StdinPipe()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	// Request pseudo terminal
	if err := session.RequestPty("xterm", 40, 80, modes); err != nil {
		log.Fatal("request for pseudo terminal failed: ", err)
	}
	if err := session.Shell(); err != nil {
		log.Fatal("failed to start shell: ", err)
	}
	// fmt.Print(session.Stderr)
	// fmt.Fprintf(stdin, "%s\n", "ll")

	session.Wait()

}
