package controllers

import (
	"github.com/jinzhu/gorm"
	"github.com/lop3ziv4n/api-user-golang-mysql/common"
)

// Struct used for maintaining HTTP Request Context
type Context struct {
	MySQLSession *gorm.DB
}

// Close mysql db
func (c *Context) Close() {
	defer c.MySQLSession.Close()
}

// Returns db context
func (c *Context) Db() *gorm.DB {
	return c.MySQLSession
}

// Create a new Context object for each HTTP request
func NewContext() *Context {
	session := common.GetDbSession()
	context := &Context{
		MySQLSession: session,
	}
	return context
}
