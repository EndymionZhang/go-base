package model

type Student struct {
	ID    int
	Name  string `gorm:"size:50"`
	Age   int
	Grade string `gorm:"size:50"`
}
