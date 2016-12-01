package main

import (
	"GO-SPA/models"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/guardian/gocapiclient"
	"github.com/guardian/gocapiclient/queries"
	"github.com/kataras/iris"
	"github.com/valyala/fasthttp"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	//"golang.org/x/tools/benchmark/parse"
)

// Declaring variables for future use
var aid string
var body string
var url string

func main() {
	apistuff()

}

func apistuff() {
	
	// Create a new Instance of Iris
	api := iris.New()
	
	// Loads the default root page
	api.Get("/", search)

	render := func(ctx *iris.Context) {
		ctx.Render("index.html", nil)
	}

	// handler registration and naming
	api.Get("/index", render)("home")
	
	// Deals with Posting comments on the forum
	api.Post("/comments", commentHandler)
	
	// Used in conjunction with a Client side Ajax call, that retrieves comments and returns to front end
	api.Get("/getcomment", getComment)

	api.Build()
	fsrv := &fasthttp.Server{Handler: api.Router}
	fsrv.ListenAndServe(":8080")
	apistuff()
}

// A struct that holds details about whats being displayed on the Client Side 
type page struct {
	Title string
	Host  string
	JObj  string
	Text  string
}

// A struct that holds details from the Guardian news API
type GuardianAPI struct {
	id     string
	title  string
	weburl string
	apiurl string
	body   string
}

// adopted from https://github.com/guardian/gocapiclient
// this method is building query string and  getting back json object with result
// result is converting and transferring into front page
func searchQuery(client *gocapiclient.GuardianContentClient, g *GuardianAPI) {

	// NewSearchQuery used to create instance of query factory
	searchQuery := queries.NewSearchQuery()

	// fields filters query result
	showParam := queries.StringParam{"q", "tech%20AND%20technology"}
	showSection := queries.StringParam{"section", "technology"}
	showPages := queries.StringParam{"page", "1"}
	showPageSize := queries.StringParam{"page-seze", "1"}
	showOrderBy := queries.StringParam{"order-by", "newest"}
	showTotal := queries.StringParam{"total", "1"}
	showFields := queries.StringParam{"show-fields", "body"}
	params := []queries.Param{showPages, showPageSize, showOrderBy, showTotal, showFields, showParam, showSection}

	// assign parameters with values for the query fields
	searchQuery.Params = params

	// get result with json object
	err := client.GetResponse(searchQuery)
	if err != nil {
		log.Fatal(err)
	}

	// debug
	fmt.Println(searchQuery.Response.Status)
	fmt.Println(searchQuery.Response.Total)

	// run through the response and assign the result to the structured names
	for i, v := range searchQuery.Response.Results {
		if i == 0 {
			g.title = v.WebTitle
			g.weburl = v.WebUrl
			g.id = v.ID
			g.apiurl = v.ApiUrl
			g.body = *v.Fields.Body
			fmt.Println(i)
		}
		// debug result
		fmt.Println(v.ID)
		fmt.Println(v.WebTitle)
	}
	// Declaring string variable
	var articleID string

	// Setting variables to API values retrieved
	articleID = g.id
	aid = g.title
	url = g.weburl
	body = g.body

	// Initialize and devlare an array of structs
	// models are used to store structs
	comments := []models.Comment{}
	//
	s := getSession()
	// Declares the database name and collection name
	// that the connection should be made to
	c := s.DB("heroku_5r938bhv").C(aid)
	// Find and all will retrieve all comments in the
	// collection stated above
	erro := c.Find(bson.M{}).All(&comments)
	if erro != nil {
		log.Fatal(err)
	}

	// connect to the MongoDB and store the current article to the db
	s = getSession()
	_c := s.DB("heroku_5r938bhv").C("Article")
	errob := _c.Insert(&models.Article{articleID, aid, url})
	println("should be inserted")
	if errob != nil {
		println(articleID + " Has already exists")
	}
}

// This handlefunc getting the content from Guardian.com web source
// and rendering the result to the index.html page
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

//http://goinbigdata.com/how-to-build-microservice-with-mongodb-in-golang/
func getComment(ctx *iris.Context) {

	comments := []models.Comment{}
	s := getSession()
	c := s.DB("heroku_5r938bhv").C(aid)
	err := c.Find(bson.M{}).All(&comments)
	if err != nil {
		log.Fatal(err)
	}
	stringVal := &comments
	commentJSON, err := json.Marshal(stringVal)
	if err != nil {
		fmt.Println(err)
		return
	}
	ctx.Data(iris.StatusOK, commentJSON) // Once requested in a GET, sends the commentJSON to the client
	ctx.Next()
}

func commentHandler(ctx *iris.Context) {
	commentVal := ctx.FormValue("userComment")
	newComment := string(commentVal)
	newComment = newComment
	nano := time.Now()
	// establish session
	s := getSession()
	// declare database and collection
	c := s.DB("heroku_5r938bhv").C(aid)
	// insert into database using model (struct)
	errs := c.Insert(&models.Comment{nano.Format("20060102150405"), newComment})
	if errs != nil {
		log.Fatal(errs)
	}
	ctx.Render("index.html", page{aid, ctx.HostString(), body, url})
	client := gocapiclient.NewGuardianContentClient("https://content.guardianapis.com/", "b1b1f668-8a1f-40ec-af20-01687425695c")
	g := &GuardianAPI{}
	searchQuery(client, g)
}
