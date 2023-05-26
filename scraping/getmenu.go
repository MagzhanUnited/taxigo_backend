package scraping

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

// type Menu struct {
// 	Id string `json:"id"`
// }

func GetMenu(g *gin.Context) {
	// var requestMenu Menu

	// if err := c.ShouldBindJSON(&requestMenu); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	var image []string
	id := g.Param("id")
	c := colly.NewCollector()
	c.AllowedDomains = []string{"astana.restoran.kz"}
	c.OnHTML(".content-page", func(e *colly.HTMLElement) {
		src := e.ChildAttr("img", "src")
		fmt.Println("Image source:", src)
		image = append(image, src)
	})
	c.Visit("https://astana.restoran.kz/cafe/" + id + "/menu")
	g.JSON(http.StatusOK, image)
}
