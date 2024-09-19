package repositorys

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type Users struct {
	database *sql.DB
}

func NewUsersRepository(database *sql.DB) *Users {
	return &Users{database}
}

func (repository Users) CreateUser(user models.User) (uint64, error) {
	statement, err := repository.database.Prepare("insert into users (name, nick, email, password) values ($1, $2, $3, $4) returning id")
	if err != nil {
		return 0, err
	}

	defer statement.Close()

	var id int64
	err = statement.QueryRow(user.Name, user.Nick, user.Email, user.Password).Scan(&id)
	if err != nil {
		return 0, err
	}

	return uint64(id), nil
}

func (repository Users) ReadUsers(nameOrNick string) ([]models.User, error) {

	lines, err := repository.database.Query(
		`select id, name, nick, email, createdin from users where "name" ~* $1 or nick ~* $2`,
		nameOrNick, nameOrNick,
	)

	if err != nil {
		return nil, err
	}

	defer lines.Close()

	var searchedUsers []models.User

	for lines.Next() {
		var searchedUser models.User
		if err := lines.Scan(
			&searchedUser.Id,
			&searchedUser.Name,
			&searchedUser.Nick,
			&searchedUser.Email,
			&searchedUser.CreatedIn,
		); err != nil {
			return nil, err
		}
		searchedUsers = append(searchedUsers, searchedUser)
	}

	return searchedUsers, nil
}

func (repository Users) ReadUser(id uint64) (models.User, error) {

	line, err := repository.database.Query(
		`select id, name, nick, email, createdin from users where id = $1`,
		id,
	)
	if err != nil {
		return models.User{}, err
	}

	defer line.Close()

	var searchedUser models.User

	if line.Next() {
		if err := line.Scan(
			&searchedUser.Id,
			&searchedUser.Name,
			&searchedUser.Nick,
			&searchedUser.Email,
			&searchedUser.CreatedIn,
		); err != nil {
			return models.User{}, err
		}
	}

	return searchedUser, nil
}

func (repository Users) UpdateUser(id uint64, user models.User) error {
	statement, err := repository.database.Prepare("update users set name = $1, nick = $2, email = $3 where id = $4")
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(user.Name, user.Nick, user.Email, id); err != nil {
		return err
	}

	return nil
}

func (repository Users) DeleteUser(id uint64) error {
	statement, err := repository.database.Prepare("delete from users where id = $1")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(id); err != nil {
		return err
	}

	return nil
}

func (repository Users) SearchEmail(email string) (models.User, error) {
	line, err := repository.database.Query("select id, password from users where email = $1", email)
	if err != nil {
		return models.User{}, err
	}
	defer line.Close()

	var user models.User

	if line.Next() {
		if err = line.Scan(&user.Id, &user.Password); err != nil {
			return models.User{}, nil
		}
	}

	return user, nil
}

func (repository Users) FollowUser(user_id, follower_id uint64) error {
	statement, err := repository.database.Prepare("insert into followers (user_id, follower_id) values ($1, $2) on conflict (user_id, follower_id) do nothing")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(user_id, follower_id); err != nil {
		return err
	}

	return nil
}

func (repository Users) UnfollowUser(user_id, follower_id uint64) error {
	statement, err := repository.database.Prepare("delete from followers where user_id = $1 and follower_id = $2")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(user_id, follower_id); err != nil {
		return err
	}

	return nil
}

func (repository Users) UserFollowers(user_id uint64) ([]models.User, error) {
	lines, err := repository.database.Query(`
		select u.id, u.name, u.nick, u.email, u.createdin 
		from users u inner join followers f 
		on u.id = f.follower_id 
		where f.user_id = $1`, user_id)
	if err != nil {
		return nil, err
	}

	defer lines.Close()

	var followers []models.User

	for lines.Next() {
		var follower models.User

		if err := lines.Scan(
			&follower.Id,
			&follower.Name,
			&follower.Nick,
			&follower.Email,
			&follower.CreatedIn,
		); err != nil {
			return nil, err
		}

		followers = append(followers, follower)
	}

	return followers, nil
}

func (repository Users) Following(follower_id uint64) ([]models.User, error) {
	lines, err := repository.database.Query(`
		select u.id, u.name, u.nick, u.email, u.createdin 
		from users u inner join followers f 
		on u.id = f.user_id 
		where f.follower_id = $1`, follower_id)
	if err != nil {
		return nil, err
	}

	defer lines.Close()

	var followers []models.User

	for lines.Next() {
		var follower models.User

		if err := lines.Scan(
			&follower.Id,
			&follower.Name,
			&follower.Nick,
			&follower.Email,
			&follower.CreatedIn,
		); err != nil {
			return nil, err
		}

		followers = append(followers, follower)
	}

	return followers, nil
}

func (repository Users) SearchPassword(user_id uint64) (string, error) {
	line, err := repository.database.Query("select password from users where id = $1", user_id)
	if err != nil {
		return "", err
	}
	defer line.Close()

	var user models.User

	if line.Next() {
		if err = line.Scan(&user.Password); err != nil {
			return "", nil
		}
	}

	return user.Password, nil
}

func (repository Users) UpdatePassword (user_id uint64, newPassword string) error {
	statement, err := repository.database.Prepare("update users set password = $1 where id = $2")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(newPassword, user_id); err != nil{
		return err
	}
	return nil
}