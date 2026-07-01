package monitoring

import "context"

type Service interface {
	Count() (int, error)

	GetAll(
		limit int,
		offset int,
	) ([]Monitoring, error)

	Import(
		ctx context.Context,
		data []Monitoring,
	) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Count() (int, error) {
	return s.repo.Count()
}

func (s *service) GetAll(
	limit int,
	offset int,
) ([]Monitoring, error) {

	return s.repo.FindAll(
		limit,
		offset,
	)
}

func (s *service) Import(
	ctx context.Context,
	data []Monitoring,
) error {

	if len(data) == 0 {
		return nil
	}

	tgl := data[0].Tgl

	exists, err := s.repo.ExistByDate(tgl)
	if err != nil {
		return err
	}

	if exists {
		err = s.repo.DeleteByDate(ctx, tgl)
		if err != nil {
			return err
		}
	}

	return s.repo.BulkInsert(ctx, data)
}
