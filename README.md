# iddate

`iddate` is a Go package that provides utilities for formatting dates in Indonesian. It allows you to convert standard library `time.Time` objects into Indonesian-formatted date strings, handling translations for day and month names (both full and abbreviated).

## Features

- **Indonesian Translations**: Full and short names for days and months.
- **Easy Formatting**: A `Format` function that works like `time.Format` but replaces English names with Indonesian equivalents.
- **Utility Functions**: Direct access to day and month names via `DayName`, `ShortDayName`, `MonthName`, and `ShortMonthName`.

## Installation

```bash
go get github.com/ardianaiqbal/iddate
```

## Usage

### Simple Formatting

The `Format` function allows you to use standard Go layout strings.

```go
package main

import (
	"fmt"
	"time"
	"github.com/ardianaiqbal/iddate"
)

func main() {
	t := time.Now()
	
	// Standard full date
	fmt.Println(iddate.Format(t, "Monday, 02 January 2006"))
	// Output: Senin, 23 April 2026 (example)

	// Short date
	fmt.Println(iddate.Format(t, "Mon, 02 Jan 06"))
	// Output: Sen, 23 Apr 26 (example)
}
```

### Utility Functions

If you only need specific parts of the date:

```go
package main

import (
	"fmt"
	"time"
	"github.com/ardianaiqbal/iddate"
)

func main() {
	t := time.Now()

	fmt.Println("Day:", iddate.DayName(t))           // e.g., Senin
	fmt.Println("Short Day:", iddate.ShortDayName(t)) // e.g., Sen
	fmt.Println("Month:", iddate.MonthName(t))       // e.g., April
	fmt.Println("Short Month:", iddate.ShortMonthName(t)) // e.g., Apr
}
```

## API Reference

### `func DayName(t time.Time) string`
Returns the full Indonesian name of the day for the given time.

### `func ShortDayName(t time.Time) string`
Returns the abbreviated Indonesian name of the day for the given time.

### `func MonthName(t time.Time) string`
Returns the full Indonesian name of the month for the given time.

### `func ShortMonthName(t time.Time) string`
Returns the abbreviated Indonesian name of the month for the given time.

### `func Format(t time.Time, layout string) string`
Formats the time according to the layout, substituting English day and month names with Indonesian ones.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
