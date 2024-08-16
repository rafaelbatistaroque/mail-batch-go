package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/rafaelbatistaroque/mail-batch-go/internal/pkg/helpers/validation"
)

func (h *campaignhandler) GetCampaignById(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := ParamURLGetId(r)
	if len(id) == 0 || !validation.IsAlphanumeric(id) {
		return nil, http.StatusBadRequest, errors.New("parameter id is invalid")
	}

	output, err := h.getcampaignbyid.Execute(id)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return output, http.StatusOK, nil
}

func ParamURLGetId(r *http.Request) string {
	urlParts := strings.Split(r.URL.Path, "/")

	return urlParts[len(urlParts)-1]
}
