package main

type MyNumber int

// Utilizando o caractere ~, podemos criar tipos genéricos que podem ser usados em qualquer lugar que eles forem usados
type Number interface {
	~int | float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, valor := range m {
		soma += valor
	}
	return soma
}

func comparar[T Number](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}


// Generics são usados para criar tipos genéricos, que podem ser usados em qualquer lugar que eles forem usados;
func main() {
	m := map[string]int{
		"Michael": 1000,
		"João":    2000,
		"Maria":   3000,
	}
	mFloat := map[string]float64{
		"Michael": 10.5,
		"João":    20.5,
		"Maria":   30.5,
	}
	mMyNumber := map[string]MyNumber{
		"Michael": 1000,
		"João":    2000,
		"Maria":   3000,
	}

	println("Soma:", Soma(m))
	println("Soma Float:", Soma(mFloat))
	println("Soma MyNumber:", Soma(mMyNumber))
	println("Comparando 1:", comparar(10, 10))
	println("Comparando 2:", comparar(10.5, 10.5))
	println("Comparando 3:", comparar(1000, 1000.0))
}
