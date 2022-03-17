package handler

import "github.com/gin-gonic/gin"

func (h *Handler) Void() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
