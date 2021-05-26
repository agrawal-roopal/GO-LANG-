package main 

import(
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
)
type Article struct{
	Id string `json:"Id"`
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}
var Articles []Article
func homepage(rw http.ResponseWriter, r *http.Request){
	fmt.Fprintf(rw, "Home page!!!")
	fmt.Println("hit: homepage")

}
func returnAllArticles(rw http.ResponseWriter, r *http.Request){
	fmt.Println("hit: returnAllArticles")
	json.NewEncoder(rw).Encode(Articles)
}
func returnSingleArticle(rw http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Fprintf(rw,"key: " + key)
	for _,article :=range Articles{
		if article.Id == key{
			json.NewEncoder(rw).Encode(article)
		}
	}
}
func createNewArticle(rw http.ResponseWriter, r *http.Request){
	reqBody,_ :=ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)
	json.NewEncoder(rw).Encode(Articles)
	// fmt.Fprintf(rw,"%+v",string(reqBody))

}
func deleteArticle(rw http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["id"]

	for index,article := range Articles{
		if article.Id == key{
			Articles = append(Articles[:index],Articles[index+1:]...)
		}
	}
}
func replaceArticle(rw http.ResponseWriter, r *http.Request){
	vars :=mux.Vars(r)
	key := vars["id"]

	reqBody,_ := ioutil.ReadAll(r.Body)
	var updatedArticle Article
	json.Unmarshal(reqBody, &updatedArticle)

	for index,article := range Articles{
		if article.Id == key{
			Articles[index] = updatedArticle
		}
	}
	json.NewEncoder(rw).Encode(Articles)

}

func updateArticle(rw http.ResponseWriter, r *http.Request){
	vars :=mux.Vars(r)
	key := vars["id"]

	reqBody,_ := ioutil.ReadAll(r.Body)
	var updatedArticle Article
	json.Unmarshal(reqBody, &updatedArticle)

	for index,article := range Articles{
		if article.Id == key{
			if updatedArticle.Id != ""{
				Articles[index].Id = updatedArticle.Id
			}
			if updatedArticle.Desc != ""{
				Articles[index].Desc = updatedArticle.Desc
			}
			if updatedArticle.Title != ""{
				Articles[index].Title = updatedArticle.Title
			}
			if updatedArticle.Content != ""{
				Articles[index].Content = updatedArticle.Content
			}
		}
	}
	json.NewEncoder(rw).Encode(Articles)

}


func handleRequests(){
	
	// http.HandleFunc("/",homepage)
	// http.HandleFunc("/return-articles",returnAllArticles)
	// log.Fatal(http.ListenAndServe(":10000",nil))
	myRouter := mux.NewRouter().StrictSlash(true);
	myRouter.HandleFunc("/",homepage)
	myRouter.HandleFunc("/articles",returnAllArticles)
	myRouter.HandleFunc("/article",createNewArticle).Methods("POST")
	myRouter.HandleFunc("/articles/{id}",deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/articles/{id}",replaceArticle).Methods("PUT")
	myRouter.HandleFunc("/update-articles/{id}",updateArticle).Methods("PUT")
	myRouter.HandleFunc("/articles/{id}",returnSingleArticle)
	log.Fatal(http.ListenAndServe(":10000",myRouter))
}

func main(){
	Articles = []Article{
		Article{Id:"1",Title:"Hello", Desc:"Article description", Content:"Article Content"},
		Article{Id:"2",Title:"Hello_2", Desc:"Article description", Content:"Article Content"},
	}
	handleRequests()
}