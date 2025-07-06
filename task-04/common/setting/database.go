package setting

type Database struct {
	Type         string
	Host         string
	Port         string
	User         string
	Password     string
	Name         string
	Conf         string
	MaxIdleConns int
	MaxOpenConns int
}
