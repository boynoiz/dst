package dst

// ArrayTypeDecorations holds decorations for ArrayType:
//
//	type R /*Start*/ [ /*Lbrack*/ 1] /*Len*/ int /*End*/
type ArrayTypeDecorations struct {
	Lbrack Decorations
	Len    Decorations
	NodeDecs
}

// AssignStmtDecorations holds decorations for AssignStmt:
//
//	*Start*/
//		i = /*Tok*/ 1 /*End*/
//
//		//
type AssignStmtDecorations struct {
	Tok Decorations
	NodeDecs
}

// BadDeclDecorations holds decorations for BadDecl:.
type BadDeclDecorations struct {
	NodeDecs
}

// BadExprDecorations holds decorations for BadExpr:.
type BadExprDecorations struct {
	NodeDecs
}

// BadStmtDecorations holds decorations for BadStmt:.
type BadStmtDecorations struct {
	NodeDecs
}

// BasicLitDecorations holds decorations for BasicLit:.
type BasicLitDecorations struct {
	NodeDecs
}

// BinaryExprDecorations holds decorations for BinaryExpr:
//
//	var P = /*Start*/ 1 /*X*/ & /*Op*/ 2 /*End*/
type BinaryExprDecorations struct {
	X  Decorations
	Op Decorations
	NodeDecs
}

// BlockStmtDecorations holds decorations for BlockStmt:
//
//	f true /*Start*/ { /*Lbrace*/
//			i++
//		} /*End*/
//
//		//
//
//	unc() /*Start*/ { /*Lbrace*/ i++ } /*End*/ ()
//
//		//
type BlockStmtDecorations struct {
	Lbrace Decorations
	NodeDecs
}

// BranchStmtDecorations holds decorations for BranchStmt:
//
//	*Start*/
//		goto /*Tok*/ A /*End*/
//
//		//
type BranchStmtDecorations struct {
	Tok Decorations
	NodeDecs
}

// CallExprDecorations holds decorations for CallExpr:
//
//	var L = /*Start*/ C /*Fun*/ ( /*Lparen*/ 0, []int{}... /*Ellipsis*/) /*End*/
type CallExprDecorations struct {
	Fun      Decorations
	Lparen   Decorations
	Ellipsis Decorations
	NodeDecs
}

// CaseClauseDecorations holds decorations for CaseClause:
//
//	witch i {
//		/*Start*/ case /*Case*/ 1: /*Colon*/
//			i++ /*End*/
//		}
//
//		//
type CaseClauseDecorations struct {
	Case  Decorations
	Colon Decorations
	NodeDecs
}

// ChanTypeDecorations holds decorations for ChanType:
//
//	type W /*Start*/ chan /*Begin*/ int /*End*/
//
//	type X /*Start*/ <-chan /*Begin*/ int /*End*/
//
//	type Y /*Start*/ chan /*Begin*/ <- /*Arrow*/ int /*End*/
type ChanTypeDecorations struct {
	Begin Decorations
	Arrow Decorations
	NodeDecs
}

// CommClauseDecorations holds decorations for CommClause:
//
//	elect {
//		/*Start*/ case /*Case*/ a := <-c /*Comm*/ : /*Colon*/
//			print(a) /*End*/
//		}
//
//		//
type CommClauseDecorations struct {
	Case  Decorations
	Comm  Decorations
	Colon Decorations
	NodeDecs
}

// CompositeLitDecorations holds decorations for CompositeLit:
//
//	var D = /*Start*/ A /*Type*/ { /*Lbrace*/ A: 0} /*End*/
type CompositeLitDecorations struct {
	Type   Decorations
	Lbrace Decorations
	NodeDecs
}

// DeclStmtDecorations holds decorations for DeclStmt:.
type DeclStmtDecorations struct {
	NodeDecs
}

// DeferStmtDecorations holds decorations for DeferStmt:
//
//	*Start*/
//		defer /*Defer*/ func() {}() /*End*/
//
//		//
type DeferStmtDecorations struct {
	Defer Decorations
	NodeDecs
}

// EllipsisDecorations holds decorations for Ellipsis:
//
//	func B(a /*Start*/ ... /*Ellipsis*/ int /*End*/) {}
type EllipsisDecorations struct {
	Ellipsis Decorations
	NodeDecs
}

// EmptyStmtDecorations holds decorations for EmptyStmt:.
type EmptyStmtDecorations struct {
	NodeDecs
}

// ExprStmtDecorations holds decorations for ExprStmt:.
type ExprStmtDecorations struct {
	NodeDecs
}

// FieldDecorations holds decorations for Field:
//
//	type A struct {
//		/*Start*/ A int /*Type*/ `a:"a"` /*End*/
//	}
type FieldDecorations struct {
	Type Decorations
	NodeDecs
}

// FieldListDecorations holds decorations for FieldList:
//
//	type A1 struct /*Start*/ { /*Opening*/
//		a, b int
//		c    string
//	} /*End*/
type FieldListDecorations struct {
	Opening Decorations
	NodeDecs
}

// FileDecorations holds decorations for File:
//
//	/*Start*/ package /*Package*/ data /*Name*/
type FileDecorations struct {
	Package Decorations
	Name    Decorations
	NodeDecs
}

// ForStmtDecorations holds decorations for ForStmt:
//
//	*Start*/
//		for /*For*/ {
//			i++
//		} /*End*/
//
//		//
//
//	*Start*/
//		for /*For*/ i < 1 /*Cond*/ {
//			i++
//		} /*End*/
//
//		//
//
//	*Start*/
//		for /*For*/ i = 0; /*Init*/ i < 10; /*Cond*/ i++ /*Post*/ {
//			i++
//		} /*End*/
//
//		//
type ForStmtDecorations struct {
	For  Decorations
	Init Decorations
	Cond Decorations
	Post Decorations
	NodeDecs
}

// FuncDeclDecorations holds decorations for FuncDecl:
//
//	Start*/
//	func /*Func*/ d /*Name*/ (d, e int) /*Params*/ {
//		return
//	} /*End*/
//
//	//
//
//	Start*/
//	func /*Func*/ TP /*Name*/ [P any] /*TypeParams*/ (a int) /*Params*/ (b P) /*Results*/ {
//		return b
//	} /*End*/
//
//	//
//
//	Start*/
//	func /*Func*/ (a *A) /*Recv*/ e /*Name*/ (d, e int) /*Params*/ {
//		return
//	} /*End*/
//
//	//
//
//	Start*/
//	func /*Func*/ (a *A) /*Recv*/ f /*Name*/ (d, e int) /*Params*/ (f, g int) /*Results*/ {
//		return
//	} /*End*/
//
//	//
type FuncDeclDecorations struct {
	Func       Decorations
	Recv       Decorations
	Name       Decorations
	TypeParams Decorations
	Params     Decorations
	Results    Decorations
	NodeDecs
}

// FuncLitDecorations holds decorations for FuncLit:
//
//	var C = /*Start*/ func(a int, b ...int) (c int) /*Type*/ { return 0 } /*End*/
type FuncLitDecorations struct {
	Type Decorations
	NodeDecs
}

// FuncTypeDecorations holds decorations for FuncType:
//
//	type T /*Start*/ func /*Func*/ (a int) /*Params*/ (b int) /*End*/
type FuncTypeDecorations struct {
	Func       Decorations
	TypeParams Decorations
	Params     Decorations
	NodeDecs
}

// GenDeclDecorations holds decorations for GenDecl:
//
//	*Start*/
//		const /*Tok*/ ( /*Lparen*/
//			a, b = 1, 2
//			c    = 3
//		) /*End*/
//
//		//
//
//	*Start*/
//		const /*Tok*/ d = 1 /*End*/
//
//		//
type GenDeclDecorations struct {
	Tok    Decorations
	Lparen Decorations
	NodeDecs
}

// GoStmtDecorations holds decorations for GoStmt:
//
//	*Start*/
//		go /*Go*/ func() {}() /*End*/
//
//		//
type GoStmtDecorations struct {
	Go Decorations
	NodeDecs
}

// IdentDecorations holds decorations for Ident:
//
//	*Start*/
//		i /*End*/ ++
//
//		//
//
//	*Start*/
//		fmt. /*X*/ Print /*End*/ ()
//
//		//
type IdentDecorations struct {
	X Decorations
	NodeDecs
}

// IfStmtDecorations holds decorations for IfStmt:
//
//	*Start*/
//		if /*If*/ a := b; /*Init*/ a /*Cond*/ {
//			i++
//		} else /*Else*/ {
//			i++
//		} /*End*/
//
//		//
type IfStmtDecorations struct {
	If   Decorations
	Init Decorations
	Cond Decorations
	Else Decorations
	NodeDecs
}

// ImportSpecDecorations holds decorations for ImportSpec:
//
//	import (
//		/*Start*/ fmt /*Name*/ "fmt" /*End*/
//	)
type ImportSpecDecorations struct {
	Name Decorations
	NodeDecs
}

// IncDecStmtDecorations holds decorations for IncDecStmt:
//
//	*Start*/
//		i /*X*/ ++ /*End*/
//
//		//
type IncDecStmtDecorations struct {
	X Decorations
	NodeDecs
}

// IndexExprDecorations holds decorations for IndexExpr:
//
//	var G = /*Start*/ []int{0} /*X*/ [ /*Lbrack*/ 0 /*Index*/] /*End*/
type IndexExprDecorations struct {
	X      Decorations
	Lbrack Decorations
	Index  Decorations
	NodeDecs
}

// IndexListExprDecorations holds decorations for IndexListExpr:
//
//	ar T4 /*Start*/ T3 /*X*/ [ /*Lbrack*/ int, string /*Indices*/] /*End*/
//
//		//
type IndexListExprDecorations struct {
	X       Decorations
	Lbrack  Decorations
	Indices Decorations
	NodeDecs
}

// InterfaceTypeDecorations holds decorations for InterfaceType:
//
//	type U /*Start*/ interface /*Interface*/ {
//		A()
//	} /*End*/
type InterfaceTypeDecorations struct {
	Interface Decorations
	NodeDecs
}

// KeyValueExprDecorations holds decorations for KeyValueExpr:
//
//	var Q = map[string]string{
//		/*Start*/ "a" /*Key*/ : /*Colon*/ "a", /*End*/
//	}
type KeyValueExprDecorations struct {
	Key   Decorations
	Colon Decorations
	NodeDecs
}

// LabeledStmtDecorations holds decorations for LabeledStmt:
//
//	/*Start*/
//	A /*Label*/ : /*Colon*/
//		print("Stmt") /*End*/
//
//		//
type LabeledStmtDecorations struct {
	Label Decorations
	Colon Decorations
	NodeDecs
}

// MapTypeDecorations holds decorations for MapType:
//
//	type V /*Start*/ map[ /*Map*/ int] /*Key*/ int /*End*/
type MapTypeDecorations struct {
	Map Decorations
	Key Decorations
	NodeDecs
}

// PackageDecorations holds decorations for Package:.
type PackageDecorations struct {
	NodeDecs
}

// ParenExprDecorations holds decorations for ParenExpr:
//
//	var E = /*Start*/ ( /*Lparen*/ 1 + 1 /*X*/) /*End*/ / 2
type ParenExprDecorations struct {
	Lparen Decorations
	X      Decorations
	NodeDecs
}

// RangeStmtDecorations holds decorations for RangeStmt:
//
//	*Start*/
//		for range /*Range*/ a /*X*/ {
//		} /*End*/
//
//		//
//
//	*Start*/
//		for /*For*/ k /*Key*/ := range /*Range*/ a /*X*/ {
//			print(k)
//		} /*End*/
//
//		//
//
//	*Start*/
//		for /*For*/ k /*Key*/, v /*Value*/ := range /*Range*/ a /*X*/ {
//			print(k, v)
//		} /*End*/
//
//		//
type RangeStmtDecorations struct {
	For   Decorations
	Key   Decorations
	Value Decorations
	Range Decorations
	X     Decorations
	NodeDecs
}

// ReturnStmtDecorations holds decorations for ReturnStmt:
//
//	unc() int {
//			/*Start*/ return /*Return*/ 1 /*End*/
//		}()
//
//		//
type ReturnStmtDecorations struct {
	Return Decorations
	NodeDecs
}

// SelectStmtDecorations holds decorations for SelectStmt:
//
//	*Start*/
//		select /*Select*/ {
//		} /*End*/
//
//		//
type SelectStmtDecorations struct {
	Select Decorations
	NodeDecs
}

// SelectorExprDecorations holds decorations for SelectorExpr:
//
//	var F = /*Start*/ tt. /*X*/ F /*End*/ ()
type SelectorExprDecorations struct {
	X Decorations
	NodeDecs
}

// SendStmtDecorations holds decorations for SendStmt:
//
//	*Start*/
//		c /*Chan*/ <- /*Arrow*/ 0 /*End*/
//
//		//
type SendStmtDecorations struct {
	Chan  Decorations
	Arrow Decorations
	NodeDecs
}

// SliceExprDecorations holds decorations for SliceExpr:
//
//	var H = /*Start*/ []int{0, 1, 2} /*X*/ [ /*Lbrack*/ 1: /*Low*/ 2: /*High*/ 3 /*Max*/] /*End*/
//
//	var H1 = /*Start*/ []int{0, 1, 2} /*X*/ [ /*Lbrack*/ 1: /*Low*/ 2 /*High*/] /*End*/
//
//	var H2 = /*Start*/ []int{0} /*X*/ [: /*Low*/] /*End*/
//
//	var H3 = /*Start*/ []int{0} /*X*/ [ /*Lbrack*/ 1: /*Low*/] /*End*/
//
//	var H4 = /*Start*/ []int{0, 1, 2} /*X*/ [: /*Low*/ 2 /*High*/] /*End*/
//
//	var H5 = /*Start*/ []int{0, 1, 2} /*X*/ [: /*Low*/ 2: /*High*/ 3 /*Max*/] /*End*/
type SliceExprDecorations struct {
	X      Decorations
	Lbrack Decorations
	Low    Decorations
	High   Decorations
	Max    Decorations
	NodeDecs
}

// StarExprDecorations holds decorations for StarExpr:
//
//	var N = /*Start*/ * /*Star*/ p /*End*/
type StarExprDecorations struct {
	Star Decorations
	NodeDecs
}

// StructTypeDecorations holds decorations for StructType:
//
//	type S /*Start*/ struct /*Struct*/ {
//		A int
//	} /*End*/
type StructTypeDecorations struct {
	Struct Decorations
	NodeDecs
}

// SwitchStmtDecorations holds decorations for SwitchStmt:
//
//	*Start*/
//		switch /*Switch*/ i /*Tag*/ {
//		} /*End*/
//
//		//
//
//	*Start*/
//		switch /*Switch*/ a := i; /*Init*/ a /*Tag*/ {
//		} /*End*/
//
//		//
type SwitchStmtDecorations struct {
	Switch Decorations
	Init   Decorations
	Tag    Decorations
	NodeDecs
}

// TypeAssertExprDecorations holds decorations for TypeAssertExpr:
//
//	var J = /*Start*/ f. /*X*/ ( /*Lparen*/ int /*Type*/) /*End*/
type TypeAssertExprDecorations struct {
	X      Decorations
	Lparen Decorations
	Type   Decorations
	NodeDecs
}

// TypeSpecDecorations holds decorations for TypeSpec:
//
//	ype (
//			/*Start*/ T1 /*Name*/ []int /*End*/
//		)
//
//		//
//
//	ype (
//			/*Start*/ T2 = /*Assign*/ T1 /*End*/
//		)
//
//		//
//
//	ype (
//			/*Start*/ T3 /*Name*/ [P any, Q any] /*TypeParams*/ []P /*End*/
//		)
//
//		//
type TypeSpecDecorations struct {
	Name       Decorations
	TypeParams Decorations
	Assign     Decorations
	NodeDecs
}

// TypeSwitchStmtDecorations holds decorations for TypeSwitchStmt:
//
//	*Start*/
//		switch /*Switch*/ f.(type) /*Assign*/ {
//		} /*End*/
//
//		//
//
//	*Start*/
//		switch /*Switch*/ g := f.(type) /*Assign*/ {
//		case int:
//			print(g)
//		} /*End*/
//
//		//
//
//	*Start*/
//		switch /*Switch*/ g := f; /*Init*/ g := g.(type) /*Assign*/ {
//		case int:
//			print(g)
//		} /*End*/
//
//		//
type TypeSwitchStmtDecorations struct {
	Switch Decorations
	Init   Decorations
	Assign Decorations
	NodeDecs
}

// UnaryExprDecorations holds decorations for UnaryExpr:
//
//	var O = /*Start*/ ^ /*Op*/ 1 /*End*/
type UnaryExprDecorations struct {
	Op Decorations
	NodeDecs
}

// ValueSpecDecorations holds decorations for ValueSpec:
//
//	ar (
//			/*Start*/ j = /*Assign*/ 1 /*End*/
//		)
//
//		//
//
//	ar (
//			/*Start*/ k, l = /*Assign*/ 1, 2 /*End*/
//		)
//
//		//
//
//	ar (
//			/*Start*/ m, n int = /*Assign*/ 1, 2 /*End*/
//		)
//
//		//
type ValueSpecDecorations struct {
	Assign Decorations
	NodeDecs
}
