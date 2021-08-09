package matematica

import (
	"fmt"
	"strconv"
)

func Media(numeros ...float64) float64 {
	total := 0.0
	for _, v := range numeros {
		total += v
	}
	media := total / float64(len(numeros))
	mediaArredondada, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)
	return mediaArredondada
}
