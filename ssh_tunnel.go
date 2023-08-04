package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"net"
)

func main() {
	// SSH 连接配置
	sshConfig := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.Password("199808"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// 连接 SSH 服务器
	sshConn, err := ssh.Dial("tcp", "192.168.232.130:22", sshConfig)
	if err != nil {
		log.Fatalf("Failed to dial: %s", err)
	}
	fmt.Println("连接 SSH 服务器成功.")

	// 监听本地端口
	localListener, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Fatalf("Failed to listen: %s", err)
	}
	fmt.Println("监听本地端口成功.")

	// 建立 SSH 隧道
	sshConn.Dial("tcp", "192.168.232.130:5236")
	fmt.Println("建立 SSH 隧道成功.")

	// 接受本地连接并转发到远程服务器
	for {
		localConn, err := localListener.Accept()
		if err != nil {
			log.Fatalf("Failed to accept: %s", err)
		}

		go func() {
			remoteConn, err := sshConn.Dial("tcp", "192.168.232.130:5236")
			if err != nil {
				log.Fatalf("Failed to dial: %s", err)
			}

			go func() {
				_, err := io.Copy(remoteConn, localConn)
				if err != nil {
					log.Fatalf("Failed to copy: %s", err)
				}
			}()

			go func() {
				_, err := io.Copy(localConn, remoteConn)
				if err != nil {
					log.Fatalf("Failed to copy: %s", err)
				}
			}()
		}()
	}
}
