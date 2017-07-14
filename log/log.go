package log

import (
	"fmt"
	"log"
	"msg"
	"runtime"
	"strings"
)

func NewError(logger *log.Logger) msg.I {
	return newMsgLogContext(logger, 1, "Error:")
}
func NewDebug(logger *log.Logger) msg.I {
	return newMsgLogContext(logger, 1, "Debug:")
}
func NewInfo(logger *log.Logger) msg.I {
	return newMsgLogContext(logger, 1, "Info:")
}

// private --------------------------------------------------------------------

type msgLog struct {
	msgLog *log.Logger // reference to log.Logger implementation
	depth  int         // given the current level of indirection (depth of stack)
}

func newMsgLog(logger *log.Logger, depth int) (ml *msgLog) {
	ml = new(msgLog)
	ml.msgLog = logger
	ml.depth = depth
	return
}
func (ml msgLog) P(a ...interface{}) {
	ml.msgLog.Output(ml.depth+3, fmt.Sprintln(a...))
}
func (ml msgLog) Pf(format string, args ...interface{}) {
	// +2 represents the stack frame of current location.
	ml.msgLog.Output(ml.depth+3, fmt.Sprintf(format+"\n", args...))
}

func caller(popcnt int) (funName string) {
	funName = "unknown"
	fpcs := []uintptr{uintptr(0)}
	// "pop" stack frames resulting from this function and Caller
	if n := runtime.Callers(popcnt+3, fpcs); n == 0 {
		return
	}
	// obtain name from function ptr
	if fun := runtime.FuncForPC(fpcs[0] - 1); fun != nil {
		funName = fun.Name()
	}
	return
}

type msgContext struct {
	fn func() string
}

func (ctx msgContext) ContextGet() string {
	return ctx.fn()
}

func newMsgLogContext(logger *log.Logger, depth int, msgtype string) (mlc *msg.Context) {
	mlc = new(msg.Context)
	mlc.I = newMsgLog(logger, depth+1)
	mc := new(msgContext)
	msgtypespc := msgtype + " "
	mc.fn = func() string {
		nameParts := strings.Split(caller(depth+3), ".")
		return msgtypespc + fmt.Sprintf("func='%s'", nameParts[len(nameParts)-1])
	}
	mlc.Contexter = mc
	return
}
