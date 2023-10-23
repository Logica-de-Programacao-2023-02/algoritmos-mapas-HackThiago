package main

func main() {
	mapa := ex01("O rato roeu a roupa do rei de roma a")
	for i, i2 := range mapa {
		println(i, " = ", i2)
	}
	println("----------------------------------")

}
