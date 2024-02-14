package handler

import (
	"fmt"
	"net/http"

	"github.com/Nivaldeir/golang-opportunities/config"
	"github.com/Nivaldeir/golang-opportunities/schemas"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandle() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}
func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}
	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}
	if err := db.Create(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error creating opening on database")
		logger.Errorf("errr creating opening: %v", err)
		return
	}
	sendSuccess(ctx, "create-opening", opening)
}
func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamsIsRequired("id", "queryParameter").Error())
		return
	}
	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not foun", id))
		return
	}
	sendSuccess(ctx, "show-opening", opening)
}
func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamsIsRequired("id", "queryParameter").Error())
		return
	}
	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusFound, fmt.Sprintf("opening with id: %s not foun", id))
		return
	}
	if err := db.Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusFound, fmt.Sprintf("error deleting opening with id: %s", id))
		return
	}
	sendSuccess(ctx, "delete-opening", opening)
}
func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}
	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		logger.Errorf("validation errr: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamsIsRequired("id", "queryParameter").Error())
		return
	}
	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}
	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}
	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("error updating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating opening")
		return
	}
	sendSuccess(ctx, "update-opening", opening)
}
func ListOpeningHandler(ctx *gin.Context) {
	opening := []schemas.Opening{}
	if err := db.Find(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}
	sendSuccess(ctx, "list-openings", opening)
}
