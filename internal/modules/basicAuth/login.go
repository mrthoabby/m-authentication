package basicAuth

type Credentials struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password"  binding:"required"`
}
