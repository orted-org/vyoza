package vault

import (
	"context"
	"database/sql"
	"errors"

	db "github.com/orted-org/vyoza/db/dao"
)

type Vault struct {
	store     db.Store
	namespace string
}

func New(s db.Store) *Vault {
	return &Vault{
		store:     s,
		namespace: "vault",
	}
}
func (v *Vault) Get(ctx context.Context, name string) (string, error) {
	if vaultKV, err := v.store.GetKeyValue(ctx, v.WithNamespace(name)); err != nil {
		return "", err
	} else {
		return vaultKV.Value, nil
	}
}
func (v *Vault) Set(ctx context.Context, name string, value string) error {
	//checking weather there is already with this key or not
	//if yes 
	_, err := v.store.GetKeyValue(ctx, v.WithNamespace(name))

	if err != nil {
		if err == sql.ErrNoRows {
			_, err := v.store.AddKeyValue(ctx, db.KeyValue{
				Key:   v.WithNamespace(name),
				Value: value,
			})
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}

	return errors.New("key is not unique")
}
func (v *Vault) Delete(ctx context.Context, name string) error {
	return v.store.DeleteKeyValue(ctx, v.WithNamespace(name))
}
func (v *Vault) WithNamespace(name string) string {
	return v.namespace + "." + name
}