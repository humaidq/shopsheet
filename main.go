package main

import (
	"gopkg.in/macaron.v1"
/*	"os"
	"fmt"
	"log"
	"git.sr.ht/~humaid/shopsheet/modules/spreadsheet"
	"git.sr.ht/~humaid/shopsheet/models"
	*/
)

func main() {
	/*wd, _ := os.Getwd()
	var csv string
	var err error
	csv, err = spreadsheet.GetCSVFromSpreadsheet(fmt.Sprintf("%s/test.xlsx", wd))
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(csv)
	fmt.Println("-----------")

	var items []models.ShopItem
	items, err = models.GetShopItemsFromCSV(csv)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(items)

	os.Exit(1)*/
	m := macaron.Classic()
	m.Use(macaron.Renderer())
	m.Get("/", func(ctx *macaron.Context) {
        ctx.HTML(200, "index")
    })

	m.Get("/new", func(ctx *macaron.Context) {
        ctx.HTML(200, "new")
    })
	m.Post("/new", func(ctx *macaron.Context) {
        ctx.HTML(200, "upload")
    })
	m.Group("/:site", func() {
		m.Get("/", func(ctx *macaron.Context) {
			ctx.HTML(200, "site/index")
		})
		m.Get("/:prod", func(ctx *macaron.Context) {
			ctx.HTML(200, "site/view_prod")
		})
		m.Get("/cart", func(ctx *macaron.Context) {
			ctx.HTML(200, "site/cart")
		})
	})
	m.Run()
}
