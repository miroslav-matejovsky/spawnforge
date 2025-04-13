package spawnforge

import (
	"embed"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

//go:embed all:test/content/file.txt
var file string

//go:embed all:test/content
var content embed.FS

func copyFs(src fs.FS, dst string) error {
	return fs.WalkDir(src, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		srcFile, err := src.Open(path)
		if err != nil {
			return err
		}
		defer func() {
			if err := srcFile.Close(); err != nil {
				panic(err)
			}
		}()

		dstPath := filepath.Join(dst, path)
		if err := os.MkdirAll(filepath.Dir(dstPath), os.ModePerm); err != nil {
			return err
		}

		dstFile, err := os.Create(dstPath)
		if err != nil {
			return err
		}
		defer func() {
			if err := dstFile.Close(); err != nil {
				panic(err)
			}
		}()

		_, err = io.Copy(dstFile, srcFile)
		return err
	})
}
