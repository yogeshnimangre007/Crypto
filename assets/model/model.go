package model

import "time"
//asset structure
type Asset struct {
	AssetName     string  `json:"asset_name"`
	AssetType     string  `json:"asset_type"`
	AssetCode     string  `json:"asset_code"`
	Id            int     `json:"id"`
	CurrentPrice  float64 `json:"current_price"`
	PreviousPrice float64 `json:"previous_price"`
	CreationDate  time.Time `json:"creation_date"`
	UpdationDate time.Time `json:"updation_date"`
}
