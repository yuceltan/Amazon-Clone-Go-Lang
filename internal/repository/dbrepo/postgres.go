package dbrepo

import (
	"bytes"
	"context"
	"database/sql"
	"log"
	"time"

	"bitbucket.org/janpavtel/site/internal/models"
)

type postgresDBRepo struct {
	DB *sql.DB
}

func (pg * postgresDBRepo) LoadAllUsers() ([]models.User, error){
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	stmt := `
	select 
		id, first_name, last_name, email
	from 
		users
	`

	var users []models.User

	rows, err := pg.DB.QueryContext(ctx, stmt)
	if err != nil {
		log.Println("Can't query users from postgres", err)
		return users, err
	}
	defer rows.Close()

	for rows.Next(){
		var user models.User
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
		)

		if err != nil{
			log.Println("Can't scan result to user model", err)
			return users, err
		}

		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		log.Println("Rows are left in error stage", err)
		return users, err
	}

	return users, nil
}


func (pg * postgresDBRepo) LoadUserById(id int) (models.User, error){
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	stmt := `
	select 
		id, first_name, last_name, email
	from 
		users
	where 
		id = $1
	`

	var user models.User

	row := pg.DB.QueryRowContext(ctx, stmt, id)

	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
	)

	if err != nil{
		log.Println("Can't scan result to user model", err)
		return user, err
	}

	return user, nil
}

func (pg * postgresDBRepo) LoadUserByEmail(email string) (models.User, error){
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	stmt := `
	select 
		u.id, u.email, u.password, r.name
	from 
		users u LEFT JOIN users_roles r ON u.id = r.user_id
	where 
		u.email = $1
	`

	var user models.User

	rows, err := pg.DB.QueryContext(ctx, stmt, email)
	if err != nil {
		log.Println("Can't query users from postgres", err)
		return user, err
	}
	defer rows.Close()

	for rows.Next(){
		var role *string
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.Password,
			&role,
		)

		if err != nil{
			log.Println("Can't scan result to user model", err)
			return user, err
		}

		if role != nil {
			user.Roles = append(user.Roles, *role)
		}
		
	}

	err = rows.Err()
	if err != nil {
		log.Println("Rows are left in error stage", err)
		return user, err
	}

	return user, nil
}

func (pg * postgresDBRepo) AddUser(user models.User) (int, error){
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()


	stmt := `insert into users (first_name, last_name, 
		email, password, 
		created_at, updated_at) 
		values ($1, $2, $3, $4, $5, $6)
		RETURNING (id)
		`

	row := pg.DB.QueryRowContext(ctx, stmt,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		time.Now(),
		time.Now(),	
	)	

	var id int 

	err := row.Scan(
		&id,
	)

	if err != nil{
		log.Println("Can't scan result id", err)
		return id, err
	}

	return id, nil
}


func (pg * postgresDBRepo) UpdateUser(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()


	stmt := `UPDATE users 
	    SET 
		first_name = $1,
		last_name = $2, 
		email = $3, 
		updated_at = $4 
		WHERE id = $5
		`

	_, err := pg.DB.ExecContext(ctx, stmt,
		user.FirstName,
		user.LastName,
		user.Email,
		time.Now(),
		user.ID,	
	)

	if err != nil{
		log.Println("Can't update user", err)
		return  err
	}

	return nil
}

func (pg *postgresDBRepo) UpdatePost(post models.Post) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()


	stmt := `UPDATE posts 
	    SET 
		title = $1,
		body = $2, 
		updated_at = $3 
		WHERE id = $4
		`

	_, err := pg.DB.ExecContext(ctx, stmt,
		post.Title,
		post.Body,
		time.Now(),
		post.ID,	
	)

	if err != nil{
		log.Println("Can't update post", err)
		return  err
	}

	return nil
}

func (pg * postgresDBRepo) AddPost(post models.Post) (int, error){
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()


	stmt := `insert into posts (title, body, owner_id, 
		created_at, updated_at) 
		values ($1, $2, $3, $4, $5)
		RETURNING (id)
		`

	row := pg.DB.QueryRowContext(ctx, stmt,
		post.Title,
		post.Body,
		post.OwnerId,
		time.Now(),
		time.Now(),	
	)	

	var id int 

	err := row.Scan(
		&id,
	)

	if err != nil{
		log.Println("Can't scan result id", err)
		return id, err
	}

	return id, nil
}

func (pg *postgresDBRepo) LoadAllPostsByOwnerID(userId int) ([] models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	stmt := `
	select 
		id, title, body, owner_id, created_at, updated_at
	from 
		posts
	where
		owner_id = $1
	`

	var posts []models.Post

	rows, err := pg.DB.QueryContext(ctx, stmt, userId)
	if err != nil {
		log.Println("Can't query posts from postgres", err)
		return posts, err
	}
	defer rows.Close()

	for rows.Next(){
		var post models.Post
		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Body,
			&post.OwnerId,
			&post.CreatedAt,
			&post.UpdatedAt,
		)

		if err != nil{
			log.Println("Can't scan result to post model", err)
			return posts, err
		}

		posts = append(posts, post)
	}

	err = rows.Err()
	if err != nil {
		log.Println("Rows are left in error stage", err)
		return posts, err
	}

	return posts, nil
}

func (pg *postgresDBRepo) LoadPostById(postId int) (models.Post, error){
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	stmt := `
	select 
		id, title, body, owner_id, created_at, updated_at
	from 
		posts
	where 
		id = $1
	`

	var post models.Post

	row := pg.DB.QueryRowContext(ctx, stmt, postId)

	err := row.Scan(
		&post.ID,
		&post.Title,
		&post.Body,
		&post.OwnerId,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil{
		log.Println("Can't scan result to post model", err)
		return post, err
	}

	return post, nil
}


func (pg *postgresDBRepo) LoadAllPosts(limit int, offset int) ([] models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	stmt := `
	SELECT 
		id, title, body, owner_id, created_at, updated_at
	FROM 
		posts
	ORDER BY created_at
	LIMIT $1 OFFSET $2
	`

	var posts []models.Post

	rows, err := pg.DB.QueryContext(ctx, stmt,
		limit,
		offset,
	)

	if err != nil {
		log.Println("Can't query posts from postgres", err)
		return posts, err
	}
	defer rows.Close()

	for rows.Next(){
		var post models.Post
		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Body,
			&post.OwnerId,
			&post.CreatedAt,
			&post.UpdatedAt,
		)

		if err != nil{
			log.Println("Can't scan result to post model", err)
			return posts, err
		}

		posts = append(posts, post)
	}

	err = rows.Err()
	if err != nil {
		log.Println("Rows are left in error stage", err)
		return posts, err
	}

	return posts, nil
}

func (pg *postgresDBRepo) UpdateUserProfilePicture(userID int, fileName string, buffer *bytes.Buffer) error{
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()


	stmt := `insert into profile_images (user_id, name, value,
		created_at, updated_at) 
		values ($1, $2, $3, $4, $5)
		ON CONFLICT (user_id) DO UPDATE
        SET name = EXCLUDED.name,
			value = EXCLUDED.value,
			updated_at = EXCLUDED.updated_at	
		`

	_, err := pg.DB.ExecContext(ctx, stmt,
		userID,
		fileName,
		buffer.Bytes(),
		time.Now(),
		time.Now(),	
	)	

	if err != nil{
		log.Println("Can't insert profile_image", err)
		return  err
	}

	return nil
}

func (pg *postgresDBRepo) LoadProfilePicture(userID int) (*bytes.Buffer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	stmt := `
	select 
		value
	from 
		profile_images
	where 
		user_id = $1
	`

	byt := []byte{}

	row := pg.DB.QueryRowContext(ctx, stmt, userID)

	err := row.Scan(
		&byt,
	)

	if err != nil{
		log.Println("Can't scan result to bytes profile picture", err)
		return nil, err
	}

	
	return bytes.NewBuffer(byt), nil


}





