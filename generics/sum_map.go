package generics

type Number interface {
	int | float64
}

func SumIntegers(integersMap map[string]int) int {
	var total int
	for _, number := range integersMap {
		total += number
	}
	return total
}

func SumFloats(floatsMap map[string]float64) float64 {
	var total float64
	for _, float := range floatsMap {
		total += float
	}
	return total
}

func SumIntegersOrFloats[K comparable, V int | float64](integersOrFloats map[K]V) V {
	var total V
	for _, value := range integersOrFloats {
		total += value
	}
	return total
}

func SumNumbers[K comparable, V Number](numbers map[K]V) V {
	var total V
	for _, number := range numbers {
		total += number
	}
	return total
}
