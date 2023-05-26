package scraping

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

type InnerProduct struct {
	Images []string `json:"images"`
	Adress string   `json:"adress"`
	Text   string   `json:"text"`
	Number string   `json:"number"`
}

func GetInner(g *gin.Context) {
	id := g.Param("id")
	var innerProduct InnerProduct
	b := colly.NewCollector()
	// fmt.Println("id,", id)
	b.AllowedDomains = []string{"astana.restoran.kz"}
	b.OnHTML(".place-page-fotorama-img", func(e *colly.HTMLElement) {
		// Extract the text content of the element
		image := e.Attr("src")
		innerProduct.Images = append(innerProduct.Images, image)
	})
	b.OnHTML("a.link.mr-3[href='#placeMap']", func(e *colly.HTMLElement) {
		text := e.Text
		// fmt.Println(text)
		innerProduct.Adress = text
	})
	b.OnHTML("#descriptionContentFader .text-content", func(e *colly.HTMLElement) {
		text := e.Text
		innerProduct.Text = text
		// fmt.Println(text)
	})
	b.OnHTML(".phone", func(e *colly.HTMLElement) {
		prefix := e.ChildText(".phone-prefix")
		code := e.ChildText(".phone-code")
		number := e.ChildText(".phone-number")

		// Concatenate the phone number components
		phoneNumber := fmt.Sprintf("%s(%s)%s", prefix, code, number)

		innerProduct.Number = phoneNumber
	})

	b.Visit("https://astana.restoran.kz/cafe/" + id)
	g.JSON(http.StatusOK, innerProduct)
}
