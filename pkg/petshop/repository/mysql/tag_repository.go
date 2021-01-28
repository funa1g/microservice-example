package mysql

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog/log"

	"github.com/funa1g/microservice-example/pkg/petshop/domain"
)

type tagRepository struct {
	Conn *sql.DB
}

func NewTagRepository(Conn *sql.DB) domain.TagRepository {
	return &tagRepository{Conn}
}

func (tr *tagRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.Tag, err error) {
	rows, err := tr.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, optimizeError(err)
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			log.Error().Msg(errRow.Error())
		}
	}()

	result = make([]domain.Tag, 0)
	for rows.Next() {
		t := domain.Tag{}
		err = rows.Scan(
			&t.ID,
			&t.Name,
		)
		result = append(result, t)
	}
	return result, nil
}

func (tr *tagRepository) Store(ctx context.Context, t *domain.Tag) error {
	query := `INSERT INTO tags (name) VALUES (?)`
	stmt, err := tr.Conn.PrepareContext(ctx, query)
	if err != nil {
		return optimizeError(err)
	}
	_, err = stmt.ExecContext(ctx, t.Name)
	return optimizeError(err)
}

func (tr *tagRepository) GetByName(ctx context.Context, name string) (t domain.Tag, err error) {
	query := `SELECT id, name FROM tags WHERE name = ?`
	list, err := tr.fetch(ctx, query, name)
	if err != nil {
		return domain.Tag{}, optimizeError(err)
	}
	if len(list) > 0 {
		t = list[0]
	} else {
		return t, domain.ErrNotFound
	}
	return t, nil
}
