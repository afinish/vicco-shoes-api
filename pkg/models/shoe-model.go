package models

import (
	// "fmt"

	// "github.com/heavenmise/vicco-shoes-api/pkg/config"
	//"github.com/gin-gonic/gin"

	"github.com/heavenmise/vicco-shoes-api/pkg/config"
	"github.com/jinzhu/gorm"
)

// var DB *gorm.DB

type Shoe struct {
	gorm.Model
	Name		string	`json:"name"`
	Article 	string	`json:"article"`
	Color 		string	`json:"color"`
	Size 		string	`json:"size"`
	Gender 		string	`json:"gender"`
	Material 	string	`json:"material"`
	Collection 	string	`json:"collection"`
	Categories 	string	`json:"categories"`
	Price		string 	`json:"price"`
}
type Pagination struct {
	Limit 		int 	`json:"limit"`
	Page		int 	`json:"page"`
	Sort		string 	`json:"sort"`
	
	MinPrice	int		`json:"minPrice"`
	MaxPrice	int		`json:"maxPrice"`

	FilterCollection	string	`json:"filterCollection"`
	FilterColor			string	`json:"filterColor"`
	FilterGender		string	`json:"filterGender"`
	FilterSize			int		`json:"filterSize"`
	
}

func GetShoes(shoe *Shoe, pagination *Pagination) (*[]Shoe, error) {
	var Shoes []Shoe
	offset := (pagination.Page - 1) * pagination.Limit
	// query := fmt.Sprintf("(price BETWEEN %s ? AND %s ?) AND ", pagination.MinPrice, pagination.MaxPrice)

	var queryBuilder *gorm.DB
	// query := fmt.Sprintf("SELECT * FROM shoes WHERE price BETWEEN %d AND %d", pagination.MinPrice, pagination.MaxPrice)

	queryBuilder = config.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)

	if pagination.FilterCollection != "" {
		queryBuilder = config.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort).Where("(price BETWEEN ? AND ?) AND collection=?", pagination.MinPrice, pagination.MaxPrice, pagination.FilterCollection)
		// query = fmt.Sprintf("%s AND collection=\"%s\"", query, pagination.FilterCollection)
	} 
	if pagination.FilterColor != "" {
		queryBuilder = config.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort).Where("(price BETWEEN ? AND ?) AND color=?", pagination.MinPrice, pagination.MaxPrice, pagination.FilterColor)
		// query = fmt.Sprintf("%s AND color=\"%s\"", query, pagination.FilterColor)
	}
	if pagination.FilterGender != "" {
		queryBuilder = config.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort).Where("(price BETWEEN ? AND ?) AND gender=?", pagination.MinPrice, pagination.MaxPrice, pagination.FilterGender)
		// query = fmt.Sprintf("%s AND gender=\"%s\"", query, pagination.FilterGender)
	}
	if pagination.FilterSize != -1 {
		queryBuilder = config.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort).Where("(price BETWEEN ? AND ?) AND size=?", pagination.MinPrice, pagination.MaxPrice, pagination.FilterSize)
		// query = fmt.Sprintf("%s AND size=%d", query, pagination.FilterSize)
	}
	
	//query = fmt.Sprintf("%s LIMIT %d OFFSET %d;", query, pagination.Limit, offset)

	// fmt.Println(query)

	//queryBuilder := config.DB.Raw(query)
	result := queryBuilder.Model(&Shoe{}).Where(Shoes).Find(&Shoes)

	if result.Error != nil {
		// fmt.Printf("error!!!!")
		msg := result.Error
		return nil, msg
	}
	return &Shoes, nil
}


// NOTE: I tried to give the user the ability to choose multiple filters at once, but could not figure it out.
// The commented out code is for that. I left them so that I can change it later.
// But for now I think the program is descent enough and I have to get some sleep...
func GetShoesCategory(shoe *Shoe, pagination *Pagination, category string) (*[]Shoe, error) {
	var Shoes []Shoe
	offset := (pagination.Page - 1) * pagination.Limit
	// query := fmt.Sprintf("(price BETWEEN %s ? AND %s ?) AND ", pagination.MinPrice, pagination.MaxPrice)
	var queryBuilder *gorm.DB
	//ch := "0000"
	// query := fmt.Sprintf("SELECT * FROM shoes WHERE price BETWEEN %d AND %d", pagination.MinPrice, pagination.MaxPrice)
	queryBuilder = config.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort).Where("categories=?", category)
	if pagination.FilterCollection != "" {
		queryBuilder = config.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort).Where("(price BETWEEN ? AND ?) AND collection=? AND categories=?", pagination.MinPrice, pagination.MaxPrice, pagination.FilterCollection, category)
		// query = fmt.Sprintf("%s AND collection=\"%s\"", query, pagination.FilterCollection)
	} 
	if pagination.FilterColor != "" {
		queryBuilder = config.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort).Where("(price BETWEEN ? AND ?) AND color=? AND categories=?", pagination.MinPrice, pagination.MaxPrice, pagination.FilterColor, category)
		// query = fmt.Sprintf("%s AND color=\"%s\"", query, pagination.FilterColor)
	}
	if pagination.FilterGender != "" {
		queryBuilder = config.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort).Where("(price BETWEEN ? AND ?) AND gender=? AND categories=?", pagination.MinPrice, pagination.MaxPrice, pagination.FilterGender, category)
		// query = fmt.Sprintf("%s AND gender=\"%s\"", query, pagination.FilterGender)
	}
	if pagination.FilterSize != -1 {
		queryBuilder = config.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort).Where("(price BETWEEN ? AND ?) AND size=? AND categories=?", pagination.MinPrice, pagination.MaxPrice, pagination.FilterSize, category)
		// query = fmt.Sprintf("%s AND size=%d", query, pagination.FilterSize)
	}
	//query = fmt.Sprintf("%s LIMIT %d OFFSET %d;", query, pagination.Limit, offset)
	//queryBuilder := config.DB.Raw(query)
	// fmt.Println(query)
	result := queryBuilder.Model(&Shoe{}).Where(Shoes).Find(&Shoes)

	if result.Error != nil {
		// fmt.Printf("error!!!!")
		msg := result.Error
		return nil, msg
	}
	return &Shoes, nil
}

func GetShoe(name string) (Shoe, error) {
	var shoe Shoe
	// fmt.Println(name)
	result := config.DB.Model(&Shoe{}).Where("name=?", name).Find(&shoe)
	if result.Error != nil {
		msg := result.Error
		return Shoe{}, msg
	}
	return shoe, nil
}

func (shoe *Shoe) CreateShoe() *Shoe {
	config.DB.NewRecord(shoe)
	config.DB.Create(&shoe)
	return shoe
}

func DeleteShoe(name string) (Shoe, error) {
	var shoe Shoe
	result := config.DB.Model(&Shoe{}).Where("name=?", name).Delete(&shoe)
	if result.Error != nil {
		msg := result.Error
		return Shoe{}, msg
	}
	return shoe, nil
}