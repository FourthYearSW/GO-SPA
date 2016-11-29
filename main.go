package main

import (
	"encoding/json"
	"fmt"
	"log"
	//"encoding/json"

	"GO-SPA/models"

	"github.com/guardian/gocapiclient"
	"github.com/guardian/gocapiclient/queries"
	"github.com/kataras/iris"
	"github.com/valyala/fasthttp"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var id int
var oid int
var aid string
var newComment string

func main() {
	//uc := controllers.NewUserController(getSession())

	apistuff()

}
func apistuff() {
	api := iris.New()
	api.Get("/", search)
	//api.Get(Register)
	// Create User
	//api.Get("/user", uc.CreateUser)

	render := func(ctx *iris.Context) {
		ctx.Render("index.html", nil)
	}

	// handler registration and naming
	api.Get("/index", render)("home")

	api.Post("/root", commentHandler)
	//api.Get("/comment", newcomment)
	api.Get("/getcomment", getComment)

	api.Build()
	fsrv := &fasthttp.Server{Handler: api.Router}
	fsrv.ListenAndServe(":8080")
	apistuff()
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
	aid = g.title
	// url = g.weburl
	// body = g.body
	comments := []models.Comment{}
	s := getSession()
	c := s.DB("heroku_5r938bhv").C(aid)
	println("com collection found")
	erro := c.Find(bson.M{}).All(&comments)
	if erro != nil {
		log.Fatal(err)
	}

	println(len(comments))

	for i := 0; i < len(comments); i++ {
		println(comments[i].Comment)
	}
	//println(comments.Comment,comments.ID)

	id = len(comments)
	oid = id
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

// https://godoc.org/gopkg.in/mgo.v2#Bulk.Insert
/*func newcomment(ctx *iris.Context) {

	// establish session
	s := getSession()
	// declare database and collection
	c := s.DB("heroku_5r938bhv").C("com")
	// insert into database using model (struct)
	err := c.Insert(&models.Comment{id,newComment})
	if err != nil {
		log.Fatal(err)
	}
	id = id+1
	ctx.Next()
}*/

//http://goinbigdata.com/how-to-build-microservice-with-mongodb-in-golang/
func getComment(ctx *iris.Context) {

	comments := []models.Comment{}
	s := getSession()
	c := s.DB("heroku_5r938bhv").C(aid)
	println("com collection found")
	err := c.Find(bson.M{}).All(&comments)
	if err != nil {
		log.Fatal(err)
	}

	// For test purposes: Prints values retrieved from Heroku, to the console.
	for i := 0; i < len(comments); i++ {
		println(comments[i].ID, comments[i].Comment)
	}

	stringVal := &comments

	commentJSON, err := json.Marshal(stringVal)
	if err != nil {
		fmt.Println(err)
		return
	}

	println(string(commentJSON)) //  For testing purposes

	ctx.Data(iris.StatusOK, commentJSON) // Once requested in a GET, sends the commentJSON to the client

	oid = id
	ctx.Next()
}

func commentHandler(ctx *iris.Context) {
	commentVal := ctx.FormValue("userComment")

	newComment := string(commentVal)

	// establish session
	s := getSession()
	// declare database and collection
	c := s.DB("heroku_5r938bhv").C(aid)
	// insert into database using model (struct)
	err := c.Insert(&models.Comment{id, newComment})
	if err != nil {
		log.Fatal(err)
	}
	id = id + 1
	ctx.Write(string(newComment))
	ctx.ResetBody()
}
