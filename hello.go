package main

import ( 
	"fmt"
	"errors"
	"math"
)

//structs -- outside functions
type person struct {
	name string
	age int
}

func main()  {
	//define struct properties
	p := person{name: "Allen", age: 27}
	fmt.Println(p);

	//dot notation
	fmt.Println(p.age);

	//defining variables := inferred type
	//var x int = 5 === x:= 5
	x := 5
	y := 7

	sum := x + y
	a := []int{5,4,3,2,1}
	a[2] = 7
	a = append(a,13)
	fmt.Println(a)
	fmt.Println(sum)

	vertices := make(map[string]int)

	vertices["triangle"] = 2
	vertices["square"] = 3
	vertices["dodcagon"] = 12

	fmt.Println(vertices)
	delete(vertices, "square")
	fmt.Println(vertices)

	//for is only loop in Go
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}	

	fmt.Println()

	//while loop
	i := 0
	
	for i < 5 {
		fmt.Println(i)
		i++
	}

	fmt.Println()
	//access index and value of an array using range
	arr := []string{"a", "b", "c"}

	for index, value := range arr {
		fmt.Println("index", index, "value", value)
	}
	fmt.Println()

	//range on map gives key instead of index
	m := make(map [string]string)
	m["a"] = "alpha"
	m["b"] = "beta"

	for key, value := range m {
		fmt.Println("key:", key, "value:", value)
	}

	result := add(2, 3)
	fmt.Println(result)

	fmt.Println()

	endResult, err := sqrt(-16)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(endResult)
	}
	
	ii := 7
	//& accesses memory location of variable
	inc(&ii) 
	fmt.Println(ii)
	
}

//pointers
func inc(x *int)  {
	*x++
}

func add(x int, y int) int {
return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}

	return math.Sqrt(x), nil
}