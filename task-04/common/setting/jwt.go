package setting

type Jwt struct {
	ExpireTime         int
	AccessTokenSecret  string
	RefreshTokenSecret string
}
