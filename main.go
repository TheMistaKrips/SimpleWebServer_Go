package main
import(
	// "encoding/json"
	"log"
	"net/http"
	"html/template"
)
func main() {

	http.HandleFunc("/", mainPage)
	
	http.HandleFunc("/users", usersPage)
	

	port := ":8080"
	println("server Listen on port:", port)
	err := http.ListenAndServe(port,nil)
	if err != nil {
		log.Fatal(err)
	}
}

type User struct {
	FirstName string `json:first_name`
	LastName string `json:last_name`
}


func mainPage(w http.ResponseWriter, r *http.Request) {
	// user := User{"Герман", "Фетисов"}
	// js, _ := json.Marshal(user)

	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w,err.Error(), 400)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w,err.Error(), 400)
		return
	}
	// w.Write(js)
}

func usersPage(w http.ResponseWriter, r *http.Request) {
	users := []User{User{"Герман", "Фетисов"}, User{"Сердар", "Бердымухаммедов"}}
	// js, _ := json.Marshal(user)

	tmpl, err := template.ParseFiles("static/users.html")
	if err != nil {
		http.Error(w,err.Error(), 400)
		return
	}

	if err := tmpl.Execute(w, users); err != nil {
		http.Error(w,err.Error(), 400)
		return
	}
	// w.Write(js)
}
