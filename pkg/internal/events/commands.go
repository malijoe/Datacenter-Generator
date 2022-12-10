package events

type Command interface {
	GetAggregateId() string
}

type BaseCommand struct {
	AggregateId string `json:"aggregateId"`
}

func NewBaseCommand(aggregateId string) BaseCommand {
	return BaseCommand{AggregateId: aggregateId}
}

func (c *BaseCommand) GetAggregateId() string {
	return c.AggregateId
}
