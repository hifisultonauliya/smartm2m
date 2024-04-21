package models

import "testing"

func TestItemNFTValidation(t *testing.T) {
	// Valid item
	validItem := ItemNFT{
		Name:         "Valid Item Name",
		Rating:       4,
		Category:     "photo",
		Image:        "https://example.com/image.jpg",
		Reputation:   600,
		Price:        100,
		Availability: 10,
	}
	err := validItem.Validate()
	if err != nil {
		t.Errorf("Validation failed for valid item: %v", err)
	}

	// Invalid item name (too short)
	shortNameItem := ItemNFT{
		Name: "Short",
		// Include other required fields...
	}
	err = shortNameItem.Validate()
	if err == nil {
		t.Error("Expected validation error for short item name, but got nil")
	}

	// Invalid item name (contains restricted word)
	invalidNameItem := ItemNFT{
		Name: "Invalid Sex Name",
		// Include other required fields...
	}
	err = invalidNameItem.Validate()
	if err == nil {
		t.Error("Expected validation error for item name containing restricted word, but got nil")
	}

	// Invalid rating (negative)
	negativeRatingItem := ItemNFT{
		Name:   "Negative Rating Item",
		Rating: -1,
		// Include other required fields...
	}
	err = negativeRatingItem.Validate()
	if err == nil {
		t.Error("Expected validation error for negative rating, but got nil")
	}

	// Invalid rating (greater than 5)
	invalidRatingItem := ItemNFT{
		Name:   "Invalid Rating Item",
		Rating: 6,
		// Include other required fields...
	}
	err = invalidRatingItem.Validate()
	if err == nil {
		t.Error("Expected validation error for rating greater than 5, but got nil")
	}

	// Invalid category
	invalidCategoryItem := ItemNFT{
		Name:     "Invalid Category Item",
		Category: "invalid",
		// Include other required fields...
	}
	err = invalidCategoryItem.Validate()
	if err == nil {
		t.Error("Expected validation error for invalid category, but got nil")
	}

	// Invalid image URL
	invalidImageItem := ItemNFT{
		Name:  "Invalid Image Item",
		Image: "invalid",
		// Include other required fields...
	}
	err = invalidImageItem.Validate()
	if err == nil {
		t.Error("Expected validation error for invalid image URL, but got nil")
	}

	// Invalid reputation (greater than 1000)
	invalidReputationItem := ItemNFT{
		Name:       "Invalid Reputation Item",
		Reputation: 1001,
		// Include other required fields...
	}
	err = invalidReputationItem.Validate()
	if err == nil {
		t.Error("Expected validation error for reputation greater than 1000, but got nil")
	}

	// Invalid price (negative)
	negativePriceItem := ItemNFT{
		Name:  "Negative Price Item",
		Price: -1,
		// Include other required fields...
	}
	err = negativePriceItem.Validate()
	if err == nil {
		t.Error("Expected validation error for negative price, but got nil")
	}

	// Invalid availability (negative)
	negativeAvailabilityItem := ItemNFT{
		Name:         "Negative Availability Item",
		Availability: -1,
		// Include other required fields...
	}
	err = negativeAvailabilityItem.Validate()
	if err == nil {
		t.Error("Expected validation error for negative availability, but got nil")
	}

	// Invalid Color Reputation
	// ....
}
