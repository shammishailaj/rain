// Package btconn provides support for dialing and accepting BitTorrent connections.
package btconn

import (
	"errors"
	"io"
	"net"
)

var (
	errInvalidInfoHash = errors.New("invalid info hash")
	ErrOwnConnection   = errors.New("dropped own connection")
	errNotEncrypted    = errors.New("connection is not encrypted")
)

type readWriter struct {
	io.Reader
	io.Writer
}

type rwConn struct {
	rw io.ReadWriter
	net.Conn
}

func (c *rwConn) Read(p []byte) (n int, err error)  { return c.rw.Read(p) }
func (c *rwConn) Write(p []byte) (n int, err error) { return c.rw.Write(p) }
