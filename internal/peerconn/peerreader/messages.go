package peerreader

import (
	"github.com/cenkalti/rain/internal/peerprotocol"
)

type Piece struct {
	peerprotocol.PieceMessage
	Data []byte
}
