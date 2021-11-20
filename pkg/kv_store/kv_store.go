package kvstore

import "errors"

var (
	ErrKeyValueNotExists = errors.New("key value does not exits")
)

type KVStore interface {
	Get(key string) (string, error)
	Set(key, value string) error
	Delete(key string) error
	Truncate() error
}
type InMemKVStore struct {
	data map[string]interface{}
}

func New() *InMemKVStore {
	return &InMemKVStore{
		data: make(map[string]interface{}),
	}
}
func (i *InMemKVStore) Get(key string) (interface{}, error) {
	if value, ok := i.data[key]; ok {
		return value, nil
	} else {
		return "", ErrKeyValueNotExists
	}
}
func (i *InMemKVStore) Set(key string, value interface{}) error {
	i.data[key] = value
	return nil
}
func (i *InMemKVStore) Delete(key string) error {
	delete(i.data, key)
	return nil
}
func (i *InMemKVStore) Truncate() error {
	i.data = make(map[string]interface{})
	return nil
}
