package booking

type Service struct {
	store BookingStore
}

func NewService(store *RedisStore) *Service {
	return &Service{store}
}

func (s *Service) Book(b Booking) error {
	return s.store.Book(b)
}

func (s *Service) ListBooking(movieID string) ([]Booking, error) {
	return s.store.ListBooking(movieID), nil
}
