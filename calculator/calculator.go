package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Data struct {
	OrigWaterLevel    float32 `json:"orig_water_level"`
	LiftingSpeedWater float32 `json:"lifting_speed_water"`
	FloodTime         float32 `json:"flood_time"`
	AnyLevel          float32 `json:"any_level"`
}

type Result struct {
	NewLevel float32 `json:"new_level"`
	NewTime  float32 `json:"new_time"`
}

var (
	data   Data
	result Result
)

func main() {
	//	port := "3000"
	r := gin.Default()

	//Загрузка шаблонов (Рендеринг HTML)
	r.LoadHTMLGlob("templates/*")
	//Маршруты:
	r.GET("/", getData)
	r.POST("/", postData)

	r.Run()
}

// 1. показать поля расчета
func getData(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

// 2. вернуть результат
func postData(c *gin.Context) {
	if err := c.Bind(&data); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		result.NewLevel = data.OrigWaterLevel + data.LiftingSpeedWater*data.FloodTime
		result.NewTime = (data.AnyLevel - data.OrigWaterLevel) / data.LiftingSpeedWater
		c.HTML(http.StatusOK, "index.html", gin.H{
			
			"new": result,
		})
	}
}
