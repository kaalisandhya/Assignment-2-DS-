1.package main

import "fmt"

func main() {

    m := make(map[string]int)

    m["k1"] = 7
    m["k2"] = 13

    fmt.Println("map:", m)

    v1 := m["k1"]
    fmt.Println("v1: ", v1)

    fmt.Println("len:", len(m))

    delete(m, "k2")
    fmt.Println("map:", m)

    _, prs := m["k2"]
    fmt.Println("prs:", prs)

    n := map[string]int{"too": 1, "bar": 2}
    fmt.Println("map:", n)
}
map: map[k1:7 k2:13]
v1:  7
len: 2
map: map[k1:7]
prs: false
map: map[bar:2 too:1]
2.package main 

import "fmt"

func main() { 

	// Creating a map 
	// Using make() function 
	var My_map = make(map[float64]string) 
	fmt.Println(My_map) 

	// As we already know that make() function 
	// always returns a map which is initialized 
	// So, we can add values in it 
	My_map[1.3] = "sandhya"
	My_map[1.5] = "kaali"
	fmt.Println(My_map) 
} 
map[]
map[1.3:sandhya 1.5:kaali]

Program exited.
3.package main

import "fmt"

func main() {
	var intSlice = []int{111, 222, 333, 444}
	var strSlice = []string{"rajampet", "kodur", "kadapa"}

	fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(intSlice), cap(intSlice))
	fmt.Printf("strSlice \tLen: %v \tCap: %v\n", len(strSlice), cap(strSlice))
}
intSlice 	Len: 4 	Cap: 4
strSlice 	Len: 3 	Cap: 3

Program exited.