package db

import (
	// "sync"
	"time"
	// "database/sql"

)

type Task struct {
	Id   int	   `json:"id"`
	Text string    `json:"text"`
	Tags []string  `json:"tags"`
	Due  time.Time `json:"deadline"` // deadline date
}


// type TaskStore struct {
// 	sync.Mutex
// 	db *sql.DB
// }