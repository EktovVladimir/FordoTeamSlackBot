package api_helper

import (
	"encoding/json"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/api"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/types"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/jinzhu/copier"
	"io"
	"net/http"
	"strconv"
)

func MapAndResponse[TModel any, TResponse any](w http.ResponseWriter, model *TModel, responseModel *TResponse) error {
	if err := copier.Copy(responseModel, model); err != nil {
		return err
	}

	resp := api.Response[TResponse]{
		Data:  *responseModel,
		Error: nil,
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		return err
	}

	return nil
}

func ValidateRequest[TRequest any](body io.Reader, requestData *TRequest) error {
	var validate = validator.New()

	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(requestData); err != nil {
		return err
	}

	if err := validate.Struct(requestData); err != nil {
		return err
	}

	return nil
}

func ValidateAndMapRequest[TRequest any, TModel any](body io.Reader, requestData *TRequest, model *TModel) error {
	if err := ValidateRequest(body, requestData); err != nil {
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

func GetUniqIdFromRoute(r *http.Request) (types.UniqId, error) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		return 0, err
	}
	return id, nil
}
