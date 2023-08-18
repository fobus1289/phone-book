package services

import (
	"context"
	"database/sql"
	"errors"

	"github.com/fobus1289/phone-book/internal/common"
	"github.com/fobus1289/phone-book/internal/model"
	"github.com/fobus1289/phone-book/internal/repository"
)

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		db: db,
	}
}

type UserService struct {
	db *sql.DB
}

func (us *UserService) Create(ctx context.Context, dto model.UserDto) (int64, error) {

	const CREATE = `
		INSERT INTO users (login, password, name, age) VALUES (?,?,?,?);	
	`

	pwdHash, err := common.GenerateFromPassword(dto.Password)

	if err != nil {
		return 0, err
	}

	row, err := us.db.ExecContext(ctx, CREATE,
		dto.Login,
		pwdHash,
		dto.Name,
		dto.Age,
	)

	if err != nil {
		return 0, err
	}

	return row.LastInsertId()
}

type OutUserFindByName struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (us *UserService) FindByName(ctx context.Context, name string) (*OutUserFindByName, error) {

	const FIND_BY_NAME = `
		SELECT id,name,age FROM users WHERE name = ? LIMIT 1;
	`

	repo := repository.NewRepositroy[OutUserFindByName](us.db, ctx)

	return repo.FindOne(FIND_BY_NAME,
		func(outUser *OutUserFindByName, scan func(...any) error) error {
			return scan(
				&outUser.Id,
				&outUser.Name,
				&outUser.Age,
			)
		},
		name,
	)
}

type OutUserFindByLoginAndPassword struct {
	Id      int
	Login   string
	PwdHash string
}

func (us *UserService) FindByLoginAndPassword(ctx context.Context, dto model.UserSignIn) (string, error) {

	const FIND_BY_LOGIN = `
		SELECT id,login,password FROM users WHERE login = ? LIMIT 1;
	`

	repo := repository.NewRepositroy[OutUserFindByLoginAndPassword](us.db, ctx)

	entity, err := repo.FindOne(FIND_BY_LOGIN,
		func(outUser *OutUserFindByLoginAndPassword, scan func(...any) error) error {
			return scan(
				&outUser.Id,
				&outUser.Login,
				&outUser.PwdHash,
			)
		},
		dto.Login,
	)

	if err != nil {
		return "", err
	}

	if !common.CompareHashAndPassword(entity.PwdHash, dto.Password) {
		return "", errors.New("incorrect login or password")
	}

	return common.GenerateJWT[model.UserPayload](&model.UserPayload{
		UserId: entity.Id,
		Login:  entity.Login,
	})

}
