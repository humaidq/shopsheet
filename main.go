package main

import (
	"bytes"
	"fmt"
	"git.sr.ht/~humaid/shopsheet/models"
	"git.sr.ht/~humaid/shopsheet/modules/spreadsheet"
	"github.com/go-macaron/binding"
	"gopkg.in/macaron.v1"
	"io/ioutil"
	"math/rand"
)

func main() {
	m := macaron.Classic()
	m.Use(macaron.Renderer())
	m.Get("/", func(ctx *macaron.Context) {
		ctx.HTML(200, "index")
	})

	m.Get("/new", func(ctx *macaron.Context) {
		ctx.HTML(200, "new")
	})
	m.Post("/new", binding.MultipartForm(models.NewForm{}), func(ctx *macaron.Context, form models.NewForm, errs binding.Errors) {
		if len(errs) > 0 {
			ctx.PlainText(400, []byte(fmt.Sprintf("Form binding error: %s", errs)))
			return
		}

		f, err := form.Spreadsheet.Open()
		if err != nil {
			ctx.PlainText(400, []byte("File uploaded cannot be opened"))
			return
		}

		dir, _ := ioutil.TempDir("", "shopsheet")
		buf := new(bytes.Buffer)
		buf.ReadFrom(f)
		ssPath := fmt.Sprintf("%s/%s", dir, form.Spreadsheet.Filename)
		ioutil.WriteFile(ssPath, buf.Bytes(), 0644)

		fmt.Println(ssPath)
		var csv string
		csv, err = spreadsheet.GetCSVFromSpreadsheet(ssPath, dir)
		if err != nil {
			ctx.PlainText(400, []byte(fmt.Sprintf("Failed to convert into CSV file: %s", err)))
			return
		}

		var items []models.ShopItem
		items, err = models.GetShopItemsFromCSV(csv)
		if err != nil {
			ctx.PlainText(400, []byte(fmt.Sprintf("Failed to parse spreadsheet file: %s", err)))
			return
		}

		instID := fmt.Sprint(rand.Intn(89999) + 10000)
		shop := models.ShopSite{
			SiteTitle:       form.SiteName,
			SiteDescription: form.SiteDescription,
			Email:           form.Email,
			LogoURL:         form.LogoURL,
			BannerURL:       form.BannerURL,
			Items:           items,
		}
		fmt.Println(form.SiteName)
		models.Shops[instID] = shop

		ctx.Redirect(fmt.Sprintf("/%s", instID))
	})

	m.Group("/:site", func() {
		m.Get("/", func(ctx *macaron.Context) {
			shop, ok := models.Shops[ctx.Params("site")]
			if !ok {
				ctx.PlainText(404, []byte("Site instance does not exist!"))
				return
			}
			fillShopData(ctx, shop, ctx.Params("site"))
			ctx.Data["IsIndex"] = 1
			ctx.HTML(200, "site/index")
		})
		m.Get("/cart", func(ctx *macaron.Context) {
			shop, ok := models.Shops[ctx.Params("site")]
			if !ok {
				ctx.PlainText(404, []byte("Site instance does not exist!"))
				return
			}
			fillShopData(ctx, shop, ctx.Params("site"))
			ctx.Data["IsCart"] = 1
			ctx.HTML(200, "site/cart")
		})
	})
	m.Run()
}

func fillShopData(ctx *macaron.Context, shop models.ShopSite, inst string) {
	ctx.Data["Title"] = shop.SiteTitle
	ctx.Data["Description"] = shop.SiteDescription
	ctx.Data["LogoURL"] = shop.LogoURL
	ctx.Data["BannerURL"] = shop.BannerURL
	ctx.Data["Email"] = shop.Email
	ctx.Data["Items"] = shop.Items
	ctx.Data["Inst"] = inst
}
