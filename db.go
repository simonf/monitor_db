package monitor_db

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

type Database struct {
	mutex     *sync.Mutex
	computers map[string]*Computer
}

// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByName []*Computer

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func NewDatabase() *Database {
	m := make(map[string]*Computer, 0)
	return &Database{computers: m, mutex: &sync.Mutex{}}
}

func (db *Database) AddComputer(item *Computer) {
	item.Updated = time.Now()
	db.mutex.Lock()
	db.computers[item.Name] = item
	db.mutex.Unlock()
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
	return nil, nil
}

func (db *Database) ListComputers() []*Computer {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	retval := make([]*Computer, 0)
	for _, cp := range db.computers {
		retval = append(retval, cp)
	}
	sort.Sort(ByName(retval))
	return retval
}

// Remove computers that have an Updated date earlier than
// specified
func (db *Database) PurgeOldComputers(minhours int) {
	max_age := time.Duration(minhours) * time.Hour

	new_db := make(map[string]*Computer, 0)

	for _, cp := range db.computers {
		age := time.Since(cp.Updated)
		if age < max_age {
			new_db[cp.Name] = cp
		}
	}
	db.computers = new_db
}

func (db *Database) PrintComputers() {
	for _, st := range db.ListComputers() {
		fmt.Println(st)
	}
}
