package api

import (
	"errors"
	"github.com/gocolly/colly"
	"github.com/nelsonp17/webdata/app/constant"
	"github.com/nelsonp17/webdata/app/database/sqlc"
	"strconv"
	"strings"
	"time"

	"context"
	"fmt"

	"github.com/antchfx/htmlquery"
	"github.com/chromedp/chromedp"
	"github.com/gofiber/fiber/v2"
	"github.com/nelsonp17/webdata/app/database/sqlc/schemas"
)

type Handler struct {
	Repo sqlc.Repo
}

const MonitorDolarVenezuela = "https://monitordolarvenezuela.com/"
const BCV = "https://www.bcv.org.ve/"

func ScrapingMonitorDolar() ([]schemas.History, error) {

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var htmlContent string
	err := chromedp.Run(ctx,
		chromedp.Navigate(MonitorDolarVenezuela),    // Reemplaza con la URL
		chromedp.WaitReady(`.col-span-3.undefined`), // Espera a que un elemento específico se cargue (opcional, pero recomendado)
		chromedp.OuterHTML(`html`, &htmlContent),    // Obtiene el HTML renderizado
	)
	if err != nil {
		fmt.Println("failed to serialize response:", err)
	}

	var p []schemas.History
	doc, err := htmlquery.Parse(strings.NewReader(htmlContent))
	if err != nil {
		fmt.Println("failed to parse HTML:", err)
		return nil, err
	}
	// encuenta todos los elementos div con la clase .col-span-3.undefined
	list := htmlquery.Find(doc, "//div[contains(@class, 'col-span-3') and contains(@class, 'undefined')]")
	for _, n := range list {
		name := htmlquery.FindOne(n, "//h3").FirstChild.Data
		// img := htmlquery.FindOne(n, "//img").Attr[0].Val
		priceString := htmlquery.FindOne(n, "//p[contains(@class, 'font-bold') and contains(@class, 'text-xl')]").FirstChild.Data
		priceString = strings.Replace(priceString, "Bs = ", "", 1)

		name = strings.Replace(name, "(Oficial)", "", -1)
		name = strings.Replace(name, "@EnParaleloVzla3", "Paralelo", -1)
		name = strings.Replace(name, "@EnParaleloVzlaVIP", "ParaleloVIP", -1)

		priceString = strings.Replace(priceString, ",", ".", -1)
		price, err := strconv.ParseFloat(priceString, 64)
		if err != nil {
			fmt.Println("failed to convert price string to float:", err)
			return nil, err
		}

		if price > 0 && name != "" {
			p = append(p, schemas.History{
				Money:     "USD",
				Source:    name,
				SourceWeb: MonitorDolarVenezuela,
				Change:    price,
			})
		}
	}
	return p, nil
}
func ScrapingBcv() ([]schemas.History, error) {

	c := colly.NewCollector()
	p := &[]schemas.History{}

	c.OnHTML("#dolar", func(e *colly.HTMLElement) {
		name := "BCV"
		price := e.ChildText("strong")
		price = strings.Replace(price, ",", ".", -1)
		priceFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("failed to convert price string to float:", err)
			return
		}
		item := schemas.History{
			Money:     "USD",
			Source:    name,
			SourceWeb: BCV,
			Change:    priceFloat,
		}
		if item.Change > 0 && item.Source != "" {
			*p = append(*p, item)
		}
	})
	c.OnHTML("#euro", func(e *colly.HTMLElement) {
		name := "BCV"
		price := e.ChildText("strong")
		price = strings.Replace(price, ",", ".", -1)
		priceFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("failed to convert price string to float:", err)
			return
		}
		item := schemas.History{
			Money:     "EUR",
			Source:    name,
			SourceWeb: BCV,
			Change:    priceFloat,
		}
		if item.Change > 0 && item.Source != "" {
			*p = append(*p, item)
		}
	})
	c.OnHTML("#yuan", func(e *colly.HTMLElement) {
		name := "BCV"
		price := e.ChildText("strong")
		price = strings.Replace(price, ",", ".", -1)
		priceFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("failed to convert price string to float:", err)
			return
		}
		item := schemas.History{
			Money:     "CNY",
			Source:    name,
			SourceWeb: BCV,
			Change:    priceFloat,
		}
		if item.Change > 0 && item.Source != "" {
			*p = append(*p, item)
		}
	})
	c.OnHTML("#lira", func(e *colly.HTMLElement) {
		name := "BCV"
		price := e.ChildText("strong")
		price = strings.Replace(price, ",", ".", -1)
		priceFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("failed to convert price string to float:", err)
			return
		}
		item := schemas.History{
			Money:     "TRY",
			Source:    name,
			SourceWeb: BCV,
			Change:    priceFloat,
		}
		if item.Change > 0 && item.Source != "" {
			*p = append(*p, item)
		}
	})
	c.OnHTML("#rublo", func(e *colly.HTMLElement) {
		name := "BCV"
		price := e.ChildText("strong")
		price = strings.Replace(price, ",", ".", -1)
		priceFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("failed to convert price string to float:", err)
			return
		}
		item := schemas.History{
			Money:     "RUB",
			Source:    name,
			SourceWeb: BCV,
			Change:    priceFloat,
		}
		if item.Change > 0 && item.Source != "" {
			*p = append(*p, item)
		}
	})

	// extract status code
	err := c.Visit("https://www.bcv.org.ve/")
	if err != nil {
		return nil, err
	}
	return *p, nil
}

//func GetDatabaseMonitorDolar(repo sqlc.Repo) ([]schemas.History, error) {
//	// var histories []schemas.History
//	// bcv, _ := repo.GetLastHistory("USD", "BCV", MonitorDolarVenezuela)
//	// binance, _ := repo.GetLastHistory("USD", "Binance", MonitorDolarVenezuela)
//	// paralelo, _ := repo.GetLastHistory("USD", "Paralelo", MonitorDolarVenezuela)
//	// paraleloVIP, _ := repo.GetLastHistory("USD", "ParaleloVIP", MonitorDolarVenezuela)
//	// dolarToday, _ := repo.GetLastHistory("USD", "DólarToday", MonitorDolarVenezuela)
//	// monitorDolarWeb, _ := repo.GetLastHistory("USD", "MonitorDolarWeb", MonitorDolarVenezuela)
//	// histories = append(histories, bcv, binance, paralelo, paraleloVIP, dolarToday, monitorDolarWeb)
//	// return histories, nil
//
//	return repo.ListHistory(MonitorDolarVenezuela, time.Now(), 9)
//}
//func GetDatabaseBcv(repo sqlc.Repo) ([]schemas.History, error) {
//	//var histories []schemas.History
//	//bcv, _ := repo.GetLastHistory("USD", "BCV", BCV)
//	//eur, _ := repo.GetLastHistory("EUR", "BCV", BCV)
//	//cny, _ := repo.GetLastHistory("CNY", "BCV", BCV)
//	//try, _ := repo.GetLastHistory("TRY", "BCV", BCV)
//	//rub, _ := repo.GetLastHistory("RUB", "BCV", BCV)
//	//histories = append(histories, bcv, eur, cny, try, rub)
//	//return histories, nil
//	return repo.ListHistory(BCV, time.Now(), 5)
//}

func GetMonitorDolar(repo sqlc.Repo) ([]schemas.History, error) {

	histories, err := repo.ListHistory(MonitorDolarVenezuela, time.Now(), 9)

	if err != nil {
		fmt.Println("error:", err)
		return histories, errors.New("error: getting data from database")
	}

	if histories == nil || len(histories) == 0 {
		fmt.Println("Histories MonitorDolar vencio")
		// Llamar a ScrapingMonitorDolar() para obtener los nuevos datos
		newData, err := ScrapingMonitorDolar()
		if err != nil {
			fmt.Println("error:", err)
			return histories, errors.New("error: scraping monitor dolar")
		}

		// Insertar los nuevos registros en la tabla history
		for _, data := range newData {
			history, err := repo.CreateHistory(data.Money, data.Source, data.SourceWeb, data.Change)
			if err != nil {
				return histories, errors.New("error: inserting new data")
			}
			histories = append(histories, history)
		}
		return histories, nil
	}

	return histories, nil
}
func GetBcv(repo sqlc.Repo) ([]schemas.History, error) {
	histories, err := repo.ListHistory(BCV, time.Now(), 5)

	if err != nil {
		fmt.Println("error:", err)
		return histories, errors.New("error: getting data from database")
	}

	// Verificar si el último registro tiene más de 8 horas de antigüedad
	if histories == nil || len(histories) == 0 {
		fmt.Println("Histories BCV vencio")
		// Llamar a ScrapingMonitorDolar() para obtener los nuevos datos
		newData, err := ScrapingBcv()
		if err != nil {
			fmt.Println("error:", err)
			return histories, errors.New("error: scraping bcv")
		}

		// Insertar los nuevos registros en la tabla history
		for _, data := range newData {
			history, err := repo.CreateHistory(data.Money, data.Source, data.SourceWeb, data.Change)
			if err != nil {
				return histories, errors.New("error: inserting new data")
			}
			histories = append(histories, history)
		}
		return histories, nil
	}

	return histories, nil
}

func (h *Handler) GetPriceDollar(c *fiber.Ctx) error {
	// Configurar chromedp
	monitorDolar, err := GetMonitorDolar(h.Repo)
	if err != nil {
		fmt.Println("failed to serialize response monitorDolar:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("failed to serialize response monitorDolar")
	}
	bcv, err := GetBcv(h.Repo)
	if err != nil {
		fmt.Println("failed to serialize response bcv:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("failed to serialize response bcv")
	}

	var clearMonitorDolar []schemas.History
	var clearBcv []schemas.History
	for _, data := range monitorDolar {
		if data.ID > 0 {
			clearMonitorDolar = append(clearMonitorDolar, data)
		}
	}
	for _, data := range bcv {
		if data.ID > 0 {
			clearBcv = append(clearBcv, data)
		}
	}
	response := constant.Response{
		Data: fiber.Map{
			"MonitorDolar": clearMonitorDolar,
			"BCV":          clearBcv,
		},
	}

	return c.JSON(response)
}
