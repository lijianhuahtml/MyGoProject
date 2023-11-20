package main

import (
	"MyGoProject/protobuf/msg"
	"google.golang.org/protobuf/proto"
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:20000")
	if err != nil {
		log.Fatal(err)
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {

		}
	}(listener)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	// 创建一个空的 MyMessage 结构
	ms := &msg.Msg{}

	for {
		// 从连接中读取数据
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Println(err)
			}
			break
		}

		// 解码 Protobuf 消息
		if err := proto.Unmarshal(buf[:n], ms); err != nil {
			log.Println("Failed to unmarshal:", err)
			continue
		}

		// 处理消息
		log.Printf("Received message: %s\n", ms.Msg)

		// 在这里可以编写回复消息的逻辑
	}
}
