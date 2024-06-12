package route

import (
	"decode/common"
	"decode/config"
	"decode/models/DatabaseStruct"
	"decode/models/ViewStruct"
	"decode/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// HandleSateFile @Summary 上传文件
// @Description 上传文件并解析元数据
// @Tags files
// @Accept  json
// @Produce  json
// @Param   input  body  ViewStruct.UploadRequest  true  "上传请求信息"
// @Success 200  {object}  map[string]interface{} "成功返回metadataId"
// @Router /upload [post]
func HandleSateFile(c *gin.Context) {
	db := config.ConnectDatabase()
	db.AutoMigrate(&DatabaseStruct.SateFY4BL1FileTab{})
	var input ViewStruct.UploadRequest
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	metadataId, err := service.ProcessFileAndSave(input.FilePath, db)
	if err != nil {
		if err == common.ErrNotEnoughParts {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "metadataId": metadataId})
}

// HandleAlertFile @Summary Say hello
// @Accept  json
// @Produce  json
// @Param   filepath     query    string     true        "Name to say hello"
// @Success 200 {object} map[string]interface{}
// @Router /xml [get]
func HandleAlertFile(c *gin.Context) {
	db := config.ConnectDatabase()
	db.AutoMigrate(&DatabaseStruct.SateFY4BL1FileTab{})
	filepath := c.Query("filepath")
	if filepath == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	metadataId, err := service.HandleAlertFile(filepath, db)
	if err != nil {
		if err == common.ErrNotEnoughParts {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "metadataId": metadataId})
}
