package z3

import (
	"unsafe"
)

// #include "go-z3.h"
import "C"

// Add creates an AST node representing adding.
//
// All AST values must be part of the same context.
func (a *AST) Add(args ...*AST) *AST {
	raws := make([]C.Z3_ast, len(args)+1)
	raws[0] = a.rawAST
	for i, arg := range args {
		raws[i+1] = arg.rawAST
	}

	return &AST{
		rawCtx: a.rawCtx,
		rawAST: C.Z3_mk_add(
			a.rawCtx,
			C.uint(len(raws)),
			(*C.Z3_ast)(unsafe.Pointer(&raws[0]))),
	}
}

// Lt creates a "less than" comparison.
//
// Maps to: Z3_mk_lt
func (a *AST) Lt(a2 *AST) *AST {
	return &AST{
		rawCtx: a.rawCtx,
		rawAST: C.Z3_mk_lt(a.rawCtx, a.rawAST, a2.rawAST),
	}
}

// Gt creates a "greater than" comparison.
//
// Maps to: Z3_mk_gt
func (a *AST) Gt(a2 *AST) *AST {
	return &AST{
		rawCtx: a.rawCtx,
		rawAST: C.Z3_mk_gt(a.rawCtx, a.rawAST, a2.rawAST),
	}
}
