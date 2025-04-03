package httpengine

import (
	"context"
	"github.com/gin-gonic/gin"
	"mainservice/internal/delivery/httpengine/handler"
)

func InitRouter(ctx context.Context, h *handler.Handler) *gin.Engine {
	router := gin.Default()

	router.GET("/user", h.GetOrCreateUser)

	router.POST("/payment/deposit", h.Deposit)
	router.POST("/payment/withdraw", h.WithDraw)

	router.POST("/ad", h.CreateAd)
	router.DELETE("/ad", h.DeleteAd)
	router.PUT("/ad", h.UpdateAd)
	router.GET("/categories", h.GetAllCategories)
	router.GET("/ads/category", h.GetAdsByCategory)
	router.GET("/ads/landlord", h.GetAdsByLandlord)

	router.GET("/rents/landlord", h.GetRentsByLandlord)
	router.GET("/rents/renter", h.GetRentsByRenter)
	router.GET("/rents/dates", h.GetRentedDates)

	return router
}
