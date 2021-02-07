package emptty

import "os"

const (
	EscString             = "\x1b"
	ClearSoftString       = EscString + "c"
	ClearScrollbackString = EscString + "[3J"
	ClearHardString       = ClearSoftString + ClearScrollbackString
)

var (
	EscBytes             = []byte(EscString)
	ClearSoftBytes       = []byte(ClearSoftString)
	ClearScrollbackBytes = []byte(ClearScrollbackString)
	ClearHardBytes       = []byte(ClearHardString)
)

func ClearScrollback() {
	_, _ = os.Stdout.Write(ClearScrollbackBytes)
}

func ClearSoft() {
	_, _ = os.Stdout.Write(ClearSoftBytes)
}

func ClearHard() {
	_, _ = os.Stdout.Write(ClearHardBytes)
}
