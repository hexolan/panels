package internal

import (
	"regexp"
	"context"
	"strconv"
	"encoding/json"
	"database/sql"
	"database/sql/driver"
	
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/go-ozzo/ozzo-validation/v4"
)

// Post Models
type Post struct {
	Id PostId `json:"id"`

	PanelId string `json:"panel_id"`
	AuthorId string `json:"author_id"`

	Title string `json:"title"`
	Content string `json:"content"`
	
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

type PostCreate struct {
	Title string `json:"title"`
	Content string `json:"content"`
}

func (p *PostCreate) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Title, validation.Required, validation.Length(3, 512), validation.Match(regexp.MustCompile("^[^_][\\w ]+[^_]$"))),
		validation.Field(&p.Content, validation.Required, validation.Length(3, 2048)),
	)
}

type PostUpdate struct {
	Title *string `json:"title,omitempty"`
	Content *string `json:"content,omitempty"`
}

func (p *PostUpdate) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Title, validation.NilOrNotEmpty, validation.Length(3, 512), validation.Match(regexp.MustCompile("^[^_][\\w ]+[^_]$"))),
		validation.Field(&p.Content, validation.NilOrNotEmpty, validation.Length(3, 2048)),
	)
}

// Service Interface Methods
type PostService interface {
	PostRepository
}

type PostRepository interface {
    CreatePost(ctx context.Context, panelId string, authorId string, data PostCreate) (*Post, error)
    GetPost(ctx context.Context, id PostId) (*Post, error)
    GetPanelPost(ctx context.Context, id PostId, panelId string) (*Post, error)
    UpdatePost(ctx context.Context, id PostId, data PostUpdate) (*Post, error)
    DeletePost(ctx context.Context, id PostId) error

    GetFeedPosts(ctx context.Context) ([]*Post, error)
    GetUserPosts(ctx context.Context, userId string) ([]*Post, error)
    GetPanelPosts(ctx context.Context, panelId string) ([]*Post, error)
}

type PostDBRepository interface {
	PostRepository

	DeletePostsByUser(ctx context.Context, userId string) ([]PostId, error)
	DeletePostsOnPanel(ctx context.Context, panelId string) ([]PostId, error)
}

// Converts IDs between int64 (base 10) internally and string (base 36) externally
type PostId struct {
	Id int64
}

func (pid *PostId) GetId() int64 {
	if pid == nil {
		return 0
	}
	return pid.Id
}

func (pid *PostId) GetReprId() string {
	if pid.GetId() == 0 {
		return ""
	}

	return strconv.FormatInt(pid.GetId(), 36)
}

func (pid *PostId) Scan(value interface{}) error {
	scnr := sql.NullInt64{}
	err := scnr.Scan(value)
	if err != nil {
		return NewServiceError(InvalidArgumentErrorCode, "failed to scan post id: must be of type int64")
	} else if scnr.Int64 < 1 {
		return NewServiceError(InvalidArgumentErrorCode, "invalid post id: value must be greater than 0")
	}

	pid.Id = scnr.Int64
	return nil
}

func (pid PostId) Value() (driver.Value, error) {
	if pid.GetId() == 0 {
		return nil, NewServiceError(InvalidArgumentErrorCode, "post id not provided (of default value)")
	}

	return driver.Value(pid.GetId()), nil
}

func (pid PostId) MarshalJSON() ([]byte, error) {
	return json.Marshal(pid.GetReprId())
}

func (pid *PostId) UnmarshalJSON(data []byte) error {
	// Attempt to unmarshal the representative id
	var repr_id string
	err := json.Unmarshal(data, &repr_id)
	if err != nil {
		return err
	}

	// Reconstruct the ID using the representative ID
	id, err := getIdFromRepr(repr_id)
	if err != nil {
		return err
	}

	pid.Id = *id
	return nil
}

func NewPostId(id int64) (*PostId, error) {
	if id < 1 {
		return nil, NewServiceError(InvalidArgumentErrorCode, "invalid post id: value must be greater than 0")
	}

	return &PostId{Id: id}, nil
}

func NewPostIdFromRepr(reprId string) (*PostId, error) {
	id, err := getIdFromRepr(reprId)
	if err != nil {
		return nil, err
	}

	return NewPostId(*id)
}

func getIdFromRepr(reprId string) (*int64, error) {
	id, err := strconv.ParseInt(reprId, 36, 64)
	return &id, err
}