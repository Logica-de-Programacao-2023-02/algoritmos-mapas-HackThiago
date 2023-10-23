package main

func ex02(map1 map[string]int, map2 map[string]int) map[string]int {
	resposta := make(map[string]int)
	for chave, valor := range map1 {
		resposta[chave] = valor
	}
	for chave, valor := range map2 {
		resposta[chave] = valor
	}
	return resposta
}
