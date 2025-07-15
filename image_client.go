package discord

import (
	"unsafe"

	"github.com/andresperezl/discordctl/core"
)

type ImageClient struct {
	manager *core.ImageManager
}

func NewImageClient(core *core.Core) *ImageClient {
	return &ImageClient{manager: core.GetImageManager()}
}

// Fetch fetches an image asynchronously (callback usage is up to the user)
func (c *ImageClient) Fetch(handle core.ImageHandle, refresh bool, callbackData, callback unsafe.Pointer) {
	c.manager.Fetch(handle, refresh, callbackData, callback)
}

// GetDimensions retrieves the dimensions of an image
func (c *ImageClient) GetDimensions(handle core.ImageHandle) (core.ImageDimensions, core.Result) {
	return c.manager.GetDimensions(handle)
}

// GetData retrieves the raw image data
func (c *ImageClient) GetData(handle core.ImageHandle, data []byte) core.Result {
	return c.manager.GetData(handle, data)
}
