package controllers

import (
	"github.com/guardian/gocapiclient/queries"
	"log"
	"fmt"
	"github.com/guardian/gocapiclient"
	"github.com/kataras/iris"
)

// Structure used to maintain the contend retrieved from www.guardian.com
type GuardianContent struct{
	id string
	title string
	weburl string
	apiurl string
	body string
}

// Page structure used for rendering (binding) content on the page
type page struct{
	Title string
	Host string
	JObj string
	Text string
}

// Was adopted from https://content.guardianapis.com/ and modified
// This function is used to retrieve the json object
// that contains header metadata such as title, id, urls, state, status and so on
// and body which is html code with article content
func searchQuery(client *gocapiclient.GuardianContentClient, g *GuardianContent) {
	// instance of query interface
	searchQuery := queries.NewSearchQuery()

	// Filter parameters. Normally that is after '?' in uri string
	//showParam := queries.StringParam{"q", "tech%20AND%20technology"}
	//showSection := queries.StringParam{"section", "technology"}
	showPages := queries.StringParam{"page", "1"}
	showPageSize := queries.StringParam{"page-seze", "1"}
	showOrderBy := queries.StringParam{"order-by", "newest"}
	showTotal := queries.StringParam{"total", "1"}
	showFields := queries.StringParam{"show-fields", "body"}

	// Gathering parameters together in to array
	params := []queries.Param{showPages, showPageSize, showOrderBy, showTotal, showFields}

	// Parsing parameters to search query
	searchQuery.Params = params

	// getting content from the source: www.guardian.com
	err := client.GetResponse(searchQuery)
	if err != nil {
		log.Fatal(err)
	}

	// console printing of state
	fmt.Println(searchQuery.Response.Status)
	fmt.Println(searchQuery.Response.Total)

	// retrieving json values and assigning into associated struct
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

// This part is responsible for handling page routing
// and rendering data (structs, values, names) on the page

// Interface declaring
type GuardianAPI struct{}

// Singlton pattern, will be called in main to get an interface
func NewGuardianAPI() *GuardianAPI{
	return &GuardianAPI{}
}

// routing page and rendering stuff on it
func (ga GuardianAPI)Search(c *iris.Context){
	client := gocapiclient.NewGuardianContentClient("https://content.guardianapis.com/", "b1b1f668-8a1f-40ec-af20-01687425695c")
	g := &GuardianContent{}
	searchQuery(client, g)

	c.Render("index.html", page{g.title, c.HostString(), g.body, g.weburl})
}
