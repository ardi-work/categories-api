package repositories

import (
	"categories-api/database"
	"categories-api/models"
	"database/sql"
	"time"
)

func GetTodayReport() (*models.DailyReport, error) {
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999999999, time.UTC)

	return GetDateRangeReportInternal(startOfDay, endOfDay)
}

func GetDateRangeReport(startDate, endDate string) (*models.DateRangeReport, error) {

	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, err
	}

	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return nil, err
	}

	endDateOnly := time.Date(end.Year(), end.Month(), end.Day(), 23, 59, 59, 999999999, time.UTC)

	report, err := GetDateRangeReportInternal(start, endDateOnly)
	if err != nil {
		return nil, err
	}

	return &models.DateRangeReport{
		DailyReport: *report,
		StartDate:   startDate,
		EndDate:     endDate,
	}, nil
}

func GetDateRangeReportInternal(startTime, endTime time.Time) (*models.DailyReport, error) {
	db := database.GetDB()

	var totalRevenue sql.NullInt64
	var totalTransactions int

	err := db.QueryRow("SELECT COALESCE(SUM(total_amount), 0) as total_revenue, COUNT(*) as total_transactions FROM transactions WHERE status = 'completed' AND created_at >= $1 AND created_at <= $2", startTime, endTime).Scan(&totalRevenue, &totalTransactions)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(`
		WITH product_sales AS (
			SELECT p.name, SUM(td.quantity) as qty_sold
			FROM transaction_details td
			JOIN transactions t ON td.transaction_id = t.id
			JOIN products p ON td.product_id = p.id
			WHERE t.status = 'completed' AND t.created_at >= $1 AND t.created_at <= $2
			GROUP BY p.id, p.name
		)
		SELECT name, qty_sold 
		FROM product_sales 
		WHERE qty_sold = (SELECT MAX(qty_sold) FROM product_sales)
	`, startTime, endTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bestSellingProducts []models.BestSellingProduct
	for rows.Next() {
		var p models.BestSellingProduct
		err := rows.Scan(&p.Name, &p.QtySold)
		if err != nil {
			continue
		}
		bestSellingProducts = append(bestSellingProducts, p)
	}

	report := &models.DailyReport{
		TotalRevenue:        int(totalRevenue.Int64),
		TotalTransactions:   totalTransactions,
		BestSellingProducts: bestSellingProducts,
	}

	return report, nil
}
