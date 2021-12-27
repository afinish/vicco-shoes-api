package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/heavenmise/vicco-shoes-api/pkg/models"
)

func GeneratePaginationFromRequest(c *gin.Context) models.Pagination {
	limit := 5
	page := 1
	sort := "created_at asc"
	minPrice := 0
	maxPrice := 2500000
	filterCollection := ""
	filterColor := ""
	filterGender := ""
	filterSize := -1

	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "sort":
			sort = queryValue
			break
		case "min_price":
			minPrice, _ = strconv.Atoi(queryValue)
			break
		case "max_price":
			maxPrice, _ = strconv.Atoi(queryValue)
			break
		case "filter_collection":
			filterCollection = queryValue
			break
		case "filter_color":
			filterColor = queryValue
			break
		case "filter_gender":
			filterGender = queryValue
			break
		case "filter_size":
			filterSize, _ = strconv.Atoi(queryValue)
			break
		}
	}
	return models.Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
		MinPrice: minPrice,
		MaxPrice: maxPrice,
		FilterCollection: filterCollection,
		FilterColor: filterColor,
		FilterGender: filterGender,
		FilterSize: filterSize,
	}

}