package controllers

import (
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// AccountController operations for AccountController
type AccountController struct {
	beego.Controller
}

// URLMapping ...
func (c *AccountController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create AccountController
// @Param	body		body 	models.AccountController	true		"body for AccountController content"
// @Success 201 {object} models.AccountController
// @Failure 403 body is empty
// @router / [post]
func (c *AccountController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get AccountController by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.AccountController
// @Failure 403 :id is empty
// @router /:id [get]
func (c *AccountController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get AccountController
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.AccountController
// @Failure 403
// @router / [get]
func (c *AccountController) GetAll() {
	bindType := c.GetString("bind_type", "")
	logs.Debug(bindType)
	var url string
	if bindType == "pdd" {
		url = "https://baidu.com"
	}
	c.Redirect(url, 302)
}

// Put ...
// @Title Put
// @Description updates the AccountController
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body 	models.AccountController	true		"body for AccountController content"
// @Success 200 {object} models.AccountController
// @Failure 403 :id is not int
// @router /:id [put]
func (c *AccountController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the AccountController
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *AccountController) Delete() {

}
