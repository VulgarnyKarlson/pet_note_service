package domain

import (
	"context"
	"time"
)

type Note struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type NoteServicePort interface {
	Create(ctx context.Context, user *User, notesChan chan *Note) (noteIDsChan chan string, doneChan chan struct{}, errChan chan error)
	ReadByID(ctx context.Context, user *User, id string) (*Note, error)
	Update(ctx context.Context, user *User, note *Note) error
	Delete(ctx context.Context, user *User, id string) (bool, error)
	Search(ctx context.Context, user *User, criteria *SearchCriteria) ([]*Note, error)
}
