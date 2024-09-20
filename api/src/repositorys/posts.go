package repositorys

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type Posts struct {
	database *sql.DB
}

func NewPostsRepository(database *sql.DB) *Posts {
	return &Posts{database}
}

func (repository Posts) CreatePost(post models.Post) (uint64, error) {
	statement, err := repository.database.Prepare("insert into posts (title, content, autor_id) values ($1, $2, $3) returning id")
	if err != nil {
		return 0, err
	}

	defer statement.Close()

	var post_id int64
	err = statement.QueryRow(post.Title, post.Content, post.Autor_id).Scan(&post_id)
	if err != nil {
		return 0, err
	}

	return uint64(post_id), nil
}

// Realiza a busca das publicações dos usuários de quem solicitou segue e também as postagens do próprio usuário que solicitou
func (repository Posts) SearchPosts(user_ID uint64) ([]models.Post, error) {
	lines, err := repository.database.Query(`
		select distinct p.*, u.nick
		from posts p
		inner join users u
		on u.id = p.autor_id
		left join followers f
		on p.autor_id = f.user_id
		where u.id = $1 or f.follower_id = $2
		order by 1 desc
	`, user_ID, user_ID)
	
	if err != nil {
		return nil, err
	}
	
	defer lines.Close()

	var posts []models.Post
	for lines.Next(){
		var post models.Post
		if err := lines.Scan(
			&post.Id,
			&post.Title,
			&post.Content,
			&post.Autor_id,
			&post.Likes,
			&post.Created_in,
			&post.Autor_nick,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}
	return posts, nil
}

// Pesquisa um a publicação específica no banco de dados
func (repository Posts) SearchPost(post_id uint64) (models.Post, error) {
	line, err := repository.database.Query(`
		select p.*, u.nick
		from posts p inner join users u
		on u.id = p.autor_id
		where p.id = $1 
	`, post_id)

	if err != nil {
		fmt.Println(err)
		return models.Post{}, err
	}

	defer line.Close()

	var post models.Post

	if line.Next() {
		if err = line.Scan(
			&post.Id,
			&post.Title,
			&post.Content,
			&post.Autor_id,
			&post.Likes,
			&post.Created_in,
			&post.Autor_nick,
		); err != nil {
			return models.Post{}, err
		}
	}

	return post, nil
}

// Atualiza uma publicação no banco de dados
func (repository Posts) UpdatePost(post_id uint64, post models.Post) error {
	statement, err := repository.database.Prepare(`update posts set title = $1, content = $2 where id = $3`)
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(post.Title, post.Content, post_id); err != nil {
		return err
	}

	return nil
}

// Deleta uma publicação do banco de dados
func (repository Posts) DeletePost(post_id uint64) error {
	statement, err := repository.database.Prepare("delete from posts where id = $1")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(post_id); err != nil {
		return err
	}

	return nil
}

func (repository Posts) SearchPostsFromUser(user_id uint64) ([]models.Post, error) {
	lines, err := repository.database.Query(`
		select p.*, u.nick
		from posts p
		join users u
		on u.id = p.autor_id
		where p.autor_id = $1
		order by 1 desc
	`, user_id)
	
	if err != nil {
		return nil, err
	}
	
	defer lines.Close()

	var posts []models.Post
	for lines.Next(){
		var post models.Post
		if err := lines.Scan(
			&post.Id,
			&post.Title,
			&post.Content,
			&post.Autor_id,
			&post.Likes,
			&post.Created_in,
			&post.Autor_nick,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}
	return posts, nil
}

func (repository Posts) Like(post_id uint64) error {
	statement, err := repository.database.Prepare(`
		update into posts set likes = likes + 1 where id = $1 
	`)
	if err != nil {
		return err
	}
	
	defer statement.Close()

	if _, err = statement.Exec(post_id); err != nil {
		return err
	}

	return nil
}

func (repository Posts) Unlike(post_id uint64) error {
	statement, err := repository.database.Prepare(`
		update into posts set likes = 
		case when likes > 0 then likes - 1
		else 0 end
		where id = $1 
	`)
	if err != nil {
		return err
	}
	
	defer statement.Close()

	if _, err = statement.Exec(post_id); err != nil {
		return err
	}
	
	return nil
}