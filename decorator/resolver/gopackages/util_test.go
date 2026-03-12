package gopackages_test

import (
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"strings"
)

func tempDir(m map[string]string) (dir string, err error) {
	if dir, err = os.MkdirTemp("", ""); err != nil {
		return dir, err
	}
	for fpathrel, src := range m {
		if strings.HasSuffix(fpathrel, "/") {
			// just a dir
			if err = os.MkdirAll(filepath.Join(dir, fpathrel), 0o777); err != nil {
				return dir, err
			}
		} else {
			fpath := filepath.Join(dir, fpathrel)
			fdir, _ := filepath.Split(fpath)
			if err = os.MkdirAll(fdir, 0o777); err != nil {
				return dir, err
			}

			var formatted []byte
			if strings.HasSuffix(fpath, ".go") {
				formatted, err = format.Source([]byte(src))
				if err != nil {
					err = fmt.Errorf("formatting %s: %w", fpathrel, err)

					return dir, err
				}
			} else {
				formatted = []byte(src)
			}

			if err = os.WriteFile(fpath, formatted, 0o666); err != nil {
				return dir, err
			}
		}
	}

	return dir, err
}
