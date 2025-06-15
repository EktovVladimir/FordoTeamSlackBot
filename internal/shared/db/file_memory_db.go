package db

import (
	"encoding/json"
	"errors"
	"os"
	"path"
)

type Entity interface {
}

type FileMemoryDb struct {
	Context     *context
	storagePath string
}

type context struct {
	Users       *entityContext[User]
	Settings    *entityContext[Setting]
	Deployments *entityContext[Deployment]
	CodeReviews *entityContext[CodeReview]
}

func NewFileMemoryDb(storagePath string) *FileMemoryDb {
	ctx := &context{
		Users:       NewEntityContext[User](),
		Settings:    NewEntityContext[Setting](),
		Deployments: NewEntityContext[Deployment](),
		CodeReviews: NewEntityContext[CodeReview](),
	}

	return &FileMemoryDb{
		Context:     ctx,
		storagePath: storagePath,
	}
}

func (db *FileMemoryDb) LoadAll() error {
	if err := db.LoadUsers(); err != nil {
		return err
	}

	if err := db.LoadSettings(); err != nil {
		return err
	}

	if err := db.LoadDeployments(); err != nil {
		return err
	}

	if err := db.LoadCodeReviews(); err != nil {
		return err
	}

	return nil
}

func (db *FileMemoryDb) SaveAll() error {
	if err := db.SaveUsers(); err != nil {
		return err
	}

	if err := db.SaveSettings(); err != nil {
		return err
	}

	if err := db.SaveDeployments(); err != nil {
		return err
	}

	if err := db.SaveCodeReviews(); err != nil {
		return err
	}

	return nil
}

func (db *FileMemoryDb) Insert(item Entity) error {
	switch i := item.(type) {
	case *User:
		db.Context.Users.maxId++
		i.Id = db.Context.Users.maxId
		db.Context.Users.Append(i)
	case *Setting:
		db.Context.Settings.maxId++
		i.Id = db.Context.Settings.maxId
		db.Context.Settings.Append(i)
	case *Deployment:
		db.Context.Deployments.maxId++
		i.Id = db.Context.Deployments.maxId
		db.Context.Deployments.Append(i)
	case *CodeReview:
		db.Context.CodeReviews.maxId++
		i.Id = db.Context.CodeReviews.maxId
		db.Context.CodeReviews.Append(i)
	default:
		return errors.New("not supported item type")
	}

	return nil
}

func (db *FileMemoryDb) GetUsers() []*User {
	return db.Context.Users.GetWithLock()
}

func (db *FileMemoryDb) LoadUsers() error {
	filePath := path.Join(db.storagePath, "users.json")
	data, err := readFormFile[User](filePath)
	if err != nil {
		return err
	}

	for _, d := range data {
		db.Context.Users.maxId = max(db.Context.Users.maxId, d.Id)
	}
	db.Context.Users.slice = data
	return nil
}

func (db *FileMemoryDb) SaveUsers() error {
	db.Context.Users.mu.Lock()
	defer db.Context.Users.mu.Unlock()
	filePath := path.Join(db.storagePath, "users.json")
	return writeToFile[User](filePath, db.Context.Users.slice)
}

func (db *FileMemoryDb) GetSettings() []*Setting {
	return db.Context.Settings.GetWithLock()
}

func (db *FileMemoryDb) LoadSettings() error {
	filePath := path.Join(db.storagePath, "settings.json")
	data, err := readFormFile[Setting](filePath)
	if err != nil {
		return err
	}

	for _, d := range data {
		db.Context.Settings.maxId = max(db.Context.Settings.maxId, d.Id)
	}
	db.Context.Settings.slice = data
	return nil
}

func (db *FileMemoryDb) SaveSettings() error {
	db.Context.Settings.mu.Lock()
	defer db.Context.Settings.mu.Unlock()
	filePath := path.Join(db.storagePath, "settings.json")
	return writeToFile[Setting](filePath, db.Context.Settings.slice)
}

func (db *FileMemoryDb) GetDeployments() []*Deployment {
	return db.Context.Deployments.GetWithLock()
}

func (db *FileMemoryDb) LoadDeployments() error {
	filePath := path.Join(db.storagePath, "deployments.json")
	data, err := readFormFile[Deployment](filePath)
	if err != nil {
		return err
	}

	for _, d := range data {
		db.Context.Deployments.maxId = max(db.Context.Deployments.maxId, d.Id)
	}
	db.Context.Deployments.slice = data
	return nil
}

func (db *FileMemoryDb) SaveDeployments() error {
	db.Context.Deployments.mu.Lock()
	defer db.Context.Deployments.mu.Unlock()
	filePath := path.Join(db.storagePath, "deployments.json")
	return writeToFile[Deployment](filePath, db.Context.Deployments.slice)
}

func (db *FileMemoryDb) GetCodeReviews() []*CodeReview {
	return db.Context.CodeReviews.GetWithLock()
}

func (db *FileMemoryDb) LoadCodeReviews() error {
	filePath := path.Join(db.storagePath, "code_reviews.json")
	data, err := readFormFile[CodeReview](filePath)
	if err != nil {
		return err
	}

	for _, d := range data {
		db.Context.CodeReviews.maxId = max(db.Context.CodeReviews.maxId, d.Id)
	}
	db.Context.CodeReviews.slice = data
	return nil
}

func (db *FileMemoryDb) SaveCodeReviews() error {
	db.Context.CodeReviews.mu.Lock()
	defer db.Context.CodeReviews.mu.Unlock()
	filePath := path.Join(db.storagePath, "code_reviews.json")
	return writeToFile[CodeReview](filePath, db.Context.CodeReviews.slice)
}

func readFormFile[T any](filePath string) ([]*T, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		if err := os.WriteFile(filePath, []byte("[]"), 0644); err != nil {
			return nil, err
		}
		return make([]*T, 0), nil
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var res []*T
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func writeToFile[T any](filePath string, data []*T) error {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, bytes, 0644)
}

func max(a UniqId, b UniqId) UniqId {
	if a > b {
		return a
	}
	return b
}
