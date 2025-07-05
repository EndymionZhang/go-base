package main

import (
	"github.com/endymion/go-base/task03/main/db"
	"github.com/endymion/go-base/task03/main/question"
)

func main() {
	gormTest()
}

func gormTest() {
	gormDb := db.GetGormDb()
	defer db.CloseGorm(gormDb)
	//question.RunStudent(gormDb)
	//question.RunTrade(gormDb)
	question.RunBlog(gormDb)
}
