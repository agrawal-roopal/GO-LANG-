package main

import(
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Tag struct{
	Id string `json:"id"`
	Name string `json:"name"`
}

func main(){
	fmt.Println("Go MySQL Tutorial")
	db, err := sql.Open("mysql", "debian-sys-maint:oDyl5neosJTVJynR@tcp(127.0.0.1:3306)/Test")
	if err != nil {
        panic(err.Error())
	}
	defer db.Close()

	//////////////////////////////////
	/////INSERT

	// insert, err := db.Query("INSERT INTO test VALUES (3, 'TEST_3')")

	// if err != nil{
	// 	panic(err.Error())
	// }

	// insert.Close()

	//////////////////////////////////////////

	///////////SELECT 
	// results, err:=db.Query("SELECT id, testcol FROM test")

	// if err != nil{
	// 	panic(err.Error())
	// }

	// for results.Next(){
	// 	var tag Tag

	// 	err = results.Scan(&tag.Id,&tag.Name)
	// 	if err != nil{
	// 		panic(err.Error())
	// 	}

	// 	log.Println(tag.Name)
	// }

	//////////////////////////
	//////////SELECT A ROW

	var tag Tag
	err = db.QueryRow("SELECT id,testcol FROM test where id = ?",2).Scan(&tag.Id, &tag.Name)

	if err != nil{
		panic(err.Error())
	}

	log.Println(tag.Id)
	log.Println(tag.Name)
}