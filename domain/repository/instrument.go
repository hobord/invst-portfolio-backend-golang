package repository

import (
	"context"

	"github.com/hobord/invst-portfolio-backend-golang/domain/entity"
)

// InstrumentRepository interface definition
// make Mock with: mockery -name=InstrumentRepository
type InstrumentRepository interface {
	// Get return entity by id
	GetByID(ctx context.Context, id int) (*entity.Instrument, error)

	// GetAll return all Instrument entity
	GetAll(ctx context.Context) ([]*entity.Instrument, error)

	// Search
	Search(ctx context.Context, keyword string) ([]*entity.Instrument, error)

	// Save is save to persistent the Instrument
	Save(ctx context.Context, entity *entity.Instrument) error

	// Delete Instrument from persistent store
	Delete(ctx context.Context, id int) error
}
