package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand/v2"
	"strings"
)

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
	var numeros_positivos []int32
	var numeros_negativos []int32
	var contagem_zeros int = 0
	numero_max_elementos := len(arr)

	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			numeros_positivos = append(numeros_positivos, arr[i])
		} else if arr[i] < 0 {
			numeros_negativos = append(numeros_negativos, arr[i])
		} else {
			contagem_zeros++
		}
	}

	max_numeros_positivos := float32(len(numeros_positivos))
	max_numeros_negativos := float32(len(numeros_negativos))
	max_zeros := float32(contagem_zeros)
	max_elementos := float32(numero_max_elementos)

	divisao_positivos := max_numeros_positivos / max_elementos
	divisao_negativos := max_numeros_negativos / max_elementos
	divisao_zeros := max_zeros / max_elementos

	fmt.Printf("%.6f\n", divisao_positivos)
	fmt.Printf("%.6f\n", divisao_negativos)
	fmt.Printf("%.6f\n", divisao_zeros)

}

func main() {
	//reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	//nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	//checkError(err)
	//n := int32(nTemp)

	//arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < 10; i++ {

		numero := int32(rand.IntN(50))

		if i%2 == 0 {
			numero = -numero
			arr = append(arr, numero)
		}

		arr = append(arr, numero)
	}

	plusMinus(arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
