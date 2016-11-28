package main

import (
	"fmt"
	"log"
	//"encoding/json"

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

// Adapted from https://github.com/swhite24/go-rest-tutorial/blob/master/models/user.go
type (
	// Comment is a struct which holds details about the articles and can be marshalled into json and bson
	Comment struct {
		ID      int    `json:"id" bson:"_id"`
		Comment string `json:"comment" bson:"comment"`
	}
)

// Adapted from https://github.com/swhite24/go-rest-tutorial/blob/master/models/user.go
type (
	// Article is a struct which holds details about the articles and can be marshalled into json and bson
	Article struct {
		ID   string `json:"id" bson:"_id"`
		Name string `json:"name" bson:"name"`
		URL  string `json:"url" bson:"url"`
	}
)

type (
	// User represents the structure of our resource
	User struct {
		Id     bson.ObjectId `json:"id" bson:"_id"`
		Name   string        `json:"name" bson:"name"`
		Gender string        `json:"gender" bson:"gender"`
		Age    int           `json:"age" bson:"age"`
	}
)

type (
	// User represents the structure of our resource
	Theuser struct {
		Id       bson.ObjectId `json:"id" bson:"_id"`
		Name     string        `json:"name" bson:"name"`
		Email    string        `json:"email" bson:"email"`
		Password string        `json:"password" bson:"password"`
	}
)

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

	comments := []Comment{}
	s := getSession()
	c := s.DB("heroku_5r938bhv").C(aid)
	println("com collection found")
	erro := c.Find(bson.M{}).All(&comments)
	if erro != nil {
		log.Fatal(err)
	}

	println(len(comments))
	id = len(comments)

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
	err := c.Insert(&Comment{id,newComment})
	if err != nil {
		log.Fatal(err)
	}
	id = id+1
	ctx.Next()
}*/

//http://goinbigdata.com/how-to-build-microservice-with-mongodb-in-golang/
func getComment(ctx *iris.Context) {

	comments := Comment{}
	s := getSession()
	c := s.DB("heroku_5r938bhv").C("com")
	println("com collection found")

	println("THIS FIRES BEFORE ERROR IN GET REQUEST....")

	err := c.Find(bson.M{"_id": oid}).One(&comments)
	if err != nil {
		log.Fatal(err)
	}

	println(comments.Comment, comments.ID)
	ctx.Write(comments.Comment, comments.ID)

	oid = id
	ctx.Next()
}

func commentHandler(ctx *iris.Context) {
	commentVal := ctx.FormValue("userComment")

	newComment := string(commentVal)

	// newComment = newComment

	// establish session
	s := getSession()
	// declare database and collection
	c := s.DB("heroku_5r938bhv").C(aid)
	// insert into database using model (struct)

	err := c.Insert(&Comment{id, newComment})
	if err != nil {
		log.Fatal(err)
	}
	id = id + 1

	ctx.Write(string(newComment))
	ctx.ResetBody()

}
