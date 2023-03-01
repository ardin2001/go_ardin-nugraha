package main

import (
	"fmt"
)

type Students struct {
	name  []string
	score []int
}

func (student Students) Average() float32 {
	AllScore := 0
	for _, score := range student.score {
		AllScore += score
	}
	avg := float32(AllScore / len(student.score))
	return avg
}
func (student Students) MinScore() (int, string) {
	min := 99999
	NameMin := ""
	for i, score := range student.score {
		if score < min {
			min = score
			NameMin = student.name[i]
		}
	}
	return min, NameMin

}
func (student Students) MaxScore() (int, string) {
	max := 0
	NameMax := ""
	for i, score := range student.score {
		if score > max {
			max = score
			NameMax = student.name[i]
		}
	}
	return max, NameMax

}

func main() {
	StoreName := make([]string, 0)
	StoreScore := make([]int, 0)

	fmt.Println("Input:")
	for i := 1; i <= 5; i++ {
		name := ""
		score := 0
		fmt.Print("input ", i, " Student's Name :")
		fmt.Scan(&name)
		fmt.Print("input ", i, " Student's Score :")
		fmt.Scan(&score)

		StoreName = append(StoreName, name)
		StoreScore = append(StoreScore, score)
	}

	AllStudent := Students{
		name:  StoreName,
		score: StoreScore,
	}

	fmt.Println("Output:")
	fmt.Println("Average Score :", AllStudent.Average())
	min, NameMin := AllStudent.MinScore()
	fmt.Println("Min Score of Students:", NameMin, "(", min, ")")
	max, NameMax := AllStudent.MaxScore()
	fmt.Println("Max Score of Students:", NameMax, "(", max, ")")

}
