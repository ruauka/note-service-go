package storage

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"

	"web/internal/domain/entities/dto"
	"web/internal/domain/entities/model"
	"web/internal/utils"
)

// tagStorage tag storage struct.
type tagStorage struct {
	db *sqlx.DB
}

// NewTagStorage tag storage func builder.
func NewTagStorage(pgDB *sqlx.DB) TagStorage {
	return &tagStorage{db: pgDB}
}

// CreateTag create tag in DB.
func (t *tagStorage) CreateTag(tag *model.Tag, userID string) (*model.Tag, error) {
	query := fmt.Sprintf("INSERT INTO %s (tagname, user_id) VALUES ($1, $2) RETURNING id", utils.TagsTable)
	if err := t.db.QueryRow(query, tag.TagName, userID).Scan(&tag.ID); err != nil {
		return nil, err
	}

	return tag, nil
}

// GetTagByID get tag by id from DB.
func (t *tagStorage) GetTagByID(tagID, userID string) (*dto.TagResp, error) {
	var tag dto.TagResp

	query := fmt.Sprintf("SELECT tagname FROM %s WHERE id=$1 AND user_id=$2", utils.TagsTable)
	if err := t.db.Get(&tag, query, tagID, userID); err != nil {
		return nil, err
	}

	return &tag, nil
}

// GetAllTagsByUser get all tags by user from DB.
func (t *tagStorage) GetAllTagsByUser(userID string) ([]dto.TagsResp, error) {
	var tags []dto.TagsResp

	query := fmt.Sprintf("SELECT id, tagname FROM %s WHERE user_id=$1", utils.TagsTable)
	if err := t.db.Select(&tags, query, userID); err != nil {
		return nil, err
	}

	return tags, nil
}

// UpdateTag update tag by id in DB.
func (t *tagStorage) UpdateTag(tag *dto.TagUpdate, tagID string) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if tag.TagName != nil {
		setValues = append(setValues, fmt.Sprintf("tagname=$%d", argID))
		args = append(args, *tag.TagName)
		argID++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", utils.TagsTable, setQuery, argID)
	args = append(args, tagID)

	_, err := t.db.Exec(query, args...)

	return err
}

// DeleteTag delete tag by id from DB.
func (t *tagStorage) DeleteTag(tagID, userID string) (int, error) {
	var id int

	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 AND user_id=$2 RETURNING id", utils.TagsTable)
	if err := t.db.QueryRow(query, tagID, userID).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
