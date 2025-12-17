package router

import "github.com/gin-gonic/gin"

func RouterApiV1(h Handler, rg *gin.RouterGroup) {
	r := rg.Group("/v1")
	AuthRouter(r, h)
	ShowtimeRouter(r, h)

}

func AuthRouter(rg *gin.RouterGroup, h Handler) {
	r := rg.Group("/auth")
	r.POST("/login", h.UserHandler.Login)
}
func ShowtimeRouter(rg *gin.RouterGroup, h Handler) {
	r := rg.Group("/showtime")
	r.Use(h.Middleware.TokenAuthMiddleware())
	r.POST("", h.UserHandler.CreateShowtime)
	r.GET("", h.UserHandler.GetAllShowtime)
	r.DELETE("/:id", h.UserHandler.DeleteShowtime)
	r.PUT("", h.UserHandler.UpdateShowtime)
	r.GET("/:id", h.UserHandler.GetShowtimeByID)

}
