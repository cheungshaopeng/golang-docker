package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-redis/redis"
)

// 定义一个全局变量
var redisclient *redis.Client

func initRedis()(err error){
	redisclient = redis.NewClient(&redis.Options{
		    Addr: "redis:6379",  // 指定
		    Password: "",
		    DB:0,		// redis一共16个库，指定其中一个库即可
	})
    _,err = redisclient.Ping().Result()
	return
}

func handler(w http.ResponseWriter, r *http.Request) {
	err := initRedis()
	var errStr = "success"
	if err != nil {
		errStr = "error"
	}

	fmt.Fprintf(
		w, `
          ##         .
    ## ## ##        ==
 ## ## ## ## ##    ===
/"""""""""""""""""\___/ ===
{                       /  ===-
\______ O           __/
 \    \         __/
  \____\_______/

	
Hello from Docker with connect to redis %s

`,errStr,
	)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", handler)

	fmt.Println("Go backend started!")
	log.Fatal(http.ListenAndServe(":80", r))
}