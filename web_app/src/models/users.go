package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"web_app/src/config"
	"web_app/src/requests"
)

type User struct {
	Id        uint64    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Nick      string    `json:"nick"`
	CreatedIn time.Time `json:"createdin"`
	Followers []User    `json:"followers"`
	Following []User    `json:"following"`
	Posts     []Post    `json:"posts"`
}

func SearchUsersCompleteData(user_id uint64, r *http.Request) (User, error) {
	userChannel := make(chan User)
	followersChannel := make(chan []User)
	followingChannel := make(chan []User)
	postsChannel := make(chan []Post)

	go SearchUsersData(userChannel, user_id, r)
	go SearchUsersFollowers(followersChannel, user_id, r)
	go SearchUsersFollowingAccounts(followingChannel, user_id, r)
	go SearchUsersPosts(postsChannel, user_id, r)

	var (
		user      User
		followers []User
		following []User
		posts     []Post
	)

	for i := 0; i < 4; i++ {
		select {
		case loadUser := <-userChannel:
			if loadUser.Id == 0 {
				return User{}, errors.New("error searching this user")
			}

			user = loadUser

		case loadFollowers := <-followersChannel:
			if loadFollowers == nil {
				return User{}, errors.New("error searching this user's followers")
			}

			followers = loadFollowers

		case loadFollowing := <-followingChannel:
			if loadFollowing == nil {
				return User{}, errors.New("error searching who this user follow")
			}

			following = loadFollowing
			
		case loadPosts := <-postsChannel:
			if loadPosts == nil {
				return User{}, errors.New("error searching this user's posts")
			}

			posts = loadPosts

		}
	}

	user.Followers = followers
	user.Following = following
	user.Posts = posts

	return user, nil
}

// Chama a Api para resgatar dados de cadastro do usuário
func SearchUsersData(channel chan<- User, user_id uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d", config.ApiURL, user_id)
	response, err := requests.RequestAuthenticated(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- User{}
		return
	}

	defer response.Body.Close()

	var user User
	if err = json.NewDecoder(response.Body).Decode(&user); err != nil {
		channel <- User{}
		return
	}

	channel <- user
}

// Chama a Api para resgatar quem segue o usuário
func SearchUsersFollowers(channel chan<- []User, user_id uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/followers", config.ApiURL, user_id)
	response, err := requests.RequestAuthenticated(r, http.MethodGet, url, nil)
	
	if err != nil {
		channel <- nil
		return
	}

	defer response.Body.Close()

	var users []User
	if err = json.NewDecoder(response.Body).Decode(&users); err != nil {
		channel <- nil
		return
	}

	if users == nil {
		channel <- make([]User, 0)
		return
	}

	channel <- users
}

// Chama a Api para resgatar quem o usuário segue
func SearchUsersFollowingAccounts(channel chan<- []User, user_id uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/following", config.ApiURL, user_id)
	response, err := requests.RequestAuthenticated(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}

	defer response.Body.Close()

	var users []User
	if err = json.NewDecoder(response.Body).Decode(&users); err != nil {
		channel <- nil
		return
	}

	if users == nil {
		channel <- make([]User, 0)
		return
	}

	channel <- users
}

// Chama a Api para resgatar todas as postagens do usuário
func SearchUsersPosts(channel chan<- []Post, user_id uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/posts", config.ApiURL, user_id)
	response, err := requests.RequestAuthenticated(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}

	defer response.Body.Close()

	var posts []Post
	if err = json.NewDecoder(response.Body).Decode(&posts); err != nil {
		channel <- nil
		return
	}

	if posts == nil {
		channel <- make([]Post, 0)
		return
	}

	channel <- posts
}
