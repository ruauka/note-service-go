package services

import (
	"web/internal/adapters/storage"
	"web/internal/domain/enteties/dto"
	"web/internal/domain/enteties/model"
)

// tagService tag service struct.
type tagService struct {
	storage storage.TagStorage
}

// NewTagService tag service func builder.
func NewTagService(tagStorage storage.TagStorage) TagService {
	return &tagService{storage: tagStorage}
}

// CreateTag create tag.
func (t *tagService) CreateTag(tag *model.Tag, userID string) (*model.Tag, error) {
	return t.storage.CreateTag(tag, userID)
}

// GetTagByID get tag by ID.
func (t *tagService) GetTagByID(tagID, userID string) (*dto.TagResp, error) {
	return t.storage.GetTagByID(tagID, userID)
}

// GetAllTagsByUser get tag by user.
func (t *tagService) GetAllTagsByUser(userID string) ([]dto.TagsResp, error) {
	return t.storage.GetAllTagsByUser(userID)
}

// UpdateTag update tag by ID.
func (t *tagService) UpdateTag(tag *dto.TagUpdate, tagID string) error {
	return t.storage.UpdateTag(tag, tagID)
}

// DeleteTag delete tag by ID.
func (t *tagService) DeleteTag(tagID, userID string) (int, error) {
	return t.storage.DeleteTag(tagID, userID)
}
