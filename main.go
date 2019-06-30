package main

import (
	"gopkg.in/macaron.v1"
	"github.com/go-macaron/binding"
	"git.sr.ht/~humaid/shopsheet/models"
	"fmt"
	"math/rand"
	/*	"os"
		"fmt"
		"log"
		"git.sr.ht/~humaid/shopsheet/modules/spreadsheet"
		"git.sr.ht/~humaid/shopsheet/models"
	*/)

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
	m.Post("/new", binding.MultipartForm(models.NewForm{}), func(ctx *macaron.Context, form models.NewForm){
		if form.Spreadsheet != nil {
		    _, err := form.Spreadsheet.Open()
			if err != nil {
				ctx.PlainText(400, []byte("File uploaded cannot be opened"))
				return
			}
		}
		instID := fmt.Sprint(rand.Intn(89999) + 10000)
		shop := models.ShopSite{
			SiteTitle: form.SiteName,
			SiteDescription: form.SiteDescription,
			Email: form.Email,
			LogoURL: form.LogoURL,
			BannerURL: form.BannerURL,
		}
		fmt.Println(form.SiteName)
		models.Shops[instID] = shop

		ctx.Redirect(fmt.Sprintf("/%s", instID))
    })
	m.Group("/:site", func() {
		m.Get("/", func(ctx *macaron.Context) {
			shop := models.Shops[ctx.Params("site")]
			fillShopData(ctx, shop)
			ctx.HTML(200, "site/index")
		})
		m.Get("/:prod", func(ctx *macaron.Context) {
			shop := models.Shops[ctx.Params("site")]
			fillShopData(ctx, shop)
			ctx.HTML(200, "site/view_prod")
		})
		m.Get("/cart", func(ctx *macaron.Context) {
			shop := models.Shops[ctx.Params("site")]
			fillShopData(ctx, shop)
			ctx.HTML(200, "site/cart")
		})
	})
	m.Run()
}

func fillShopData(ctx *macaron.Context, shop models.Shop) {
	ctx.Data["Title"] = shop.SiteTitle
	ctx.Data["Description"] = shop.Description
	ctx.Data["LogoURL"] = shop.LogoURL
	ctx.Data["BannerURL"] = shop.BannerURL
	ctx.Data["Items"] = shop.Items
}
