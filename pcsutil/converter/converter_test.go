package converter_test

import (
	"github.com/Luis0001/BaiduPCS-Go/pcsutil/converter"
	"strings"
	"testing"
)

func TestTrimPathInvalidChars(t *testing.T) {
	trimed := converter.TrimPathInvalidChars("ksjadfi*/?adf")
	if strings.Compare(trimed, "ksjadfiadf") != 0 {
		t.Fatalf("trimed: %s\n", trimed)
	}
	return
}
