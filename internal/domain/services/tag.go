package services

import (
	"web/internal/adapters/storage"
	"web/internal/domain/enteties/dto"
	"web/internal/domain/enteties/model"
)

type tagService struct {
	storage storage.TagStorage
	// logger
}

func NewTagService(db storage.TagStorage) TagService {
	return &tagService{storage: db}
}

func (t *tagService) CreateTag(tag *model.Tag, userID string) (*model.Tag, error) {
	return t.storage.CreateTag(tag, userID)
}

func (t *tagService) GetTagByID(tagID, userID string) (*dto.TagResp, error) {
	return t.storage.GetTagByID(tagID, userID)
}

func (t *tagService) GetAllTagsByUser(userID string) ([]dto.TagsResp, error) {
	return t.storage.GetAllTagsByUser(userID)
}

func (t *tagService) UpdateTag(tag *dto.TagUpdate, tagID string) error {
	return t.storage.UpdateTag(tag, tagID)
}

func (t *tagService) DeleteTag(tagID, userID string) (int, error) {
	return t.storage.DeleteTag(tagID, userID)
}
