package daos

import (
	"fmt"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/models"
	"github.com/tealeg/xlsx"
)

type ExcelFileDAO struct {
	filename string
	data     []models.Birthday
}

func (e *ExcelFileDAO) GetData() []models.Birthday {
	return e.data
}

func NewExcelFileDAO(filename string) *ExcelFileDAO {
	return &ExcelFileDAO{filename: filename}
}

func (e *ExcelFileDAO) GetFromFile() (int, error) {
	// открываем файл
	var rows int = 0
	xlFile, err := xlsx.OpenFile(e.filename)
	if err != nil {
		fmt.Println(err.Error())
		return rows, err
	}
	readByUsername, err := NewUserDAO().ReadByUsername("Golobar")
	if err != nil {
		return 0, err
	}
	// Перемещаем страницу листа, чтобы прочитать
	for _, sheet := range xlFile.Sheets {
		fmt.Println("sheet name: ", sheet.Name)
		// Обходим строки для чтения
		rows = len(sheet.Rows)
		for row_num, row := range sheet.Rows {
			var birthday = models.Birthday{}
			// Обходим столбцы каждой строки для чтения
			for i, cell := range row.Cells {
				switch i {
				case 0:
					birthday.FullName = cell.String()
				case 1:
					birthday.BirthDate, err = cell.GetTime(false)
					if err != nil {
						fmt.Println(err.Error())
						return row_num, err
					}
				case 2:
					birthday.PhoneNumber = cell.String()
				}
			}
			//FOR TESTING
			birthday.UserID = readByUsername.ID
			//FOR TESTING
			e.data = append(e.data, birthday)
		}
	}
	return rows, nil
}

func (e *ExcelFileDAO) SetToFile(fielename string) error {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("student_list")
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}

	//add data
	for _, birthday := range e.data {
		row := sheet.AddRow()
		nameCell := row.AddCell()
		nameCell.Value = birthday.FullName

		genderCell := row.AddCell()
		genderCell.Value = birthday.BirthDate.Format("02.01.2006")

		phoneCell := row.AddCell()
		phoneCell.Value = birthday.PhoneNumber

	}
	err = file.Save(fielename)
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}
	fmt.Println("\n\nexport success")
	return nil
}
