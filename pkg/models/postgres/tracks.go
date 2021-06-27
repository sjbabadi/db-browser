package postgres

import (
	"database/sql"

	"github.com/sjbabadi/db-browser/pkg/models"
)

type TrackModel struct {
	DB *sql.DB
}

func (t *TrackModel) GetAll() ([]*models.Track, error) {
	stmt := `SELECT t.id, t.name, a.title, mt.name, g.name, composer, milliseconds, bytes, unit_price
		FROM tracks t
		INNER JOIN albums a ON a.id = t.album_id
		INNER JOIN media_types mt ON mt.id = t.media_type_id
		INNER JOIN genres g ON g.id = t.genre_id		
		`

	rows, err := t.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tracks := []*models.Track{}

	for rows.Next() {
		track := &models.Track{}

		err := rows.Scan(
			&track.ID,
			&track.Name,
			&track.Album,
			&track.MediaType,
			&track.Genre,
			&track.Composer,
			&track.Milliseconds,
			&track.Bytes,
			&track.UnitPrice,
		)
		if err != nil {
			return nil, err
		}

		tracks = append(tracks, track)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tracks, nil
}
