package models

type Product struct {
	Id               int        `json:"Id"`
	Name             string     `json:"Name"`
	Image            string     `json:"Image"`
	ShortDescription string     `json:"ShortDescription"`
	Description      string     `json:"Description"`
	Price            int        `json:"Price"`
	SalePrice        int        `json:"SalePrice"`
	BrandId          int        `json:"BrandId"`
	Brand            Brand      `gorm:" ForeignKey:BrandId "`
	CategoryId       int        `json:"CategoryId"`
	Categories       Category `gorm:" ForeignKey:CategoryId "`
}
