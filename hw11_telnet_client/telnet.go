package main

import (
	fmt "fmt"
	"io"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	Close() error
	Send() error
	Receive() error
}

type telnetClient struct {
	address string
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer
	conn    net.Conn
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &telnetClient{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}

func (tc *telnetClient) Connect() error {
	conn, err := net.DialTimeout("tcp", tc.address, tc.timeout)
	if err != nil {
		return fmt.Errorf("dial: %w", err)
	}

	tc.conn = conn

	return nil
}

func (tc *telnetClient) Close() error {
	return tc.conn.Close()
}

func (tc *telnetClient) Send() error {
	_, err := io.Copy(tc.conn, tc.in)

	return err
}

func (tc *telnetClient) Receive() error {
	_, err := io.Copy(tc.out, tc.conn)

	return err
}
