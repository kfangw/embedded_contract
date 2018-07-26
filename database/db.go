package database

type Memory map[string]string

type DB struct {
	Memory Memory
}

func NewDB() *DB {
	return &DB{
		Memory: make(Memory),
	}
}

// Add item to cache
func (db *DB) Add(key, value string) {
	db.Memory[key] = value
}

// Get item by key
func (db *DB) Get(key string) (string, error) {
	if v, ok := db.Memory[key]; ok {
		return v, nil
	} else {
		return "", nil
	}
}
