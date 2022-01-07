package users

import (
	"context"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/leomirandadev/db-golang/entities"
)

type repoSqlx struct {
	writer *sqlx.DB
	reader *sqlx.DB
}

func NewSqlxRepository(writer, reader *sqlx.DB) UserRepository {
	return &repoSqlx{writer: writer, reader: reader}
}

func (repo *repoSqlx) Create(ctx context.Context, newUser entities.User) error {

	_, err := repo.writer.ExecContext(ctx, `
		INSERT INTO users (nick_name,name,email,password) VALUES (?, ?, ?, ?)
	`, newUser.NickName, newUser.Name, newUser.Email, newUser.Password)

	if err != nil {
		fmt.Println(err)
		return errors.New("usuário não cadastrado")
	}

	return nil
}

func (repo *repoSqlx) GetByID(ctx context.Context, ID int64) (*entities.UserWithoutPassword, error) {

	var user entities.UserWithoutPassword

	err := repo.reader.GetContext(ctx, &user, `
		SELECT 
			id,
			nick_name,
			name,
			email,
			created_at,
			updated_at
		FROM users 
		WHERE id=?
	`, ID)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("usuário não encontrado")
	}

	return &user, nil

}

func (repo *repoSqlx) GetAll(ctx context.Context) ([]entities.UserWithoutPassword, error) {

	var user []entities.UserWithoutPassword

	err := repo.reader.SelectContext(ctx, &user, `
		SELECT 
			id,
			nick_name,
			name,
			email,
			created_at,
			updated_at
		FROM users 
	`)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("usuário não encontrado")
	}

	return user, nil

}

func (repo *repoSqlx) GetUserByEmail(ctx context.Context, email string) ([]entities.User, error) {
	var user []entities.User

	err := repo.reader.GetContext(ctx, &user, `
		SELECT 
			id,
			nick_name,
			name,
			email,
			password,
			created_at,
			updated_at
		FROM users 
		WHERE email = ?
	`, email)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("usuário não encontrado")
	}

	return user, nil
}
