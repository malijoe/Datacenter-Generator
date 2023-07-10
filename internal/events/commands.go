package events

type Command interface {
	GetAggregateID() string
}

type BaseCommand struct {
	AggregateId string `json:"aggregateId"`
}

func NewBaseCommand(aggregateId string) BaseCommand {
	return BaseCommand{AggregateId: aggregateId}
}

func (c *BaseCommand) GetAggregateID() string {
	return c.AggregateId
}
