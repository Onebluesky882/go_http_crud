package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"Table:users"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()"`
	Name          string    `bun:"name,nullzero,notnull"`
	Email         string    `bun:"email,nullzero,notnull,unique"`
	CreatedAt     time.Time `bun:"created_at,nullzero,notnull,default:current_time"`
	UpdatedAt     time.Time `bun:"updated_at,default:current_time"`
	DeletedAt     time.Time `bun:"deleted_at,nullzero,soft_delete"`
}
