package main

import (
	"fmt"
	"myproject/simplecalc"
)

type Pname struct { /////struct
	fname string
	lname string
}

type Patient struct {
	age     int
	disease string
	patient Pname
}

func main() {
	f := Pname{
		fname: "Murari",
		lname: "Datta",
	}
	g := Patient{
		age:     25,
		disease: "fever",
		patient: f,
	}
	fmt.Printf("%s%s has the problem %s\n", g.patient.fname, g.patient.lname, g.disease)
	var arr [5]int = [5]int{6, 7, 8, 9, 10} /// arrays
	for q := 0; q < len(arr); q++ {
		fmt.Println(arr[q])
	}
	day := 5
	switch day { //////Switch Case
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursay")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	default:
		fmt.Println("Invalid")

	}
	for i := 0; i < 10; i++ { //////for loop
		fmt.Println("Murari")
	}
	num := 6
	if num%2 == 0 { ///// if condition
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	//for {
	//fmt.Println("world")   /////infinite for loop
	//}
	a, b := 15.0, 5.0
	defer fmt.Println("Katragadda") //defer
	defer fmt.Println("Datta")      //defer
	fmt.Println("Murari")
	fmt.Println(simplecalc.Add(a, b))
	fmt.Println(simplecalc.Sub(a, b))

	var p *float64 //pointers
	j := 23.5
	p = &j
	*p = 15
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(&p)
}
