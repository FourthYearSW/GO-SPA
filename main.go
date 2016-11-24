package main

import (
	"fmt"
	"log"

	"github.com/guardian/gocapiclient"
	"github.com/guardian/gocapiclient/queries"
	"github.com/kataras/iris"
	"github.com/valyala/fasthttp"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	//uc := controllers.NewUserController(getSession())

	api := iris.New()
	api.Get("/", search)

	// When posted, routes to here: The commentHandler function is initialised at the bottom
	api.Post("/comment", commentHandler)

	// Create User
	//api.Get("/user", uc.CreateUser)
	api.Get("/user", newuser)

	api.Build()
	fsrv := &fasthttp.Server{Handler: api.Router}
	fsrv.ListenAndServe(":9999")

	//iris.Listen(":9999")
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

// addopted from https://github.com/swhite24/go-rest-tutorial/blob/master/server.go
// getSession creates a new mongo session and panics if connection error occurs
func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("test:test@ds029585.mlab.com:29585/heroku_5r938bhv")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}

	// Deliver session
	return s
}

type User struct {
	id   string
	name string
}

// Possibly need to create a user to in order to add comments to that user, for MongoDB entries etc
type UserComment struct {
	Comment string `json:"name"`
}

func newuser(ctx *iris.Context) {
	s := getSession()
	c := s.DB("heroku_5r938bhv").C("testcollection")
	err := c.Insert(&User{"001", "Andrej"},
		&User{"002", "Christy"})
	if err != nil {
		log.Fatal(err)
	}

	result := User{}
	err = c.Find(bson.M{"id": "001"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("name: ", result.name)
	ctx.Next()
}

// Takes the POST parameter from the form, saves it in the a local variable and writes back out to screen
func commentHandler(ctx *iris.Context) {
	comment := ctx.FormValue("userComment")
	ctx.Write("%s", comment)
}
