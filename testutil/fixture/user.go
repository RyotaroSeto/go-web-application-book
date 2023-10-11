package fixture

import (
	"go_todo_app/entity"
	"math/rand"
	"strconv"
	"time"
)

// https://engineering.mercari.com/blog/entry/20220411-42fc0ba69c/

// https://engineering.mercari.com/blog/entry/20230630-32980896a4/ 同じ人が最近フィクスチャの投稿
func User(u *entity.User) *entity.User {
	result := &entity.User{
		ID:       entity.UserID(rand.Int()),
		Name:     "budougumi" + strconv.Itoa(rand.Int())[:5],
		Password: "password",
		Role:     "admin",
		Created:  time.Now(),
		Modified: time.Now(),
	}
	if u == nil {
		return result
	}
	if u.ID != 0 {
		result.ID = u.ID
	}
	if u.Name != "" {
		result.Name = u.Name
	}
	if u.Password != "" {
		result.Password = u.Password
	}
	if u.Role != "" {
		result.Role = u.Role
	}
	if !u.Created.IsZero() {
		result.Created = u.Created
	}
	if !u.Modified.IsZero() {
		result.Modified = u.Modified
	}
	return result
}
