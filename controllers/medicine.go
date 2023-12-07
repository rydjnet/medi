package controllers

import (
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
func FindRow(c *gin.Context) {
	log.Printf("Start FindRow")
	var i models.Medicine
	i.Name = c.Param("name")
	err := GetMedRow(&i)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": i})

}
func UpdateRow(c *gin.Context) {
	var i models.Medicine
	i.Name = c.Param("name")
	if err := GetMedRow(&i); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if err := c.ShouldBindJSON(&i); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if err := UpdateMed(&i); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": i})

}
func DelRow(c *gin.Context) {
	var i models.Medicine
	i.Name = c.Param("name")
	if err := DeleteMed(i.Name); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": i})
}

//-----------------

func AddMed(m *models.Medicine) error {
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
func GetMedRow(m *models.Medicine) error {
	err := models.DB.QueryRow("SELECT id, name, docmed FROM medi where name=$1", m.Name).Scan(&m.Id, &m.Name, &m.Docmed)
	if err != nil {
		return err
	}
	return nil
}
func UpdateMed(m *models.Medicine) error {
	_, err := models.DB.Exec("UPDATE medi SET name=$1, docmed=$2 WHERE id=$3", m.Name, m.Docmed, m.Id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteMed(s string) error {
	_, err := models.DB.Exec("DELETE FROM medi WHERE name=$1", s)
	if err != nil {
		return err
	}
	return nil
}
