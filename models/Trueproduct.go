package models

type Trueproduct struct {
	Id            int           `json:"Id"`
	ProductId     int           `json:"ProductId"`
	Product       Product       `gorm:" ForeignKey:ProductId "`
	DetailId      int           `json:"DetailId"`
	Detailproduct Detailproduct `gorm:" ForeignKey:DetailId "`
	Quantity      int           `json:"Quantity"`
}
