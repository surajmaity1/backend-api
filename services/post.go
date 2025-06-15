package services

import (
	"context"
	"database/sql"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/surajmaity1/backend-api/config"
	"github.com/surajmaity1/backend-api/constants"
	"github.com/surajmaity1/backend-api/dtos"
)

func CreatePost(context context.Context, post dtos.PostRequest) (int64, error) {
	_, db := config.DBConnect()
	result, err := db.ExecContext(context, constants.POST_INSERT_QUERY, post.PostContent, post.PostImage, post.PostedBy)
	if err != nil {
		logrus.Error(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		logrus.Error(err)
	}

	return id, err
}

func GetPost(context context.Context, postId int64) (*dtos.Post, error) {
	_, db := config.DBConnect()
	var post dtos.Post
	err := db.QueryRowContext(context, constants.POST_SINGLE_FETCH_QUERY, postId).Scan(&post.Id, &post.PostContent, &post.PostImage, &post.PostedBy)

	if err != nil {
		logrus.Error(err)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
	}

	return &post, nil
}

func EditPost(context context.Context, postId int64, post dtos.PostRequest) (dtos.Post, error) {
	_, db := config.DBConnect()
	_, err := db.ExecContext(context, constants.POST_SINGLE_UPDATE_QUERY, post.PostContent, post.PostImage, post.PostedBy, postId)

	if err != nil {
		logrus.Error(err)
	}

	responseBody := dtos.Post{
		Id:          postId,
		PostContent: post.PostContent,
		PostImage:   post.PostImage,
		PostedBy:    post.PostedBy,
	}

	return responseBody, err
}
