package btf

import (
	"strings"
	"testing"

	"github.com/psanford/btf/internal"
)

func TestParseExtInfoBigRecordSize(t *testing.T) {
	rd := strings.NewReader("\xff\xff\xff\xff\x00\x00\x00\x000709171295166016")
	table := stringTable("\x00")

	if _, err := parseFuncInfos(rd, internal.NativeEndian, table); err == nil {
		t.Error("Parsing func info with large record size doesn't return an error")
	}

	if _, err := parseLineInfos(rd, internal.NativeEndian, table); err == nil {
		t.Error("Parsing line info with large record size doesn't return an error")
	}
}
