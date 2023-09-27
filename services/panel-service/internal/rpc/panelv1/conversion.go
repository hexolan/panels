package panelv1

import (
	"encoding/json"

	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	"github.com/hexolan/panels/panel-service/internal"
)

// Panel -> Protobuf 'Panel'
func PanelToProto(panel *internal.Panel) *Panel {
	proto := Panel{
		Id: internal.StringifyPanelId(panel.Id),
		Name: panel.Name,
		Description: panel.Description,
		CreatedAt: timestamppb.New(panel.CreatedAt.Time),
	}

	// convert nullable attributes to PB form (if present)
	if panel.UpdatedAt.Valid == true {
		proto.UpdatedAt = timestamppb.New(panel.UpdatedAt.Time)
	}

	return &proto
}

// Protobuf 'Panel' -> Panel
func PanelFromProto(proto *Panel) (*internal.Panel, error) {
	marshalled, err := json.Marshal(proto)
	if err != nil {
		return nil, err
	}

	var panel internal.Panel
	err = json.Unmarshal(marshalled, &panel)
	if err != nil {
		return nil, err
	}

	return &panel, nil
}

// Protobuf 'PanelMutable' -> PanelCreate
func PanelCreateFromProto(proto *PanelMutable) internal.PanelCreate {
	return internal.PanelCreate{
		Name: proto.GetName(),
		Description: proto.GetDescription(),
	}
}

// Protobuf 'PanelMutable' -> PanelUpdate
func PanelUpdateFromProto(proto *PanelMutable) internal.PanelUpdate {
	return internal.PanelUpdate{
		Name: proto.Name,
		Description: proto.Description,
	}
}