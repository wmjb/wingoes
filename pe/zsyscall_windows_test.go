// Code generated by 'go generate'; DO NOT EDIT.

package pe

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
	errERROR_EINVAL     error = syscall.EINVAL
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return errERROR_EINVAL
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	moddbghelp  = windows.NewLazySystemDLL("dbghelp.dll")
	modimagehlp = windows.NewLazySystemDLL("imagehlp.dll")

	procImageDirectoryEntryToDataEx = moddbghelp.NewProc("ImageDirectoryEntryToDataEx")
	procImageNtHeader               = moddbghelp.NewProc("ImageNtHeader")
	procSymSrvGetFileIndexInfoW     = moddbghelp.NewProc("SymSrvGetFileIndexInfoW")
	procImageEnumerateCertificates  = modimagehlp.NewProc("ImageEnumerateCertificates")
	procImageGetCertificateData     = modimagehlp.NewProc("ImageGetCertificateData")
)

func imageDirectoryEntryToDataEx(base uintptr, mappedAsImage byte, directoryEntry uint16, size *uint32, foundHeader *SectionHeader) (ret uintptr, err error) {
	r0, _, e1 := syscall.Syscall6(procImageDirectoryEntryToDataEx.Addr(), 5, uintptr(base), uintptr(mappedAsImage), uintptr(directoryEntry), uintptr(unsafe.Pointer(size)), uintptr(unsafe.Pointer(foundHeader)), 0)
	ret = uintptr(r0)
	if ret == 0 {
		err = errnoErr(e1)
	}
	return
}

func imageNtHeader(base uintptr) (ret *_IMAGE_NT_HEADERS_FIXED, err error) {
	r0, _, e1 := syscall.Syscall(procImageNtHeader.Addr(), 1, uintptr(base), 0, 0)
	ret = (*_IMAGE_NT_HEADERS_FIXED)(unsafe.Pointer(r0))
	if ret == nil {
		err = errnoErr(e1)
	}
	return
}

func symSrvGetFileIndexInfoW(file *uint16, info *_SYMSRV_INDEX_INFO, flags uint32) (err error) {
	err = procSymSrvGetFileIndexInfoW.Find()
	if err != nil {
		return
	}
	r1, _, e1 := syscall.Syscall(procSymSrvGetFileIndexInfoW.Addr(), 3, uintptr(unsafe.Pointer(file)), uintptr(unsafe.Pointer(info)), uintptr(flags))
	if int32(r1) == 0 {
		err = errnoErr(e1)
	}
	return
}

func imageEnumerateCertificates(fileHandle windows.Handle, typeFilter WIN_CERT_TYPE, certificateCount *uint32, indices *uint32, indexCount uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procImageEnumerateCertificates.Addr(), 5, uintptr(fileHandle), uintptr(typeFilter), uintptr(unsafe.Pointer(certificateCount)), uintptr(unsafe.Pointer(indices)), uintptr(indexCount), 0)
	if int32(r1) == 0 {
		err = errnoErr(e1)
	}
	return
}

func imageGetCertificateData(fileHandle windows.Handle, certificateIndex uint32, certificate *byte, requiredLength *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procImageGetCertificateData.Addr(), 4, uintptr(fileHandle), uintptr(certificateIndex), uintptr(unsafe.Pointer(certificate)), uintptr(unsafe.Pointer(requiredLength)), 0, 0)
	if int32(r1) == 0 {
		err = errnoErr(e1)
	}
	return
}