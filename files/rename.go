package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// get path file
	// create another file
	// rename this file
	// remove the old file

	src := "abc/def/ghj/src.txt"

	dir := filepath.Dir(src)

	dst := dir + "/dst.txt"

	fmt.Println(src)
	fmt.Println(dst)

	// ext := filepath.Ext(dst)
	// fmt.Println("ext", ext)

	// bkp := strings.TrimSuffix(dst, ext) + "_bkb" + ext

	// fmt.Println("TrimSuffix", strings.TrimSuffix(dst, ext))

	os.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)

	// copy(dst, bkp)

	copy(src, dst)

	if _, err := os.Stat(src); err != nil {
		// The source does not exist or some other error accessing the source
		log.Fatal("Error to check the source file:", err)
	}
	if _, err := os.Stat(dst); err != nil {
		// The destination exists or some other error accessing the destination
		log.Fatal("Erro to check the temporary file:", err)
	}
	if err := os.Rename(dst, src); err != nil {
		log.Fatal(err)
	}
}

func copy(src, dst string) error {

	checkStatFile, err := os.Stat(src)

	if err != nil {
		return err
	}

	if !checkStatFile.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)

	if err != nil {
		return err
	}

	defer source.Close()

	destination, err := os.Create(dst)

	if err != nil {
		return err
	}

	defer destination.Close()

	io.Copy(destination, source)

	return nil
}
