package models

import (
	"strings"
	"errors"
	"fmt"
	"encoding/csv"
	"strconv"
)


// ShopSite represents an shop site instance.
type ShopSite struct {
	InstanceID string
	SiteTitle string
	SiteDescription string
	Email string
	LogoURL string
	BannerURL string
	Items []ShopItem
}

// ShopItem represents an item in the shop.
type ShopItem struct {
	ID string
	Name string
	Description string
	Price float64
	ImageURL string
}

// GetShopItemsFromCSV marshalls the shop items listed in a CSV
// file into an array of ShopItem.
func GetShopItemsFromCSV(csvText string) ([]ShopItem, error) {
	var items []ShopItem
	r := csv.NewReader(strings.NewReader(csvText))
	recs, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	for i, rec := range recs {
		if i == 0 {
			if len(rec) != 5 || rec[0] != "ID" || rec[1] != "Name" ||
				rec[2] != "Description" || rec[3] != "Price" ||
				rec[4] != "Image URL" {
					return nil, errors.New(fmt.Sprintf("Invalid spreadsheet schema: %s", rec))
			}
			continue
		}

		if len(rec) < 4 {
			continue // Not enough data to make shop item
		}

		var imgURL string
		if len(rec) < 5 {
			imgURL = "https://example.com/placeholder.png"
		} else {
			imgURL = rec[4]
		}

		var price float64
		price, err = strconv.ParseFloat(rec[3], 64)
		if err != nil {
			return nil, err
		}

		items = append(items, ShopItem{
			ID: rec[0],
			Name: rec[1],
			Description: rec[2],
			Price: price,
			ImageURL: imgURL,
		})
	}

	return items, nil
}
