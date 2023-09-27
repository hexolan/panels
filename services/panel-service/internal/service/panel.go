package service

import (
	"context"

	"github.com/hexolan/panels/panel-service/internal"
	"github.com/hexolan/panels/panel-service/internal/kafka"
)

type panelService struct {
	events kafka.PanelEventProducer

	repo internal.PanelRepository
}

func NewPanelService(events kafka.PanelEventProducer, repo internal.PanelRepository) internal.PanelService {
	return panelService{
		events: events,
		repo: repo,
	}
}

func (srv panelService) GetPanelIdFromName(ctx context.Context, name string) (*int64, error) {
	return srv.repo.GetPanelIdFromName(ctx, name)
}

func (srv panelService) CreatePanel(ctx context.Context, data internal.PanelCreate) (*internal.Panel, error) {
	// Validate the data
	err := data.Validate()
	if err != nil {
		return nil, internal.NewServiceErrorf(internal.InvalidArgumentErrorCode, "invalid argument: %s", err.Error())
	}

	// Create the panel
	panel, err := srv.repo.CreatePanel(ctx, data)

	// Dispatch panel created event
	if err == nil {
		srv.events.DispatchCreatedEvent(*panel)
	}

	return panel, err
}

func (srv panelService) GetPanel(ctx context.Context, id int64) (*internal.Panel, error) {
	return srv.repo.GetPanel(ctx, id)
}

func (srv panelService) GetPanelByName(ctx context.Context, name string) (*internal.Panel, error) {
	// Get the panel ID from the provided name
	id, err := srv.GetPanelIdFromName(ctx, name)
	if err != nil {
		return nil, err
	}
	
	// Pass to service method for GetPanel (by id).
	return srv.GetPanel(ctx, *id)
}

func (srv panelService) UpdatePanel(ctx context.Context, id int64, data internal.PanelUpdate) (*internal.Panel, error) {
	// Validate the data.
	if data == (internal.PanelUpdate{}) {
		return nil, internal.NewServiceError(internal.InvalidArgumentErrorCode, "no data values provided")
	}

	err := data.Validate()
	if err != nil {
		return nil, internal.NewServiceErrorf(internal.InvalidArgumentErrorCode, "invalid argument: %s", err.Error())
	}

	// Perform some checks if the target is a primary panel
	if id == 1 {
		if data.Name != nil && *data.Name != "" {
			return nil, internal.NewServiceError(internal.ForbiddenErrorCode, "cannot modify name of primary panel")
		}
	}

	// Update the panel
	panel, err := srv.repo.UpdatePanel(ctx, id, data)

	// Dispatch panel updated event
	if err == nil {
		srv.events.DispatchUpdatedEvent(*panel)
	}

	return panel, err
}

func (srv panelService) UpdatePanelByName(ctx context.Context, name string, data internal.PanelUpdate) (*internal.Panel, error) {
	// Get the panel ID from the provided name
	id, err := srv.GetPanelIdFromName(ctx, name)
	if err != nil {
		return nil, err
	}
	
	// Pass to service method for UpdatePanel (by id).
	return srv.UpdatePanel(ctx, *id, data)
}

func (srv panelService) DeletePanel(ctx context.Context, id int64) error {
	// Ensure the target is not the primary panel
	if id == 1 {
		return internal.NewServiceError(internal.ForbiddenErrorCode, "cannot delete primary panel")
	}

	// Delete the panel.
	err := srv.repo.DeletePanel(ctx, id)

	// Dispatch panel deleted event
	if err == nil {
		srv.events.DispatchDeletedEvent(id)
	}

	return err
}

func (srv panelService) DeletePanelByName(ctx context.Context, name string) error {
	// Get the panel ID from the provided name
	id, err := srv.GetPanelIdFromName(ctx, name)
	if err != nil {
		return err
	}

	// Pass to service method for DeletePanel (by id).
	return srv.DeletePanel(ctx, *id)
}