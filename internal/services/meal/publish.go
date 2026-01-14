package meal

import "context"

func (s *service) Publish(ctx context.Context) error {
	return s.producer.Publish()
}
