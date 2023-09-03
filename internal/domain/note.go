package domain

import (
	"time"
)

type Note struct {
	id        string
	userID    string
	title     string
	content   string
	createdAt time.Time
	updatedAt time.Time
}

func NewNote(noteID, userID, title, content string) (*Note, error) {
	note := &Note{
		id:      noteID,
		userID:  userID,
		title:   title,
		content: content,
	}
	return note, nil
}

func (n *Note) Copy() *Note {
	return &Note{
		id:        n.id,
		userID:    n.userID,
		title:     n.title,
		content:   n.content,
		createdAt: n.createdAt,
		updatedAt: n.updatedAt,
	}
}

func (n *Note) ID() string {
	return n.id
}

func (n *Note) SetID(id string) {
	n.id = id
}

func (n *Note) UserID() string {
	return n.userID
}

func (n *Note) SetUserID(userID string) {
	n.userID = userID
}

func (n *Note) Title() string {
	return n.title
}

func (n *Note) SetTitle(title string) {
	n.title = title
}

func (n *Note) Content() string {
	return n.content
}

func (n *Note) SetContent(content string) {
	n.content = content
}

func (n *Note) CreatedAt() time.Time {
	return n.createdAt
}

func (n *Note) SetCreatedAt(createdAt time.Time) {
	n.createdAt = createdAt
}

func (n *Note) UpdatedAt() time.Time {
	return n.updatedAt
}

func (n *Note) SetUpdatedAt(updatedAt time.Time) {
	n.updatedAt = updatedAt
}

type SearchCriteria struct {
	Title    string
	Content  string
	FromDate time.Time
	ToDate   time.Time
}

type CreateNoteResult struct {
	ID  string
	Err error
}
