package httpserver

import (
	"iteraciones/backend/internal/sebr/knowledge"
	"iteraciones/backend/internal/sebr/service"
	"iteraciones/backend/internal/sebr/store"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	ruleStore := store.NewRuleStore(knowledge.Rules())
	ruleService := service.NewRuleService(ruleStore)
	sebrHandler := NewSEBRHandler(ruleService)
	router.GET("/sebr/rules", sebrHandler.ListRules)
	router.POST("/sebr/rules", sebrHandler.CreateRule)
	router.POST("/sebr/infer", sebrHandler.Infer)

	return router
}
