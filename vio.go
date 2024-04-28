package vio

import "io/fs"

// Virtual function for os.ReadFile
type ReadFunc func(string) ([]byte, error)

// Virtual function for os.WriteFile
type WriteFunc func(string, []byte, fs.FileMode) error

// Virtual function for os.Remove/os.RemoveAll
type RemoveFunc func(string) error

// Virtual function for os.Mkdir/os.MkdirAll
type DirFunc func(string, fs.FileMode) error
