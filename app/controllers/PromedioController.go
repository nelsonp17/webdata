package controllers

import (
    "encoding/json"
    "fmt"
    "github.com/gocolly/colly"
    "io/ioutil"
    "log"
    "net/http"
    "strconv"
    "strings"

    //"github.com/Nelson2017-8/webdata/promedio"
)




type Promedio struct {
    name string `json:"name"`
    price string `json:"price"`
    id int `json:"id"`
    update string `json:"update"`
    imgTitle string `json:"imgTitle"`
    imgSrc string `json:"imgSrc"`
    nameAlt string `json:"nameAlt"`
} 

type PromedioInter interface{
    NameAltPromedio(c *Promedio) bool
    UpdatePromedio(c* Promedio) bool
}
func (c *Promedio) UpdatePromedio(small string) bool {
    if small == "www.monitordolarvenezuela.com" {
        return false
    }
    c.price = small
    return true
}
func (c *Promedio) NameAltPromedio(nameAlt string) bool {
    switch nameAlt {
    case "BCV (Oficial)":
        c.name = "BCV"
    case "@EnParaleloVzla3":
        c.name = "Paralelo"
    case "@DolarToday":
        c.name = "DolarToday"
    case "@MonitorDolarWeb":
        c.name = "Monitor Dolar"
    case "@EnParaleloVzlaVip":
        c.name = "ParaleloVIP"
    case "Binance P2P":
        c.name = "BinanceP2P"
    case "AIRTM":
        c.name = "AIRTM"
    default:
        return false
    }
    return true
}


type pageData struct {
    StatusCode int `json:"status_code"`
    Data      map[string]interface{} `json:"data"`
}
type Array struct {
    array map[string]interface{}
}

func PromedioView(w http.ResponseWriter, r *http.Request) {
    c := colly.NewCollector()
    //var promedios Promedio[max]
    count := 0

    //var promedios []*Promedio
    // On every a element which has href attribute call callback
    //
    p := &pageData{Data: make(map[string]interface{})}
    c.OnHTML("#promedios .row div.col-lg-2 div", func(e *colly.HTMLElement) {
        if(count < 7){

            imgTitle := e.ChildAttr("img", "alt")
            srcTitle := e.ChildAttr("img", "src")
            title := e.ChildText("h4.title-prome")
            update := e.ChildText("small")
            price := e.ChildText("p")


            // filtrando
            price = strings.Replace(price, "Bs = ", "", 1)
            price2 := strings.Replace(price, ",", ".", 1)
            update = strings.Replace(update, "www.monitordolarvenezuela.comactualizado: ", "", 1)
            update = strings.Replace(update, "PM", "PM ", 1)
            update = strings.Replace(update, "AM", "AM ", 1)

            fmt.Println(price)
            fmt.Println(update)

            promedio := new(Promedio)
            promedio.id = count +1
            promedio.imgTitle = imgTitle
            promedio.imgSrc = srcTitle
            promedio.price = price
            promedio.NameAltPromedio(title)
            promedio.UpdatePromedio(update)


            priceF, _ := strconv.ParseFloat(price2, 2)
            var items = struct {
                Precio float64 `json:"precio"`
                PrecioStr string `json:"precio_str"`
                Title  string  `json:"title"`
                UpdatedAt string `json:"updated_at"`
            }{priceF, price, title, update}

            p.Data[promedio.name] = items

        }
        count += 1
        //link := e.Attr("href")

        //fmt.Printf("Link found: %q -> %s\n", e.Text, link)
        //c.Visit(e.Request.AbsoluteURL(link))
    })

    // extract status code
    c.OnResponse(func(r *colly.Response) {
        log.Println("response received", r.StatusCode)
        p.StatusCode = 200
    })
    c.OnError(func(r *colly.Response, err error) {
        log.Println("error:", r.StatusCode, err)
        p.StatusCode = 400
    })


    // Start scraping on https://hackerspaces.org
    c.Visit("https://monitordolarvenezuela.com/")
    //fmt.Println(p)


    // dump results
    b, err := json.Marshal(p)
    if err != nil {
       log.Println("failed to serialize response:", err)
       return
    }
    // Crear un JSON con los datos
    err = ioutil.WriteFile("data/data.json", b, 0644)

    w.Header().Add("Content-Type", "application/json")
    w.Write(b)

    //fmt.Println(b)


}