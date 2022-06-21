package utils

import "fmt"

type formatter struct {
}

var Formatter formatter

func (formatter) FormatPrice(price float32) string {
	var format string = "%f\n"

	switch {
	case price <= 0.0001:
		return "0.001\n"
	case price < 0.01:
		format = "%.6f\n"
		break
	case price < 1:
		format = "%.4f\n"
		break
	case price < 100:
		format = "%.2f\n"
		break
	default:
		format = "%.0f\n"
		break
	}

	return fmt.Sprintf(format, price)
}

func (formatter) PrintPrice(price float32) {
	fmt.Print(Formatter.FormatPrice(price))
}
