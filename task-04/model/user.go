package model

type User struct {
	BaseEntity
	// form 传参数
	Username   string `gorm:"type:varchar(20);not null" json:"username" validate:"min:6,max:20" form:"username"`
	Password   string `gorm:"type:varchar(20);not null" json:"password" validate:"min:6,max:20" form:"password"`
	Email      string `gorm:"type:varchar(255);not null" json:"email" validate:"email" form:"email"`
	PostsCount int    `gorm:"type:int;default:0" json:"posts_count"`
}

func (user *User) CreateUser() error {
	return db.Create(user).Error
}

func IsUserNameExist(user *User) (exist bool) {
	tx := db.Where("username = ?", user.Username).First(&User{})
	if tx.RowsAffected == 1 {
		return true
	}
	return false
}

func FindUserByName(username string) (user *User, exist bool) {
	user = &User{}
	tx := db.Where("username = ?", username).First(&user)
	exist = tx.RowsAffected == 1
	return
}
