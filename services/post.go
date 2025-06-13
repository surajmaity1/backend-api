package services

import (
	"context"
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

//
//func GetPost(postId int64) {
//	var posts []dtos.Post
//	_, db := config.DBConnect()
//
//	rows, err := db.Query("select * from posts where id = ?", postId)
//
//	if err != nil {
//		logrus.Error("Error while executing Query query")
//	}
//
//	for rows.Next() {
//		var post dtos.Post
//		if err := rows.Scan(&post.Id, &post.PostContent, &post.PostImage, &post.PostedBy); err != nil {
//			logrus.Error("Error while executing Scan query")
//		}
//		posts = append(posts, post)
//	}
//
//	fmt.Println(posts)
//
//	if err := rows.Err(); err != nil {
//		logrus.Error("Error while executing Err query")
//	}
//	defer rows.Close()
//}
//
//func AddPost(post dtos.Post) (int64, error) {
//	db := sql.DB{}
//	result, err := db.Exec(constants.POST_INSERT_QUERY, post.PostContent, post.PostImage, post.PostedBy)
//	if err != nil {
//		return 0, fmt.Errorf("addPost: %v", err)
//	}
//
//	id, err := result.LastInsertId()
//	if err != nil {
//		return 0, fmt.Errorf("addPost: %v", err)
//	}
//	return id, nil
//}
