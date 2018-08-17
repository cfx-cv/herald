package herald

type Herald struct {
	Broker
}

func NewHerald(broker Broker) *Herald {
	return &Herald{Broker: broker}
}
