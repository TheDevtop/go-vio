package vio

import (
	"fmt"
	"io/fs"
	"os"
	"testing"
	"time"
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

// Virtual stat function
func vStat(path string) (fs.FileInfo, error) {
	vvfi := MakeVirtualFileInfo(path, 0, 0664, time.Now(), false, nil)
	return vvfi, nil
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

func Test_StatFunc(t *testing.T) {
	// Preparation
	var (
		path    = "/tmp/data"
		getStat = func(fn StatFunc) (os.FileInfo, error) {
			return fn(path)
		}
		fi  os.FileInfo
		err error
	)

	// Execution
	fi, err = getStat(vStat) // Can be replaced with os.Stat for comparison

	// Assertion
	if err != nil || fi.Name() != path {
		t.Fatal(err)
	}
}
