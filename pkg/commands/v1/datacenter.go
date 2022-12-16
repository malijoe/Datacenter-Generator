package v1

import (
	"context"
	"errors"

	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/malijoe/DatacenterGenerator/pkg/aggregates/datacenterAggregate"
	"github.com/malijoe/DatacenterGenerator/pkg/aggregates/deviceAggregate"
	"github.com/malijoe/DatacenterGenerator/pkg/aggregates/deviceTemplateAggregate"
	"github.com/malijoe/DatacenterGenerator/pkg/aggregates/podAggregate"
	"github.com/malijoe/DatacenterGenerator/pkg/aggregates/rackAggregate"
	"github.com/malijoe/DatacenterGenerator/pkg/internal/events"
	"github.com/malijoe/DatacenterGenerator/pkg/internal/logger"
)

type InitDatacenterCommand struct {
	events.BaseCommand
	Site      string
	Building  string
	Room      string
	Providers map[string]string
}

func NewInitDatacenterCommand(aggregateId string, site string, building string, room string, providers map[string]string) *InitDatacenterCommand {
	return &InitDatacenterCommand{BaseCommand: events.NewBaseCommand(aggregateId), Site: site, Building: building, Room: room, Providers: providers}
}

type InitDatacenterCmdHandler interface {
	Handle(ctx context.Context, cmd *InitDatacenterCommand) error
}

type initDatacenterCmdHandler struct {
	store events.AggregateStore
	log   logger.Logger
}

func NewInitDatacenterHandler(store events.AggregateStore, log logger.Logger) *initDatacenterCmdHandler {
	return &initDatacenterCmdHandler{store: store, log: log}
}

func (h *initDatacenterCmdHandler) Handle(ctx context.Context, cmd *InitDatacenterCommand) error {
	dc := datacenterAggregate.NewDatacenterAggregateWithId(cmd.GetAggregateId())
	err := h.store.Exists(ctx, dc.GetId())
	if err != nil && !errors.Is(err, esdb.ErrStreamNotFound) {
		return err
	}

	if err = dc.CreateDatacenter(ctx, cmd.Site, cmd.Building, cmd.Room, cmd.Providers); err != nil {
		return err
	}

	return h.store.Save(ctx, dc)
}

type CreateRackCommand struct {
	events.BaseCommand
	Name         string
	Size         int
	DatacenterId string
}

func NewCreateRackCommand(aggregateId string, name string, size int, datacenterId string) *CreateRackCommand {
	return &CreateRackCommand{BaseCommand: events.NewBaseCommand(aggregateId), Name: name, Size: size, DatacenterId: datacenterId}
}

type CreateRackCmdHandler interface {
	Handle(ctx context.Context, cmd *CreateRackCommand) error
}

type createRackCmdHandler struct {
	store events.AggregateStore
	log   logger.Logger
}

func NewCreateRackCmdHandler(store events.AggregateStore, log logger.Logger) *createRackCmdHandler {
	return &createRackCmdHandler{store: store, log: log}
}

func (h *createRackCmdHandler) Handle(ctx context.Context, cmd *CreateRackCommand) error {
	rack := rackAggregate.NewRackAggregateWithId(cmd.GetAggregateId())

	err := h.store.Exists(ctx, rack.GetId())
	if err != nil && !errors.Is(err, esdb.ErrStreamNotFound) {
		return err
	}

	if err = rack.CreateRack(ctx, cmd.Name, cmd.Size, cmd.DatacenterId); err != nil {
		return err
	}

	return h.store.Save(ctx, rack)
}

type DatacenterAddRackCommand struct {
	events.BaseCommand
	RackId string
}

func NewDatacenterAddRackCommand(aggregateId string, rackId string) *DatacenterAddRackCommand {
	return &DatacenterAddRackCommand{BaseCommand: events.NewBaseCommand(aggregateId), RackId: rackId}
}

type DatacenterAddRackCmdHandler interface {
	Handle(ctx context.Context, cmd *DatacenterAddRackCommand) error
}

type datacenterAddRackCmdHandler struct {
	store events.AggregateStore
	log   logger.Logger
}

func NewDatacenterAddRackCmdHandler(store events.AggregateStore, log logger.Logger) *datacenterAddRackCmdHandler {
	return &datacenterAddRackCmdHandler{store: store, log: log}
}

func (h *datacenterAddRackCmdHandler) Handle(ctx context.Context, cmd *DatacenterAddRackCommand) error {
	dc, err := datacenterAggregate.LoadDatacenterAggregate(ctx, h.store, cmd.GetAggregateId())
	if err != nil {
		return err
	}

	if err = dc.AddRack(ctx, cmd.RackId); err != nil {
		return err
	}

	return h.store.Save(ctx, dc)
}

type CreatePodCommand struct {
	events.BaseCommand
	Function     string
	DatacenterId string
}

func NewCreatePodCommand(aggregateId string, function string, datacenterId string) *CreatePodCommand {
	return &CreatePodCommand{BaseCommand: events.NewBaseCommand(aggregateId), Function: function, DatacenterId: datacenterId}
}

type CreatePodCmdHandler interface {
	Handle(ctx context.Context, cmd *CreatePodCommand) error
}

type createPodCmdHandler struct {
	store events.AggregateStore
	log   logger.Logger
}

func NewCreatePodCmdHandler(store events.AggregateStore, log logger.Logger) *createPodCmdHandler {
	return &createPodCmdHandler{store: store, log: log}
}

func (h *createPodCmdHandler) Handle(ctx context.Context, cmd *CreatePodCommand) error {
	pod := podAggregate.NewPodAggregateWithId(cmd.GetAggregateId())

	err := h.store.Exists(ctx, pod.GetId())
	if err != nil && !errors.Is(err, esdb.ErrStreamNotFound) {
		return err
	}

	dc, err := datacenterAggregate.LoadDatacenterAggregate(ctx, h.store, cmd.DatacenterId)
	if err != nil {
		return err
	}

	if err = pod.CreatePod(ctx, cmd.Function, dc.Datacenter); err != nil {
		return err
	}

	return h.store.Save(ctx, pod)
}

type DatacenterAddPodCommand struct {
	events.BaseCommand
	PodId string
}

func NewDatacenterAddPodCommand(aggregateId string, podId string) *DatacenterAddPodCommand {
	return &DatacenterAddPodCommand{BaseCommand: events.NewBaseCommand(aggregateId), PodId: podId}
}

type DatacenterAddPodCmdHandler interface {
	Handle(ctx context.Context, cmd *DatacenterAddPodCommand) error
}

type datacenterAddPodCmdHandler struct {
	store events.AggregateStore
	log   logger.Logger
}

func NewDatacenterAddPodCmdHandler(store events.AggregateStore, log logger.Logger) *datacenterAddPodCmdHandler {
	return &datacenterAddPodCmdHandler{store: store, log: log}
}

func (h *datacenterAddPodCmdHandler) Handle(ctx context.Context, cmd *DatacenterAddPodCommand) error {
	dc, err := datacenterAggregate.LoadDatacenterAggregate(ctx, h.store, cmd.GetAggregateId())
	if err != nil {
		return err
	}

	if err = dc.AddPod(ctx, cmd.PodId); err != nil {
		return err
	}

	return h.store.Save(ctx, dc)
}

type CreateDeviceCommand struct {
	events.BaseCommand
	TemplateId  string
	Elevation   int
	RackId      string
	Cluster     int
	Designation string
	PodId       string
}

func NewCreateDeviceCommand(aggregateId string, templateId string, elevation int, rackId string, cluster int, designation string, podId string) *CreateDeviceCommand {
	return &CreateDeviceCommand{BaseCommand: events.NewBaseCommand(aggregateId), TemplateId: templateId, Elevation: elevation, RackId: rackId, Cluster: cluster, Designation: designation, PodId: podId}
}

type CreateDeviceCmdHandler interface {
	Handle(ctx context.Context, cmd *CreateDeviceCommand) error
}

type createDeviceCmdHandler struct {
	store events.AggregateStore
	log   logger.Logger
}

func (h *createDeviceCmdHandler) Handle(ctx context.Context, cmd *CreateDeviceCommand) error {
	device := deviceAggregate.NewDeviceAggregateWithId(cmd.GetAggregateId())

	err := h.store.Exists(ctx, device.GetId())
	if err != nil && !errors.Is(err, esdb.ErrStreamNotFound) {
		return err
	}

	return nil
}

type CreateDeviceTemplateCommand struct {
	events.BaseCommand
	ModelId          string
	Variant          string
	Categories       []string
	HostnameTemplate string
	Alias            string
	Function         string
}

func NewCreateDeviceTemplateCommand(aggregateId string, modelId, variant string, categories []string, hostnameTemplate, alias, function string) *CreateDeviceTemplateCommand {
	return &CreateDeviceTemplateCommand{BaseCommand: events.NewBaseCommand(aggregateId), ModelId: modelId, Variant: variant, Categories: categories, Alias: alias, Function: function}
}

type CreateDeviceTemplateCmdHandler interface {
	Handle(ctx context.Context, cmd *CreateDeviceTemplateCommand) error
}

type createDeviceTemplateCmdHandler struct {
	store events.AggregateStore
	log   logger.Logger
}

func NewCreateDeviceTemplateCmdHandler(store events.AggregateStore, log logger.Logger) *createDeviceTemplateCmdHandler {
	return &createDeviceTemplateCmdHandler{store: store, log: log}
}

func (h *createDeviceTemplateCmdHandler) Handle(ctx context.Context, cmd *CreateDeviceTemplateCommand) error {
	deviceTemplate := deviceTemplateAggregate.NewDeviceTemplateAggregateWithId(cmd.GetAggregateId())

	err := h.store.Exists(ctx, deviceTemplate.GetId())
	if err != nil && !errors.Is(err, esdb.ErrStreamNotFound) {
		return err
	}

	return nil
}
