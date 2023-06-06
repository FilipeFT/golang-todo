package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/utils"
	"os"
	"time"
)

// Item struct represents an item of a to-do list
type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

// List represents a list of to-do items
type List []item

// Add creates a new todo item and appends it to the list
func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time(),
	}
	*l = append(*l, t)
}

// Complete method marks a to-do item as completed by
// setting Done = true and CompletedAt to the current time
func (l *List) Complete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exist", i)
	}
	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()

	return nil
}
