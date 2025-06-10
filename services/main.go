package services

import (
	"backend-api/constants"
	"backend-api/dtos"
	"database/sql"
	"fmt"
)

func addPost(post dtos.Post) (int64, error) {
	db := sql.DB{}
	result, err := db.Exec(constants.POST_INSERT_QUERY, post.PostContent, post.PostImage, post.PostedBy)
	if err != nil {
		return 0, fmt.Errorf("addPost: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addPost: %v", err)
	}
	return id, nil
}
