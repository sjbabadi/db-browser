package postgres

import (
	"database/sql"

	"github.com/sjbabadi/db-browser/pkg/models"
)

type GenreModel struct {
	DB *sql.DB
}

func (g *GenreModel) GetAll() ([]*models.Genre, error) {
	stmt := `SELECT * FROM genres`

	rows, err := g.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	genres := []*models.Genre{}

	for rows.Next() {
		genre := &models.Genre{}

		err := rows.Scan(&genre.ID, &genre.Name)
		if err != nil {
			return nil, err
		}

		genres = append(genres, genre)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return genres, nil
}
