package main

import (
	"math/rand"
)

func diagonalDifference(arr [][]int32) int32 {
	var elementos_diagnoal_1 []int32
	var elementos_diagnoal_2 []int32
	i := 0
	for j := len(arr) - 1; j >= 0; j-- {
		elementos_diagnoal_2 = append(elementos_diagnoal_2, arr[j][i])
		i++
	}

	i = 0
	for j := 0; j < len(arr); j++ {
		if i == j {
			elementos_diagnoal_1 = append(elementos_diagnoal_1, arr[i][i])
			i++
		}
	}

	var soma_diagonal_1 int32 = 0
	var soma_diagnoal_2 int32 = 0

	for i := 0; i < len(elementos_diagnoal_1); i++ {
		soma_diagonal_1 += elementos_diagnoal_1[i]
		soma_diagnoal_2 += elementos_diagnoal_2[i]
	}

	resultado := soma_diagonal_1 - soma_diagnoal_2

	if resultado < 0 {
		resultado = -resultado
	}

	return resultado

}

func main() {
	var array [][]int //array bidimensional

	for i := 0; i < 2; i++ {
		linha := []int{} // cria uma nova linha

		for j := 0; j < 2; j++ {
			linha = append(linha, rand.Intn(10)) // adiciona à linha
		}

		array = append(array, linha) // adiciona a linha completa ao array
	}

}
