package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
)

func main() {

	students := map[int]Student{
		1: Student{
			Name: "Iman Tumorang",
			Age:  17,
		},
		2: Student{
			Name: "Christin Tumorang",
			Age:  18,
		},
		3: Student{
			Name: "Idola Manurung",
			Age:  18,
		},
	}

	sample := &ClassRoom{
		Name:     "D3TI2-Classmate",
		Students: students,
	}

	jbyt, err := json.Marshal(sample)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jbyt))

}

type ClassRoom struct {
	Name     string          `json:"name"`
	Students map[int]Student `json:"students"`
}
type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (s *ClassRoom) MarshalJSON() ([]byte, error) {
	arrStudent := []Student{}
	arrKey := []int{}
	for k, _ := range s.Students {
		arrKey = append(arrKey, k)
	}
	sort.Ints(arrKey)

	for _, pos := range arrKey {
		arrStudent = append(arrStudent, s.Students[pos])
	}
	return json.Marshal(struct {
		Name     string    `json:"name"`
		Students []Student `json:"students"`
	}{
		Students: arrStudent,
		Name:     s.Name,
	})
}
