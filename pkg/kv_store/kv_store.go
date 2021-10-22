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
	data map[string]string
}

func New() *InMemKVStore {
	return &InMemKVStore{
		data: make(map[string]string),
	}
}
func (i *InMemKVStore) Get(key string) (string, error) {
	if value, ok := i.data[key]; ok {
		return value, nil
	} else {
		return "", ErrKeyValueNotExists
	}
}
func (i *InMemKVStore) Set(key, value string) error {
	i.data[key] = value
	return nil
}
func (i *InMemKVStore) Delete(key string) error {
	delete(i.data, key)
	return nil
}
func (i *InMemKVStore) Truncate() error {
	i.data = make(map[string]string)
	return nil
}
