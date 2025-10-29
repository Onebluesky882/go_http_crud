package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/onebluesky882/go-http-crud/cusErr"
	m "github.com/onebluesky882/go-http-crud/models"
	"github.com/uptrace/bun"
)

type User struct {
	db bun.IDB
}

func Service(db bun.IDB) *User {
	return &User{
		db: db,
	}
}

/*

Scan()	struct	SELECT, INSERT/UPDATE/DELETE + RETURNING
Exec()	sql.Result	UPDATE, DELETE ที่ไม่ต้องเอาค่ากลับ

ถ้าคุณใช้ Bun + Fiber CRUD app:

Create → Scan() (คืนค่าที่ insert)

FindById, FindAll → Scan()

Update → ขึ้นอยู่กับว่าจะคืนค่าไหม ถ้าไม่ → Exec()

Delete → ถ้าไม่เอาค่า → Exec()

*/

// ✅ จำเป็นต้องใช้ pointer เพื่อให้ ORM อัปเดตค่า ID, timestamp
// create user
func (s User) Create(ctx context.Context, user *m.User) (*m.User, error) {
	err := s.db.NewInsert().Model(user).Returning("*").Scan(ctx, user)
	if err != nil {
		// wrap original error
		return nil, cusErr.Handle(err, "failed to insert user")
	}

	return user, nil
}

// find by id

func (s User) FindById(ctx context.Context, id uuid.UUID) (*m.User, error) {
	user := new(m.User)
	err := s.db.NewSelect().Model(user).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return user, err
}

// find all

func (s User) FindAll(ctx context.Context) ([]*m.User, error) {
	var users []*m.User
	err := s.db.NewSelect().Model(&users).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return users, err
}

// delete by id
func (s User) DeleteUserById(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.NewDelete().Model(&m.User{}).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return cusErr.Handle(err, "failed to delete user")
	}
	return nil
}

// update by id
func (s User) UpdateUser(ctx context.Context, id uuid.UUID, user *m.User) (err error) {
	r, err := s.db.NewUpdate().Model(user).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return cusErr.Handle(err, "failed to update user")
	}

	rowsAffected, err := r.RowsAffected()
	if err != nil {
		return cusErr.Handle(err, "failed to get rows affected")
	}

	if rowsAffected == 0 {
		return cusErr.Handle(nil, "no user found to update")
	}
	return nil
}
