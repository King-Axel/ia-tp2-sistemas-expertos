package httpserver

import (
	"net/http"

	"iteraciones/backend/internal/sebr/knowledge"
	"iteraciones/backend/internal/sebr/service"
	"iteraciones/backend/internal/sebr/store"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(corsMiddleware())

	ruleStore := store.NewRuleStore(knowledge.Rules())
	ruleService := service.NewRuleService(ruleStore)
	sebrHandler := NewSEBRHandler(ruleService)
	router.GET("/sebr/rules", sebrHandler.ListRules)
	router.POST("/sebr/rules", sebrHandler.CreateRule)
	router.POST("/sebr/infer", sebrHandler.Infer)

	return router
}

func corsMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		context.Header("Access-Control-Allow-Headers", "Content-Type")

		if context.Request.Method == http.MethodOptions {
			context.AbortWithStatus(http.StatusNoContent)
			return
		}

		context.Next()
	}
}
