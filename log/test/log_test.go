package logtst

import (
	"log"
	"msg"
	mlog "msg/log"
	"os"
	"testing"
)

func TestDebug(t *testing.T) {
	logger := log.New(os.Stdout, "", log.Lshortfile|log.Ldate|log.Ltime)
	dbg := mlog.NewDebug(logger)
	dbg.P("Hello")
	anotherLevel(dbg)
	dbg.P("Hello2")
	dbg.Pf("hi %s, %d", "t", 5)
}
func anotherLevel(dbg msg.I) {
	dbg.Pf("Hello %s", "There!")
}
