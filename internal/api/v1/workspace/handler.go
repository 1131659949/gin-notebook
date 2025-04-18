package workspace

import (
	"gin-notebook/internal/http/message"
	"gin-notebook/internal/http/response"
	"gin-notebook/internal/pkg/dto"
	"gin-notebook/internal/service/note"
	"gin-notebook/internal/service/workspace"
	"gin-notebook/pkg/logger"
	validator "gin-notebook/pkg/utils/validatior"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetWorkspaceListApi(c *gin.Context) {
	userID := c.MustGet("userID").(int64)
	responseCode, data := workspace.GetWorkspaceList(userID)
	if data != nil {
		data = map[string]interface{}{
			"workspaces": WorkspaceListSerializer(c, data.(*[]dto.WorkspaceListDTO)),
			"total":      len(*data.(*[]dto.WorkspaceListDTO)),
		}
	}
	c.JSON(http.StatusOK, response.Response(responseCode, data))
}

func CreateWorkspaceApi(c *gin.Context) {
	userID := c.MustGet("userID").(int64)
	params := &dto.WorkspaceValidation{
		Owner: userID,
	}

	if err := c.ShouldBindJSON(params); err != nil {
		log.Printf("params %s", err)
		c.JSON(http.StatusInternalServerError, response.Response(message.ERROR_INVALID_PARAMS, nil))
		return
	}

	err := validator.ValidateStruct(params)
	if err != nil {
		c.JSON(http.StatusOK, response.Response(message.ERROR_WORKSPACE_VALIDATE, nil))
		return
	}

	responseCode, data := workspace.CreateWorkspace(params)
	c.JSON(http.StatusOK, response.Response(responseCode, data))
}

func GetWorkspaceApi(c *gin.Context) {
	userID := c.MustGet("userID").(int64)
	workspaceID := c.Query("workspace_id")
	logger.LogDebug("workspaceID: ", map[string]interface{}{
		"workspace_id": workspaceID,
		"user_id":      userID,
	})
	responseCode, data := workspace.GetWorkspace(workspaceID, userID)
	if data != nil {
		data = WorkspaceSerializer(c, data.(*dto.WorkspaceListDTO))
	}
	c.JSON(http.StatusOK, response.Response(responseCode, data))
}

func GetWorkspaceNotesApi(c *gin.Context) {
	workspaceID := c.Query("workspace_id")
	offset, limit := c.DefaultQuery("offset", "0"), c.DefaultQuery("limit", "20")

	noteOffset, err := strconv.Atoi(offset)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response(message.ERROR_INVALID_PARAMS, nil))
		return
	}

	noteLimit, err := strconv.Atoi(limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response(message.ERROR_INVALID_PARAMS, nil))
		return
	}

	if workspaceID == "" {
		c.JSON(http.StatusBadRequest, response.Response(message.ERROR_INVALID_PARAMS, nil))
		return
	}

	userID, err := strconv.ParseInt(c.DefaultQuery("user_id", "0"), 10, 64)
	if err != nil || userID <= 0 {
		userID = c.MustGet("userID").(int64)
	}

	responseCode, data := note.GetWorkspaceNotesList(workspaceID, userID, noteLimit, noteOffset)
	if data != nil {
		data = map[string]interface{}{
			"notes": data,
			"total": len(*data.(*[]dto.WorkspaceNoteDTO)),
		}
	}
	c.JSON(http.StatusOK, response.Response(responseCode, data))

}

func GetWorkspaceNotesCategoryApi(c *gin.Context) {
	wid := c.Query("workspace_id")
	if wid == "" {
		c.JSON(http.StatusBadRequest, response.Response(message.ERROR_INVALID_PARAMS, nil))
		return
	}

	workspaceID, err := strconv.ParseInt(wid, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response(message.ERROR_INVALID_PARAMS, nil))
		return
	}

	responseCode, data := note.GetWorkspaceNotesCategory(workspaceID)
	c.JSON(http.StatusOK, response.Response(responseCode, data))

}

func UpdateWorkspaceNoteApi(c *gin.Context) {
	userID := c.MustGet("userID").(int64)
	params := &dto.UpdateWorkspaceNoteValidator{
		UserID: userID,
	}

	if err := c.ShouldBindJSON(params); err != nil {
		log.Printf("params %s", err)
		c.JSON(http.StatusInternalServerError, response.Response(message.ERROR_INVALID_PARAMS, nil))
		return
	}

	if err := validator.ValidateStruct(params); err != nil {
		c.JSON(http.StatusOK, response.Response(message.ERROR_WORKSPACE_NOTE_VALIDATE, nil))
		return
	}

	responseCode, data := note.UpdateNote(params)
	if responseCode != message.SUCCESS {
		c.JSON(http.StatusInternalServerError, response.Response(responseCode, nil))
		return
	}
	c.JSON(http.StatusOK, response.Response(responseCode, data))
}

func UpdateWorkspaceCategoryApi(c *gin.Context) {
	userID := c.MustGet("userID").(int64)
	params := &dto.UpdateWorkspaceNoteCategoryDTO{}

	if err := c.ShouldBindJSON(params); err != nil {
		log.Printf("params %s", err)
		c.JSON(http.StatusInternalServerError, response.Response(message.ERROR_INVALID_PARAMS, nil))
		return
	}
	params.OwnerID = &userID
	if err := validator.ValidateStruct(params); err != nil {
		c.JSON(http.StatusOK, response.Response(message.ERROR_WORKSPACE_NOTE_VALIDATE, nil))
		return
	}

	responseCode, data := note.UpdateNoteCategory(params)
	if responseCode != message.SUCCESS {
		c.JSON(http.StatusInternalServerError, response.Response(responseCode, nil))
		return
	}
	c.JSON(http.StatusOK, response.Response(responseCode, data))
}
