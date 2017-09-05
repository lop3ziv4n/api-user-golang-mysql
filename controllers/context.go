package controllers

import (
	"github.com/jinzhu/gorm"
	"github.com/lop3ziv4n/api-user-golang-mysql/common"
)

// Struct used for maintaining HTTP Request Context
type Context struct {
	DB *gorm.DB
}

// Close mysql db
func (c *Context) close() {
	c.DB.Close()
}

// Returns db context
func (c *Context) db() *gorm.DB {
	return c.DB
}

// Create a new Context object for each HTTP request
func newContext() *Context {
	session := common.GetDbSession()
	context := &Context{
		DB: session,
	}
	return context
}
