//go:build windows
// +build windows

package discordcgo

/*
#include <windows.h>
*/
import "C"

func getCurrentThreadID() uint64 {
	return uint64(C.GetCurrentThreadId())
}
