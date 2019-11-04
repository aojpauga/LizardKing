package main

import (
	"bufio"
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"golang.org/x/crypto/pbkdf2"

	crand "crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
)

func handlePlayerConnection(db *sql.DB, conn net.Conn, inputs chan<- InputEvent, player *Player) {
	//time.Sleep(30 * time.Second)
	fmt.Fprintf(conn, "Welcome %s! ", player.Name)
	scanner := bufio.NewScanner(conn)
	fmt.Fprintln(conn, "Enter a command or type 'quit' to quit.")
	go func() {
		for {
			outputevent, more := <-player.Outputs
			if more {
				fmt.Fprintf(conn, outputevent.Text)
			} else {
				log.Print("Output channel closed for user: ", player.Name)
				conn.Close()
				return
			}
		}
	}()
	for scanner.Scan() {
		input := scanner.Text()
		//fmt.Println(command) // Println will add back the final '\n'
		//fmt.Printf("Fields are: %q", strings.Fields(command))
		commandFields := strings.Fields(input)
		//fmt.Printf("%T\n", commandFields)
		if len(commandFields) == 0 {
			fmt.Fprintln(conn, "Please Enter a command")
		} else if commandFields[0] == "quit" || commandFields[0] == "Quit" {
			inputevent := InputEvent{player, commandFields, true, false}
			inputs <- inputevent
		} else {
			// pass commands to Main goroutine from here
			inputevent := InputEvent{player, commandFields, false, false}
			inputs <- inputevent
			//commandHandler(player, commandFields)
		}

		//commands["look"](input)

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	}
}
func checkForUser(tx *sql.Tx, username string) (bool, int, error) {
	var name string
	var id int
	dbuser := tx.QueryRow("SELECT name,id FROM players WHERE name=?", username)
	switch err := dbuser.Scan(&name, &id); err {
	case sql.ErrNoRows:
		fmt.Printf("User %v does not exist\n", username)
		return false, -1, tx.Commit()
	case nil:
		fmt.Printf("User %v found!\n", username)
		return true, id, tx.Commit()
	default:
		fmt.Printf("Error checking for user.\n")
		tx.Rollback()
		return false, -1, err
	}
}
func loginUser(tx *sql.Tx, password string, username string) (bool, error) {
	dbplayer, err := tx.Query("SELECT salt,hash FROM players WHERE name=?", username)
	if err != nil {
		tx.Rollback()
		return false, err
	}
	defer dbplayer.Close()
	var salt64, hash64 string
	for dbplayer.Next() {

		if err := dbplayer.Scan(&salt64, &hash64); err != nil {
			tx.Rollback()
			return false, err
		}
	}
	salt, err := base64.StdEncoding.DecodeString(salt64)
	if err != nil {
		tx.Rollback()
		return false, err
	}
	hash, err := base64.StdEncoding.DecodeString(hash64)
	if err != nil {
		tx.Rollback()
		return false, err
	}
	inputhash := pbkdf2.Key(
		[]byte(password),
		salt,
		64*1024,
		32,
		sha256.New)
	if subtle.ConstantTimeCompare(hash, inputhash) != 1 {
		return false, tx.Commit()
	} else {
		return true, tx.Commit()
	}
}
func getCharacters(tx *sql.Tx, username string) ([]*Character, error) {
	dbcharacters, err := tx.Query("SELECT name, class, level FROM characters WHERE player_name=?", username)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	defer dbcharacters.Close()
	var characters []*Character
	for dbcharacters.Next() {
		var level int
		var name, class string
		if err := dbcharacters.Scan(&name, &class, &level); err != nil {
			tx.Rollback()
			return nil, err
		}
		character := Character{name, class, level}
		characters = append(characters, &character)
	}
	if err := dbcharacters.Err(); err != nil {
		tx.Rollback()
		return nil, err
	}
	return characters, tx.Commit()
}
func createUser(
	tx *sql.Tx, password string, username string) error {
	salt := make([]byte, 32)
	_, err := crand.Read(salt)
	if err != nil {
		log.Print("Error creating salt: ", err)
	}
	salt64 := base64.StdEncoding.EncodeToString(salt)
	hash := pbkdf2.Key(
		[]byte(password),
		salt,
		64*1024,
		32,
		sha256.New)
	hash64 := base64.StdEncoding.EncodeToString(hash)

	tx.Exec("INSERT INTO players (name,salt,hash) VALUES (?,?,?)", username, salt64, hash64)
	return tx.Commit()
}
