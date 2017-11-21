package trace

import (
    "bytes"
    "testing"
)


func TestNew(t *testing.T) {
    //t.Error("We haven't written our test yet")
    var buf bytes.Buffer
    tracer := New(&buf)

    if tracer == nil {
        t.Error("Return from new should not be nil")
    } else {
        tracer.Trace("Hello trace package.")
        if buf.String() != "Hello trace package.\n" {
            t.Error("Trace should not write '%s'", buf.String())
        }
    } 
}
