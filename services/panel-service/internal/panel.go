package internal

import (
	"context"
	"regexp"
	"strconv"

	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jackc/pgx/v5/pgtype"
)

// Panel Model
type Panel struct {
	Id int64 `json:"id"`

	Name string `json:"name"`
	Description string `json:"description"`
	
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

func StringifyPanelId(id int64) string {
	return strconv.FormatInt(id, 10)
}

func DestringifyPanelId(reprId string) (int64, error) {
	id, err := strconv.ParseInt(reprId, 10, 64)
	if err != nil || id < 1 {
		return 0, NewServiceError(InvalidArgumentErrorCode, "invalid panel id")
	}
	return id, nil
}

// Model for creating panels
type PanelCreate struct {
	Name string `json:"name"`
	Description string `json:"description"`
}

func (p *PanelCreate) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Name, validation.Required, validation.Length(3, 32), validation.Match(regexp.MustCompile("^[^_]\\w+[^_]$"))),
		validation.Field(&p.Description, validation.Required, validation.Length(3, 512)),
	)
}

// Model for updating a panel
type PanelUpdate struct {
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

func (p *PanelUpdate) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Name, validation.NilOrNotEmpty, validation.Length(3, 32), validation.Match(regexp.MustCompile("^[^_]\\w+[^_]$"))),
		validation.Field(&p.Description, validation.NilOrNotEmpty, validation.Length(3, 512)),
	)
}

// Interface methods
type PanelService interface {
	PanelRepository

	GetPanelByName(ctx context.Context, name string) (*Panel, error)
	UpdatePanelByName(ctx context.Context, name string, data PanelUpdate) (*Panel, error)
	DeletePanelByName(ctx context.Context, name string) error
}

type PanelRepository interface {
	CreatePanel(ctx context.Context, data PanelCreate) (*Panel, error)
	GetPanel(ctx context.Context, id int64) (*Panel, error)
	GetPanelIdFromName(ctx context.Context, name string) (*int64, error)
	UpdatePanel(ctx context.Context, id int64, data PanelUpdate) (*Panel, error)
	DeletePanel(ctx context.Context, id int64) error
}