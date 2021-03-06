// +build !android,arm,!js

package raylib

/*
#include "raylib.h"
#include <stdlib.h>
*/
import "C"

import (
	"os"
	"unsafe"
)

// InitWindow - Initialize Window and OpenGL Graphics
func InitWindow(width int32, height int32, t interface{}) {
	cwidth := (C.int)(width)
	cheight := (C.int)(height)

	title, ok := t.(string)
	if ok {
		ctitle := C.CString(title)
		cptitle := unsafe.Pointer(ctitle)
		defer C.free(cptitle)
		C.InitWindow(cwidth, cheight, cptitle)
	}
}

// SetCallbackFunc - Sets callback function
func SetCallbackFunc(func(unsafe.Pointer)) {
	return
}

// SetMainLoop - Sets main loop function
func SetMainLoop(f func(), fps, simulateInfiniteLoop int) {
}

// ShowCursor - Shows cursor
func ShowCursor() {
	C.ShowCursor()
}

// HideCursor - Hides cursor
func HideCursor() {
	C.HideCursor()
}

// IsCursorHidden - Returns true if cursor is not visible
func IsCursorHidden() bool {
	ret := C.IsCursorHidden()
	v := bool(int(ret) == 1)
	return v
}

// EnableCursor - Enables cursor
func EnableCursor() {
	C.EnableCursor()
}

// DisableCursor - Disables cursor
func DisableCursor() {
	C.DisableCursor()
}

// IsFileDropped - Check if a file have been dropped into window
func IsFileDropped() bool {
	return false
}

// GetDroppedFiles - Retrieve dropped files into window
func GetDroppedFiles(count *int32) (files []string) {
	return
}

// ClearDroppedFiles - Clear dropped files paths buffer
func ClearDroppedFiles() {
	return
}

// OpenAsset - Open asset
func OpenAsset(name string) (Asset, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	return f, nil
}
