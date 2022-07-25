package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/m4rc0nd35/db-golang/repositories/users"
)

type Container struct {
	User users.IUserRepository
}

type Options struct {
	WriterGorm *gorm.DB
	ReaderGorm *gorm.DB
}

func New(opts Options) *Container {
	return &Container{
		User: users.NewGormRepository(opts.WriterGorm, opts.ReaderGorm),
	}
}
