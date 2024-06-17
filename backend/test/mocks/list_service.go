package mocks

import (
	"cine/entity/model"
	"cine/service"
	"context"
	"github.com/google/uuid"
)

var _ service.ListService = (*ListServiceMock)(nil)

type ListServiceMock struct {
	CreateListFn           func(ctx context.Context, ownerID uuid.UUID, title string) (*model.List, error)
	DeleteListFn           func(ctx context.Context, ownerID uuid.UUID, id uuid.UUID) error
	UpdateListFn           func(ctx context.Context, ownerID uuid.UUID, id uuid.UUID, listU *model.ListU) (*model.List, error)
	AddMemberToListFn      func(ctx context.Context, ownerID uuid.UUID, listID uuid.UUID, userID uuid.UUID) error
	RemoveMemberFromListFn func(ctx context.Context, ownerID uuid.UUID, listID uuid.UUID, userID uuid.UUID) error
	GetAllListsFn          func(ctx context.Context, memberID uuid.UUID) ([]*model.DetailedList, error)
	GetPublicListsFn       func(ctx context.Context, userID uuid.UUID) ([]*model.DetailedList, error)
	GetDetailedListFn      func(ctx context.Context, memberID uuid.UUID, id uuid.UUID) (*model.DetailedList, error)
	AddMovieToListFn       func(ctx context.Context, memberID uuid.UUID, listID uuid.UUID, ref int) error
	RemoveMovieFromListFn  func(ctx context.Context, memberID uuid.UUID, listID uuid.UUID, ref int) error
	AddShowToListFn        func(ctx context.Context, memberID uuid.UUID, listID uuid.UUID, ref int) error
	RemoveShowFromListFn   func(ctx context.Context, memberID uuid.UUID, listID uuid.UUID, ref int) error
}

func NewListService() *ListServiceMock {
	return &ListServiceMock{}
}

func (m *ListServiceMock) CreateList(ctx context.Context, ownerID uuid.UUID, title string) (*model.List, error) {
	if m.CreateListFn != nil {
		return m.CreateListFn(ctx, ownerID, title)
	}
	return &model.List{}, nil
}

func (m *ListServiceMock) DeleteList(ctx context.Context, ownerID uuid.UUID, id uuid.UUID) error {
	if m.DeleteListFn != nil {
		return m.DeleteListFn(ctx, ownerID, id)
	}
	return nil
}

func (m *ListServiceMock) UpdateList(ctx context.Context, ownerID uuid.UUID, id uuid.UUID, listU *model.ListU) (*model.List, error) {
	if m.UpdateListFn != nil {
		return m.UpdateListFn(ctx, ownerID, id, listU)
	}
	return &model.List{}, nil
}

func (m *ListServiceMock) AddMemberToList(ctx context.Context, ownerID uuid.UUID, listID uuid.UUID, userID uuid.UUID) error {
	if m.AddMemberToListFn != nil {
		return m.AddMemberToListFn(ctx, ownerID, listID, userID)
	}
	return nil
}

func (m *ListServiceMock) RemoveMemberFromList(ctx context.Context, ownerID uuid.UUID, listID uuid.UUID, userID uuid.UUID) error {
	if m.RemoveMemberFromListFn != nil {
		return m.RemoveMemberFromListFn(ctx, ownerID, listID, userID)
	}
	return nil
}

func (m *ListServiceMock) GetAllLists(ctx context.Context, memberID uuid.UUID) ([]*model.DetailedList, error) {
	if m.GetAllListsFn != nil {
		return m.GetAllListsFn(ctx, memberID)
	}
	return []*model.DetailedList{}, nil
}

func (m *ListServiceMock) GetPublicLists(ctx context.Context, userID uuid.UUID) ([]*model.DetailedList, error) {
	if m.GetPublicListsFn != nil {
		return m.GetPublicListsFn(ctx, userID)
	}
	return []*model.DetailedList{}, nil
}

func (m *ListServiceMock) GetDetailedList(ctx context.Context, memberID uuid.UUID, id uuid.UUID) (*model.DetailedList, error) {
	if m.GetDetailedListFn != nil {
		return m.GetDetailedListFn(ctx, memberID, id)
	}
	return &model.DetailedList{}, nil
}

func (m *ListServiceMock) AddMovieToList(ctx context.Context, memberID uuid.UUID, listID uuid.UUID, ref int) error {
	if m.AddMovieToListFn != nil {
		return m.AddMovieToListFn(ctx, memberID, listID, ref)
	}
	return nil
}

func (m *ListServiceMock) RemoveMovieFromList(ctx context.Context, memberID uuid.UUID, listID uuid.UUID, ref int) error {
	if m.RemoveMovieFromListFn != nil {
		return m.RemoveMovieFromListFn(ctx, memberID, listID, ref)
	}
	return nil
}

func (m *ListServiceMock) AddShowToList(ctx context.Context, memberID uuid.UUID, listID uuid.UUID, ref int) error {
	if m.AddShowToListFn != nil {
		return m.AddShowToListFn(ctx, memberID, listID, ref)
	}
	return nil
}

func (m *ListServiceMock) RemoveShowFromList(ctx context.Context, memberID uuid.UUID, listID uuid.UUID, ref int) error {
	if m.RemoveShowFromListFn != nil {
		return m.RemoveShowFromListFn(ctx, memberID, listID, ref)
	}
	return nil
}
