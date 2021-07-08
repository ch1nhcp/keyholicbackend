package models

type Orderitem struct {
	Id        int     `json:"Id"`
	OrderId   int     `json:"OrderId"`
	Order     Order   `gorm:" ForeignKey:OrderId "`
	ProductId int     `json:"ProductId"`
	Product   Product `gorm:" ForeignKey:ProductId "`
	Quantity  int     `json:"Quantity"`
}
