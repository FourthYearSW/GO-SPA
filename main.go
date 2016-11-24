package main

import (
	"fmt"
	"log"

	"github.com/guardian/gocapiclient"
	"github.com/guardian/gocapiclient/queries"
	"github.com/kataras/iris"
	"github.com/valyala/fasthttp"
)

func main() {
	api := iris.New()
	api.Get("/", search)

	api.Build()
	fsrv := &fasthttp.Server{Handler: api.Router}
	fsrv.ListenAndServe(":9999")
}

type page struct {
	Title string
	Host  string
	JObj  string
	Text  string
}

type GuardianAPI struct {
	id     string
	title  string
	weburl string
	apiurl string
	body   string
}

type User struct {
	name     string
	username string
	email    string
	password string
}

func searchQuery(client *gocapiclient.GuardianContentClient, g *GuardianAPI) {
	searchQuery := queries.NewSearchQuery()

	//showParam := queries.StringParam{"q", "tech%20AND%20technology"}
	//showSection := queries.StringParam{"section", "technology"}
	showPages := queries.StringParam{"page", "1"}
	showPageSize := queries.StringParam{"page-seze", "1"}
	showOrderBy := queries.StringParam{"order-by", "newest"}
	showTotal := queries.StringParam{"total", "1"}
	showFields := queries.StringParam{"show-fields", "body"}
	params := []queries.Param{showPages, showPageSize, showOrderBy, showTotal, showFields}

	searchQuery.Params = params

	err := client.GetResponse(searchQuery)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(searchQuery.Response.Status)
	fmt.Println(searchQuery.Response.Total)

	for i, v := range searchQuery.Response.Results {
		if i == 0 {
			g.title = v.WebTitle
			g.weburl = v.WebUrl
			g.id = v.ID
			g.apiurl = v.ApiUrl
			g.body = *v.Fields.Body
			fmt.Println(i)
		}
		fmt.Println(v.ID)
		fmt.Println(v.WebTitle)
	}
}

func search(ctx *iris.Context) {
	client := gocapiclient.NewGuardianContentClient("https://content.guardianapis.com/", "b1b1f668-8a1f-40ec-af20-01687425695c")
	g := &GuardianAPI{}
	searchQuery(client, g)

	ctx.Render("index.html", page{g.title, ctx.HostString(), g.body, g.weburl})
}
