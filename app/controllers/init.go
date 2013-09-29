package controllers

import (
	"github.com/toshipon/revmgo"
	"github.com/robfig/revel"
)

func (c Admin) checkUser() revel.Result {
	_, loggedIn := c.Session["username"]
	if loggedIn == false {
		c.Flash.Error("Please log in first")
		return c.Redirect(App.Login)
	} else {
		isAdmin := c.Session["role"] == "admin"
		if isAdmin == false {
			c.Flash.Error("Please log in as an admin first")
			return c.Redirect(App.Login)
		}
	}
	return nil
}

func init() {
	revmgo.ControllerInit()
	revel.InterceptMethod(Admin.checkUser, revel.BEFORE)
}
