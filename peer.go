package main

import (
	"net"
)

type Peer struct {
	conn    net.Conn
	msgChan chan []byte
}

func NewPeer(conn net.Conn, msgCh chan []byte) *Peer {
	return &Peer{
		conn:    conn,
		msgChan: msgCh,
	}
}

func (p *Peer) readLoop() error {
	buf := make([]byte, 1024)

	for {
		n, err := p.conn.Read(buf)
		if err != nil {
			return err
		}
		msgBuf := make([]byte, n)
		copy(msgBuf, buf[:n])
		p.msgChan <- msgBuf
	}
}
