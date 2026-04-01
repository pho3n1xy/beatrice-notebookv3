package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	//The blank identifier '_' means we import it for its side effects. (registering the driver)
	_ "github.com/mattn/go-sqlite3"
)

// App struct now holds our database connection
type App struct {
	ctx context.Context
	db *sql.DB
}


// Dream represents a single dream entry.
// The `json: "..." tags tell Go how to format this when sending it to Svelte
type Dream struct {
		ID 				int 		`json: "id"`
		Title 		string 	`json: "title"`
		Content		string 	`json: "content"`
		Date 			string 	`json: "date"`
		MoonPhase string  `json: "moonphase"`
		CreatedAt string 	`json: "createdAt"`
}


// & (address-of operator) returns the memory address of a variable
// * (dereference operator) returns the value that a pointer points to
// used to access the actual data
func AppInstance() *App{
		return &App{}
}


// Wails has a startup function. This is where we want to find the user's home
// directory, create a hidden folder (~/.dream-journal), and open the SQLite
// database.

//(app *App) is a reciever declaration.
//*App means it's a pointer reciever - allow modification of original struct
func (app *App) startup(ctx context.Context){
	app.ctx = ctx

	// Find the user's home directory securely
	homeDir, err := os.UserHomeDir()
	if err != nil{
		log.Fatal("Could not get home directory: %v", err)
	}

	// Create the hidden .dream-journal folder if it doesn't exist
	appDir := filepath.Join(homeDir, ".dream-journal")
	if err := os.MkdirAll(appDir, 0755); err != nil {
		log.Fatal("Could not create app directory: %v", err)
	}

	// Connect to the SQLite database
	dbPath := filepath.Join(appDir, "data.db")
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Fatalf("Could not open database: %v", err)
	}

	app.db = db
	app.initDB()

}


// initDB creates the neccessary tables if they don't already exist
func (app *App) initDB(){
	query := `
	CREATE TABLE IF NOT EXISTS dreams(
		id 					INTEGER PRIMARY KEY AUTOINCREMENT,
		title 			TEXT NOT NULL,
		content 		TEXT NOT NULL,
		date 				TEXT NOT NULL,
		moonphase 	TEXT NOT NULL,
		created_at 	DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := app.db.Exec(query)
	if err != nil{
			log.Fatal("Could not create tables: %v", err)
	}
}


// the "Bindings" (Functions Svelte will call)
// we add the functions that actually save and retrieve data.
// Notice they start with capital letters (SaveDream and GetDreams)
// Wails will detect these and generate JavaScript
func (app *App) SaveDream(title, content, date string) (Dream, error) {
	if title == ""{
		title = "Untitle Dream"
	}

	query := `INSERT INTO dreams (title, content, date, moonphase) VALUES (?, ? , ?, ?)`
	result, err := a.db.Exec(query, title, content, date, moonphase)
	if err != nil {
		return Dream{}, fmt.Errorf("failed to save dream: %v", err)
	}

	id, err := res.LastInsertId()
	if err != nil{
			return Dream{}, fmt.Error("failed to get inserted ID: %v", err)
	}

	//Return the save dream object so the frontend can display it immediately.
	return Dream{
		ID: 					int(id),
		Title: 				title,
		Content: 			content,
		Date: 				date,
		Moonphase: 		moonphase,
		CreatedAt: 		time.Now().Format(time.RFC3339),
	}, nil
}

// GetDreamsDesc retrieves all dreams ordered by date descending
func (app *App) GetDreamsDesc() ([]Dream, error){
	query := `SELECT id, title, content, date, moonphase, created_at FROM dreams ORDER BY date DESC, id DESC`
	rows, err := app.db.Query(query)

	if err != nil{
		return nil, fmt.Errorf("failed to query dreams: %v", err)
	}

	defer rows.Close()

	var dreams []Dream
	for rows.Next(){
		var dream Dream
		if err := rows.Scan(&dream.ID,&dream.Title, &dream.Content, &dream.Date, &dream.Moonphase, &dream.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan dream row: %v", err)
		}
		dreams = append(dreams, dream)
	}

	//if no dreams, return an empty array instead of null for the svelte frontend
	if dreams == nil{
		dreams = []Dream{}
	}
	return dreams, nil
}

//GetDreamsAESC retrieves all dreams ordered by date ascending
func (app *App) GetDreamsAsc() ([]Dream, error){
	query := `SELECT id, title, content, date, moonphase, create_at FROM dreams ORDER BY date ASC, id ASC`
	rows, error := app.db.Query(query)

	if err != nil{
		return nil, fmt.Errorf("failed to query dreams: %v", err)
	}

	defer rows.Close()

	var dreams []Dream
	for rows.Next(){
		var dream Dream
		if err := rows.Scan(&dream.ID, &dream.Title, &dream.Content, &dream.Date, &dream.Moonphase, &dream.CreatedAt); err != nil{
			return nil, fmt.Errorf("failed to scan dream row: %v", err)
		}
		dreams = append(dreams, dream)
	}

	//if no dreams, return an empty array instead of null for the svelte frontend
	if dreams == nil{
		dreams = []Dream{}
	}
	return dreams, nil
}
