package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	tempDir, err := ioutil.TempDir("", "test-symlinks")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	wd, _ := os.Getwd()
	defer func() {
		os.Chdir(wd)
	}()

	must := func(err error) {
		if err != nil {
			panic(err)
		}
	}

	dir1Dir := filepath.Join(tempDir, "dir1")
	dir1Real := filepath.Join(dir1Dir, "real")
	dir2Dir := filepath.Join(tempDir, "dir2")

	must(os.MkdirAll(filepath.Join(dir1Real), 0777))
	must(os.MkdirAll(filepath.Join(dir2Dir), 0777))

	must(os.Chdir(dir1Dir))
	must(os.Symlink("real", "symlinked"))
	must(os.Chdir(dir1Real))
	must(os.Symlink("../real", "cyclic"))
	os.Chdir(dir2Dir)
	must(os.Symlink("../dir1/real/cyclic", "dir2real"))

	_, err = filepath.EvalSymlinks(filepath.Join(dir2Dir, "dir2real", "cyclic"))
	must(err)

}
