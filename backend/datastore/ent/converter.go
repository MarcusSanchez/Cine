package ent

import (
	"cine/datastore"
	"cine/datastore/ent/ent"
	"cine/entity/model"
	"errors"
)

var c converter

type converter struct{}

func (c converter) user(user *ent.User) *model.User {
	if user != nil {
		return &model.User{
			ID:             user.ID,
			DisplayName:    user.DisplayName,
			Username:       user.Username,
			Email:          user.Email,
			Password:       user.Password,
			ProfilePicture: user.ProfilePicture,
			CreatedAt:      user.CreatedAt,
			UpdatedAt:      user.UpdatedAt,
		}
	}
	return nil
}

func (c converter) users(users []*ent.User) []*model.User {
	result := make([]*model.User, 0, len(users))
	for _, user := range users {
		result = append(result, c.user(user))
	}
	return result
}

func (c converter) session(session *ent.Session) *model.Session {
	if session != nil {
		return &model.Session{
			ID:         session.ID,
			UserID:     session.UserID,
			CSRF:       session.Csrf,
			Token:      session.Token,
			Expiration: session.Expiration,
			CreatedAt:  session.CreatedAt,
			UpdatedAt:  session.UpdatedAt,
		}
	}
	return nil
}

func (c converter) sessions(sessions []*ent.Session) []*model.Session {
	result := make([]*model.Session, 0, len(sessions))
	for _, session := range sessions {
		result = append(result, c.session(session))
	}
	return result
}

func (c converter) comment(comment *ent.Comment) *model.Comment {
	if comment != nil {
		return &model.Comment{
			ID:           comment.ID,
			UserID:       comment.UserID,
			MediaID:      comment.MediaID,
			ReplyingToID: comment.ReplyingToID,
			Content:      comment.Content,
			CreatedAt:    comment.CreatedAt,
			UpdatedAt:    comment.UpdatedAt,
		}
	}
	return nil
}

func (c converter) comments(comments []*ent.Comment) []*model.Comment {
	result := make([]*model.Comment, 0, len(comments))
	for _, comment := range comments {
		result = append(result, c.comment(comment))
	}
	return result
}

func (c converter) like(like *ent.Like) *model.Like {
	if like != nil {
		return &model.Like{
			ID:        like.ID,
			UserID:    like.UserID,
			CommentID: like.CommentID,
			CreatedAt: like.CreatedAt,
			UpdatedAt: like.UpdatedAt,
		}
	}
	return nil
}

func (c converter) likes(likes []*ent.Like) []*model.Like {
	result := make([]*model.Like, 0, len(likes))
	for _, like := range likes {
		result = append(result, c.like(like))
	}
	return result
}

func (c converter) review(review *ent.Review) *model.Review {
	if review != nil {
		return &model.Review{
			ID:        review.ID,
			UserID:    review.UserID,
			MediaID:   review.MediaID,
			Content:   review.Content,
			Rating:    review.Rating,
			CreatedAt: review.CreatedAt,
			UpdatedAt: review.UpdatedAt,
		}
	}
	return nil
}

func (c converter) reviews(reviews []*ent.Review) []*model.Review {
	result := make([]*model.Review, 0, len(reviews))
	for _, review := range reviews {
		result = append(result, c.review(review))
	}
	return result
}

func (c converter) list(list *ent.List) *model.List {
	if list != nil {
		return &model.List{
			ID:        list.ID,
			OwnerID:   list.OwnerID,
			Title:     list.Title,
			Public:    list.Public,
			CreatedAt: list.CreatedAt,
			UpdatedAt: list.UpdatedAt,
		}
	}
	return nil
}

func (c converter) lists(lists []*ent.List) []*model.List {
	result := make([]*model.List, 0, len(lists))
	for _, list := range lists {
		result = append(result, c.list(list))
	}
	return result
}

func (c converter) media(media *ent.Media) *model.Media {
	if media != nil {
		return &model.Media{
			ID:           media.ID,
			Ref:          media.Ref,
			MediaType:    model.MediaType(media.MediaType),
			Overview:     media.Overview,
			BackdropPath: media.BackdropPath,
			Language:     media.Language,
			PosterPath:   media.PosterPath,
			ReleaseDate:  media.ReleaseDate,
			Title:        media.Title,
			CreatedAt:    media.CreatedAt,
			UpdatedAt:    media.UpdatedAt,
		}
	}
	return nil
}

func (c converter) medias(medias []*ent.Media) []*model.Media {
	result := make([]*model.Media, 0, len(medias))
	for _, media := range medias {
		result = append(result, c.media(media))
	}
	return result
}

func (c converter) error(err error) error {
	if err != nil {
		var (
			notFound    *ent.NotFoundError
			constraint  *ent.ConstraintError
			notLoaded   *ent.NotLoadedError
			validation  *ent.ValidationError
			notSingular *ent.NotSingularError
		)
		switch {
		case errors.As(err, &notFound):
			return datastore.ErrNotFound
		case errors.As(err, &validation):
			return datastore.Wrap(datastore.ErrValidation, err.Error())
		case errors.As(err, &constraint):
			return datastore.Wrap(datastore.ErrConstraint, err.Error())
		case errors.As(err, &notSingular), errors.As(err, &notLoaded):
			return datastore.Wrap(datastore.ErrInternal, err.Error())
		default:
			return datastore.Wrap(datastore.ErrInternal, err.Error())
		}
	}
	return nil
}
