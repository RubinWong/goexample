package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

func main() {
	go ServerMain()
	defer clientMain()
}

func ServerMain() {
	listen, err := net.Listen("tcp", ":30000")
	if err != nil {
		fmt.Println("listen failed, er: ", err)
		return
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, er: ", err)
			continue
		}
		go process(conn)
	}
}

func process(c net.Conn) {
	defer c.Close()
	reader := bufio.NewReader(c)
	// var buf [4096]byte
	for {
		n, err := Decode(reader)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read failed, er: ", err)
			break
		}
		recvStr := string(n)
		fmt.Println("recv: ", recvStr)
	}
}

func clientMain() {
	conn, err := net.Dial("tcp", ":30000")
	if err != nil {
		fmt.Println("dial failed, er: ", err)
		return
	}

	defer conn.Close()
	for i := 0; i < 10; i++ {
		msg := "Hello, hello. How are u?"
		pkg, err := Encode([]byte(msg))
		if err != nil {
			fmt.Println("encode failed, er: ", err)
			break
		}
		conn.Write(pkg)
	}
}

func Encode(data []byte) ([]byte, error) {
	buffer := new(bytes.Buffer)
	length := len(data)
	if err := binary.Write(buffer, binary.LittleEndian, int32(length)); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.LittleEndian, data); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func Decode(reader *bufio.Reader) ([]byte, error) {
	lengthbyte, err := reader.Peek(4)
	if err != nil {
		return nil, err
	}
	lengthbuffer := bytes.NewBuffer(lengthbyte)
	var length int32
	err = binary.Read(lengthbuffer, binary.LittleEndian, &length)
	if err != nil {
		fmt.Println("read length failed, er: ", err)
		return nil, err
	}
	if reader.Buffered() < int(length)+4 {
		return nil, fmt.Errorf("buffer not enough, length: %d, buffered: %d", length, reader.Buffered())
	}

	buffer := make([]byte, length+4)
	_, err = reader.Read(buffer)
	if err != nil {
		return nil, err
	}

	return buffer[4:], nil
}

type Header struct {
	URI uint32
	Len uint16
	Cid uint16
}

type Packet struct {
	Header
	Data    []byte
	Padding []byte
}
