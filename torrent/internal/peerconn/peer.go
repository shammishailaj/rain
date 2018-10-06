package peerconn

import (
	"net"

	"github.com/cenkalti/rain/internal/logger"
	"github.com/cenkalti/rain/torrent/internal/bitfield"
	"github.com/cenkalti/rain/torrent/internal/peerconn/peerprotocol"
	"github.com/cenkalti/rain/torrent/internal/peerconn/peerreader"
	"github.com/cenkalti/rain/torrent/internal/peerconn/peerwriter"
	"github.com/cenkalti/rain/torrent/internal/pieceio"
)

type Peer struct {
	conn          net.Conn
	id            [20]byte
	FastExtension bool
	reader        *peerreader.PeerReader
	writer        *peerwriter.PeerWriter
	log           logger.Logger
	closeC        chan struct{}
	closedC       chan struct{}
}

func New(conn net.Conn, id [20]byte, extensions *bitfield.Bitfield, l logger.Logger) *Peer {
	fastExtension := extensions.Test(61)
	extensionProtocol := extensions.Test(43)
	return &Peer{
		conn:          conn,
		id:            id,
		FastExtension: fastExtension,
		reader:        peerreader.New(conn, l, fastExtension, extensionProtocol),
		writer:        peerwriter.New(conn, l),
		log:           l,
		closeC:        make(chan struct{}),
		closedC:       make(chan struct{}),
	}
}

func (p *Peer) ID() [20]byte {
	return p.id
}

func (p *Peer) String() string {
	return p.conn.RemoteAddr().String()
}

func (p *Peer) Close() {
	close(p.closeC)
	<-p.closedC
}

func (p *Peer) Logger() logger.Logger {
	return p.log
}

func (p *Peer) Messages() <-chan interface{} {
	return p.reader.Messages()
}

func (p *Peer) SendMessage(msg peerprotocol.Message) {
	p.writer.SendMessage(msg)
}

func (p *Peer) SendPiece(msg peerprotocol.RequestMessage, pi *pieceio.Piece) {
	p.writer.SendPiece(msg, pi)
}

// Run reads and processes incoming messages after handshake.
func (p *Peer) Run() {
	defer close(p.closedC)
	p.log.Debugln("Communicating peer", p.conn.RemoteAddr())

	readerDone := make(chan struct{})
	go func() {
		p.reader.Run(p.closeC)
		close(readerDone)
	}()

	writerDone := make(chan struct{})
	go func() {
		p.writer.Run(p.closeC)
		close(writerDone)
	}()

	select {
	case <-p.closeC:
		p.conn.Close()
		<-readerDone
		<-writerDone
	case <-readerDone:
		p.conn.Close()
		<-writerDone
	case <-writerDone:
		p.conn.Close()
		<-readerDone
	}
}