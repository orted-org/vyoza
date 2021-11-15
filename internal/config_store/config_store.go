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
func (c *Config) Get(ctx context.Context, name string) (string, error) {
	if config, err := c.store.GetKeyValue(ctx, c.WithNamespace(name)); err != nil {
		return "", err
	} else {
		return config.Value, nil
	}
}
func (c *Config) Set(ctx context.Context, name string, value string) error {
	_, err := c.store.UpdateKeyValue(ctx, db.KeyValue{
		Key:   c.WithNamespace(name),
		Value: value,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			_, err := c.store.AddKeyValue(ctx, db.KeyValue{
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
func (c *Config) Delete(ctx context.Context, name string) error {
	return c.store.DeleteKeyValue(ctx, c.WithNamespace(name))
}
func (c *Config) WithNamespace(name string) string {
	return c.namespace + "." + name
}
