package arrays

func Soma(numeros []int) (soma int) {
	for _, numero := range numeros {
		soma += numero
	}
	return
}

func SomaTudo(numerosParaSomar ...[]int) (somas []int) {
	for _, numeros := range numerosParaSomar {
		somas = append(somas, Soma(numeros))
	}
	return
}
