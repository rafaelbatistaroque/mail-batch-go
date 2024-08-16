package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/usecase/searchcampaign"
)

func (h *campaignhandler) SearchCampaign(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var input searchcampaign.Input

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return nil, http.StatusBadRequest, err
	}

	output, err := h.searchcampaign.Execute(&input)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return output, http.StatusOK, nil
}
