package wordnet

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path/filepath"
)

type Result struct {
	Word       string
	Definition string
	Type       string
}

type WordNet struct {
	*sql.DB
}

func (db *WordNet) Init(dbpath string) error {
	var err error
	if dbpath == "" {
		db.DB, err = sql.Open("sqlite3", filepath.Join(os.Getenv("GOPATH"), "/src/github.com/godwhoa/wordnet-go/wordnet.db"))
	} else {
		db.DB, err = sql.Open("sqlite3", dbpath)
	}
	return err
}

var norows = sql.ErrNoRows

func (db *WordNet) ByWord(word string, limit int) ([]Result, error) {
	results := []Result{}
	var err error
	var stmt *sql.Stmt
	var rows *sql.Rows

	stmt, err = db.Prepare("SELECT word,definition,type FROM words WHERE word = ? LIMIT ?;")
	if err != nil {
		return results, fmt.Errorf("Prepare statement error: %v", err)
	}

	rows, err = stmt.Query(word, limit)
	if err == sql.ErrNoRows {
		return results, norows
	}
	if err != nil {
		return results, fmt.Errorf("Query error: %v", err)
	}

	for rows.Next() {
		result := Result{}
		err = rows.Scan(&result.Word, &result.Definition, &result.Type)
		if err != nil {
			return results, fmt.Errorf("Row Scan error: %v", err)
		}
		results = append(results, result)
	}

	rows.Close()
	stmt.Close()
	return results, nil
}

func (db *WordNet) ByType(word string, mtype string, limit int) ([]Result, error) {
	results := []Result{}
	if !strIn([]string{"noun", "verb", "adj", "adv"}, mtype) {
		return results, fmt.Errorf("Invalid type. Valid types: adj, adv, noun, verb")
	}
	var err error
	var stmt *sql.Stmt
	var rows *sql.Rows

	stmt, err = db.Prepare("SELECT word,definition,type FROM words WHERE word = ? AND type = ? LIMIT ?;")
	if err != nil {
		return results, fmt.Errorf("Prepare statement error: %v", err)
	}

	rows, err = stmt.Query(word, mtype, limit)
	if err != nil {
		return results, fmt.Errorf("Query error: %v", err)
	}

	for rows.Next() {
		result := Result{}
		err = rows.Scan(&result.Word, &result.Definition, &result.Type)
		if err != nil {
			return results, fmt.Errorf("Row Scan error: %v", err)
		}
		results = append(results, result)
	}

	rows.Close()
	stmt.Close()
	return results, nil
}
