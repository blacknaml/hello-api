package translation

import (
	"context"
	"fmt"

	"github.com/blacknaml/hello-api/config"
	"github.com/blacknaml/hello-api/handlers/rest"
	"github.com/redis/go-redis/v9"
)

var _ rest.Translator = &Database{}

type Database struct {
	conn *redis.Client
}

func NewDatabaseService(cfg config.Configuration) *Database {
	rdb := redis.NewClient(
		&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", cfg.DatabaseURL, cfg.DatabasePort),
			Password: "",
			DB:       0,
		})
	return &Database{
		conn: rdb,
	}
}

func (db *Database) Close() error {
	return db.conn.Close()
}

func (db *Database) Translate(word, language string) string {
	out := db.conn.Get(context.Background(), fmt.Sprintf("%s:%s", word, language))
	return out.Val()
}
