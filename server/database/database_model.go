package database

// Database abstraction
type Database interface {
	Set(key string, value string) ([]byte, error)
	Get(key string) ([]byte, error)
}

// Factory looks up acording to the databaseName the database implementation
func Factory(databaseName string) (Database, error) {
	switch databaseName {
	case "redis":
		return createRedisDatabase()
	default:
		return nil, &NotImplementedDatabaseError{databaseName}
	}
}
