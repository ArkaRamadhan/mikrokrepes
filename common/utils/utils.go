package utils

func getColumn(row []string, index int) string {
	if index >= len(row) {
		return ""
	}
	return row[index]
}

// Helper function to return nil if the string is empty
func getStringOrNil(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}
