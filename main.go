package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func check(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func getDistrict() string  {
	districts := [8]string{"colombo", "kandy", "galle", "ampara", "anuradhapura", "badulla", "batticaloa", "gampaha"}
	fmt.Println("\nPlease select a district")
	for i, dis := range districts {
		fmt.Println("\t", i+1, ". ", dis)
	}
	fmt.Print("Enter district: ")
	var district int
	_, err := fmt.Scanf("%d", &district)
	check(err)
	switch district {
	case 1:
		return districts[district-1]
	case 2:
		return districts[-district-1]
	case 3:
		return districts[-district-1]
	default:
		return ""
	}
}

func getCategory() string {
	categories := [9]string{"electronics", "vehicles", "property", "animals", "other", "agriculture", "education", "jobs", "services"}
	fmt.Println("\nPlease select a category")
	for i, dis := range categories {
		fmt.Println("\t", i+1, ". ", dis)
	}
	fmt.Print("Enter category: ")
	var category int
	_, err := fmt.Scanf("%d", &category)
	check(err)
	switch category {
	case 1:
		return categories[category-1]
	case 2:
		return categories[category-1]
	case 3:
		return categories[category-1]
	default:
		return ""
	}
}

func main() {
	fmt.Println("Web Scraper 2.0")
	fmt.Println("===============")

	district := getDistrict()
	category := getCategory()

	col := colly.NewCollector()

	col.OnHTML(".gtm-normal-ad", func(element *colly.HTMLElement) {
		fmt.Println("Data: {")
		fmt.Println("\tTitle: ", element.ChildText(".heading--2eONR"))
		fmt.Println("\tDescription: ", element.ChildText(".description--2-ez3"))
		fmt.Println("\tPrice: ", element.ChildText(".price--3SnqI"))
		fmt.Println("}")
	})

	col.OnRequest(func(r *colly.Request) {
		fmt.Println("\nVisiting", r.URL.String())
	})
	col.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
		fmt.Println("")
	})
	col.OnScraped(func(r *colly.Response) {
		fmt.Println("\nFinished", r.Request.URL)
	})
	_ = col.Visit("https://ikman.lk/en/ads/"+district+"/"+category)
}