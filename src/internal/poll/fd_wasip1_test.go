//go:build wasip1

package poll_test

import (
	"internal/poll"
	"testing"
	"time"
)

func TestWaitWriteExpiredDeadline(t *testing.T) {
	fd := &poll.FD{Sysfd: 1, IsStream: true}
	_ = fd.Init("file", true)
	if err := fd.SetWriteDeadline(time.Now().Add(-time.Second)); err != nil {
		t.Fatal(err)
	}
	if err := fd.WaitWrite(); err != poll.ErrDeadlineExceeded {
		t.Fatalf("WaitWrite with expired deadline: got %v, want %v", err, poll.ErrDeadlineExceeded)
	}
}
