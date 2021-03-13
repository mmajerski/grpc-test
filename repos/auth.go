package repos

import (
	"crypto/ed25519"
	"errors"
	"time"

	"github.com/pascaldekloe/jwt"
	"github.com/userq11/grpc-test/models"
	"xorm.io/xorm"
)

var (
	prvKey ed25519.PrivateKey
	pbKey  ed25519.PublicKey
)

func init() {
	var err error
	pbKey, prvKey, err = ed25519.GenerateKey(nil)
	if err != nil {
		panic(err)
	}
}

type AuthRepo interface {
	GetNewClaims(subject string, set map[string]interface{}) *jwt.Claims
	GetSignedToken(claims *jwt.Claims) (string, error)
	GetDataFromToken(token string) (*models.User, error)
}

// NewAuthRepo returns new auth repo
func NewAuthRepo(db *xorm.Engine) AuthRepo {
	return &authRepo{db: db}
}

type authRepo struct {
	db *xorm.Engine
}

func (a authRepo) GetNewClaims(subject string, set map[string]interface{}) *jwt.Claims {
	return &jwt.Claims{
		Registered: jwt.Registered{
			Subject: subject,
		},
		Set: set,
	}
}

func (a authRepo) GetSignedToken(claims *jwt.Claims) (string, error) {
	now := time.Now().Round(time.Second)

	claims.Registered.Issued = jwt.NewNumericTime(now)
	claims.Registered.Expires = jwt.NewNumericTime(now.Add(7 * time.Hour * 24))
	claims.NotBefore = jwt.NewNumericTime(now.Add(-time.Second))

	token, err := claims.EdDSASign(prvKey)
	if err != nil {
		return "", err
	}

	return string(token), nil
}

func (a authRepo) GetDataFromToken(token string) (*models.User, error) {
	claims, err := jwt.EdDSACheck([]byte(token), pbKey)
	if err != nil {
		return nil, err
	}

	userData, ok := claims.Set["user"].(map[string]interface{})
	if !ok {
		return nil, errors.New("user data missing in token")
	}

	user := new(models.User)

	id, ok := userData["id"].(float64)
	if !ok {
		return nil, errors.New("user data missing in token")
	}
	user.ID = int64(id)

	email, ok := userData["email"].(string)
	if !ok {
		return nil, errors.New("user data missing in token")
	}
	user.Email = email

	visible, ok := userData["visible"].(bool)
	if !ok {
		return nil, errors.New("user data missing in token")
	}
	user.Visible = visible

	return user, nil
}
