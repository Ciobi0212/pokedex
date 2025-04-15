package utils

import "strings"

func CleanInput(textInput string) []string {
	textInput = strings.ToLower(textInput)
	split := strings.Fields(textInput)

	return split
}

func CalculateCatchChanceTiered(baseExp float64) float64 {
	if baseExp < 60 {
		return 0.85
	} // Very Easy
	if baseExp < 120 {
		return 0.60
	} // Easy
	if baseExp < 200 {
		return 0.35
	} // Medium
	if baseExp < 300 {
		return 0.15
	} // Hard
	return 0.05 // Very Hard
}
