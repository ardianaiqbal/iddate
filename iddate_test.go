package iddate

import (
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	// 2006-01-02T15:04:05Z
	// Monday, 02 January 2006
	dt := time.Date(2023, time.August, 17, 10, 0, 0, 0, time.UTC) // Thursday, 17 August 2023

	tests := []struct {
		name     string
		layout   string
		expected string
	}{
		{
			name:     "Full day and month",
			layout:   "Monday, 02 January 2006",
			expected: "Kamis, 17 Agustus 2023",
		},
		{
			name:     "Short day and short month",
			layout:   "Mon, 02 Jan 2006",
			expected: "Kam, 17 Agt 2023",
		},
		{
			name:     "May edge case",
			layout:   "Monday, 02 January 2006",
			expected: "Senin, 01 Mei 2023",
		},
	}

	dtMay := time.Date(2023, time.May, 1, 10, 0, 0, 0, time.UTC) // Monday, 1 May 2023

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			currDt := dt
			if tt.name == "May edge case" {
				currDt = dtMay
			}

			got := Format(currDt, tt.layout)
			if got != tt.expected {
				t.Errorf("Format() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestIndividualFunctions(t *testing.T) {
	dt := time.Date(2023, time.August, 17, 10, 0, 0, 0, time.UTC) // Thursday, 17 August 2023

	if got := DayName(dt); got != "Kamis" {
		t.Errorf("DayName() = %v, want %v", got, "Kamis")
	}

	if got := ShortDayName(dt); got != "Kam" {
		t.Errorf("ShortDayName() = %v, want %v", got, "Kam")
	}

	if got := MonthName(dt); got != "Agustus" {
		t.Errorf("MonthName() = %v, want %v", got, "Agustus")
	}

	if got := ShortMonthName(dt); got != "Agt" {
		t.Errorf("ShortMonthName() = %v, want %v", got, "Agt")
	}
}
