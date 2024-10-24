package db

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type Student struct {
	gorm.Model
	Name   string `json:"name"`
	CPF    int    `json:"cpf"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	Active bool   `json:"registration"`
}

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("student.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&Student{}) // db gerenciar a criação e atualização da tabela Student

	return db
}

func AddStudent(student Student) {
	db := Init()

	if result := db.Create(&student); result.Error != nil {
		fmt.Println("Error to create student")
	}

	fmt.Println("Student created")
}
