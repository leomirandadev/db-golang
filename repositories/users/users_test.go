package users

import (
	"context"
	"testing"

	"github.com/leomirandadev/db-golang/configs"
)

func BenchmarkGetAllGorm(b *testing.B) {
	ctx := context.Background()
	connection := NewGormRepository(configs.GetReaderGorm(), configs.GetWriterGorm())
	for i := 0; i < b.N; i++ {
		connection.GetAll(ctx)
	}
}

func BenchmarkGetAllSqlx(b *testing.B) {
	ctx := context.Background()
	connection := NewSqlxRepository(configs.GetReaderSqlx(), configs.GetWriterSqlx())
	for i := 0; i < b.N; i++ {
		connection.GetAll(ctx)
	}
}
