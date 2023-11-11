package main

import (
	"chethan0707/basic_number_filtering/mathworks"
	"fmt"
)

func main() {

	evenNumbers := mathworks.EvenNumbers([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Println("Even Numbers: ", evenNumbers)

	oddNumbers := mathworks.OddNumbers([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Println("Odd Numbers: ", oddNumbers)

	primeNumbers := mathworks.PrimeNumbers([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Println("Odd Numbers: ", primeNumbers)

	oddPrimeNmbers := mathworks.OddPrimeNumbers([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Println("Odd Prime Numbers: ", oddPrimeNmbers)

	multiplesOfFiveAndEven := mathworks.EvenMultiplesOf5NumberFilter([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20})
	fmt.Println("Multiples of 5 with Even filter: ", multiplesOfFiveAndEven)

	multiplesOf3NumerFilter := mathworks.MultiplesOf3NumerFilter([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20})
	fmt.Println("Multiples of 3 with Odd filter: ", multiplesOf3NumerFilter)

	odd := func(n int) bool { return n%2 != 0 }
	greaterThanN := func(n int) mathworks.Condition { return func(m int) bool { return m > n } }
	greaterThan5 := greaterThanN(5)
	
	multiplesOf := func(n int) mathworks.Condition { return func(m int) bool { return m%n == 0 } }
	multiplesOf3 := multiplesOf(3)
	allFilter := mathworks.Filter([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, []mathworks.Condition{odd, greaterThan5, multiplesOf3}...)
	fmt.Println("All filter: ", allFilter)
}
