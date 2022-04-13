package compressiontools

import (
	"testing"
)

func Test_Decompress(t *testing.T) {
	fullCompressed := "QlpoNjFBWSZTWWbAHREAAEF/wPsE0AAAA7ACCUAIgD7v34AwAAAAoABqGiJVP/aqp/5qaqh/6qg/2iqP9VP/RpUfhkDFRpVR/6qn/qqf+9VVH/qqP/yqoH/5VU//VUH/lVP/VU//VVP/aqocd2MfojYcPbKS9HcwUbMSnlNuZB2WcxOAOAtUdprIOXQcZiDmuhGvt+dzFQmwR0uFZwpNtvaXPJeqBUKOBKs7Bt3Ci+ljXf4XckU4UJBmwB0R"

	var out CompressableStruct
	testDecomp := NewDecompressor()
	testDecomp.Decompress([]byte(fullCompressed), out)
	t.Log(out)
}

func Test_BasicRound(t *testing.T) {
	compressable := CompressableStruct{
		Field1: "Hello",
		Field2: "World",
	}
	cx := NewCompressor(false, false)
	compd, err := cx.Compress(compressable)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	decomp := NewDecompressor()
	out := CompressableStruct{}
	out = decomp.Decompress(compd, out).(CompressableStruct)
	t.Log(out)
}

func Test_DecompBZip(t *testing.T) {
	compressable := CompressableStruct{
		Field1: "Hello",
		Field2: "World",
	}
	cx := NewCompressor(true, false)
	compd, err := cx.Compress(compressable)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	decomp := NewDecompressor()
	out := CompressableStruct{}
	out = decomp.Decompress(compd, out).(CompressableStruct)
	t.Log(out)
}
func Test_DecompBZipAndB64(t *testing.T) {
	compressable := CompressableStruct{
		Field1: "Hello",
		Field2: "World",
	}
	cx := NewCompressor(true, true)
	compd, err := cx.Compress(compressable)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	decomp := NewDecompressor()
	out := CompressableStruct{}
	out = decomp.Decompress(compd, out).(CompressableStruct)
	t.Log(out)
}

func Test_DecompB64(t *testing.T) {
	compressable := CompressableStruct{
		Field1: "Hello",
		Field2: "World",
	}
	cx := NewCompressor(true, true)
	compd, err := cx.Compress(compressable)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	decomp := NewDecompressor()
	out := CompressableStruct{}
	out = decomp.Decompress(compd, out).(CompressableStruct)
	t.Log(out)
}
