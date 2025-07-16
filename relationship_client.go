package discord

import (
	"unsafe"

	"github.com/andresperezl/discordgamesdk-go/core"
)

type RelationshipClient struct {
	manager *core.RelationshipManager
}

func NewRelationshipClient(core *core.Core) *RelationshipClient {
	return &RelationshipClient{manager: core.GetRelationshipManager()}
}

// Filter filters relationships using a callback function
func (c *RelationshipClient) Filter(filterData, filter unsafe.Pointer) {
	c.manager.Filter(filterData, filter)
}

// Count returns the number of relationships
func (c *RelationshipClient) Count() (int32, core.Result) {
	return c.manager.Count()
}

// Get retrieves a relationship by user ID
func (c *RelationshipClient) Get(userID int64) (*core.Relationship, core.Result) {
	return c.manager.Get(userID)
}

// GetAt retrieves a relationship by index
func (c *RelationshipClient) GetAt(index uint32) (*core.Relationship, core.Result) {
	return c.manager.GetAt(index)
}
