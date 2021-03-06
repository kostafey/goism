package compiler

import (
	"bytes"
	"dt"
	"ir"
	"sexp"
)

type Compiler struct {
	buf         bytes.Buffer
	cvec        *dt.ConstPool
	st          *dt.DataStack
	lastLabelID uint16
}

func New() *Compiler {
	return &Compiler{
		cvec: &dt.ConstPool{},
	}
}

func (cl *Compiler) CompileFunc(f *sexp.Func) *ir.Object {
	cl.reset(f.Params)

	compileStmtList(cl, f.Body.Forms)

	return &ir.Object{
		StackUsage: cl.st.MaxLen(),
		Code:       cl.buf.Bytes(),
		ConstVec:   cl.cvec,
	}
}

// Prepare compiler for re-use.
func (cl *Compiler) reset(bindings []string) {
	cl.buf.Truncate(0)
	cl.cvec.Clear()
	cl.st = dt.NewDataStack(bindings)
	cl.lastLabelID = 0
}
