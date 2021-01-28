package mysql

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog/log"

	"github.com/funa1g/microservice-example/pkg/petshop/domain"
)

type petRepository struct {
	Conn *sql.DB
}

func NewPetRepository(Conn *sql.DB) domain.PetRepository {
	return &petRepository{Conn}
}

func (pr *petRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.Pet, err error) {
	rows, err := pr.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, optimizeError(err)
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			log.Error().Msg(errRow.Error())
		}
	}()

	result = make([]domain.Pet, 0)
	for rows.Next() {
		p := domain.Pet{}
		err = rows.Scan(
			&p.ID,
			&p.Name,
			&p.Tag,
		)
		result = append(result, p)
	}
	return result, nil
}

func (pr *petRepository) GetList(ctx context.Context, limit int) ([]domain.Pet, error) {
	query := `SELECT p.id, p.name, t.name FROM pets p JOIN tags t ON t.id = p.tag_id LIMIT ?`
	list, err := pr.fetch(ctx, query, limit)
	return list, optimizeError(err)
}

func (pr *petRepository) Store(ctx context.Context, p *domain.PetOrigin) (int64, error) {
	query := `INSERT INTO pets (name, tag_id) VALUES (?, ?)`
	stmt, err := pr.Conn.PrepareContext(ctx, query)
	if err != nil {
		return 0, optimizeError(err)
	}
	res, err := stmt.ExecContext(ctx, p.Name, p.TagId)
	if err != nil {
		return 0, optimizeError(err)
	}
	return res.LastInsertId()
}

func (pr *petRepository) GetById(ctx context.Context, id int) (p domain.Pet, err error) {
	query := `SELECT p.id, p.name, t.name FROM pets p JOIN tags t ON t.id = p.tag_id WHERE p.id = ?`
	list, err := pr.fetch(ctx, query, id)
	if err != nil {
		return domain.Pet{}, optimizeError(err)
	}
	if len(list) > 0 {
		p = list[0]
	} else {
		return p, domain.ErrNotFound
	}
	return p, nil
}
