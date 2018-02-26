package godiag

import (
	"eaciit/diagnostic"
	"errors"
	"testing"

	"github.com/dewanggasurya/godiag/tasks"
)

func TestDiagnostic(t *testing.T) {
	d := diagnostic.NewDiagnostic()

	if e := d.RegisterFunc("SuccessfulTask", func() error {
		return nil
	}); e != nil {
		t.Error(e)
	}
	t.Log("SuccessfulTask is registered perfectly")

	if e := d.RegisterFunc("FailedTask", func() error {
		return errors.New("Wew, something happend...")
	}); e != nil {
		t.Error(e)
	}
	t.Log("FailedTask is also registered perfectly")

	//=== Adding pre-defined task and check if nginx is running
	if e := d.Register("Nginx", tasks.IsProcessRunning("nginx")); e != nil {
		t.Error(e)
	}
	t.Log("Nginx is also registered perfectly")

	t.Log("Diagnostic process is running")
	result := d.Run()
	t.Log("Here is the result", result)
}
