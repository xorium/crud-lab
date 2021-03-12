package postgres

import (
	"context"
	"crud-lab/pkg/domain/entities"
	"crud-lab/pkg/filters"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type User struct {
	ID          int `gorm:"primaryKey"`
	UserName    string
	Password    string
	FirstName   string
	LastName    string
	PhoneNumber string
	Email       string
	Image       string
	Settings    datatypes.JSON
	CompanyID   int
	//Company       *Company `gorm:"foreignKey:CompanyID"`
	LocationID int
	//Location      *Location       `gorm:"foreignKey:LocationID"`
	//Roles         []*Role         `gorm:"many2many:user_role;joinForeignKey:UsersID;JoinReferences:roles_id"`
	//Subscriptions []*Subscription `gorm:"foreignKey:UserID"`
}

type UserRepo struct {
	db *gorm.DB
}

func (u *UserRepo) List(
	ctx context.Context, filters filters.Filter,
) (users []*entities.User, err error) {
	// ...
	return users, err
}

func (u *UserRepo) Get(ctx context.Context, id string) (user *entities.User, err error) {
	// ...
	return user, err
}

func (u *UserRepo) Create(ctx context.Context, user *entities.User) (err error) {
	// ...
	return err
}

func (u *UserRepo) Update(ctx context.Context, user *entities.User) (err error) {
	// ...
	return err
}

func (u *UserRepo) Delete(ctx context.Context, id string) (err error) {
	// ...
	return err
}
