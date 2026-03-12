package gobuild

import (
	"go/build"

	"github.com/boynoiz/dst/decorator/resolver"
)

func New(dir string) *RestorerResolver {
	return &RestorerResolver{Dir: dir}
}

func WithContext(dir string, context *build.Context) *RestorerResolver {
	return &RestorerResolver{Dir: dir, Context: context}
}

func WithHints(dir string, hints map[string]string) *RestorerResolver {
	return &RestorerResolver{Dir: dir, Hints: hints}
}

type RestorerResolver struct {
	FindPackage func(ctxt *build.Context, importPath, fromDir string, mode build.ImportMode) (*build.Package, error)
	Context     *build.Context
	Hints       map[string]string
	Dir         string
}

func (r *RestorerResolver) ResolvePackage(importPath string) (string, error) {
	if name, ok := r.Hints[importPath]; ok {
		return name, nil
	}

	fp := r.FindPackage
	if fp == nil {
		fp = (*build.Context).Import
	}

	bc := r.Context
	if bc == nil {
		bc = &build.Default
	}

	p, err := fp(bc, importPath, r.Dir, 0)
	if err != nil {
		return "", err
	}

	if p == nil {
		return "", resolver.ErrPackageNotFound
	}

	return p.Name, nil
}
