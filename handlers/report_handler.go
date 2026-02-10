package handlers

import (
	"encoding/json"
	"net/http"

	"categories-api/repositories"
)

// @Summary		Get today's report
// @Description	Get sales report for today including total revenue, total transactions, and best selling products
// @Tags			reports
// @Accept			json
// @Produce		json
// @Success		200	{object}	models.DailyReport	"Success"
// @Failure		500	{object}	map[string]string	"Internal Server Error"
// @Router			/api/report/hari-ini [get]
func TodayReportHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	report, err := repositories.GetTodayReport()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(report)
}

// @Summary		Get date range report
// @Description	Get sales report for a specific date range including total revenue, total transactions, and best selling products
// @Tags			reports
// @Accept			json
// @Produce		json
// @Param			start_date	query		string					true	"Start date (YYYY-MM-DD)"
// @Param			end_date	query		string					true	"End date (YYYY-MM-DD)"
// @Success		200			{object}	models.DateRangeReport	"Success"
// @Failure		400			{object}	map[string]string		"Bad Request - Invalid date format"
// @Failure		500			{object}	map[string]string		"Internal Server Error"
// @Router			/api/report [get]
func DateRangeReportHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	if startDate == "" || endDate == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "start_date and end_date are required"})
		return
	}

	report, err := repositories.GetDateRangeReport(startDate, endDate)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid date format, use YYYY-MM-DD"})
		return
	}

	json.NewEncoder(w).Encode(report)
}
