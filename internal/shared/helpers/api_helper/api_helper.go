package api_helper

import (
	"encoding/json"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/api"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/db"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
	"net/http"
	"strconv"
)

func MapAndResponse[TModel any, TResponse any](c *gin.Context, model *TModel, responseModel *TResponse) error {
	if err := copier.Copy(responseModel, model); err != nil {
		return err
	}

	resp := api.Response[TResponse]{
		Data:  *responseModel,
		Error: nil,
	}

	c.JSON(http.StatusOK, resp)
	return nil
}

func ValidateRequest[TRequest any](c *gin.Context, requestData *TRequest) error {
	var validate = validator.New()

	decoder := json.NewDecoder(c.Request.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(requestData); err != nil {
		return err
	}

	if err := validate.Struct(requestData); err != nil {
		return err
	}

	return nil
}

func ValidateAndMapRequest[TRequest any, TModel any](c *gin.Context, requestData *TRequest, model *TModel) error {
	if err := ValidateRequest(c, requestData); err != nil {
		return err
	}

	if err := copier.Copy(model, requestData); err != nil {
		return err
	}

	return nil
}

func MapIgnoreEmpty[T1 any, T2 any](from *T1, to *T2) error {
	return copier.CopyWithOption(to, from,
		copier.Option{
			IgnoreEmpty: true,
		})
}

func GetUniqIdFromRoute(c *gin.Context) (db.UniqId, error) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		return 0, err
	}
	return id, nil
}
