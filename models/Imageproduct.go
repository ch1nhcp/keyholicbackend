package models

type Imageproduct struct {
	Id        int     `json:"Id"`
	ProductId int     `json:"ProductId"`
	Product   Product `gorm:" ForeignKey:ProductId "`
	Image     string  `json:"Image"`
}
