package http

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type contextKey int

const (
	ContextKeyDB contextKey = iota
)

//	func (s *Services) db(ctx context.Context) *gorm.DB {
//		if db, ok := ctx.Value(ContextKeyDB).(*gorm.DB); ok {
//			return db
//		}
//		return nil
//	}
func Middleware(conn *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c)
		// DB goroutine-safe
		ctx := context.WithValue(c.Request.Context(), ContextKeyDB, conn.WithContext(c.Request.Context()))
		c.Request = c.Request.WithContext(ctx)
	}
}
