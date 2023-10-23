package main

import "strings"

func ex01(s string) map[string]int {
	aux := strings.Fields(s)
	resposta := make(map[string]int)
	for _, palavra := range aux {
		resposta[palavra]++
	}
	return resposta
}
