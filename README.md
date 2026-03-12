# dst — Decorated Syntax Tree

> **This is a maintained fork of [`github.com/dave/dst`](https://github.com/dave/dst).**
>
> The upstream repository is unmaintained and incompatible with modern Go versions (Go 1.21+).
> This fork modernizes the codebase, cherry-picks accepted upstream PRs, and adds support for
> newer Go language features.
>
> **Module path:** `github.com/boynoiz/dst`

## Changes from upstream

| Version | Change |
|---------|--------|
| fork | Module renamed from `github.com/dave/dst` to `github.com/boynoiz/dst` |
| fork | Fixed struct literal field ordering and missing `errors` imports |
| fork | Fix: restorer only updates space type if new value is larger (upstream PR #86) |
| fork | Feat: `GoVersion` field on `dst.File` (Go 1.21+) |
| fork | Feat: Go 1.24 generic type alias support (`TypeSpec.Assign` token ordering) |
| fork | Modernize: `ioutil` → `os`/`io`, octal literals, `os.DirEntry` API |

See [CHANGELOG.md](./CHANGELOG.md) for the full history.

## Installation

```bash
go get github.com/boynoiz/dst
```

---

## Original Documentation

The `dst` package enables manipulation of a Go syntax tree with high fidelity. Decorations (e.g.
comments and line spacing) remain attached to the correct nodes as the tree is modified.

### Where does `go/ast` break?

The `go/ast` package wasn't created with source manipulation as an intended use-case. Comments are
stored by their byte offset instead of attached to nodes, so re-arranging nodes breaks the output.
See [this Go issue](https://github.com/golang/go/issues/20744) for more information.

Consider this example where we want to reverse the order of the two statements. As you can see the
comments don't remain attached to the correct nodes:

```go
code := `package a

func main(){
	var a int    // foo
	var b string // bar
}
`
fset := token.NewFileSet()
f, err := parser.ParseFile(fset, "", code, parser.ParseComments)
if err != nil {
	panic(err)
}

list := f.Decls[0].(*ast.FuncDecl).Body.List
list[0], list[1] = list[1], list[0]

if err := format.Node(os.Stdout, fset, f); err != nil {
	panic(err)
}

//Output:
//package a
//
//func main() {
//	// foo
//	var b string
//	var a int
//	// bar
//}
```

Here's the same example using `dst`:

```go
code := `package a

func main(){
	var a int    // foo
	var b string // bar
}
`
f, err := decorator.Parse(code)
if err != nil {
	panic(err)
}

list := f.Decls[0].(*dst.FuncDecl).Body.List
list[0], list[1] = list[1], list[0]

if err := decorator.Print(f); err != nil {
	panic(err)
}

//Output:
//package a
//
//func main() {
//	var b string // bar
//	var a int    // foo
//}
```

### Usage

Parsing a source file to `dst` and printing the results after modification can be accomplished with
several `Parse` and `Print` convenience functions in the
[decorator](https://pkg.go.dev/github.com/boynoiz/dst/decorator) package.

For more fine-grained control you can use
[Decorator](https://pkg.go.dev/github.com/boynoiz/dst/decorator#Decorator)
to convert from `ast` to `dst`, and
[Restorer](https://pkg.go.dev/github.com/boynoiz/dst/decorator#Restorer)
to convert back again.

### Comments

Comments are added at decoration attachment points. [See here](https://github.com/boynoiz/dst/blob/master/decorations-types-generated.go)
for a full list of these points, along with demonstration code of where they are rendered in the
output.

The decoration attachment points have convenience functions `Append`, `Prepend`, `Replace`, `Clear`
and `All` to accomplish common tasks. Use the full text of your comment including the `//` or `/**/`
markers. When adding a line comment, a newline is automatically rendered.

```go
code := `package main

func main() {
	println("Hello World!")
}`
f, err := decorator.Parse(code)
if err != nil {
	panic(err)
}

call := f.Decls[0].(*dst.FuncDecl).Body.List[0].(*dst.ExprStmt).X.(*dst.CallExpr)

call.Decs.Start.Append("// you can add comments at the start...")
call.Decs.Fun.Append("/* ...in the middle... */")
call.Decs.End.Append("// or at the end.")

if err := decorator.Print(f); err != nil {
	panic(err)
}

//Output:
//package main
//
//func main() {
//	// you can add comments at the start...
//	println /* ...in the middle... */ ("Hello World!") // or at the end.
//}
```

### Spacing

The `Before` property marks the node as having a line space (new line or empty line) before the node.
These spaces are rendered before any decorations attached to the `Start` decoration point. The `After`
property is similar but rendered after the node (and after any `End` decorations).

```go
code := `package main

func main() {
	println(a, b, c)
}`
f, err := decorator.Parse(code)
if err != nil {
	panic(err)
}

call := f.Decls[0].(*dst.FuncDecl).Body.List[0].(*dst.ExprStmt).X.(*dst.CallExpr)

call.Decs.Before = dst.EmptyLine
call.Decs.After = dst.EmptyLine

for _, v := range call.Args {
	v := v.(*dst.Ident)
	v.Decs.Before = dst.NewLine
	v.Decs.After = dst.NewLine
}

if err := decorator.Print(f); err != nil {
	panic(err)
}

//Output:
//package main
//
//func main() {
//
//	println(
//		a,
//		b,
//		c,
//	)
//
//}
```

### Decorations

The common decoration properties (`Start`, `End`, `Before` and `After`) occur on all nodes, and can be
accessed with the `Decorations()` method on the `Node` interface:

```go
code := `package main

func main() {
	var i int
	i++
	println(i)
}`
f, err := decorator.Parse(code)
if err != nil {
	panic(err)
}

list := f.Decls[0].(*dst.FuncDecl).Body.List

list[0].Decorations().Before = dst.EmptyLine
list[0].Decorations().End.Append("// the Decorations method allows access to the common")
list[1].Decorations().End.Append("// decoration properties (Before, Start, End and After)")
list[2].Decorations().End.Append("// for all nodes.")
list[2].Decorations().After = dst.EmptyLine

if err := decorator.Print(f); err != nil {
	panic(err)
}

//Output:
//package main
//
//func main() {
//
//	var i int  // the Decorations method allows access to the common
//	i++        // decoration properties (Before, Start, End and After)
//	println(i) // for all nodes.
//
//}
```

#### dstutil.Decorations

While debugging, it is often useful to have a list of all decorations attached to a node. The
[dstutil](https://github.com/boynoiz/dst/tree/master/dstutil) package provides a helper function `Decorations` which
returns a list of the attachment points and all decorations for any node:

```go
code := `package main

// main comment
// is multi line
func main() {

	if true {

		// foo
		println( /* foo inline */ "foo")
	} else if false {
		println /* bar inline */ ("bar")

		// bar after

	} else {
		// empty block
	}
}`

f, err := decorator.Parse(code)
if err != nil {
	panic(err)
}

dst.Inspect(f, func(node dst.Node) bool {
	if node == nil {
		return false
	}
	before, after, points := dstutil.Decorations(node)
	var info string
	if before != dst.None {
		info += fmt.Sprintf("- Before: %s\n", before)
	}
	for _, point := range points {
		if len(point.Decs) == 0 {
			continue
		}
		info += fmt.Sprintf("- %s: [", point.Name)
		for i, dec := range point.Decs {
			if i > 0 {
				info += ", "
			}
			info += fmt.Sprintf("%q", dec)
		}
		info += "]\n"
	}
	if after != dst.None {
		info += fmt.Sprintf("- After: %s\n", after)
	}
	if info != "" {
		fmt.Printf("%T\n%s\n", node, info)
	}
	return true
})
```

### Newlines

The `Before` and `After` properties cover the majority of cases, but occasionally a newline needs to
be rendered inside a node. Simply add a `\n` decoration to accomplish this.

### Clone

Re-using an existing node elsewhere in the tree will panic when the tree is restored to `ast`. Instead,
use the `Clone` function to make a deep copy of the node before re-use:

```go
code := `package main

var i /* a */ int`

f, err := decorator.Parse(code)
if err != nil {
	panic(err)
}

cloned := dst.Clone(f.Decls[0]).(*dst.GenDecl)

cloned.Decs.Before = dst.NewLine
cloned.Specs[0].(*dst.ValueSpec).Names[0].Name = "j"
cloned.Specs[0].(*dst.ValueSpec).Names[0].Decs.End.Replace("/* b */")

f.Decls = append(f.Decls, cloned)

if err := decorator.Print(f); err != nil {
	panic(err)
}

//Output:
//package main
//
//var i /* a */ int
//var j /* b */ int
```

### Apply

The [dstutil](https://github.com/boynoiz/dst/tree/master/dstutil) package is a fork of `golang.org/x/tools/go/ast/astutil`,
and provides the `Apply` function with similar semantics.

### Imports

The decorator can automatically manage the `import` block, which is a non-trivial task.

Use [NewDecoratorWithImports](https://pkg.go.dev/github.com/boynoiz/dst/decorator#NewDecoratorWithImports)
and [NewRestorerWithImports](https://pkg.go.dev/github.com/boynoiz/dst/decorator#NewRestorerWithImports)
to create an import-aware decorator / restorer.

During decoration, remote identifiers are normalised — `*ast.SelectorExpr` nodes that represent
qualified identifiers are replaced with `*dst.Ident` nodes with the `Path` field set to the path of
the imported package.

When adding a qualified identifier node, there is no need to use `*dst.SelectorExpr` — just add a
`*dst.Ident` and set `Path` to the imported package path. The restorer will wrap it in a
`*ast.SelectorExpr` where appropriate when converting back to ast, and also update the import block.

### Load

The [Load](https://pkg.go.dev/github.com/boynoiz/dst/decorator#Load) convenience function uses
`go/packages` to load packages and decorate all loaded ast files, with import management enabled:

```go
pkgs, err := decorator.Load(&packages.Config{Dir: dir, Mode: packages.LoadSyntax}, "root")
if err != nil {
	panic(err)
}
p := pkgs[0]
f := p.Syntax[0]

b := f.Decls[0].(*dst.FuncDecl).Body
b.List = append(b.List, &dst.ExprStmt{
	X: &dst.CallExpr{
		Fun:  &dst.Ident{Path: "fmt", Name: "Println"},
		Args: []dst.Expr{
			&dst.BasicLit{Kind: token.STRING, Value: strconv.Quote("Hello, World!")},
		},
	},
})

r := decorator.NewRestorerWithImports("root", gopackages.New(dir))
if err := r.Print(p.Syntax[0]); err != nil {
	panic(err)
}

//Output:
//package main
//
//import "fmt"
//
//func main() { fmt.Println("Hello, World!") }
```

### Mappings

The decorator exposes `Dst.Nodes` and `Ast.Nodes` which map between `ast.Node` and `dst.Node`. This
enables systems that refer to `ast` nodes (such as `go/types`) to be used.

### Resolvers

Two separate interfaces are defined by the
[resolver package](https://github.com/boynoiz/dst/tree/master/decorator/resolver)
for automatically managing the imports block.

**DecoratorResolver** — resolves the package path of any `*ast.Ident`:
- `gotypes` — full dot-import compatibility using `go/types.Info`
- `goast` — simplified, scans a single ast file (no dot-import support)

**RestorerResolver** — resolves the name of a package given its path:
- `gopackages` — full Go modules compatibility via `golang.org/x/tools/go/packages`
- `gobuild` — legacy `go/build` system, better performance, not modules-aware
- `guess` — guesses package name from the last path segment
- `simple` — resolves only from a provided map

## Status

This fork is actively maintained for internal use. The API is compatible with `dave/dst`.
Known upstream bugs are documented in [CLAUDE.md](./CLAUDE.md).

## Original project

This library is based on [`github.com/dave/dst`](https://github.com/dave/dst) by
[@dave](https://github.com/dave). Original generics support by
[@hawkinsw](https://github.com/hawkinsw).
