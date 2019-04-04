package main

import ("fmt"
	"sort"
)
type people []string

func (a people) Len() int           { return len(a) }
func (a people) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a people) Less(i, j int) bool { return a[i] < a[j] }

func main() {

	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	sort.Strings(studyGroup)
	fmt.Println(studyGroup)
	sort.Reverse(studyGroup)



	s := []string{"Zeno", "John", "Al", "Jenny"}
	sort.StringSlice(s).Sort()
	fmt.Println(s)
	//sort.Reverse(s)


}

//Use https://godoc.org/sort to sort the following:
//
//(1)
//type people []string
//studyGroup := people{"Zeno", "John", "Al", "Jenny"}
//
//(2)
//s := []string{"Zeno", "John", "Al", "Jenny"}
//
//(3)
//n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
//
//Also sort the above in reverse order