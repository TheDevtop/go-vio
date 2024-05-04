package vio

import (
	"io/fs"
	"time"
)

// Virtual file information object
type VirtualFileInfo struct {
	iName    string
	iSize    int64
	iMode    fs.FileMode
	iModTime time.Time
	iIsDir   bool
	iSys     any
}

// Implementation of os.FileInfo interface
func (vfi VirtualFileInfo) Name() string       { return vfi.iName }
func (vfi VirtualFileInfo) Size() int64        { return vfi.iSize }
func (vfi VirtualFileInfo) Mode() fs.FileMode  { return vfi.iMode }
func (vfi VirtualFileInfo) ModTime() time.Time { return vfi.iModTime }
func (vfi VirtualFileInfo) IsDir() bool        { return vfi.iIsDir }
func (vfi VirtualFileInfo) Sys() any           { return vfi.iSys }

// Allocation function for VirtualFileInfo
func MakeVirtualFileInfo(n string, si int64, m fs.FileMode, mt time.Time, i bool, sy any) VirtualFileInfo {
	return VirtualFileInfo{
		iName:    n,
		iSize:    si,
		iMode:    m,
		iModTime: mt,
		iIsDir:   i,
		iSys:     sy,
	}
}
