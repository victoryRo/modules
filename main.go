package main

import (
	"fmt"
	"github.com/victoryRo/modules/slices"
	quote "rsc.io/quote/v3"
	"strings"
)

func main() {
	list := []string{"Afortunado", "Ganador", "Vencedor", "Ganancia", "Victorioso", "Acertivo", quote.HelloV3()}

	slices.Filter(list, func(item string) bool {
		return strings.HasPrefix(strings.ToLower(item), "g")
	})

	fmt.Println(quote.Concurrency())
}
