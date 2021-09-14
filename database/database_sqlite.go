package database


import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DbConnection *sql.DB

type Person struct {
	Name string
	Age  int
}

func database() {
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()
	cmd := `CREATE TABLE IF NOT EXISTS person(
		name STRING,
		age INT)`
	//結果がいらない時にexecを使用して抜ける
	_, err := DbConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	//データベース追加
	cmd = "INSERT INTO person (name,age) VALUES (?,?)"
	_, err = DbConnection.Exec(cmd, "Miku", 20)
	if err != nil {
		log.Fatalln(err)
	}

	//データを取得する
	//マルチセレクト
	cmd = "SELECT * FROM person"
	rows, _ := DbConnection.Query(cmd)
	defer rows.Close()

	//ここはパターンとして暗記する
	var pp []Person
	for rows.Next() {
		var p Person
		err := rows.Scan(&p.Name, &p.Age)
		if err != nil {
			log.Println(err)
		}
		pp = append(pp, p)
	}

	err = rows.Err()
	if err != nil {
		log.Fatalln(err)
	}

	for _, p := range pp {
		fmt.Println(p.Name, p.Age)
	}

	//シングルセレクト

	cmd = "SELECT * FROM person where age=?"
	row := DbConnection.QueryRow(cmd, 20)
	var p Person
	err = row.Scan(&p.Name, &p.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No row")
		} else {
			log.Println(err)
		}
	}
	tableName :="person"
	//tableを入れる場合は?をつけない
	cmd =fmt.Sprint("SELECT * FROM %s",tableName)

}
