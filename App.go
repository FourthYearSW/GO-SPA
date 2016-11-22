package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"encoding/json"
	"io/ioutil"
	"github.com/valyala/fasthttp"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/kataras/iris"
	"github.com/guardian/gocapiclient"
	"github.com/guardian/gocapiclient/queries"
)

// all session code adapted from http://www.gorillatoolkit.org/pkg/sessions
var store = sessions.NewCookieStore([]byte("secret"))
var tpl *template.Template
var mongoConnection, err = newMongoConnection()

func init() {
	tpl = template.Must(template.ParseGlob("public/templates/index.html"))
}

func main() {

	port := os.Getenv("PORT")

	r := mux.NewRouter()
	r.HandleFunc("/", display)
	r.HandleFunc("/register", Register)
	r.HandleFunc("/login", loginHandler)
	http.Handle("/", r)
	http.HandleFunc("/css/", serveResource)

	if port == "" {
		http.ListenAndServe(":4000", nil)
	} else {
		http.ListenAndServe(":"+port, nil)
	}

	// Andrej
		api := iris.New()

	api.Static("/*", "./public/*", 1)

	api.Get("/", search)

	api.Get("/mypath", func(ctx *iris.Context){
		ctx.Write("Hello from the server on path /mypath")
	})

	api.HandleFunc("GET", "/get", myhandler)

	api.API("/users", UserApi{}, myUsersMiddleware1, myUsersMiddleware2)

	api.API("/redirect", HackerNews{}, myUsersMiddleware1, myUsersMiddleware2)

	api.Get("/search", getpage)



	//client := gocapiclient.NewGuardianContentClient("https://content.guardianapis.com/", "b1b1f668-8a1f-40ec-af20-01687425695c")
	//searchQuery(client, GuardianAPI{})

	//searchQueryPaged(client)
	//itemQuery(client)

	// Handler API


	// to use a custom server you have to call .Build after
	// route, sessions, templates, websockets, ssh... before server's listen
	api.Build()

	/*
	ln, err := net.Listen("tcp4", "0.0.0.0:9999")
	if err != nil{
		panic(err)
	}

	iris.Serve(ln)
	*/

	// create our custom fasthttp server and assign the Handler/Router
	fsrv := &fasthttp.Server{Handler: api.Router}
	fsrv.ListenAndServe(":9999")
	// Andrej

}
// Andrej

type page struct{
	Title string
	Host string
	JObj string
	Text string
}

type GuardianAPI struct{
	id string
	title string
	weburl string
	apiurl string
	body string
}

type JsonObj struct{
	by string
	descendants string
	id int
	kids []int
	score int
	text string
	time int
	title string
	Type string
	url string
}

func getpage(ctx *iris.Context){

	// Retrieving json object from HackerNews
	resp, err := http.Get("https://hacker-news.firebaseio.com/v0/item/160705.json?print=pretty")
	if err != nil{panic(err.Error())}
	body, err := ioutil.ReadAll(resp.Body)
	jsonstring := fmt.Sprintf("%s", body)
	var f interface{}
	json.Unmarshal(body, &f)

	m := f.(map[string]interface{})

	println("m is -> ", m)

	j := JsonObj{}
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
			if k == "title" {j.title = vv}
			if k == "type" {j.Type = vv}
			if k == "by" {j.by = vv}
			if k == "descendants" {j.descendants = vv}
			if k == "text" {j.text = vv}
			if k == "url" {j.url = vv}
		case int:
			fmt.Println(k, "is int", vv)
			if k == "score" {j.score = vv}
			if k == "time" {j.time = vv}
			if k == "id" {j.id = vv}
		case []interface{}:
			fmt.Println(k, "is an array:")
			for u := range vv {
				fmt.Println(u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle", vv)
		}
	}

	println("Print something")
	println(j.title)
	for u := range j.kids{
		fmt.Println("j.kids -> ", u)
	}
	ctx.Next()

	ctx.Render("index.html", page{j.title, ctx.HostString(), jsonstring, j.text})
}

func myhandler(c *iris.Context){
	c.Write("From %s. Implementation of handlerFunction", c.PathString())
}

type UserApi struct{*iris.Context}
// GET /users
func (u UserApi) Get(){
	//u.Write("Get from /users")
	u.HTML(iris.StatusOK, "<h3>Get all from users</h3>")
	//u.Redirect("https://hacker-news.firebaseio.com/v0/item/121003.json?print=pretty", iris.StatusOK)
}

func myUsersMiddleware1(ctx *iris.Context){
	println("From User middleware 1")
	ctx.Next()
}

func myUsersMiddleware2(ctx *iris.Context){
	println("From User middleware 2")
	ctx.Next()
}

// Retrieving json object from HackerNews and printing it into page
// 1) initialize structure with
type HackerNews struct{*iris.Context}
func (u HackerNews) Get(){
	//u.Write("Get from /users")
	//u.HTML(iris.StatusOK, "<h3>Get all from users</h3>")
	//u.Redirect("https://hacker-news.firebaseio.com/v0/item/121003.json?print=pretty", iris.StatusOK)
	//u.Request.SetRequestURI("https://hacker-news.firebaseio.com/v0/item/121003.json?print=pretty")
	resp, err := http.Get("http://content.guardianapis.com/search?q=tech%20AND%20technology&section=technology&page=1&page-size=1&order-by=newest&api-key=b1b1f668-8a1f-40ec-af20-01687425695c")
	if err != nil{panic(err.Error())}
	body, err := ioutil.ReadAll(resp.Body)
	u.Write("%s", body)
}

/*
func searchQueryPaged(client *gocapiclient.GuardianContentClient) {
	searchQuery := queries.NewSearchQuery()
	searchQuery.PageOffset = int64(10)

	showParam := queries.StringParam{"q", "tech%20AND%20technology"}
	showSection := queries.StringParam{"section", "technology"}
	showPages := queries.StringParam{"page", "1"}
	showPageSize := queries.StringParam{"page-seze", "1"}
	showOrderBy := queries.StringParam{"order-by", "newest"}
	params := []queries.Param{&showParam, showSection, showPages, showPageSize, showOrderBy}

	searchQuery.Params = params

	iterator := client.SearchQueryIterator(searchQuery)

	for page := range iterator {
		fmt.Println("Page: " + strconv.FormatInt(int64(page.SearchResponse.CurrentPage), 10))
		for _, v := range page.SearchResponse.Results {
			fmt.Println("searchQueryPaged ==> ", v.ID)
		}
	}
}
*/

func searchQuery(client *gocapiclient.GuardianContentClient, g *GuardianAPI) {
	searchQuery := queries.NewSearchQuery()

	showParam := queries.StringParam{"q", "tech%20AND%20technology"}
	showSection := queries.StringParam{"section", "technology"}
	showPages := queries.StringParam{"page", "1"}
	showPageSize := queries.StringParam{"page-seze", "1"}
	showOrderBy := queries.StringParam{"order-by", "newest"}
	showTotal := queries.StringParam{"total", "1"}
	showFields := queries.StringParam{"show-fields", "body"}
	params := []queries.Param{&showParam, showSection, showPages, showPageSize, showOrderBy, showTotal, showFields}

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

/*
func itemQuery(client *gocapiclient.GuardianContentClient) {
	itemQuery := queries.NewItemQuery("technology/2016/aug/12/no-mans-sky-review-hello-games")

	showParam := queries.StringParam{"show-fields", "all"}
	params := []queries.Param{&showParam}

	itemQuery.Params = params

	err := client.GetResponse(itemQuery)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(itemQuery.Response.Status)
	fmt.Println(itemQuery.Response.Content.WebTitle)
}
*/

func search(ctx *iris.Context){
	client := gocapiclient.NewGuardianContentClient("https://content.guardianapis.com/", "b1b1f668-8a1f-40ec-af20-01687425695c")
	g := &GuardianAPI{}
	searchQuery(client, g)

	ctx.Render("index.html", page{g.title, ctx.HostString(), g.body, g.weburl})
}
// Andrej

func display(w http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(w, "index.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}

}

func newMongoConnection() (*mgo.Session, error) {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://test:test@ds029585.mlab.com:29585/heroku_5r938bhv")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return s, err
}

type (
	User struct {
		Name     string
		Username string
		Password string
		Email    string
	}
)

func Register(w http.ResponseWriter, req *http.Request) {

	u := req.FormValue("username")
	p := req.FormValue("password")
	e := req.FormValue("email")
	n := req.FormValue("name")
	err := tpl.ExecuteTemplate(w, "index.html", User{u, p, e, n})
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}

	a := User{Username: u, Password: p, Email: e, Name: n}
	if a.Username != "" || a.Password != "" || a.Email != "" || a.Name != "" {
		insert(a)
	}
}
func loginHandler(w http.ResponseWriter, req *http.Request) {
	username := req.FormValue("username")
	password := req.FormValue("password")
	/* WE NEED TO ADD MONGO CHECKING HERE AS WELL */
	//if err := session.DB(authDB).Login(user, pass); err == nil {
	if loginValidation(username, password) {
		session, err := store.Get(req, "session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		session.Values["username"] = username
		session.Values["password"] = password
		session.Save(req, w)
		//}
		http.Redirect(w, req, "/", 302)
	} else {
		fmt.Println("Invalid login")
		// TODO: notify user of invalid username password
	}
}

func loginValidation(username string, password string) bool {
	c := mongoConnection.DB("heroku_5r938bhv").C("Users")
	result := User{}
	err = c.Find(bson.M{"username": username}).Select(bson.M{"username": 1, "password": 1, "_id": 0}).One(&result)
	if err != nil {
		// TODO: This exits the Script if the query fails to find the user, needs to be changed
		log.Fatal(err)
	}
	if result.Username == username && result.Password == password {
		fmt.Println("Connection succesful")
		return true
	} else {
		return false
	}
}

func logoutHandler(w http.ResponseWriter, req *http.Request) {
	session, err := store.Get(req, "session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["username"] = ""
	if err := session.Save(req, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, req, "/", 302)
}

//http://stackoverflow.com/questions/36323232/golang-css-files-are-being-sent-with-content-type-text-plain
func serveResource(w http.ResponseWriter, req *http.Request) {
	path := "public" + req.URL.Path
	var contentType string
	if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	}

	f, err := os.Open(path)

	if err == nil {
		defer f.Close()
		w.Header().Add("Content-Type", contentType)

		br := bufio.NewReader(f)
		br.WriteTo(w)
	} else {
		w.WriteHeader(404)
	}
}

//adapted from https://stevenwhite.com/building-a-rest-service-with-golang-3
// used to make connection to mongoDB database

func insert(a User) {
	c := mongoConnection.DB("heroku_5r938bhv").C("Users")
	err = c.Insert(&User{a.Name, a.Username, a.Password, a.Email})
	if err != nil {
		log.Fatal(err)
	}
}
