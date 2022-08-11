package daos

import (
	"fmt"
	"log"
	"testing"
)

func TestExcelFileDAO_GetFromFile(t *testing.T) {
	dao := NewExcelFileDAO("Birthdays.xlsx")
	num, err := dao.GetFromFile()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(num)
}
func TestExcelFileDAO_SetToFile(t *testing.T) {
	dao := NewExcelFileDAO("Birthdays.xlsx")
	num, err := dao.GetFromFile()
	if err != nil {
		log.Fatal(err)
	}
	err = dao.SetToFile("example.xlsx")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(num)
}
