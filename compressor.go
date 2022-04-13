package compressiontools

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"io"

	"github.com/dsnet/compress/bzip2"
)

type Compressor struct {
	wOut    io.Writer
	buffer  *bytes.Buffer
	enc     *gob.Encoder
	useB64  bool
	useBzip bool
}

func NewCompressor(bzip, b64 bool) *Compressor {
	var wOut io.Writer
	var err error
	buf := &bytes.Buffer{}
	if bzip {
		wOut, err = bzip2.NewWriter(buf, nil)
		if err != nil {
			panic(err)
		}
	} else {
		wOut = buf
	}

	enc := gob.NewEncoder(wOut)

	return &Compressor{
		wOut:    wOut,
		buffer:  buf,
		enc:     enc,
		useB64:  b64,
		useBzip: bzip,
	}
}

func (c *Compressor) Compress(compressVal interface{}) ([]byte, error) {
	gob.Register(compressVal)
	err := c.enc.Encode(&compressVal)
	if err != nil {
		return []byte(""), err
	}
	if c, ok := c.wOut.(io.Closer); ok {
		c.Close()
	}
	encVal := c.buffer.Bytes()
	if c.useB64 {
		return []byte(base64.RawStdEncoding.EncodeToString(encVal)), nil
	}

	return encVal, nil
}
