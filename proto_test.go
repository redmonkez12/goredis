package main

import "testing"

func TestProtocol(t *testing.T) {
	msg := "*3\r\n$3\r\nSET\r\n$3\r\nmykey\r\n$3\r\nbar\r\n"
	parseCommand(msg)
}
