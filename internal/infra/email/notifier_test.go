package email

import (
	"bytes"
	"log"
	"testing"
)

func TestSMTPEmailNotifierSend(t *testing.T) {
	buf := &bytes.Buffer{}
	log.SetOutput(buf)
	defer log.SetOutput(nil)

	n := NewSMTPEmailNotifier()
	if err := n.Send("to", "sub", "body"); err != nil {
		t.Fatalf("send error: %v", err)
	}
	if buf.Len() == 0 {
		t.Error("no log output")
	}
}
