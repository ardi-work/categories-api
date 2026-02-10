package models

type BestSellingProduct struct {
	Name    string `json:"nama"`
	QtySold int    `json:"qty_terjual"`
}

type DailyReport struct {
	TotalRevenue        int                  `json:"total_revenue"`
	TotalTransactions   int                  `json:"total_transaksi"`
	BestSellingProducts []BestSellingProduct `json:"produk_terlaris"`
}

type DateRangeReport struct {
	DailyReport
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}
