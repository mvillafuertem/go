package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type UserController struct {
	Router *gin.Engine
}

func (u *UserController) V1() {
	u.getUser()
	u.ping()
	u.pingPost()
}

func (u *UserController) getUser() {

	u.Router.GET("/user/:userId", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"userId":  context.Param("userId"),
			"name":    "Pepe",
			"surname": "Pipo",
		})
	})
}

func (u *UserController) ping() {

	u.Router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

func (u *UserController) pingPost() {

	lookup := StartActor()
	output := StartActor()
	query := StartActor()
	post := StartActor()


	u.Router.POST("/ping/:service", func(context *gin.Context) {
		action := Action{path: "http://localhost:9001/ping/"}
		service := context.Param("service")

		switch service {
		case "lookup":
			fmt.Println("lookup")
			lookup.receive(action)
		case "output":
			fmt.Println("output")
			output.receive(action)
		case "query":
			fmt.Println("output")
			query.receive(action)
		case "post":
			fmt.Println("output")
			post.receive(action)
		}

		select {
		case a := <-lookup.action:
			context.JSON(200, gin.H{
				"action":  "lookup",
				"message": a.payload,
			})

		case b := <-output.action:
			context.JSON(200, gin.H{
				"action":  "output",
				"message": b.payload,
			})
		case c := <-query.action:
			context.JSON(200, gin.H{
				"action":  "query",
				"message": c.payload,
			})
		case d := <-post.action:
			context.JSON(200, gin.H{
				"action":  "output",
				"message": d.payload,
			})
		}
	})
}

func ClientHttp() *http.Client {

	return &http.Client{}
}

type Actor struct {
	action chan Action
}

type Action struct {
	payload string
	path    string
}

func (a *Actor) send() {

	for {
		invoker(a)
	}

}

func (a *Actor) receive(action Action) {
	fmt.Println("Receive action ", action)
	a.action <- action
}

func invoker(actor *Actor) {
	a := <-actor.action

	request, _ := http.NewRequest("GET", a.path, nil)

	resp, _ := ClientHttp().Do(request)

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	a.payload = string(body)

	actor.action <- a
}

func StartActor() *Actor {
	actor := &Actor{action: make(chan Action, 1)}
	go actor.send()
	return actor
}
