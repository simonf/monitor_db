package monitor_db

import (
	"fmt"
	"sync"
)

type Database struct {
	mutex     *sync.Mutex
	computers map[string]*Computer
}

func NewDatabase() *Database {
	m := make(map[string]*Computer, 0)
	return &Database{computers: m, mutex: &sync.Mutex{}}
}

func (db *Database) AddComputer(item *Computer) {
	db.mutex.Lock()
	db.computers[item.Name] = item
	db.mutex.Unlock()
	fmt.Println("Added")
}

func (db *Database) GetComputer(name string) (*Computer, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	c := db.computers[name]
	if c == nil {
		return nil, fmt.Errorf("Not found")
	} else {
		return c, nil
	}
}

func (db *Database) ListComputers() []*Computer {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	retval := make([]*Computer, 0)
	for _, cp := range db.computers {
		retval = append(retval, cp)
	}
	fmt.Println("Returning")
	return retval
}

func (db *Database) PrintComputers() {
	for _, st := range db.ListComputers() {
		fmt.Println(st)
	}
}