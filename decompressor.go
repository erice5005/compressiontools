package compressiontools

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"io"
	"log"

	"github.com/dsnet/compress/bzip2"
)

type Decompressor struct {
	rd       io.Reader
	rootBuff *bytes.Buffer
}

func NewDecompressor() *Decompressor {
	return &Decompressor{}
}

func (d *Decompressor) parseB64(val []byte) ([]byte, error) {
	return base64.RawStdEncoding.DecodeString(string(val))
}

func (d *Decompressor) parseB64ToReader(val []byte) error {
	var rd io.Reader
	var rootBuff *bytes.Buffer
	deB64, err := d.parseB64(val)
	if err != nil {
		rd = bytes.NewReader(val)
		rootBuff = bytes.NewBuffer(val)
	} else {
		rd = bytes.NewReader(deB64)
		rootBuff = bytes.NewBuffer(deB64)
	}
	d.rd = rd
	d.rootBuff = rootBuff
	return nil
}

func (d *Decompressor) bZipReader() error {
	zipRd, err := bzip2.NewReader(d.rd, nil)
	if err != nil {
		log.Printf("bzip err: %v\n", err.Error())
		d.rd = bytes.NewReader(d.rootBuff.Bytes())
		return nil
		// return err
	}
	d.rd = zipRd
	return nil
}

func (d *Decompressor) Decompress(inp []byte, out interface{}) interface{} {
	gob.Register(out)
	d.parseB64ToReader(inp)
	d.bZipReader()
	dec := gob.NewDecoder(d.rd)

	err := dec.Decode(&out)
	if err != nil {
		if err.Error() == "bzip2: corrupted input: invalid stream magic" {
			d.rd = bytes.NewReader(d.rootBuff.Bytes())
			dec = gob.NewDecoder(d.rd)
			err = dec.Decode(&out)
			if err != nil {
				panic(err)
			}
			return out
		}
		panic(err)
	}
	return out
}
