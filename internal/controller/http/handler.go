package httpController

import "github.com/gin-gonic/gin"

func (h *handler) AddUser(ctx *gin.Context) {
	dto, err := ParseBodyDto[AddUserDto](ctx)
	if err != nil {
		ctx.JSON(
			400,
			gin.H{"error": "DTO validation error"},
		)
		return
	}

	ID, err := h.service.AddUser(ctx.Request.Context(), dto.Email)
	if err != nil {
		ctx.JSON(
			500,
			gin.H{"error": "Internal Server Error"},
		)
		return
	}

	ctx.JSON(
		201,
		gin.H{"id": ID},
	)
}

func (h *handler) AddProvider(ctx *gin.Context) {
	dto, err := ParseBodyDto[AddProviderDto](ctx)
	if err != nil {
		ctx.JSON(
			400,
			gin.H{"error": "DTO validation error"},
		)
		return
	}

	ID, err := h.service.AddProvider(ctx.Request.Context(), dto.Email)
	if err != nil {
		ctx.JSON(
			500,
			gin.H{"error": "Internal Server Error"},
		)
		return
	}

	ctx.JSON(
		201,
		gin.H{"id": ID},
	)
}

func (h *handler) FirstBuy(ctx *gin.Context) {
	dto, err := ParseBodyDto[FirstBuyDto](ctx)
	if err != nil {
		ctx.JSON(
			400,
			gin.H{"error": "DTO validation error"},
		)
		return
	}

	if err := h.service.FirstBuy(ctx.Request.Context(), dto.ID); err != nil {
		ctx.JSON(
			500,
			gin.H{"error": "Internal Server Error"},
		)
		return
	}

	ctx.Status(200)
}

func (h *handler) RepeatBuy(ctx *gin.Context) {
	dto, err := ParseBodyDto[RepeatBuyDto](ctx)
	if err != nil {
		ctx.JSON(
			400,
			gin.H{"error": "DTO validation error"},
		)
		return
	}

	if err := h.service.RepeatBuy(ctx.Request.Context(), dto.ID); err != nil {
		ctx.JSON(
			500,
			gin.H{"error": "Internal Server Error"},
		)
		return
	}

	ctx.Status(200)
}
