package submission

type (
	Broker interface {
		Subscribe(event string, callback func(msg []byte) error)
	}

	EventService interface {
	}

	Event struct {
	}

	EventHandler struct {
		event   string
		service EventService
	}
)

func NewEventHandler(event string, service EventService) *EventHandler {
	return &EventHandler{event: event, service: service}
}

func (e *EventHandler) Subscribe(b Broker) {
	b.Subscribe(e.event, e.Read)
}

func (e *EventHandler) Read(msg []byte) error {
	return nil
}
