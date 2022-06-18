package formatutil

import "fmt"

func PrintPrice(price float32) {
		var format string = "%f\n"

		switch {
		case price < 0.0001:
			format = "0.001\n"
			break
		case price < 0.01:
			format = "%.6f\n"
			break
		case price < 1:
			format = "%.4f\n"
			break
		case price < 100:
			format = "%.2f\n"
			break
		case price < 1000:
			format = "%.0f\n"
		}

		fmt.Printf(format, price)
}
