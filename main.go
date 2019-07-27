package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"session"
	"time"
)

var locale Lang
var appSettings Settings
var globalSessions *session.Manager

func main() {
	var err error
	globalSession = openSession()
	appSettings = loadSettings()
	globalSessions, err = session.NewManager("memory", "gosessionid", 3600)
	checkErr(err)

	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

	var langFile string
	switch appSettings.Locale {
	case "ru":
		{
			langFile = "assets/lang/ru.json"
			break
		}
	case "no":
		{
			langFile = "assets/lang/no.json"
			break
		}
	default:
		{
			langFile = "assets/lang/en.json"
		}
	}
	localeLoad(langFile)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/student", studentHandler)
	http.HandleFunc("/teacher", teacherHandler)
	http.HandleFunc("/input", inputData)
	http.HandleFunc("/testing", testing)
	http.HandleFunc("/locale", changeLocale)
	http.HandleFunc("/subjects", subjects)
	http.HandleFunc("/admin", admin)
	http.HandleFunc("/tests", tests)
	http.HandleFunc("/users", users)
	http.HandleFunc("/settings", settings)

	srv := &http.Server{
		Addr:         ":3000",
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Println("Listening on port ", ":3000")
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	<-stopChan
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	err = srv.Shutdown(ctx)
	checkErr(err)
	defer cancel()
	log.Println("Server gracefully stopped!")
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err) //respond with error page or message
	}
}

func localeLoad(path string) {
	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully opened locale file")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var lang Lang

	if err = json.Unmarshal(byteValue, &lang); err != nil {
		fmt.Println(err.Error())
	}

	locale = lang
}

func changeLocale(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	lang := query["locale"][0]
	appSettings.Locale = lang
	changeDBLocale()

	var langFile string
	switch appSettings.Locale {
	case "ru":
		{
			langFile = "assets/lang/ru.json"
			break
		}
	case "no":
		{
			langFile = "assets/lang/no.json"
			break
		}
	default:
		{
			langFile = "assets/lang/en.json"
		}
	}
	localeLoad(langFile)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	globalSessions.SessionStart(w, r)

	t, err := template.ParseFiles("static/index.html", "static/header.html", "static/footer.html", "static/language.html")
	if err != nil {
		_, err = fmt.Fprintf(w, err.Error())
		if err != nil {
			panic(err.Error())
		}
	}
	data := Data{nil, locale}
	err = t.ExecuteTemplate(w, "test", data)
	if err != nil {
		panic(err.Error())
	}
}

func studentHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/student.html", "static/header.html", "static/footer.html", "static/language.html")
	if err != nil {
		_, err = fmt.Fprintf(w, err.Error())
		if err != nil {
			panic(err.Error())
		}
	}
	data := Data{nil, locale}
	err = t.ExecuteTemplate(w, "student", data)
	if err != nil {
		panic(err.Error())
	}
}

func teacherHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/teachersignin.html", "static/header.html", "static/footer.html", "static/language.html")
	if err != nil {
		_, err = fmt.Fprintf(w, err.Error())
		if err != nil {
			panic(err.Error())
		}
	}
	data := Data{nil, locale}
	err = t.ExecuteTemplate(w, "teachersignin", data)
	if err != nil {
		panic(err.Error())
	}
}

func inputData(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/inputData.html", "static/header.html", "static/footer.html", "static/language.html")
	if err != nil {
		_, err = fmt.Fprintf(w, err.Error())
		if err != nil {
			panic(err.Error())
		}
	}
	groups := getGroups()
	data := Data{groups, locale}
	err = t.ExecuteTemplate(w, "inputData", data)
	if err != nil {
		panic(err.Error())
	}
}

func testing(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/testing.html", "static/header.html", "static/footer.html", "static/language.html")
	if err != nil {
		_, err = fmt.Fprintf(w, err.Error())
		if err != nil {
			panic(err.Error())
		}
	}
	data := Data{nil, locale}
	err = t.ExecuteTemplate(w, "testing", data)
	if err != nil {
		panic(err.Error())
	}
}

func admin(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/adminPanel.html", "static/header.html", "static/footer.html", "static/language.html")
	if err != nil {
		_, err = fmt.Fprintf(w, err.Error())
		if err != nil {
			panic(err.Error())
		}
	}
	data := Data{nil, locale}
	err = t.ExecuteTemplate(w, "admin", data)
	if err != nil {
		panic(err.Error())
	}
}

func tests(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/adminThings/tests.html")
	if err != nil {
		_, err = fmt.Fprintf(w, err.Error())
		if err != nil {
			panic(err.Error())
		}
	}
	data := Data{nil, locale}
	err = t.ExecuteTemplate(w, "tests", data)
	if err != nil {
		panic(err.Error())
	}
}

func users(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/adminThings/users.html")
	if err != nil {
		_, err = fmt.Fprintf(w, err.Error())
		if err != nil {
			panic(err.Error())
		}
	}
	data := Data{nil, locale}
	err = t.ExecuteTemplate(w, "users", data)
	if err != nil {
		panic(err.Error())
	}
}

func settings(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/adminThings/settings.html")
	if err != nil {
		_, err = fmt.Fprintf(w, err.Error())
		if err != nil {
			panic(err.Error())
		}
	}
	data := Data{nil, locale}
	err = t.ExecuteTemplate(w, "settings", data)
	if err != nil {
		panic(err.Error())
	}
}

func subjects(w http.ResponseWriter, r *http.Request) {

	subjects := getSubjects()

	id := ""

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if id == "" {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		err := json.NewEncoder(w).Encode(subjects)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
		}
	} else {
		if subject, err := getSubject(id); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			err := json.NewEncoder(w).Encode(subject)
			if err != nil {
				w.WriteHeader(http.StatusForbidden)
			}
		}
	}

}
