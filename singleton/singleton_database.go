package singleton

import "sync"

type Database interface {
	GetPopulation(name string) int
}

type singletonDatabase struct {
	capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

// sync.Once or init() - thread safety
// laziness - only create when needed

var once sync.Once
var instance Database

func GetSingletonDatabase() Database {
	once.Do(func() {
		db := singletonDatabase{}
		caps, err := ReadData(".\\capitals.txt")
		if err != nil {
			db.capitals = caps
		}
		instance = &db
	})
	return instance
}

func GetTotalPopulation(db Database, cities []string) int {
	result := 0
	for _, city := range cities {
		result += db.GetPopulation(city)
	}
	return result
}
