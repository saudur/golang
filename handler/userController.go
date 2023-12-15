package userController

import (
	"encoding/json"
	"net/http"
	"sekolah/config"
	"sekolah/models"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	var coba []models.User

	config.DB.Find(&coba)
	c.JSON(http.StatusOK, gin.H{"data": coba})
}

func Create(c *gin.Context) {

	var coba models.User

	if err := c.ShouldBindJSON(&coba); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	config.DB.Create(&coba)
	c.JSON(http.StatusOK, gin.H{"data": coba})
}

func Update(c *gin.Context) {
	var coba models.User
	id := c.Param("id")

	if err := c.ShouldBindJSON(&coba); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if config.DB.Model(&coba).Where("id = ?", id).Updates(&coba).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate Data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})

}

func Delete(c *gin.Context) {

	var coba models.User

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if config.DB.Delete(&coba, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
