package models

import (
	"github.com/jinzhu/gorm"
)

type Item struct {
	gorm.Model
	ListID      uint   `gorm:"index"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

// List struct stores a list of items
type List struct {
	gorm.Model
	Name   string `json:"name"`
	Items  []Item `json:"items"`
	UserID uint   `gorm:"index"`
}

// Add takes in a description and adds it to list
func (l *List) Add(desc string) *List {
	l.Items = append(l.Items, *l.NewItem(desc))
	return l
}

// User stores the name and a collection of lists
type User struct {
	gorm.Model
	Name  string `json:"name" gorm:"index"`
	Lists []List `json:"lists"`
}

func NewUser(name string) *User {
	u := &User{Name: name}
	db.Create(u)
	return u
}

func (u *User) NewList(name string) *List {
	l := &List{Name: name, UserID: u.ID}
	db.Create(l)
	return l
}

func (l *List) NewItem(desc string) *Item {
	i := &Item{Description: desc, ListID: l.ID}
	db.Create(i)
	return i
}

// Add adds a list to a user and is chainable
func (u *User) Add(list List) *User {
	u.Lists = append(u.Lists, list)
	return u
}
func (u *User) Save() *User {
	db.Update(u)
	return u
}
func FindUser(id int) *User {
	var u User
	db.Preload("Lists").Preload("Lists.Items").First(&u, uint(id))
	return &u
}
