package auth

type PasswordHasher interface {
	Hash(password string) (string, error)
	Compare(hashedPassword, plainPassword string) bool
}
