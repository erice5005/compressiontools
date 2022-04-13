package compressiontools

import (
	"encoding/binary"
	"testing"
)

type Compressable struct {
	Input int
	Stin  string
}

var sampleCompress = []Compressable{
	{Input: 100},
	{Input: 20000},
	{Input: 20000, Stin: "hello World"},
	{Input: 2000000, Stin: "hello World!!!"},
}

type CompressableStruct struct {
	Field1 string
	Field2 string
}

func Test_BasicCompression(t *testing.T) {
	testcompressor := NewCompressor(false, false)
	out, err := testcompressor.Compress(
		CompressableStruct{
			Field1: "Hello",
			Field2: "World",
		},
	)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(string(out))
}

func Test_CompressWithBzip(t *testing.T) {
	testcompressor := NewCompressor(true, false)
	out, err := testcompressor.Compress(
		CompressableStruct{
			Field1: "Hello",
			Field2: "World",
		},
	)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(string(out))
}

func Test_CompressWithBzipAndB64(t *testing.T) {
	testcompressor := NewCompressor(true, true)
	out, err := testcompressor.Compress(
		CompressableStruct{
			Field1: "Hello",
			Field2: "World",
		},
	)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(string(out))
}

func BenchmarkCompression(b *testing.B) {
	tests := []struct {
		name    string
		bzipVal bool
		b64Val  bool
	}{
		{
			name:    "basic",
			bzipVal: false,
			b64Val:  false,
		},
		{
			name:    "withBzipOnly",
			bzipVal: true,
			b64Val:  false,
		},
		{
			name:    "withB64Only",
			bzipVal: false,
			b64Val:  true,
		},
		{
			name:    "withBoth",
			bzipVal: true,
			b64Val:  true,
		},
	}

	for _, bt := range tests {
		for _, sx := range sampleCompress {
			b.Run(bt.name, func(b *testing.B) {

				for i := 0; i < b.N; i++ {
					cx := NewCompressor(bt.bzipVal, bt.b64Val)
					out, err := cx.Compress(sx)
					if err != nil {
						b.Error(err)
						b.FailNow()
					}
					b.ReportMetric(float64(binary.Size(out)), "bytes")
				}
			})
		}
	}
}
