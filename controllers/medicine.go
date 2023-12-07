package controllers

import (
	"fmt"
	"log"
	"medi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var input models.Medicine
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	medicine := models.Medicine{Name: input.Name, Docmed: input.Docmed}
	if err := AddMed(&medicine); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": medicine})
}
func GetList(c *gin.Context) {
	log.Printf("Start GetList")
	list, err := GetMedList()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	}
	c.JSON(http.StatusOK, gin.H{"data": list})
}

//-----------------

func AddMed(m *models.Medicine) error {
	fmt.Println("start AddMed: ", m)
	sqlStatement := "INSERT INTO public.medi (id, name, docmed) VALUES (uuid_generate_v4(), $1, $2)"
	_, err := models.DB.Exec(sqlStatement, m.Name, m.Docmed)
	if err != nil {
		return err
	}
	log.Printf("New medicine %s add successfuly", m.Name)
	return nil
}

func GetMedList() ([]models.Medicine, error) {
	m := []models.Medicine{}
	rows, err := models.DB.Query("SELECT id,name,docmed FROM public.medi")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var r models.Medicine
		err := rows.Scan(&r.Id, &r.Name, &r.Docmed)
		if err != nil {
			return nil, err
		}
		m = append(m, r)
	}

	return m, nil

}
