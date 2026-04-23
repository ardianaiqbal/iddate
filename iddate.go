package iddate

import (
	"strings"
	"time"
)

var days = [...]string{
	"Minggu",
	"Senin",
	"Selasa",
	"Rabu",
	"Kamis",
	"Jumat",
	"Sabtu",
}

var months = [...]string{
	"Januari",
	"Februari",
	"Maret",
	"April",
	"Mei",
	"Juni",
	"Juli",
	"Agustus",
	"September",
	"Oktober",
	"November",
	"Desember",
}

var shortDays = [...]string{
	"Min",
	"Sen",
	"Sel",
	"Rab",
	"Kam",
	"Jum",
	"Sab",
}

var shortMonths = [...]string{
	"Jan",
	"Feb",
	"Mar",
	"Apr",
	"Mei",
	"Jun",
	"Jul",
	"Agt",
	"Sep",
	"Okt",
	"Nov",
	"Des",
}

// DayName returns the Indonesian name of the day for the given time.
func DayName(t time.Time) string {
	return days[t.Weekday()]
}

// ShortDayName returns the abbreviated Indonesian name of the day for the given time.
func ShortDayName(t time.Time) string {
	return shortDays[t.Weekday()]
}

// MonthName returns the Indonesian name of the month for the given time.
func MonthName(t time.Time) string {
	return months[t.Month()-1]
}

// ShortMonthName returns the abbreviated Indonesian name of the month for the given time.
func ShortMonthName(t time.Time) string {
	return shortMonths[t.Month()-1]
}

// Format formats the time according to the layout, substituting English day and month names with Indonesian ones.
// It uses standard time.Format under the hood and performs string replacement on the result.
func Format(t time.Time, layout string) string {
	formatted := t.Format(layout)

	// Replace full day names
	formatted = strings.Replace(formatted, "Sunday", days[0], -1)
	formatted = strings.Replace(formatted, "Monday", days[1], -1)
	formatted = strings.Replace(formatted, "Tuesday", days[2], -1)
	formatted = strings.Replace(formatted, "Wednesday", days[3], -1)
	formatted = strings.Replace(formatted, "Thursday", days[4], -1)
	formatted = strings.Replace(formatted, "Friday", days[5], -1)
	formatted = strings.Replace(formatted, "Saturday", days[6], -1)

	// Replace full month names
	formatted = strings.Replace(formatted, "January", months[0], -1)
	formatted = strings.Replace(formatted, "February", months[1], -1)
	formatted = strings.Replace(formatted, "March", months[2], -1)
	formatted = strings.Replace(formatted, "April", months[3], -1)
	formatted = strings.Replace(formatted, "May", months[4], -1)
	formatted = strings.Replace(formatted, "June", months[5], -1)
	formatted = strings.Replace(formatted, "July", months[6], -1)
	formatted = strings.Replace(formatted, "August", months[7], -1)
	formatted = strings.Replace(formatted, "September", months[8], -1)
	formatted = strings.Replace(formatted, "October", months[9], -1)
	formatted = strings.Replace(formatted, "November", months[10], -1)
	formatted = strings.Replace(formatted, "December", months[11], -1)

	// Replace short day names
	formatted = strings.Replace(formatted, "Sun", shortDays[0], -1)
	formatted = strings.Replace(formatted, "Mon", shortDays[1], -1)
	formatted = strings.Replace(formatted, "Tue", shortDays[2], -1)
	formatted = strings.Replace(formatted, "Wed", shortDays[3], -1)
	formatted = strings.Replace(formatted, "Thu", shortDays[4], -1)
	formatted = strings.Replace(formatted, "Fri", shortDays[5], -1)
	formatted = strings.Replace(formatted, "Sat", shortDays[6], -1)

	// Replace short month names
	formatted = strings.Replace(formatted, "Jan", shortMonths[0], -1)
	formatted = strings.Replace(formatted, "Feb", shortMonths[1], -1)
	formatted = strings.Replace(formatted, "Mar", shortMonths[2], -1)
	formatted = strings.Replace(formatted, "Apr", shortMonths[3], -1)
	// May is handled in the full name replacement above as it shares the same name
	formatted = strings.Replace(formatted, "Jun", shortMonths[5], -1)
	formatted = strings.Replace(formatted, "Jul", shortMonths[6], -1)
	formatted = strings.Replace(formatted, "Aug", shortMonths[7], -1)
	formatted = strings.Replace(formatted, "Sep", shortMonths[8], -1)
	formatted = strings.Replace(formatted, "Oct", shortMonths[9], -1)
	formatted = strings.Replace(formatted, "Nov", shortMonths[10], -1)
	formatted = strings.Replace(formatted, "Dec", shortMonths[11], -1)

	return formatted
}
