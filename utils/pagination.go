package utils

func Paginate[T any](data []T, page, limit int) []T {
	start := (page - 1) * limit
	end := start + limit

	if start > len(data) {
		return []T{}
	}
	if end > len(data) {
		end = len(data)
	}
	return data[start:end]
}
