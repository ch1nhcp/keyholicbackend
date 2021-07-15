package repository

import (
	"errors"
	"finalbackend/database"
	"finalbackend/models"
	"math"
	"net/http"
	"strconv"
	"strings"
)

func FindPage(writer *http.ResponseWriter, request *http.Request) int {
	params := request.URL.Query()
	if pages, ok := params["page"]; ok {
		page := strings.Split(pages[0], ",")

		// var arrConvert []int
		value, _ := strconv.Atoi(page[0])
		if value <= 1 {
			value = 1
		} else {
			value = value

		}
		return value
	}
	return 1
}

type paginate struct {
	Product  []models.Product
	Total    int
	Page     int
	Lastpage float64
}

func GetAllProduct(page int) paginate {
	var total int
	limit := 2
	offset := (page - 1) * limit
	var product []models.Product
	database.DB.Raw("SELECT * FROM `products` LIMIT ? OFFSET ?", limit, offset).Scan(&product)
	database.DB.Raw("SELECT COUNT(*) FROM `products`").Scan(&total)
	paginate := paginate{
		Product:  product,
		Total:    total,
		Page:     page,
		Lastpage: math.Ceil(float64(total/limit) + 1),
	}
	return paginate
}

type productDetail struct {
	Products models.Product
	Image    []string
}

func FindProductByName(name string) (productDetail, error) {
	var product models.Product
	var imageproduct []string
	productDetail := productDetail{}
	database.DB.Raw("SELECT * FROM `products` WHERE name =?", name).Scan(&product)
	if product.Id > 0 {
		database.DB.Raw("SELECT image FROM `imageproducts` WHERE product_id =?", product.Id).Scan(&imageproduct)
		productDetail.Products = product
		productDetail.Image = imageproduct
		return productDetail, nil
	}
	return productDetail, errors.New("not")

}

func FindProductByCategory(name string, page int) paginate {
	var total int
	limit := 2
	offset := (page - 1) * limit
	var product []models.Product
	var idcategory int
	database.DB.Raw("SELECT id FROM `categories` WHERE name = ?", name).Scan(&idcategory)
	database.DB.Raw("SELECT * FROM `products` WHERE category_id = ? LIMIT ? OFFSET ?", idcategory, limit, offset).Scan(&product)
	database.DB.Raw("SELECT COUNT(*) FROM `products` WHERE category_id = ?", idcategory).Scan(&total)
	paginate := paginate{
		Product:  product,
		Total:    total,
		Page:     page,
		Lastpage: math.Ceil(float64(total/limit) + 1),
	}
	return paginate
}
func FindProductByBrand(name []string, page int) paginate {
	var total int
	limit := 8
	offset := (page - 1) * limit
	var product []models.Product
	var idbrand []int
	database.DB.Raw("SELECT id FROM `brands` WHERE name in (?)", name).Scan(&idbrand)
	database.DB.Raw("SELECT * FROM `products` WHERE brand_id in (?) LIMIT ? OFFSET ?", idbrand, limit, offset).Scan(&product)
	database.DB.Raw("SELECT COUNT(*) FROM `products` WHERE brand_id in (?)", idbrand).Scan(&total)
	paginate := paginate{
		Product:  product,
		Total:    total,
		Page:     page,
		Lastpage: math.Ceil(float64(total/limit) + 1),
	}
	return paginate
}
func CheckProductExist(name string) error {
	var Amount int
	database.DB.Raw("SELECT COUNT(*) FROM `brands` WHERE name =?", name).Scan(&Amount)
	if Amount > 0 {
		return errors.New("name exist")
	}
	return nil
}

func GetProductLatest() []productDetail {
	productlatest := []productDetail{}
	productDetail := productDetail{}
	var product []models.Product
	var imageproduct []string
	database.DB.Raw("SELECT * FROM `products` ORDER BY id DESC LIMIT 8").Scan(&product)
	for _, v := range product {
		if v.Id > 0 {
			database.DB.Raw("SELECT image FROM `imageproducts` WHERE product_id =?", v.Id).Scan(&imageproduct)

		}
		productDetail.Products = v
		productDetail.Image = imageproduct
		productlatest = append(productlatest, productDetail)
	}
	return productlatest
}
func GetProductHot() []productDetail {
	productlatest := []productDetail{}
	productDetail := productDetail{}
	var product []models.Product
	var imageproduct []string
	database.DB.Raw("SELECT * FROM `products` ORDER BY RAND() LIMIT 8").Scan(&product)
	for _, v := range product {
		if v.Id > 0 {
			database.DB.Raw("SELECT image FROM `imageproducts` WHERE product_id =?", v.Id).Scan(&imageproduct)
			productDetail.Products = v
			productDetail.Image = imageproduct
			productlatest = append(productlatest, productDetail)
		}
	}
	return productlatest
}
