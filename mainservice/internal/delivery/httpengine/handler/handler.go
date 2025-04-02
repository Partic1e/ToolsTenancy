package handler

import (
	"github.com/gin-gonic/gin"
	ad "mainservice/api/ad"
	"mainservice/internal/lib/grpcclient"
	"net/http"
	"strconv"
)

type Handler struct {
	userClient    *grpcclient.GrpcClient
	paymentClient *grpcclient.GrpcClient
	adClient      *grpcclient.GrpcClient
}

func NewHandler(userClient, paymentClient, adClient *grpcclient.GrpcClient) *Handler {
	return &Handler{
		userClient:    userClient,
		paymentClient: paymentClient,
		adClient:      adClient,
	}
}

func (h *Handler) GetOrCreateUser(c *gin.Context) {
	tgIdStr := c.Query("tg_id")
	if tgIdStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no tg_id provided"})
		return
	}

	tgId, err := strconv.ParseInt(tgIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid tg_id"})
		return
	}

	resp, err := h.userClient.GetOrCreateUser(c.Request.Context(), tgId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tg_id": resp.TgId, "balance": resp.Balance})
}

func (h *Handler) Deposit(c *gin.Context) {
	userIdStr := c.Query("user_id")
	amount := c.Query("amount")

	if userIdStr == "" || amount == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing parameters"})
		return
	}

	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id"})
		return
	}

	success, err := h.paymentClient.Deposit(c.Request.Context(), userId, amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": success})
}

func (h *Handler) WithDraw(c *gin.Context) {
	userIdStr := c.Query("user_id")
	amount := c.Query("amount")

	if userIdStr == "" || amount == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing parameters"})
		return
	}

	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id"})
		return
	}

	success, err := h.paymentClient.WithDraw(c.Request.Context(), userId, amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": success})
}

func (h *Handler) CreateAd(c *gin.Context) {
	var req ad.CreateAdRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	resp, err := h.adClient.CreateAd(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) DeleteAd(c *gin.Context) {
	name := c.Query("name")
	landlordIdStr := c.Query("landlord_id")

	if name == "" || landlordIdStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing parameters"})
		return
	}

	landlordId, err := strconv.ParseInt(landlordIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid landlord_id"})
		return
	}

	req := &ad.DeleteAdRequest{
		Name:       name,
		LandlordId: landlordId,
	}

	success, err := h.adClient.DeleteAd(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": success})
}

func (h *Handler) UpdateAd(c *gin.Context) {
	var req ad.Ad
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	resp, err := h.adClient.UpdateAd(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetAllCategories(c *gin.Context) {
	resp, err := h.adClient.GetAllCategories(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetAdsByCategory(c *gin.Context) {
	categoryIdStr := c.Query("category_id")
	if categoryIdStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing category_id"})
		return
	}

	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category_id"})
		return
	}

	resp, err := h.adClient.GetAdsByCategory(c.Request.Context(), categoryId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetAdsByLandlord(c *gin.Context) {
	landlordIdStr := c.Query("landlord_id")
	if landlordIdStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing landlord_id"})
		return
	}

	landlordId, err := strconv.ParseInt(landlordIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid landlord_id"})
		return
	}

	resp, err := h.adClient.GetAdsByLandlord(c.Request.Context(), landlordId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
