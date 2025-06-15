package constants

const POST_INSERT_QUERY = `INSERT INTO posts (post_content, post_image, posted_by) VALUES (?, ?, ?)`
const POST_SINGLE_FETCH_QUERY = `SELECT * FROM posts WHERE id = ?`
const POST_SINGLE_UPDATE_QUERY = `UPDATE posts SET post_content = ?, post_image = ?, posted_by = ?  WHERE id = ?`
