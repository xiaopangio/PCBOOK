// Package service  @Author xiaobaiio 2023/2/23 9:24:00
package service

import (
	"github.com/xiaopangio/pcbook/orm/dal"
	"github.com/xiaopangio/pcbook/orm/model"
	"gorm.io/gorm"
	"sync"
)

type UserStore interface {
	Save(user *User) error
	Find(username string) (*User, error)
}
type InMemoryUserStore struct {
	mutex sync.RWMutex
	users map[string]*User
}

func NewInMemoryUserStore() *InMemoryUserStore {
	return &InMemoryUserStore{
		users: make(map[string]*User),
	}
}
func (store *InMemoryUserStore) Save(user *User) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	if store.users[user.Username] != nil {
		return ErrAlreadyExists
	}
	store.users[user.Username] = user.Clone()
	return nil
}

func (store *InMemoryUserStore) Find(username string) (*User, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()
	user := store.users[username]
	if user == nil {
		return nil, nil
	}
	return user.Clone(), nil
}

type DBUserStore struct {
	db *gorm.DB
}

func (store *DBUserStore) Save(user *User) error {
	u := &model.User{}
	u.Username = user.Username
	u.HashedPassword = user.HashedPassword
	role, err := dal.Role.Where(dal.Role.RoleName.Eq(user.Role)).First()
	u.RoleID = role.RoleID
	if err != nil {
		return err
	}
	err = dal.User.Omit().Create(u)
	return err
}

func (store *DBUserStore) Find(username string) (*User, error) {
	user, err := dal.User.Where(dal.User.Username.Eq(username)).Preload(dal.User.Role).First()
	u := &User{
		Username:       user.Username,
		HashedPassword: user.HashedPassword,
		Role:           user.Role.RoleName,
	}
	return u, err
}

func NewDBUserStore(DB *gorm.DB) *DBUserStore {
	dal.SetDefault(DB)
	return &DBUserStore{db: DB}
}
