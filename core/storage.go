package core

import (
	"unsafe"

	dcgo "github.com/andresperezl/discordctl/discordcgo"
)

// StorageManager provides access to storage-related functionality
type StorageManager struct {
	manager unsafe.Pointer
}

// Read reads data from storage
func (s *StorageManager) Read(name string, data []byte) (int, Result) {
	if s.manager == nil {
		return 0, ResultInternalError
	}
	var read uint32
	res := dcgo.RunOnDispatcherSync(func() int32 {
		return dcgo.StorageManagerReadGo(s.manager, name, data, &read)
	})
	return int(read), Result(res)
}

// ReadAsync reads data from storage asynchronously
func (s *StorageManager) ReadAsync(name string, callback func(result Result, data []byte)) {
	if s.manager == nil {
		if callback != nil {
			callback(ResultInternalError, nil)
		}
		return
	}
	cname := dcgo.GoStringToCChar(name)
	defer dcgo.FreeCChar(cname)
	dcgo.StorageManagerReadAsync(s.manager, cname, func(result int32, data []byte) {
		if callback != nil {
			callback(Result(result), data)
		}
	})
}

// ReadAsyncPartial reads partial data from storage asynchronously
func (s *StorageManager) ReadAsyncPartial(name string, offset, length uint64, callback func(result Result, data []byte)) {
	// Not implemented: would require cgo callback trampoline
	if callback != nil {
		callback(ResultInternalError, nil)
	}
}

// Write writes data to storage
func (s *StorageManager) Write(name string, data []byte) Result {
	if s.manager == nil {
		return ResultInternalError
	}
	res := dcgo.RunOnDispatcherSync(func() int32 {
		return dcgo.StorageManagerWriteGo(s.manager, name, data)
	})
	return Result(res)
}

// WriteAsync writes data to storage asynchronously
func (s *StorageManager) WriteAsync(name string, data []byte, callback func(result Result)) {
	if s.manager == nil {
		if callback != nil {
			callback(ResultInternalError)
		}
		return
	}
	cname := dcgo.GoStringToCChar(name)
	defer dcgo.FreeCChar(cname)
	var datPtr unsafe.Pointer
	if len(data) > 0 {
		datPtr = unsafe.Pointer(&data[0])
	}
	dcgo.StorageManagerWriteAsync(s.manager, cname, datPtr, uint32(len(data)), func(result int32) {
		if callback != nil {
			callback(Result(result))
		}
	})
}

// Delete deletes a file from storage
func (s *StorageManager) Delete(name string) Result {
	if s.manager == nil {
		return ResultInternalError
	}
	res := dcgo.RunOnDispatcherSync(func() int32 {
		return dcgo.StorageManagerDeleteGo(s.manager, name)
	})
	return Result(res)
}

// Exists checks if a file exists in storage
func (s *StorageManager) Exists(name string) (bool, Result) {
	if s.manager == nil {
		return false, ResultInternalError
	}
	var exists bool
	res := dcgo.RunOnDispatcherSync(func() int32 {
		return dcgo.StorageManagerExistsGo(s.manager, name, &exists)
	})
	return exists, Result(res)
}

// Count gets the count of files in storage
func (s *StorageManager) Count() (int32, Result) {
	if s.manager == nil {
		return 0, ResultInternalError
	}
	var count int32
	dcgo.RunOnDispatcherSync(func() any {
		dcgo.StorageManagerCount(s.manager, unsafe.Pointer(&count))
		return nil
	})
	return count, ResultOk
}

// Stat gets file statistics
func (s *StorageManager) Stat(name string) (*FileStat, Result) {
	if s.manager == nil {
		return nil, ResultInternalError
	}
	var cstat dcgo.DiscordFileStat
	res := dcgo.RunOnDispatcherSync(func() int32 {
		return dcgo.StorageManagerStatGo(s.manager, name, unsafe.Pointer(&cstat))
	})
	if res != 0 {
		return nil, Result(res)
	}
	return convertFileStat(&cstat), ResultOk
}

// StatAt gets file statistics at index
func (s *StorageManager) StatAt(index int32) (*FileStat, Result) {
	if s.manager == nil {
		return nil, ResultInternalError
	}
	var cstat dcgo.DiscordFileStat
	res := dcgo.RunOnDispatcherSync(func() int32 {
		return dcgo.StorageManagerStatAtGo(s.manager, index, unsafe.Pointer(&cstat))
	})
	if res != 0 {
		return nil, Result(res)
	}
	return convertFileStat(&cstat), ResultOk
}

// GetPath gets the storage path
func (s *StorageManager) GetPath() (string, Result) {
	if s.manager == nil {
		return "", ResultInternalError
	}
	var path [4096]byte
	res := dcgo.RunOnDispatcherSync(func() int32 {
		return dcgo.StorageManagerGetPathGo(s.manager, unsafe.Pointer(&path[0]))
	})
	if res != 0 {
		return "", Result(res)
	}
	return string(path[:]), ResultOk
}

// Helper for FileStat conversion
func convertFileStat(cstat *dcgo.DiscordFileStat) *FileStat {
	return &FileStat{
		Filename:     dcgo.GetDiscordFileStatFilename(cstat),
		Size:         dcgo.GetDiscordFileStatSize(cstat),
		LastModified: dcgo.GetDiscordFileStatLastModified(cstat),
	}
}
