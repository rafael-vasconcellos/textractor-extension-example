package main

import (
	"C"
	"errors"
	"unsafe"
)

// Struct Go que replica layout C++
type InfoForExtension struct {
	name  *C.char //*byte
	value int64
}

// SentenceInfo contém um array de InfoForExtension.
type SentenceInfo struct {
	InfoArray []InfoForExtension
}

func cStringToGoString(cstr *byte) string {
	if cstr == nil {
		return ""
	}
	var bytes []byte
	for p := cstr; *p != 0; p = (*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 1)) {
		bytes = append(bytes, *p)
	}
	return string(bytes)
}

// Get retorna o valor associado ao propertyName, ou retorna um erro se não encontrar.
func (s *SentenceInfo) Get(propertyName string) (int64, error) {
	for _, info := range s.InfoArray {
		if info.name == nil {
			break // Marca o fim do array, como no C++
		}
		goName := C.GoString(info.name)
		if propertyName == goName {
			return info.value, nil
		}
	}
	return 0, errors.New("property not found in SentenceInfo")
}

// Skip lança SkipError, similar ao throw SKIP() em C++.
func Skip() error {
	var SkipError = errors.New("skip called")
	return SkipError
}

func NewSentenceInfo(info *InfoForExtension) SentenceInfo {
	return SentenceInfo{
		InfoArray: []InfoForExtension{*info},
	}
}
