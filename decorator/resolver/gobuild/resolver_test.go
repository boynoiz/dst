package gobuild_test

import (
	"errors"
	"path/filepath"
	"testing"

	"github.com/boynoiz/dst/decorator/resolver"
	"github.com/boynoiz/dst/decorator/resolver/gobuild"
)

func TestRestorerResolver(t *testing.T) {
	type tc struct{ importPath, fromDir, expectName string }
	tests := []struct {
		resolve func() (end func(), root string, r *gobuild.RestorerResolver)
		name    string
		cases   []tc
		skip    bool
		solo    bool
	}{
		{
			name: "gobuild.Resolver",
			resolve: func() (end func(), root string, r *gobuild.RestorerResolver) {
				src := map[string]string{
					"main1/vendor/a/a.go": "package a1 \n\n func A(){}",
					"main1/main1.go":      "package main \n\n func main(){}",
					"main2/main2.go":      "package main \n\n func main(){}",
					"a/a.go":              "package a2 \n\n func A(){}",
				}
				bc, err := buildContext(src)
				if err != nil {
					t.Fatal(err)
				}
				r = &gobuild.RestorerResolver{Context: bc}
				root = "/gopath/src"

				return end, root, r
			},
			cases: []tc{
				{"a", "/main1", "a1"},
				{"a", "/main2", "a2"},
			},
		},
	}
	var solo bool
	for _, test := range tests {
		if test.solo {
			solo = true

			break
		}
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if solo && !test.solo {
				t.Skip()
			}
			if test.skip {
				t.Skip()
			}
			for _, c := range test.cases {
				end, root, r := test.resolve()
				fromDir := filepath.Join(root, c.fromDir)
				r.Dir = fromDir
				name, err := r.ResolvePackage(c.importPath)
				if end != nil {
					end() // delete temp dir if created
				}
				if errors.Is(err, resolver.ErrPackageNotFound) {
					name = ""
				} else if err != nil {
					t.Errorf("error resolving path %s from dir %s: %v", c.importPath, fromDir, err)
				}
				if name != c.expectName {
					t.Errorf("package %s, dir %s - expected %s, got %s", c.importPath, c.fromDir, c.expectName, name)
				}
			}
		})
	}
}
