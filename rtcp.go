package rtcp

import "io"

type ReadStream interface {
	io.ReadCloser
}

type WriteStream interface {
	io.Writer
}

type Session interface {
	AcceptStream() (ReadStream, uint32, error)
	OpenReadStream(ssrc uint32) (ReadStream, error)
	OpenWriteStream() (WriteStream, error)
}
