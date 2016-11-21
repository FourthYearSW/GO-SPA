package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// mongo ds029585.mlab.com:29585/heroku_5r938bhv -u <dbuser> -p <dbpassword>
// all session code adapted from http://www.gorillatoolkit.org/pkg/sessions
var store = sessions.NewCookieStore([]byte("secret"))
var tpl *template.Template
var mongoConnection, err = newMongoConnection()

func init() {
	tpl = template.Must(template.ParseGlob("public/index.html"))
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

}

func display(w http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(w, "index.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}

}

func newMongoConnection() (*mgo.Session, error) {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://tester:tester@ds029585.mlab.com:29585/heroku_5r938bhv")

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
