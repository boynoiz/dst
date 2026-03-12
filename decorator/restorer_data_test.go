package decorator

import "testing"

func TestData(t *testing.T) {
	testPackageRestoresCorrectlyWithImports(t, "github.com/boynoiz/dst/gendst/data")
}
