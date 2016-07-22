package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jlopezr/routeinjector"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
	fmt.Printf("%+v \n", r)
	fmt.Printf("%+v \n", ps)
}

//Person is an example schema
type Person struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Name      string        `description:"Soy el nombre"`
	Phone     string        `description:"Soy el numero de telefono"`
	Timestamp time.Time
}

func main() {
	/*
		m1 := routeinjector.Model{}
		m1.Name = "User"
		m1.Plural = "Users"
		fmt.Printf("%+v\n", m1)

		m2 := routeinjector.NewModel("Person")
		fmt.Printf("%+v\n", m2)

		r1 := routeinjector.Route{}
		r1.Path = "/hello1"
		r1.Method = "GET"
		r1.Handler = index

		r2 := routeinjector.Route{}
		r2.Path = "/hello2"
		r2.Method = "GET"
		r2.Handler = index

		m1.Routes = append(m1.Routes, r1)
		m1.AddRoute(r2)
	*/
	ri := routeinjector.NewInjector()
	ri.Start()
	defer ri.Stop()

	ri.RegisterModel(Person{})

	router := httprouter.New()
	router.GET("/", index)
	router.GET("/hello/:name", hello)

	//m1.ProcessRoutes(router)

	//routeinjector.TestMongo()

	log.Fatal(http.ListenAndServe(":8080", router))
}
