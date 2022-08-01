package storage

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"

	"web/internal/domain/enteties/dto"
	"web/internal/domain/enteties/model"
	"web/internal/utils"
)

type tagStorage struct {
	db *sqlx.DB
}

func NewTagStorage(pgDB *sqlx.DB) TagStorage {
	return &tagStorage{db: pgDB}
}

func (t *tagStorage) CreateTag(tag *model.Tag, userID string) (*model.Tag, error) {
	query := fmt.Sprintf("INSERT INTO %s (tagname, user_id) VALUES ($1, $2) RETURNING id", utils.TagsTable)
	if err := t.db.QueryRow(query, tag.TagName, userID).Scan(&tag.ID); err != nil {
		return nil, err
	}

	return tag, nil
}

func (t *tagStorage) GetTagByID(tagID, userID string) (*dto.TagResp, error) {
	var tag dto.TagResp

	query := fmt.Sprintf("SELECT tagname FROM %s WHERE id=$1 AND user_id=$2", utils.TagsTable)
	if err := t.db.Get(&tag, query, tagID, userID); err != nil {
		return nil, err
	}

	return &tag, nil
}

func (t *tagStorage) GetAllTagsByUser(userID string) ([]dto.TagsResp, error) {
	var tags []dto.TagsResp

	query := fmt.Sprintf("SELECT id, tagname FROM %s WHERE user_id=$1", utils.TagsTable)
	if err := t.db.Select(&tags, query, userID); err != nil {
		return nil, err
	}

	return tags, nil
}

func (t *tagStorage) UpdateTag(tag *dto.TagUpdate, tagID string) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if tag.TagName != nil {
		setValues = append(setValues, fmt.Sprintf("tagname=$%d", argId))
		args = append(args, *tag.TagName)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", utils.TagsTable, setQuery, argId)
	args = append(args, tagID)

	_, err := t.db.Exec(query, args...)
	return err
}

func (t *tagStorage) DeleteTag(tagID, userID string) (int, error) {
	var id int

	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 AND user_id=$2 RETURNING id", utils.TagsTable)
	if err := t.db.QueryRow(query, tagID, userID).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
