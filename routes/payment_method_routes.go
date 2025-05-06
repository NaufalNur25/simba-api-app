package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/naufal/simba-api-app/controllers"
)

func PaymentMethodRoute(router *gin.Engine) {
	paymentMethodGroup := router.Group("/payment-methods")
	{
		paymentMethodGroup.GET("/", controllers.GetPaymentMethods)
		paymentMethodGroup.POST("/", controllers.CreatePaymentMethod)
		paymentMethodGroup.GET("/:id", controllers.GetPaymentMethodDetail)
		paymentMethodGroup.PUT("/:id", controllers.UpdatePaymentMethod)
		paymentMethodGroup.DELETE("/:id", controllers.DeletePaymentMethod)
	}
}
