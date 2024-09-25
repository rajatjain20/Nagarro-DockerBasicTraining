package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/welcome", func(res http.ResponseWriter, wr *http.Request) {
		res.Write([]byte("Welcome to your first docker demo!!!!!"))
	})

	http.HandleFunc("/", getRoot)
	http.HandleFunc("/demo", getDemo)
	http.HandleFunc("/check_db", checkdbConnection)
	http.HandleFunc("/execSql", execSql)
	http.HandleFunc("/updateSql", updateSql)

	fmt.Println("Server started and Listening at port :3333")
	http.ListenAndServe(":3333", nil)
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ENV_NAME=%s\n", getEnv())
	fmt.Fprintf(w, "Hello! Welcome to first docker demo.\n")
}

func getDemo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ENV_NAME=%s\n", getEnv())

	val := r.URL.Query()
	if val.Has("name") {
		fmt.Fprintf(w, "Hello! %s", string(val["name"][0]))
	} else {
		fmt.Fprintf(w, "Hello! Please provide name.")
	}
}

func getEnv() string {
	return os.Getenv("ENV_NAME")
}
