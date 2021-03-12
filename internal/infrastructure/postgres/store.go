package postgres

import (
	"crud-lab/pkg/config"
	"crud-lab/pkg/domain/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repo struct {
	UsersRepo repository.UsersRepo
}

func NewRepo(cfg config.Config) (*Repo, error) {
	rdb, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  cfg.RdbDsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	repo := &Repo{UsersRepo: &UserRepo{rdb}}
	if err := repo.migrate(rdb); err != nil {
		return nil, err
	}

	return repo, nil
}

func (r *Repo) migrate(rdb *gorm.DB) error {
	if err := rdb.AutoMigrate(&User{}); err != nil {
		return err
	}
	return nil
}

func (r *Repo) Users() repository.UsersRepo {
	return r.UsersRepo
}
