package spares

import (
	"log"
	"net/http"
	"time"

	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/dgrijalva/jwt-go"
)

const (
	SecretKey = "Muuha"
	ValidUser = "dhc"
	ValidPass = "1234"

	AuthRedirectUrl = "/account"
)

// User model
type LoginUserModel struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// Field validator
func (u *LoginUserModel) Validate(errors *binding.Errors, req *http.Request) {

	// if len(u.Username) < 4 {
	// 	errors.Fields["username"] = "Too short; minimum 4 characters"
	// }

	// if len(u.Password) < 6 {
	// 	errors.Fields["password"] = "Too short; minimum 4 characters"
	// }
}

func AccountHandlers(m *martini.ClassicMartini) {

	m.Use(BearerTokenValidation)
	m.Get("/account", func(r render.Render) {
		r.HTML(200, "account/login", nil)
	})
	// Authenticate user
	m.Post("/account", binding.Bind(LoginUserModel{}), func(user LoginUserModel, r *http.Request, render render.Render) {

		log.Println("user", user.Username)
		log.Println("pass", user.Password)

		if user.Username == ValidUser && user.Password == ValidPass {

			// Create JWT token
			token := jwt.New(jwt.GetSigningMethod("HS256"))
			token.Claims["username"] = user.Username
			// Expire in 5 mins
			token.Claims["exp"] = time.Now().Add(time.Minute * 5).Unix()
			tokenString, err := token.SignedString([]byte(SecretKey))

			if err == nil {
				data := map[string]string{"token": tokenString}
				render.JSON(201, data)
				return
			}

			r.Header.Add("Authorization", tokenString)

		}

		render.Redirect("/")
	})
}

func BearerTokenValidation(w http.ResponseWriter, r *http.Request, render render.Render) {

	authHeader := r.Header["Authorization"]
	log.Println(authHeader)

	if r.URL.Path != AuthRedirectUrl {
		//render.Redirect(AuthRedirectUrl, 302)
	}

}

func loginHandler(r render.Render) {
	r.HTML(200, "account/login", nil)
}

func logoutHandler() (int, string) {
	return 200, "Hello, World!!!!"
}

func validateToken(token string) bool {
	t, err := jwt.Parse(token, func(token *jwt.Token) ([]byte, error) {
		return []byte(SecretKey), nil
	})

	if err == nil && t.Valid {
		return true
	}

	return false
}
