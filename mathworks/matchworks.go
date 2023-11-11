package mathworks

func EvenNumbers(numbers []int) []int {
	var result []int

	for _, number := range numbers {
		if number%2 == 0 {
			result = append(result, number)
		}
	}

	return result
}

func OddNumbers(numbers []int) []int {
	var result []int

	for _, number := range numbers {
		if number%2 != 0 {
			result = append(result, number)
		}
	}
	return result
}

func IsPrimeNume(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func PrimeNumbers(numbers []int) []int {
	var output []int
	for _, val := range numbers {
		if val == 1 {
			continue
		}
		if IsPrimeNume(val) {
			output = append(output, val)
		}
	}
	return output
}
func OddPrimeNumbers(numbers []int) []int {
	var output []int
	for _, val := range numbers {
		if val == 1 {
			continue
		}
		if IsPrimeNume(val) {
			output = append(output, val)
		}
	}
	output = OddNumbers(output)

	return output
}

func isMultipleOfFive(num int) bool {
	return num%5 == 0

}
func EvenMultiplesOf5NumberFilter(numbers []int) []int {
	var output []int
	for _, val := range numbers {
		if isMultipleOfFive(val) {
			output = append(output, val)
		}
	}
	output = EvenNumbers(output)

	return output
}

func MultiplesOf3NumerFilter(numbers []int) []int {
	var output []int
	for _, val := range numbers {
		if val%3 == 0 && val > 10 {
			output = append(output, val)
		}
	}
	output = OddNumbers(output)

	return output
}

type Condition func(n int) bool

func Filter(numbers []int, conditions ...Condition) []int {
	var output []int
	for _, val := range numbers {
		flag := true
		for _, condition := range conditions {
			if !condition(val) {
				flag = false
			}
		}
		if flag {
			output = append(output, val)
		}
	}
	return output
}

func FilterAny(numbers []int, conditions ...Condition) []int {
	var output []int
	for _, val := range numbers {
		flag := false
		for _, condition := range conditions {
			if condition(val) {
				flag = true
				continue
			}
		}
		if flag {
			output = append(output, val)
		}
	}
	return output
}
