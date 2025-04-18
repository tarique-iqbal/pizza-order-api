package auth

type JWTService interface {
	GenerateToken(userID uint) (string, error)
	ParseToken(tokenString string) (*Claims, error)
}

type Claims struct {
	UserID uint `json:"user_id"`
}
