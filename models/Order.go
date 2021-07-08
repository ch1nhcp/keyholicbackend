package models

type Order struct {
	Id            int    `json:"Id"`
	UserId        int    `json:"UserId"`
	User          User   `gorm:" ForeignKey:UserId "`
	Name          string `json:"Name"`
	Phone         int    `json:"Phone"`
	Address       string `json:"UsAddresserId"`
	TotalProducts int    `json:"TotalProducts"`
	Price         int    `json:"UsPriceerId"`
	Sale          int    `json:"Sale"`
	Ship          int    `json:"Ship"`
	TotalPrice    int    `json:"TotalPrice"`
}
