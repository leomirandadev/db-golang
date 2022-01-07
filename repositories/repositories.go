package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
	"github.com/leomirandadev/db-golang/repositories/users"
)

// Container modelo para exportação dos repositórios instanciados
type Container struct {
	User users.UserRepository
}

// Options struct de opções para a criação de uma instancia dos serviços
type Options struct {
	WriterSqlx *sqlx.DB
	ReaderSqlx *sqlx.DB
	WriterGorm *gorm.DB
	ReaderGorm *gorm.DB
}

// New cria uma nova instancia dos repositórios
func New(opts Options) *Container {
	return &Container{
		User: users.NewSqlxRepository(opts.WriterSqlx, opts.ReaderSqlx),
	}
}
