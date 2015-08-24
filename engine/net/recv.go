package net

import (
	"bytes"
	"bufio"
	libNet "net"
)

func RecvLoop(conn libNet.Conn, cb func(*bufio.Reader, error)) {
	ch := make(chan []byte)
	eCh := make(chan error)

	go func() {
		for {
			data := make([]byte, 512)
			l, err := conn.Read(data)
			if err != nil {
				eCh <- err
				return
			}
			ch <- data[:l]
		}
	}()

	var buffer bytes.Buffer
	reader := bufio.NewReader(&buffer)
	for {
		select {
		case data := <-ch:
			buffer.Write(data)
			cb(reader, nil)
		case err := <-eCh:
			cb(nil, err)
			break
		}
	}

	close(ch)
	close(eCh)
}

