package ggin

import (
	"context"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lsq51201314/go-utils/gexit"
	"github.com/lsq51201314/go-utils/gjwt"
)

func TestGgin(t *testing.T) {
	g := New("/api", 22345, true, func(r *gin.RouterGroup) {
		r.GET("/a", nil)
		r.POST("/b", nil)
		r.PUT("/c", nil)
		r.DELETE("/d", nil)
		j := gjwt.New("123456")
		r.Use(Validate(j))
		{
			r.GET("/e", nil)
			r.POST("/f", nil)
			r.PUT("/g", nil)
			r.DELETE("/h", nil)
		}
	})
	g.Run()
	gexit.WaitExit(func(ctx context.Context) {
		g.Stop(ctx)
	})
}
