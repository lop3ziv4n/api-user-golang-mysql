package controllers

import (
	"github.com/jinzhu/gorm"
	"github.com/lop3ziv4n/api-user-golang-mysql/common"
)

// Context Struct used for maintaining HTTP Request Context
type Context struct {
	MySQLSession *gorm.DB
}

// Close mysql db
func (c *Context) Close() {
	c.MySQLSession.Close()
}

// Db Returns db context
func (c *Context) Db() *gorm.DB {
	return c.MySQLSession
}

// NewContext Create a new Context object for each HTTP request
func NewContext() *Context {
	session := common.GetDbSession()
	context := &Context{
		MySQLSession: session,
	}
	return context
}
