package vio

import (
	"fmt"
	"io/fs"
	"testing"
)

// Virtual read function
func vRead(path string) ([]byte, error) {
	str := fmt.Sprintf("Contents of %s\n", path)
	return []byte(str), nil
}

// Virtual write function
func vWrite(path string, content []byte, perm fs.FileMode) error {
	return nil
}

// Virtual remove function
func vRemove(path string) error {
	return nil
}

// Virtual make directory function
func vMkDir(path string, perm fs.FileMode) error {
	return nil
}

func Test_ReadFunc(t *testing.T) {
	// Preparation
	var (
		loadConf = func(fn ReadFunc) ([]byte, error) {
			return fn("/tmp/config")
		}
		buf []byte
		err error
	)

	// Execution
	buf, err = loadConf(vRead) // Can be replaced with os.ReadFile for comparison

	// Assertion
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(string(buf))
	}
}

func Test_WriteFunc(t *testing.T) {
	// Preparation
	var (
		storeConf = func(fn WriteFunc) error {
			return fn("/tmp/config", []byte{0x00}, 0644)
		}
		err error
	)

	// Execution
	err = storeConf(vWrite) // Can be replaced with os.WriteFile for comparison

	// Assertion
	if err != nil {
		t.Fatal(err)
	}
}

func Test_RemoveFunc(t *testing.T) {
	// Preparation
	var (
		removeConf = func(fn RemoveFunc) error {
			return fn("/tmp/config")
		}
		err error
	)

	// Execution
	err = removeConf(vRemove) // Can be replaced with os.Remove for comparison

	// Assertion
	if err != nil {
		t.Fatal(err)
	}
}

func Test_DirFunc(t *testing.T) {
	// Preparation
	var (
		mkDataDir = func(fn DirFunc) error {
			return fn("/tmp/data", 0644)
		}
		err error
	)

	// Execution
	err = mkDataDir(vMkDir) // Can be replaced with os.Mkdir for comparison

	// Assertion
	if err != nil {
		t.Fatal(err)
	}
}
