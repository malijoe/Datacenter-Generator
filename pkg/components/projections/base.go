package projections

import "time"

// BaseProjection defines metadata for projections.
type BaseProjection struct {
	// when the projection was created
	CreatedAt time.Time `json:"createdAt" bson:"createdAt,omitempty"`
	// when the projection was updated
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt,omitempty"`
	// when the projection was deleted
	DeletedAt time.Time `json:"deletedAt" bson:"deltedAt,omitempty"`
}

type projectionOptions func(*BaseProjection)

func newBaseProjection(opts ...projectionOptions) BaseProjection {
	var projection BaseProjection
	for _, op := range opts {
		op(&projection)
	}
	return projection
}

func baseProjectionCreated(p *BaseProjection) {
	p.CreatedAt = time.Now()
}

func baseProjectionUpdated(p *BaseProjection) {
	p.UpdatedAt = time.Now()
}

func baseProjectionDeleted(p *BaseProjection) {
	p.DeletedAt = time.Now()
}

func NewCreatedProjection() BaseProjection {
	return newBaseProjection(baseProjectionCreated)
}

func NewUpdatedProjection() BaseProjection {
	return newBaseProjection(baseProjectionUpdated)
}

func NewDeletedProjection() BaseProjection {
	return newBaseProjection(baseProjectionDeleted)
}
