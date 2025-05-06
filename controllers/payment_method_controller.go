package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/naufal/simba-api-app/models"
	"github.com/naufal/simba-api-app/services"
)

func GetPaymentMethods(c *gin.Context) {
	isActiveQuery := c.Query("active")

	var isActive *bool

	if isActiveQuery != "" {
		val, err := strconv.ParseBool(isActiveQuery)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value for 'active'. Use true or false."})
			return
		}
		isActive = &val
	}

	paymentMethods, err := services.GetAllPaymentMethod(isActive)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch payment methods"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": paymentMethods})
}

func CreatePaymentMethod(c *gin.Context) {
	var input models.PaymentMethod
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.CreatePaymentMethod(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create payment method"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": result})
}

func GetPaymentMethodDetail(c *gin.Context) {
	id := c.Param("id")

	paymentMethod, err := services.GetPaymentMethodByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Payment method not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": paymentMethod,
	})
}

func UpdatePaymentMethod(c *gin.Context) {
	id := c.Param("id")
	var input models.PaymentMethod

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.UpdatePaymentMethod(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update payment method"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

func DeletePaymentMethod(c *gin.Context) {
	id := c.Param("id")

	err := services.DeletePaymentMethod(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete payment method"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment method deleted"})
}
