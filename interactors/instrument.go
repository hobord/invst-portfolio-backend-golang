package interactor

import (
	"context"

	"github.com/hobord/invst-portfolio-backend-golang/domain/entity"
	"github.com/hobord/invst-portfolio-backend-golang/domain/repository"
)

type InstrumentInteractorInterface interface {
	GetByID(ctx context.Context, id int) (*entity.Instrument, error)
	GetAll(ctx context.Context) ([]*entity.Instrument, error)
	Search(ctx context.Context, keyword string) ([]*entity.Instrument, error)
	Save(ctx context.Context, entity *entity.Instrument) error
	Delete(ctx context.Context, id int) error
}

// InstrumentInteractor provides an implementation of InstrumentInteractorInterface
type InstrumentInteractor struct {
	InstrumentRepository repository.InstrumentRepository
}

// CreateInstrumentInteractor is create a new example "service" / "interactor"
func CreateInstrumentInteractor(repository repository.InstrumentRepository) *InstrumentInteractor {
	return &InstrumentInteractor{
		InstrumentRepository: repository,
	}
}

// GetByID return entity by id
func (i *InstrumentInteractor) GetByID(ctx context.Context, id int) (*entity.Instrument, error) {
	return i.InstrumentRepository.GetByID(ctx, id)
}

// GetAll return all entity
func (i *InstrumentInteractor) GetAll(ctx context.Context) ([]*entity.Instrument, error) {
	return i.InstrumentRepository.GetAll(ctx)
}

// Search return entities by keyword
func (i *InstrumentInteractor) Search(ctx context.Context, keyword string) ([]*entity.Instrument, error) {
	return i.InstrumentRepository.Search(ctx, keyword)
}

// Save is save to persistent the entity
func (i *InstrumentInteractor) Save(ctx context.Context, entity *entity.Instrument) error {
	return i.InstrumentRepository.Save(ctx, entity)
}

// Delete entity from persistent store
func (i *InstrumentInteractor) Delete(ctx context.Context, id int) error {
	return i.InstrumentRepository.Delete(ctx, id)
}
