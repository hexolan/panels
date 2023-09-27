package rpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"

	"github.com/hexolan/panels/panel-service/internal"
	pb "github.com/hexolan/panels/panel-service/internal/rpc/panelv1"
)

type panelServer struct {
	pb.UnimplementedPanelServiceServer
	
	service internal.PanelService
}

func NewPanelServer(service internal.PanelService) panelServer {
	return panelServer{service: service}
}

func (svr *panelServer) CreatePanel(ctx context.Context, request *pb.CreatePanelRequest) (*pb.Panel, error) {
	// Ensure the required args are provided
	if request.GetData() == nil {
		return nil, status.Error(codes.InvalidArgument, "malformed request")
	}

	// Convert to business model
	data := pb.PanelCreateFromProto(request.GetData())
	
	// Attempt to create the panel
	panel, err := svr.service.CreatePanel(ctx, data)
	if err != nil {
		return nil, err
	}

	return pb.PanelToProto(panel), nil
}

func (svr *panelServer) GetPanel(ctx context.Context, request *pb.GetPanelByIdRequest) (*pb.Panel, error) {
	// Ensure the required args are provided
	if request.GetId() == "" {
		return nil, status.Error(codes.InvalidArgument, "panel id not provided")
	}

	// Convert to business model
	id, err := internal.DestringifyPanelId(request.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid panel id")
	}

	// Attempt to get the panel
	panel, err := svr.service.GetPanel(ctx, id)
	if err != nil {
		return nil, err
	}
	return pb.PanelToProto(panel), nil
}

func (svr *panelServer) GetPanelByName(ctx context.Context, request *pb.GetPanelByNameRequest) (*pb.Panel, error) {
	// Ensure the required args are provided
	var name string = request.GetName()
	if request.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid panel name")
	}

	// Attempt to delete the panel
	panel, err := svr.service.GetPanelByName(ctx, name)
	if err != nil {
		return nil, err
	}
	return pb.PanelToProto(panel), nil
}

func (svr *panelServer) UpdatePanel(ctx context.Context, request *pb.UpdatePanelByIdRequest) (*pb.Panel, error) {
	// Ensure the required args are provided
	if request.GetId() == "" {
		return nil, status.Error(codes.InvalidArgument, "panel id not provided")
	}

	// Convert to ID to business model
	id, err := internal.DestringifyPanelId(request.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid panel id")
	}

	// Ensure that the data values have been provided
	if request.GetData() == nil {
		return nil, status.Error(codes.InvalidArgument, "malformed request")
	}

	// Convert data to business model
	data := pb.PanelUpdateFromProto(request.GetData())

	// Attempt to update the panel
	panel, err := svr.service.UpdatePanel(ctx, id, data)
	if err != nil {
		return nil, err
	}

	return pb.PanelToProto(panel), nil
}

func (svr *panelServer) UpdatePanelByName(ctx context.Context, request *pb.UpdatePanelByNameRequest) (*pb.Panel, error) {
	// Ensure the required args are provided
	var name string = request.GetName()
	if name == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid panel name")
	}

	if request.GetData() == nil {
		return nil, status.Error(codes.InvalidArgument, "malformed request")
	}

	// Convert data to business model
	data := pb.PanelUpdateFromProto(request.GetData())

	// Attempt to update the panel
	panel, err := svr.service.UpdatePanelByName(ctx, name, data)
	if err != nil {
		return nil, err
	}

	return pb.PanelToProto(panel), nil
}

func (svr *panelServer) DeletePanel(ctx context.Context, request *pb.DeletePanelByIdRequest) (*emptypb.Empty, error) {
	// Ensure the required args are provided
	if request.GetId() == "" {
		return nil, status.Error(codes.InvalidArgument, "panel id not provided")
	}

	// Convert id to business model
	id, err := internal.DestringifyPanelId(request.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid panel id")
	}

	// Attempt to delete the panel
	err = svr.service.DeletePanel(ctx, id)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (svr *panelServer) DeletePanelByName(ctx context.Context, request *pb.DeletePanelByNameRequest) (*emptypb.Empty, error) {
	// Ensure the required args are provided
	var name string = request.GetName()
	if name == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid panel name")
	}

	// Attempt to delete the panel
	err := svr.service.DeletePanelByName(ctx, name)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}