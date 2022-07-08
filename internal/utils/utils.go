package utils

import (
	"fmt"
	"math"
)

type formatter struct {
}

var Formatter formatter

func (formatter) FormatPrice(price float32) string {
	var format string = "%f\n"
  absVal := math.Abs(float64(price));

	switch {
	case absVal <= 0.0001:
		return "0.001\n"
	case absVal < 0.01:
		format = "%.6f\n"
		break
	case absVal < 1:
		format = "%.4f\n"
		break
	case absVal < 100:
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
