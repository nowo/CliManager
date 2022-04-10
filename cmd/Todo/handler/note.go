package handler

import (
	"CliTaskManager/cmd/Todo/model"
	"CliTaskManager/database"
	"fmt"
)

func CreateNewNote(content string) {
	db := database.DB
	note := model.Note{Content: content, IsCompleted: false}
	db.Create(&note)
}

func GetAllNotes() {
	db := database.DB
	var notes []model.Note
	db.Find(&notes)
	for _, value := range notes {
		fmt.Printf("%d : %s \t tamamlandi mi : %t \n", value.ID, value.Content, value.IsCompleted)
	}
}

func DeleteNote(index int) {
	db := database.DB
	note := &model.Note{}
	err := db.Table("notes").Where("id = ?", index).First(note).Error
	fmt.Println(index)
	if err != nil {
		fmt.Println(err)
	}
	rowsEffected := db.Table("notes").Where("id = ?", index).Delete(note).RowsAffected

	if rowsEffected == 0 {
		fmt.Println("etkilenen row yok")
	}
}
