package herald

type Broker interface {
	Publish(topic, message string) error
	Request(topic, message string) (interface{}, error)
}
