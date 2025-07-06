package model

func InitDb() error {
	err := db.AutoMigrate(&User{}, &Post{}, &Comment{})
	return err
}
