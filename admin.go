package spares

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/codegangsta/martini-contrib/render"
	"log"
	"net/http"

	"github.com/spares/model"
	"github.com/spares/storage"
)

// User model
type UserModel struct {
	Fullname string `form:"fullname" binding:"required"`
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
	Role     int    `form:"role" binding:"required"`
}

// Field validator
func (u *UserModel) Validate(errors *binding.Errors, req *http.Request) {

	if len(u.Fullname) < 6 {
		errors.Fields["fullname"] = "Too short; minimum 6 characters"
	}

	if len(u.Password) < 6 {
		errors.Fields["password"] = "Too short; minimum 6 characters"
	}

	if u.Role <= 0 {
		errors.Fields["role"] = "A role must be provieded."
	}

}

func AdministationHandlers(m *martini.ClassicMartini) {
	log.Println("Initializing Administration pages.")

	m.Get("/admin", adminIndexHandler)

	// accounts.
	m.Post("/api/admin/account", binding.Bind(UserModel{}), adminNewAccount)
}

func adminIndexHandler(render render.Render) {
	render.HTML(200, "admin/index", nil)
}

func adminNewAccount(usr UserModel) {

	var repo = new(storage.AccountRepository)
	repo.Post(usr)

}
