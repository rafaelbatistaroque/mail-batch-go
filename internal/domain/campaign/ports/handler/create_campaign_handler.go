package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/usecase/createcampaign"
)

func (h *campaignhandler) CreateCampaign(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var input createcampaign.Input

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return nil, http.StatusBadRequest, err
	}

	output, err := h.createcampaign.Execute(&input)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return output.Id, http.StatusCreated, nil
}
