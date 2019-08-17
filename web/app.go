package main

import ( 
	"fmt"
        "net/http"
        "log"
        "encoding/json"
)

func home_handler(w http.ResponseWriter, r *http.Request) {
        out := "<h1>Welcome to Linux Interface Rest Api</h1>"
        fmt.Fprintf(w,"%s" ,out)
}

func show_handler(w http.ResponseWriter, r *http.Request) {
        var i Interface
	out, _ := i.show()
        jout, _ := json.Marshal(out)
	fmt.Fprintf(w, "%s", jout)
}

func main() {
    http.HandleFunc("/interfaces/", home_handler)
    http.HandleFunc("/interfaces/show", show_handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
