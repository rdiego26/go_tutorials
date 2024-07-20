package main

import "fmt"

func main() {
	var intArr [3]int32
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	newIntArr := [3]int32{1, 2, 3}
	fmt.Println(newIntArr)

	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("\nThe length is %v with capacity %v\n", len(intSlice), cap(intSlice))

	var intSliece2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSliece2...)
	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 3, 10)
	fmt.Println(intSlice3)

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Person1": 10, "Person2": 20}
	fmt.Println(myMap2["Person1"])

	var age, ok = myMap2["Person3"]
	if ok {
		fmt.Println(age)
	} else {
		fmt.Println("Person3 is not in the map")
	}

	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v\n", name, age)
	}

}
