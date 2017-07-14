package msg

/*
Interface to output messages.

  P  - Expects message content is already completely rendered.
  Pf - Message content requires further rendering before its issued.

Their output behavior should be completely identical.  For example if P's
implementation inserts a newline, then Pf's does too.
*/

type I interface {
	Pf(format string, args ...interface{})
	P(a ...interface{})
}

func NewDiscard() I {
	d := new(discard)
	return d
}

type Contexter interface {
	ContextGet() string
}

type Context struct {
	I
	Contexter
}

func (mc Context) P(a ...interface{}) {
	arglist := make([]interface{}, 0, len(a)+1)
	arglist = append(arglist, mc.ContextGet())
	arglist = append(arglist, a...)
	mc.I.P(arglist...)
}
func (mc Context) Pf(format string, args ...interface{}) {
	arglist := make([]interface{}, 0, len(args)+1)
	arglist = append(arglist, mc.ContextGet())
	arglist = append(arglist, args...)
	mc.I.Pf("%s "+format, arglist...)
}

// private --------------------------------------------------------------------

type discard struct {
}

func (discard) Pf(string, ...interface{}) {}
func (discard) P(...interface{})          {}
