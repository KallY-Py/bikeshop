package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"vue-go-backend/repository"
	"vue-go-backend/utils"

	"github.com/gorilla/mux"
)

type ReportHandler struct {
	repo *repository.ReportRepository
}

func NewReportHandler(repo *repository.ReportRepository) *ReportHandler {
	return &ReportHandler{repo: repo}
}

func (h *ReportHandler) GetReports(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reports, err := h.repo.GetAll()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, reports)
}

func (h *ReportHandler) UpdateReportStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid report ID")
		return
	}

	var input struct {
		Status string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil || input.Status == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Status is required")
		return
	}

	if err := h.repo.UpdateStatus(id, input.Status); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"message": "Report status updated successfully",
		"id":      id,
		"status":  input.Status,
	})
}
