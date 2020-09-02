package main

import (
	"easyCache/lru"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var cache *lru.Cache

func init() {
	cache = lru.NewCache(1000)
}
func main() {

	http.HandleFunc("/get", getHandler) //get
	http.HandleFunc("/set", setHandler) //post
	http.HandleFunc("/del", delHandler) //get
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func getHandler(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)
	var data map[string]interface{}

	query := request.URL.Query()
	key := query.Get("key")
	if value, ok := cache.Get(key); ok {
		data = map[string]interface{}{
			"code":    200,
			"message": "ok",
			"data":    value,
		}
		json.NewEncoder(writer).Encode(data)
		return
	} else {
		data = map[string]interface{}{
			"code":    200,
			"message": "not found",
			"data":    nil,
		}
		json.NewEncoder(writer).Encode(data)
		return
	}
}

type setDataPackage struct {
	Key   string
	Value string
	Time  int64
}

//{"key":"","value":""}
func setHandler(writer http.ResponseWriter, request *http.Request) {

	if request.Method == "POST" {

		var p setDataPackage
		body, _ := ioutil.ReadAll(request.Body)
		json.Unmarshal(body, &p)

		fmt.Println("k:"+p.Key, "v:"+p.Value, "time:", p.Time)

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(200)
		var data map[string]interface{}

		if p.Key != "" && p.Value != "" {
			cache.Set(p.Key, p.Value, p.Time)
			data = map[string]interface{}{
				"code":    200,
				"message": "ok",
				"data":    nil,
			}
		} else {
			data = map[string]interface{}{
				"code":    200,
				"message": "参数不正确",
				"data":    nil,
			}
		}
		json.NewEncoder(writer).Encode(data)
		return
	} else {
		writer.WriteHeader(404)
		writer.Write([]byte("404 page not found"))
		return
	}
}

func delHandler(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)
	var data map[string]interface{}
	query := request.URL.Query()
	key := query.Get("key")
	if value, ok := cache.Get(key); ok {
		cache.Remove(key)
		data = map[string]interface{}{
			"code":    200,
			"message": "ok",
			"data":    value,
		}
		json.NewEncoder(writer).Encode(data)
		return
	} else {
		data = map[string]interface{}{
			"code":    200,
			"message": "not found",
			"data":    nil,
		}
		json.NewEncoder(writer).Encode(data)
		return
	}
}
