package config

import (
	"errors"
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber"
	. "github.com/itsursujit/fiber-boilerplate/app"
	"time"
)

type AuthConfiguration struct {
	App_Jwt_Secret string
	Api_Jwt_Secret string
	Jwt_Expire     int
}

type Token struct {
	Hash   string
	Expire int64
}

var AuthConfig *AuthConfiguration //nolint:gochecknoglobals

var Auth *PermissionMiddleware

func LoadAuthConfig() {
	loadDefaultAuthConfig()
	ViperConfig.Unmarshal(&AuthConfig)
}

func loadDefaultAuthConfig() {
	ViperConfig.SetDefault("APP_JWT_SECRET", "SECRET_APP")
	ViperConfig.SetDefault("API_JWT_SECRET", "SECRET_API")
	ViperConfig.SetDefault("JWT_EXPIRE", 60*60)
}

func SetupPermission() { //nolint:whitespace
	LoadAuthConfig()
	var err error
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/", DBConfig.DB_User, DBConfig.DB_Pass, DBConfig.DB_Host, DBConfig.DB_Port) //nolint:wsl,lll
	PermissionAdapter, err = gormadapter.NewAdapter(DBConfig.DB_Driver, connectionString)

	if err != nil {
		panic(fmt.Sprintf("failed to initialize casbin adapter: %v", err))
	}
	Enforcer, _ = casbin.NewEnforcer("rbac_model.conf", PermissionAdapter) //nolint:wsl
	Auth = &PermissionMiddleware{
		Enforcer:      Enforcer, //nolint:gofmt
		PolicyAdapter: PermissionAdapter,
		Lookup: func(ctx *fiber.Ctx) string {
			return "sujit"
		},
		Unauthorized: func(c *fiber.Ctx) {
			var err fiber.Error
			err.Code = fiber.StatusUnauthorized
			CustomErrorHandler(c, &err)
		},
		Forbidden: func(c *fiber.Ctx) {
			var err fiber.Error
			err.Code = fiber.StatusForbidden
			CustomErrorHandler(c, &err)
		},
	}
}

//CreateToken authenticates the user
func CreateToken(c *fiber.Ctx, userID uint, secret string) (Token, error) {
	var t Token
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = userID
	expiresIn := time.Now().Add(time.Duration(AuthConfig.Jwt_Expire) * time.Second).Unix()
	claims["exp"] = expiresIn

	tokenHash, err := token.SignedString([]byte(secret))

	if err != nil {
		return t, err
	}

	c.Cookie(&fiber.Cookie{
		Name:     "fiber-boilerplate-Token",
		Value:    tokenHash,
		Secure:   false,
		HTTPOnly: true,
	})
	t.Hash = tokenHash
	t.Expire = expiresIn
	return t, nil
}

//ParseToken returns the users id or error
func ParseToken(c *fiber.Ctx, secret string) (uint, error) {
	tokenString := c.Cookies("fiber-boilerplate-Token")

	if tokenString == "" {
		return 0, errors.New("Empty auth cookie")
	}

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return 0, err
	}

	//Checks if the token is valid if it is not then it deletes it
	err2 := claims.Valid()

	if err2 != nil {
		DeleteToken(c)
		return 0, err2
	}

	return uint(claims["id"].(float64)), nil
}

//DeleteToken deletes the jwt token
func DeleteToken(c *fiber.Ctx) {
	c.ClearCookie("fiber-boilerplate-Token")
}

//RefreshToken refreshes the token
func RefreshToken(c *fiber.Ctx, secret string) {
	u, err := ParseToken(c, secret)

	if err != nil {
		return
	}

	CreateToken(c, u, secret)
}
