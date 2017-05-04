package sexpconv

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"sexp"
)

func Stmt(info *types.Info, node ast.Stmt) sexp.Form {
	switch node := node.(type) {
	case *ast.IfStmt:
		return IfStmt(info, node)
	case *ast.ReturnStmt:
		return ReturnStmt(info, node)
	case *ast.BlockStmt:
		return BlockStmt(info, node)
	case *ast.DeclStmt:
		return DeclStmt(info, node)
	case *ast.AssignStmt:
		return AssignStmt(info, node)

	default:
		panic(fmt.Sprintf("unexpected stmt: %#v\n", node))
	}
}

func IfStmt(info *types.Info, node *ast.IfStmt) *sexp.If {
	if node.Init != nil {
		panic("unimplemented")
	}

	test := Expr(info, node.Cond)
	then := BlockStmt(info, node.Body)
	form := &sexp.If{Test: test, Then: then}
	if node.Else != nil {
		form.Else = Stmt(info, node.Else)
	}

	return form
}

func ReturnStmt(info *types.Info, node *ast.ReturnStmt) *sexp.Return {
	return &sexp.Return{Results: exprList(info, node.Results)}
}

func BlockStmt(info *types.Info, node *ast.BlockStmt) *sexp.Block {
	return &sexp.Block{
		Forms: stmtList(info, node.List),
		Scope: info.Scopes[node],
	}
}

func DeclStmt(info *types.Info, node *ast.DeclStmt) sexp.Form {
	decl := node.Decl.(*ast.GenDecl)

	switch decl.Tok {
	case token.VAR:
		return varDecl(info, decl)
	}

	panic("unimplemented")
}

func varDecl(info *types.Info, node *ast.GenDecl) *sexp.FormList {
	forms := make([]sexp.Form, 0, 1)

	for _, spec := range node.Specs {
		spec := spec.(*ast.ValueSpec)

		for i, ident := range spec.Names {
			forms = append(forms, &sexp.Bind{
				Name: ident.Name,
				Init: Expr(info, spec.Values[i]),
			})
		}
	}

	return &sexp.FormList{Forms: forms}
}

func AssignStmt(info *types.Info, node *ast.AssignStmt) *sexp.FormList {
	forms := make([]sexp.Form, 0, 1)

	for i, lhs := range node.Lhs {
		switch lhs := lhs.(type) {
		case *ast.Ident:
			def := info.Defs[lhs]
			expr := Expr(info, node.Rhs[i])
			name := lhs.Name

			if def == nil {
				forms = append(forms, &sexp.Assign{Name: name, Expr: expr})
			} else {
				forms = append(forms, &sexp.Bind{Name: name, Init: expr})
			}

		default:
			panic("unimplemented")
		}
	}

	return &sexp.FormList{Forms: forms}
}
