package workspace

import (
	"gin-notebook/internal/http/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.RouterGroup) {
	authGroup := r.Group("/workspace")
	authGroup.Use(middleware.JWTAuth())
	{
		authGroup.GET("/list", GetWorkspaceListApi)
		authGroup.POST("", CreateWorkspaceApi)
		authGroup.GET("", GetWorkspaceApi)
		authGroup.GET("/notes/", GetWorkspaceNotesApi)
		authGroup.PUT("/notes/", UpdateWorkspaceNoteApi)
		authGroup.GET("/notes/category/", GetWorkspaceNotesCategoryApi)
		authGroup.PUT("/notes/category/", UpdateWorkspaceCategoryApi)
	}
}
