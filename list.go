package glisp

// List represents a list of Atom's
type List []Atom

// NewList will return a new list
func NewList(ts *Tokens) (l List, err error) {
	var (
		token Token
		ok    bool
	)

	for {
		if token, ok = ts.Shift(); !ok {
			err = ErrUnexpectedEOF
			return
		}

		if token == ")" {
			return
		}

		var e Expression
		if e, err = toExpression(ts, token); err != nil {
			return
		}

		l = append(l, e)
	}
}
