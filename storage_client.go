package discord

import (
	"context"
	"fmt"

	core "github.com/andresperezl/discordctl/core"
)

// StorageClient provides Go-like interfaces for storage management
type StorageClient struct {
	manager *core.StorageManager
	core    *core.Core
}

// Read reads data from storage
func (sc *StorageClient) Read(name string) ([]byte, error) {
	if sc.manager == nil {
		return nil, fmt.Errorf("storage manager not available")
	}
	buf := make([]byte, 4096) // Default buffer size
	n, result := sc.manager.Read(name, buf)
	if result != core.ResultOk {
		return nil, fmt.Errorf("failed to read: %v", result)
	}
	return buf[:n], nil
}

// ReadAsync reads data from storage asynchronously
func (sc *StorageClient) ReadAsync(name string) (<-chan []byte, <-chan error) {
	dataChan := make(chan []byte, 1)
	errChan := make(chan error, 1)
	if sc.manager == nil {
		errChan <- fmt.Errorf("storage manager not available")
		close(dataChan)
		close(errChan)
		return dataChan, errChan
	}
	sc.manager.ReadAsync(name, func(result core.Result, data []byte) {
		if result != core.ResultOk {
			errChan <- fmt.Errorf("failed to read async: %v", result)
			close(dataChan)
			close(errChan)
			return
		}
		dataChan <- data
		close(dataChan)
		close(errChan)
	})
	return dataChan, errChan
}

// Write writes data to storage
func (sc *StorageClient) Write(name string, data []byte) error {
	if sc.manager == nil {
		return fmt.Errorf("storage manager not available")
	}
	result := sc.manager.Write(name, data)
	if result != core.ResultOk {
		return fmt.Errorf("failed to write: %v", result)
	}
	return nil
}

// WriteAsync writes data to storage asynchronously
func (sc *StorageClient) WriteAsync(name string, data []byte) <-chan error {
	errChan := make(chan error, 1)
	if sc.manager == nil {
		errChan <- fmt.Errorf("storage manager not available")
		close(errChan)
		return errChan
	}
	sc.manager.WriteAsync(name, data, func(result core.Result) {
		if result != core.ResultOk {
			errChan <- fmt.Errorf("failed to write async: %v", result)
		} else {
			errChan <- nil
		}
		close(errChan)
	})
	return errChan
}

// WriteWithContext writes data to storage, respecting context cancellation and timeout.
//
// Example usage:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
//	defer cancel()
//	err := client.Storage().WriteWithContext(ctx, "file.txt", []byte("hello"))
//	if err != nil {
//	    log.Fatalf("failed to write: %v", err)
//	}
//
// Returns an error if the context is cancelled, deadline exceeded, or the write fails.
func (sc *StorageClient) WriteWithContext(ctx context.Context, name string, data []byte) error {
	if sc.manager == nil {
		return fmt.Errorf("storage manager not available")
	}
	errChan := make(chan error, 1)

	sc.manager.WriteAsync(name, data, func(result core.Result) {
		if result != core.ResultOk {
			errChan <- fmt.Errorf("failed to write async: %v", result)
		} else {
			errChan <- nil
		}
	})

	select {
	case err := <-errChan:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Delete deletes data from storage
func (sc *StorageClient) Delete(name string) error {
	if sc.manager == nil {
		return fmt.Errorf("storage manager not available")
	}
	result := sc.manager.Delete(name)
	if result != core.ResultOk {
		return fmt.Errorf("failed to delete: %v", result)
	}
	return nil
}

// Exists checks if data exists in storage
func (sc *StorageClient) Exists(name string) (bool, error) {
	if sc.manager == nil {
		return false, fmt.Errorf("storage manager not available")
	}
	exists, result := sc.manager.Exists(name)
	if result != core.ResultOk {
		return false, fmt.Errorf("failed to check existence: %v", result)
	}
	return exists, nil
}

// Count returns the number of files in storage
func (sc *StorageClient) Count() (int32, error) {
	if sc.manager == nil {
		return 0, fmt.Errorf("storage manager not available")
	}
	count, result := sc.manager.Count()
	if result != core.ResultOk {
		return 0, fmt.Errorf("failed to count: %v", result)
	}
	return count, nil
}

// Stat gets file statistics
func (sc *StorageClient) Stat(name string) (*core.FileStat, error) {
	if sc.manager == nil {
		return nil, fmt.Errorf("storage manager not available")
	}
	stat, result := sc.manager.Stat(name)
	if result != core.ResultOk {
		return nil, fmt.Errorf("failed to stat: %v", result)
	}
	return stat, nil
}

// GetPath returns the storage path
func (sc *StorageClient) GetPath() (string, error) {
	if sc.manager == nil {
		return "", fmt.Errorf("storage manager not available")
	}
	path, result := sc.manager.GetPath()
	if result != core.ResultOk {
		return "", fmt.Errorf("failed to get path: %v", result)
	}
	return path, nil
}

// ReadWithContext reads data from storage, respecting context cancellation and timeout.
//
// Example usage:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
//	defer cancel()
//	data, err := client.Storage().ReadWithContext(ctx, "file.txt")
//	if err != nil {
//	    log.Fatalf("failed to read: %v", err)
//	}
//	fmt.Printf("Read %d bytes\n", len(data))
//
// Returns the data or error if the context is cancelled, deadline exceeded, or the read fails.
func (sc *StorageClient) ReadWithContext(ctx context.Context, name string) ([]byte, error) {
	if sc.manager == nil {
		return nil, fmt.Errorf("storage manager not available")
	}
	dataChan := make(chan []byte, 1)
	errChan := make(chan error, 1)

	sc.manager.ReadAsync(name, func(result core.Result, data []byte) {
		if result != core.ResultOk {
			errChan <- fmt.Errorf("failed to read async: %v", result)
			return
		}
		dataChan <- data
	})

	select {
	case data := <-dataChan:
		return data, nil
	case err := <-errChan:
		return nil, err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
