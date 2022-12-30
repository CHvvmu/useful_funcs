package main

import (
	"fmt"
	"sort"
)

type StructureName struct {
	Field1 string 
	Field2 string 
	Field3 int    
}

// сортировка слайса структур по выбранному полю
type SortingByField []StructureName             //указать структуру
func (c SortingByField) Len() int           { return len(c) }
func (с SortingByField) Swap(i, j int)      { с[i], с[j] = с[j], с[i] }
//изменить поле
func (c SortingByField) Less(i, j int) bool { return c[i].Field3 < c[j].Field3 }
func sortStructure(std []StructureName) []StructureName {
	sort.Stable(SortingByField(std)) //результат
	return std
}

func main() {
	std := []StructureName{
		{"RU", "Gmail", 522},
		{"RU", "Yahoo", 383},
		{"RU", "Live", 89},
		{"US", "MSN", 213},
		{"US", "Orange", 165},
		{"US", "AOL", 25},
		{"GB", "Hotmail", 495},
		{"GB", "RediffMail", 321},
		{"GB", "Comcast", 16},		
	}
	sortStructure(std) //результат
	fmt.Println("После сортировки...")		
	for _, row := range std {
		fmt.Printf("%v \n", row)
	}
}
