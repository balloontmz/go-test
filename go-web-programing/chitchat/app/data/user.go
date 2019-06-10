package data

type (
	User struct {
		Password string
	}
)

func UserByEmail(email string) (user User, err error) {
	return
}

func Encrypt(passwd string) (passwdHash string) {
	return
}
