package users

import (
	"context"
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/leomirandadev/db-golang/entities"
)

type repoGorm struct {
	writer *gorm.DB
	reader *gorm.DB
}

func NewGormRepository(writer, reader *gorm.DB) UserRepository {
	return &repoGorm{writer: writer, reader: reader}
}

func (repo *repoGorm) Create(ctx context.Context, newUser entities.User) error {
	repo.writer.Table("users").Create(&newUser)

	if repo.writer.Error != nil {
		fmt.Println(repo.writer.Error)
		return errors.New("usuário não cadastrado")
	}

	return nil
}

func (repo *repoGorm) GetUserByEmail(ctx context.Context, email string) ([]entities.User, error) {
	var user []entities.User
	repo.reader.Table("users").Where("email = ?", email).Find(&user)

	if repo.reader.Error != nil {
		fmt.Println(repo.reader.Error)
		return nil, errors.New("usuário não encontrado")
	}

	return user, nil
}

func (repo *repoGorm) GetByID(ctx context.Context, ID int64) (*entities.UserWithoutPassword, error) {
	var user entities.UserWithoutPassword
	repo.reader.Table("users").Where("ID = ?", ID).First(&user)

	if repo.reader.Error != nil {
		fmt.Println(repo.reader.Error)
		return nil, errors.New("usuário não encontrado")
	}

	return &user, nil
}

func (repo *repoGorm) GetAll(ctx context.Context) ([]entities.UserWithoutPassword, error) {
	var user []entities.UserWithoutPassword
	repo.reader.Table("users").Find(&user)

	if repo.reader.Error != nil {
		fmt.Println(repo.reader.Error)
		return nil, errors.New("usuário não encontrado")
	}

	return user, nil
}
