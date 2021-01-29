package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocolly/colly"
	"os"
)

func check(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func getDistrict() string {
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
		return districts[district-1]
	case 3:
		return districts[district-1]
	case 4:
		return districts[district-1]
	case 5:
		return districts[district-1]
	case 6:
		return districts[district-1]
	case 7:
		return districts[district-1]
	case 8:
		return districts[district-1]
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
	case 4:
		return categories[category-1]
	case 5:
		return categories[category-1]
	case 6:
		return categories[category-1]
	case 7:
		return categories[category-1]
	case 8:
		return categories[category-1]
	case 9:
		return categories[category-1]
	default:
		return ""
	}
}

func main() {
	db, _ := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/WebScraper")
	fmt.Println("Web Scraper 2.0")
	fmt.Println("===============")

	district := getDistrict()
	if district == "" {
		fmt.Println("Invalid Choice!")
		os.Exit(0)
	}
	category := getCategory()
	if category == "" {
		fmt.Println("Invalid Choice!")
		os.Exit(0)
	}

	col := colly.NewCollector()

	var (
		title           string
		description     string
		price           string
		url             string
		postedOn        string
		forSaleBy       string
		meta            string
		fullDescription string
	)

	col.OnHTML(".gtm-normal-ad", func(element *colly.HTMLElement) {
		title = element.ChildText(".heading--2eONR")
		description = element.ChildText(".description--2-ez3")
		price = element.ChildText(".price--3SnqI")
		url = element.ChildAttr(".card-link--3ssYv", "href")
		fmt.Println("Data: {")
		fmt.Println("\tTitle: ", title)
		fmt.Println("\tDescription: ", description)
		fmt.Println("\tPrice: ", price)
		fmt.Println("\tURL: ", url)
		err := element.Request.Visit(url)
		check(err)
	})

	col.OnHTML(".sub-title--37mkY", func(element *colly.HTMLElement) {
		postedOn = element.Text
		fmt.Println("\tPosted On: ", postedOn)
	})

	col.OnHTML(".contact-name--m97Sb", func(element *colly.HTMLElement) {
		forSaleBy = element.Text
		fmt.Println("\tFor Sale By: ", forSaleBy)
	})

	col.OnHTML(".ad-meta--17Bqm", func(element *colly.HTMLElement) {
		meta = element.Text
		fmt.Println("\tMeta: ", meta)
	})

	col.OnHTML(".description-section--oR57b > div > .description--1nRbz", func(element *colly.HTMLElement) {
		fullDescription = element.Text
		fmt.Println("\tFull Description: ", fullDescription)
		fmt.Println("}")
	})

	col.OnRequest(func(r *colly.Request) {
		//fmt.Println("\nVisiting", r.URL.String())
	})
	col.OnResponse(func(r *colly.Response) {
		//fmt.Println("Visited", r.Request.URL)
		//fmt.Println("")
	})
	col.OnScraped(func(r *colly.Response) {
		//fmt.Println("\nFinished", r.Request.URL)
		insert, err := db.Query("INSERT INTO Advertisement (district, category, title, description, price, url, postedOn, forSaleBy, meta, fullDescription) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
			district, category, title, description, price, url, postedOn, forSaleBy, meta, fullDescription)
		check(err)
		defer insert.Close()
	})
	_ = col.Visit("https://ikman.lk/en/ads/" + district + "/" + category)
}
