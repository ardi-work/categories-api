package models

type BestSellingProduct struct {
	Name    string `json:"nama" example:"Indomie Goreng"`
	QtySold int    `json:"qty_terjual" example:"15"`
}

type DailyReport struct {
	TotalRevenue        int                  `json:"total_revenue" example:"50000"`
	TotalTransactions   int                  `json:"total_transaksi" example:"5"`
	BestSellingProducts []BestSellingProduct `json:"produk_terlaris"`
}

type DateRangeReport struct {
	DailyReport
	StartDate string `json:"start_date" example:"2026-01-01"`
	EndDate   string `json:"end_date" example:"2026-02-01"`
}
