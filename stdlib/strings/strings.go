package strings

import (
	"github.com/itsmontoya/glisp/types"
	"github.com/itsmontoya/glisp/utils"
)

// Add will add a series of strings
func Add(sc types.Scope, args types.List) (out types.Expression, err error) {
	var (
		val types.String
		str types.String
	)

	for _, exp := range args {
		if str, err = utils.GetString(sc, exp); err != nil {
			return
		}

		val += str
	}

	out = val
	return
}

// LessThan will return if a is less than b
func LessThan(sc types.Scope, as types.String, b types.Atom) (out types.Expression, err error) {
	var bs types.String
	if bs, err = utils.GetString(sc, b); err != nil {
		return
	}

	if as < bs {
		out = "true"
	}

	return
}

// GreaterThan will return if a is greater than b
func GreaterThan(sc types.Scope, as types.String, b types.Atom) (out types.Expression, err error) {
	var bs types.String
	if bs, err = utils.GetString(sc, b); err != nil {
		return
	}

	if as > bs {
		out = "true"
	}

	return
}