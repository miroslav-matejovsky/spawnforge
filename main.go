package main

import (
	"embed"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

//go:embed all:test/content
var content embed.FS

func main() {

	dst := ".tmp"

	println("I copy files to", dst)

	err := copyFs(content, dst)
	if err != nil {
		println("Error copying files:", err)
	} else {
		println("Files copied successfully.")
	}
}

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
		defer srcFile.Close()

		dstPath := filepath.Join(dst, path)
		if err := os.MkdirAll(filepath.Dir(dstPath), os.ModePerm); err != nil {
			return err
		}

		dstFile, err := os.Create(dstPath)
		if err != nil {
			return err
		}
		defer dstFile.Close()

		_, err = io.Copy(dstFile, srcFile)
		return err
	})
}
