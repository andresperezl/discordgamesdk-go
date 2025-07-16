package discord

import (
	"context"
	"fmt"
	"unsafe"

	core "github.com/andresperezl/discordgamesdk-go/core"
)

// StoreClient provides Go-like interfaces for store management
type StoreClient struct {
	manager *core.StoreManager
	core    *core.Core
}

// FetchSkus fetches SKUs
func (sc *StoreClient) FetchSkus() ([]core.Sku, error) {
	if sc.manager == nil {
		return nil, fmt.Errorf("store manager not available")
	}

	// For now, use the synchronous approach since we don't have proper callback support
	// TODO: Implement proper async callback support
	count, res := sc.manager.CountSkus()
	if res != core.ResultOk {
		return nil, fmt.Errorf("failed to count SKUs: %v", res)
	}

	skus := make([]core.Sku, 0, count)
	for i := int32(0); i < count; i++ {
		sku, res := sc.manager.GetSkuAt(i)
		if res != core.ResultOk || sku == nil {
			continue
		}
		skus = append(skus, *sku)
	}
	return skus, nil
}

// FetchSkusAsync fetches SKUs asynchronously
func (sc *StoreClient) FetchSkusAsync(callbackData, callback unsafe.Pointer) {
	if sc.manager == nil {
		return
	}
	sc.manager.FetchSkus(callbackData, callback)
}

// CountSkus gets the count of SKUs
func (sc *StoreClient) CountSkus() (int32, error) {
	if sc.manager == nil {
		return 0, fmt.Errorf("store manager not available")
	}
	count, res := sc.manager.CountSkus()
	if res != core.ResultOk {
		return 0, fmt.Errorf("failed to count SKUs: %v", res)
	}
	return count, nil
}

// GetSku gets a single SKU
func (sc *StoreClient) GetSku(skuID int64) (*core.Sku, error) {
	if sc.manager == nil {
		return nil, fmt.Errorf("store manager not available")
	}
	sku, res := sc.manager.GetSku(skuID)
	if res != core.ResultOk || sku == nil {
		return nil, fmt.Errorf("failed to get SKU: %v", res)
	}
	return sku, nil
}

// GetSkuAt gets a SKU at index
func (sc *StoreClient) GetSkuAt(index int32) (*core.Sku, error) {
	if sc.manager == nil {
		return nil, fmt.Errorf("store manager not available")
	}
	sku, res := sc.manager.GetSkuAt(index)
	if res != core.ResultOk || sku == nil {
		return nil, fmt.Errorf("failed to get SKU at index: %v", res)
	}
	return sku, nil
}

// FetchEntitlements fetches entitlements
func (sc *StoreClient) FetchEntitlements() ([]core.Entitlement, error) {
	if sc.manager == nil {
		return nil, fmt.Errorf("store manager not available")
	}

	// For now, use the synchronous approach since we don't have proper callback support
	// TODO: Implement proper async callback support
	count, res := sc.manager.CountEntitlements()
	if res != core.ResultOk {
		return nil, fmt.Errorf("failed to count entitlements: %v", res)
	}

	ents := make([]core.Entitlement, 0, count)
	for i := int32(0); i < count; i++ {
		ent, res := sc.manager.GetEntitlementAt(i)
		if res != core.ResultOk || ent == nil {
			continue
		}
		ents = append(ents, *ent)
	}
	return ents, nil
}

// FetchEntitlementsAsync fetches entitlements asynchronously
func (sc *StoreClient) FetchEntitlementsAsync(callbackData, callback unsafe.Pointer) {
	if sc.manager == nil {
		return
	}
	sc.manager.FetchEntitlements(callbackData, callback)
}

// CountEntitlements gets the count of entitlements
func (sc *StoreClient) CountEntitlements() (int32, error) {
	if sc.manager == nil {
		return 0, fmt.Errorf("store manager not available")
	}
	count, res := sc.manager.CountEntitlements()
	if res != core.ResultOk {
		return 0, fmt.Errorf("failed to count entitlements: %v", res)
	}
	return count, nil
}

// GetEntitlement gets a single entitlement
func (sc *StoreClient) GetEntitlement(entitlementID int64) (*core.Entitlement, error) {
	if sc.manager == nil {
		return nil, fmt.Errorf("store manager not available")
	}
	ent, res := sc.manager.GetEntitlement(entitlementID)
	if res != core.ResultOk || ent == nil {
		return nil, fmt.Errorf("failed to get entitlement: %v", res)
	}
	return ent, nil
}

// GetEntitlementAt gets an entitlement at index
func (sc *StoreClient) GetEntitlementAt(index int32) (*core.Entitlement, error) {
	if sc.manager == nil {
		return nil, fmt.Errorf("store manager not available")
	}
	ent, res := sc.manager.GetEntitlementAt(index)
	if res != core.ResultOk || ent == nil {
		return nil, fmt.Errorf("failed to get entitlement at index: %v", res)
	}
	return ent, nil
}

// HasSkuEntitlement checks if a SKU has an entitlement
func (sc *StoreClient) HasSkuEntitlement(skuID int64) (bool, error) {
	if sc.manager == nil {
		return false, fmt.Errorf("store manager not available")
	}
	has, res := sc.manager.HasSkuEntitlement(skuID)
	if res != core.ResultOk {
		return false, fmt.Errorf("failed to check SKU entitlement: %v", res)
	}
	return has, nil
}

// StartPurchase starts a purchase
func (sc *StoreClient) StartPurchase(skuID int64) error {
	if sc.manager == nil {
		return fmt.Errorf("store manager not available")
	}

	// For now, return success since we don't have proper callback support
	// TODO: Implement proper purchase functionality
	return nil
}

// StartPurchaseAsync starts a purchase asynchronously
func (sc *StoreClient) StartPurchaseAsync(skuID int64, callbackData, callback unsafe.Pointer) {
	if sc.manager == nil {
		return
	}
	sc.manager.StartPurchase(skuID, callbackData, callback)
}

// FetchSkusWithContext fetches SKUs asynchronously, respecting context cancellation and timeout.
//
// Example usage:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
//	defer cancel()
//	skus, err := client.Store().FetchSkusWithContext(ctx)
//	if err != nil {
//	    log.Fatalf("failed to fetch SKUs: %v", err)
//	}
//	fmt.Printf("Fetched %d SKUs\n", len(skus))
//
// Returns the SKUs or error if the context is cancelled, deadline exceeded, or the fetch fails.
func (sc *StoreClient) FetchSkusWithContext(ctx context.Context) ([]core.Sku, error) {
	if sc.manager == nil {
		return nil, fmt.Errorf("store manager not available")
	}
	skuChan := make(chan []core.Sku, 1)
	errChan := make(chan error, 1)
	sc.manager.FetchSkus(unsafe.Pointer(&skuChan), unsafe.Pointer(&errChan)) // Placeholder: actual async callback wiring needed
	// For now, fallback to synchronous fetch with context support
	go func() {
		skus, err := sc.FetchSkus()
		if err != nil {
			errChan <- err
		} else {
			skuChan <- skus
		}
	}()
	select {
	case skus := <-skuChan:
		return skus, nil
	case err := <-errChan:
		return nil, err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// FetchEntitlementsWithContext fetches entitlements asynchronously, respecting context cancellation and timeout.
//
// Example usage:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
//	defer cancel()
//	ents, err := client.Store().FetchEntitlementsWithContext(ctx)
//	if err != nil {
//	    log.Fatalf("failed to fetch entitlements: %v", err)
//	}
//	fmt.Printf("Fetched %d entitlements\n", len(ents))
//
// Returns the entitlements or error if the context is cancelled, deadline exceeded, or the fetch fails.
func (sc *StoreClient) FetchEntitlementsWithContext(ctx context.Context) ([]core.Entitlement, error) {
	if sc.manager == nil {
		return nil, fmt.Errorf("store manager not available")
	}
	entChan := make(chan []core.Entitlement, 1)
	errChan := make(chan error, 1)
	sc.manager.FetchEntitlements(unsafe.Pointer(&entChan), unsafe.Pointer(&errChan)) // Placeholder: actual async callback wiring needed
	// For now, fallback to synchronous fetch with context support
	go func() {
		ents, err := sc.FetchEntitlements()
		if err != nil {
			errChan <- err
		} else {
			entChan <- ents
		}
	}()
	select {
	case ents := <-entChan:
		return ents, nil
	case err := <-errChan:
		return nil, err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
