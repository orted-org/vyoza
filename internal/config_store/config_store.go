package configstore

import (
	"context"
	"database/sql"

	db "github.com/orted-org/vyoza/db/dao"
)

type Config struct {
	store     db.Store
	namespace string
}

func New(s db.Store) *Config {
	return &Config{
		store:     s,
		namespace: "config",
	}
}
func (c *Config) Get(name string) (string, error) {
	if config, err := c.store.GetKeyValue(context.Background(), c.WithNamespace(name)); err != nil {
		return "", err
	} else {
		return config.Value, nil
	}
}
func (c *Config) Set(name string, value string) error {
	_, err := c.store.UpdateKeyValue(context.Background(), db.KeyValue{
		Key:   c.WithNamespace(name),
		Value: value,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			_, err := c.store.AddKeyValue(context.Background(), db.KeyValue{
				Key:   c.WithNamespace(name),
				Value: value,
			})
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	return nil
}
func (c *Config) Delete(name string) error {
	return c.store.DeleteKeyValue(context.Background(), c.WithNamespace(name))
}
func (c *Config) WithNamespace(name string) string {
	return c.namespace + "." + name
}
