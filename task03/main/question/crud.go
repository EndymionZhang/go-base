package question

import (
	"fmt"
	"github.com/endymion/go-base/task03/main/model"
	"gorm.io/gorm"
)

type Student = model.Student

func RunStudent(db *gorm.DB) {
	//createStudent(db)
	findStudent(db)
	updateStudent(db)
	deleteStudent(db)
}

func createStudent(db *gorm.DB) {
	err := db.Debug().AutoMigrate(&Student{})
	if err != nil {
		fmt.Println("AutoMigrate error", err)
	}
	student := &Student{Name: "张三", Age: 20, Grade: "三年级"}
	db.Create(student)
}

func findStudent(db *gorm.DB) {
	var students []Student
	// 查询所有
	db.Debug().Where("age > ?", 18).Find(&students)
	for _, student := range students {
		fmt.Println(student)
	}
}

func updateStudent(db *gorm.DB) {
	var student Student
	student.ID = 1
	student.Name = "张三"
	db.Debug().Where(student).First(&student)
	student.Grade = "五年级"
	db.Debug().Save(&student)
}

func deleteStudent(db *gorm.DB) {
	var student Student
	db.Debug().Where("age < ?", 15).Delete(&student)
}
