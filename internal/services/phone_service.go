package services

import (
	"context"
	"database/sql"
	"errors"

	"github.com/fobus1289/phone-book/internal/model"
	"github.com/fobus1289/phone-book/internal/repository"
)

func NewPhoneService(db *sql.DB) *PhoneService {
	return &PhoneService{
		db: db,
	}
}

type PhoneService struct {
	db *sql.DB
}

func (ps *PhoneService) Create(ctx context.Context, dto model.PhoneDto) (int64, error) {

	const CREATE = `
		INSERT INTO phones (phone, is_fax, description, user_id) VALUES (?,?,?,?);	
	`

	row, err := ps.db.ExecContext(ctx, CREATE,
		dto.Phone,
		dto.IsFax,
		dto.Description,
		dto.UserId,
	)

	if err != nil {
		return 0, err
	}

	return row.LastInsertId()
}

func (ps *PhoneService) Update(ctx context.Context, dto model.Phone) error {

	const UPDATE = `
		UPDATE phones SET phone = ?, is_fax = ?, description = ? 
			WHERE id = ? AND user_id = ?;
	`

	result, err := ps.db.ExecContext(ctx, UPDATE,
		dto.Phone,
		dto.IsFax,
		dto.Description,

		dto.Id,
		dto.UserId,
	)

	if err != nil {
		return err
	}

	aff, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if aff == 0 {
		return errors.New("record not found")
	}

	return nil
}

func (ps *PhoneService) FindByNumber(ctx context.Context, number string) ([]model.PhoneDto, error) {

	const FIND_BY_NUMBER = `
		SELECT 
			user_id,
			phone,
			description,
			is_fax
			FROM phones
		WHERE phone LIKE '%' || ? || '%';
	`

	repo := repository.NewRepositroy[model.PhoneDto](ps.db, ctx)

	return repo.Find(FIND_BY_NUMBER,
		func(phone *model.PhoneDto, scan func(...any) error) error {
			return scan(
				&phone.UserId,
				&phone.Phone,
				&phone.Description,
				&phone.IsFax,
			)
		},
		number,
	)
}

func (ps *PhoneService) Delete(ctx context.Context, dto model.PhoneDeleteDto) error {

	const DELETE = `
		DELETE FROM phones WHERE id = ? AND user_id = ?;
	`

	result, err := ps.db.ExecContext(ctx, DELETE, dto.Id, dto.UserId)

	if err != nil {
		return err
	}

	aff, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if aff == 0 {
		return errors.New("record not found")
	}

	return nil
}
