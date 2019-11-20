package main

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestCreateUser(t *testing.T) {
	db, err := sql.Open("sqlite3", "./world.db")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}
	err2 := createUser(tx, "TESTPW", "TESTUSER")
	if err2 != nil {
		t.Fatal(err)
	}
	tx1, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}
	userExists, err := checkForUser(tx1, "TESTUSER")
	if !userExists {
		t.Errorf("User input user failed, got: %v, want: %v.", userExists, !userExists)
	}
	tx2, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}
	tx2.Exec("DELETE FROM players WHERE name=?", "TESTUSER")
	err = tx2.Commit()
	if err != nil {
		t.Fatal(err)
	}
}
