package countries

import (
	"fmt"
)

// FormatAddress tries to make sense of the typical input fields given for addresses,
// and format them into a proper address the way it's usually shown in the country.
func FormatAddress(line1, line2, zip, city, state, cc string) string {
	var a string
	switch cc {
	case "NO", "SE":
		a = fmt.Sprintf("%s\n%s %s\n%s", line1, zip, city, Country(cc))
	default:
		a = fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s", line1, line2, zip, city, state, Country(cc))
	}
	return a
}
