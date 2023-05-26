package scraping

import (
	"fmt"
	"net/http"

	// "regexp"
	// "strings"

	//"strings"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

type Data struct {
	Image       string `json:"image"`
	Name        string `json:"name"`
	Price       string `json:"price"`
	Style       string `json:"style"`
	Description string `json:"description"`
	Id          string `json:"id"`
	// Menu        []string `json:"menu"`
}

func ScrapedData(g *gin.Context) {
	var cafeProducts []Data
	c := colly.NewCollector()
	c.AllowedDomains = []string{"astana.restoran.kz"}

	c.OnHTML(".place-list-card", func(e *colly.HTMLElement) {

		cafeProduct := Data{}
		cafeProduct.Name = e.ChildText(".link-inherit-color")
		//cafeProduct.Price = e.ChildText(".d-flex.mr-5.mb-3")
		id := e.Attr("data-site-id")
		cafeProduct.Id = id

		//Extract the style text
		styleText := e.ChildText("div.list-unstyled.mb-4 li:nth-of-type(1)")
		var style *string
		if styleText != "" {
			style = &styleText
		}

		// Extract the price text
		priceText := e.ChildText("div.list-unstyled.mb-4 li:nth-of-type(2)")
		// fmt.Println("priceText:", priceText)
		var price *string
		if priceText != "" {
			price = &priceText
		}

		// Extract the description text
		descriptionText := e.ChildText("div.list-unstyled.mb-4 li:nth-of-type(3)")
		var description *string
		if descriptionText != "" {
			description = &descriptionText
		}

		// Process or store the extracted data as needed
		if style != nil {
			cafeProduct.Style = *style
			// fmt.Println("Style:", *style)
		}
		if price != nil {
			cafeProduct.Price = *price
			// fmt.Println("Price:", *price)
		}
		if description != nil {
			cafeProduct.Description = *description
			// fmt.Println("Description:", *description)
		}
		img := e.ChildAttr("img", "data-src")
		fmt.Println("image: ", img)
		cafeProduct.Image = img
		cafeProducts = append(cafeProducts, cafeProduct)

	})
	c.Visit("https://astana.restoran.kz/cafe")
	// fmt.Println(cafeProducts)
	g.JSON(http.StatusOK, cafeProducts)
}
