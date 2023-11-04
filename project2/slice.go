package project2

import "fmt"

func Slice() {
	var arr [5]int = [5]int{6, 7, 8, 9, 10}
	a := arr[1:4]
	a = append(a, 21)

	fmt.Println(a)
	var names [5]string = [5]string{"john", "luke", "murari", "mike", "ted"}
	fmt.Println(names)
	b := names[0:2]
	c := names[1:3]
	fmt.Println(b, c)

	b[0] = "xyz"
	fmt.Println(b, c)
	fmt.Println(names)

	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4, 5)
	fmt.Println(x)

	f := map[string]float64{
		"q": 77.5,
		"w": 55.6,
		"e": 66.8,
	}
	f["r"] = 87.9
	wvalue := f["w"]
	fmt.Println(wvalue)
	f["e"] = 89.5
	fmt.Println(f)
	delete(f, "e")
	fmt.Println(f)

	for student, grade := range f {
		fmt.Println(student, grade)
	}

	if grade, ok := f["w"]; ok {
		fmt.Println(grade, ok)
	} else {
		fmt.Println("invalid")
	}

	u := make(map[string]int)

	u["x"] = 56
	u["y"] = 91
	u["z"] = 88

	for student, marks := range u {
		fmt.Println(student, marks)
	}

}
