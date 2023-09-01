package models

import (
	"database/sql"
	"errors"
	"time"
)

type Snap struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnapModel struct {
	DB *sql.DB
}

func (m *SnapModel) Insert(title, content string, expires int) (int, error) {
	queryStatement := `INSERT INTO snaps (title, content, created, expires)
						VALUES (?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(queryStatement, title, content, expires)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *SnapModel) Get(id int) (*Snap, error) {
	queryStatement := `SELECT id, title, content, created, expires FROM snaps
						WHERE expires > UTC_TIMESTAMP() AND id = ?`
	row := m.DB.QueryRow(queryStatement, id)

	s := &Snap{}

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}

func (m *SnapModel) Latest() ([]*Snap, error) {
	queryStatement := `SELECT id, title, content, created, expires FROM snaps 
						WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10`

	rows, err := m.DB.Query(queryStatement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	snaps := []*Snap{}

	for rows.Next() {
		s := &Snap{}

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snaps = append(snaps, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return snaps, nil
}
