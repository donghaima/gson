package gson

import "fmt"
import "testing"

var _ = fmt.Sprintf("dummy text")

func TestJsonEmpty(t *testing.T) {
	config := NewDefaultConfig()
	cbr := config.NewCbor(make([]byte, 10), 0)
	jsn := config.NewJson(make([]byte, 10), 0)
	clt := config.NewCollate(make([]byte, 10), 0)

	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected panic")
			}
		}()
		jsn.Tovalue()
	}()
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected panic")
			}
		}()
		jsn.Tocbor(cbr)
	}()
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected panic")
			}
		}()
		jsn.Tocollate(clt)
	}()
}
