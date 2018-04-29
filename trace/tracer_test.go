package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("newからの戻り値がnilです")
	} else {
		tracer.Trace("hello trace package")
		if buf.String() != "hello trace package\n" {
			t.Errorf("'%s'という誤った文字列が出力されました", buf.String())
		}
	}
}
