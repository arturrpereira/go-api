package handler

import (
	"net/http"

	"github.com/arturrpereira/go-api/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List opening
// @Description List all job openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} CreateOpeningResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}
