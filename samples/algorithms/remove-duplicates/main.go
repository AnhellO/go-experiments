package main

import (
	"fmt"
)

func main() {
	studentsCities := []string{
		"Mumbai",
		"Delhi",
		"Ahmedabad",
		"Mumbai",
		"Bangalore",
		"Delhi",
		"Kolkata",
		"Pune",
	}

	studentsIDs := []int{11, 33, 22, 1, 33, 12, 25, 11}
	studentsAssessments := []float32{98.3, 89.2, 95.5, 95.5, 82.21, 81.89, 90, 77.6, 90}

	fmt.Println(removeDuplicateStr(studentsCities))
	fmt.Println(removeDuplicateInt(studentsIDs))
	fmt.Println(removeDuplicate(studentsAssessments))

	fmt.Println("***Generic Solution***")

	// Generic solution
	fmt.Println(removeDuplicate(studentsCities))
	fmt.Println(removeDuplicate(studentsIDs))
	fmt.Println(removeDuplicate(studentsAssessments))

}
func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func removeDuplicateInt(intSlice []int) []int {
	allKeys := make(map[int]bool)
	list := []int{}
	for _, item := range intSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func removeDuplicateFloat32(floatSlice []float32) []float32 {
	allKeys := make(map[float32]bool)
	list := []float32{}
	for _, item := range floatSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

// Generic solution
func removeDuplicate[T string | int | float32 | float64](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
