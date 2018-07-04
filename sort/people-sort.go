package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  uint8
}

type People []Person

func (p People) Len() int {
	return len(p)
}

func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p People) Less(i, j int) bool {
	return p[i].Name < p[j].Name
}

func main() {
	tmp := []Person{
		{"Nanoha", 9},
		{"Hayate", 9},
		{"Fate", 9},
	}
	var people People = tmp

	fmt.Println("before sort")
	fmt.Println(people)
	//	スライスの中身をコピーで入れ替え

	//	昇順
	sort.Sort(people)
	fmt.Println("after sort [name asc]")
	fmt.Println(people)
	//	降順
	fmt.Println("after sort [name desc]")
	sort.Sort(sort.Reverse(people))
	fmt.Println(people)
}
