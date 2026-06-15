package httpController

import "github.com/gin-gonic/gin"

// Conveniently parses request body to dto by type
func ParseBodyDto[dtoType any](c *gin.Context) (dtoType, error) {
	var dto dtoType
	if err := c.ShouldBindBodyWithJSON(&dto); err != nil {
		var nilAnswer dtoType
		return nilAnswer, err
	}
	return dto, nil
}

type AddUserDto struct {
	Email string `json:"email" binding:"required,email"`
}

type AddProviderDto struct {
	Email string `json:"email" binding:"required,email"`
}

type FirstBuyDto struct {
	ID int `json:"id" binding:"gte=1"`
}

type RepeatBuyDto struct {
	ID int `json:"id" binding:"gte=1"`
}
