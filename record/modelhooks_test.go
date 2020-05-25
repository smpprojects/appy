package record

import (
	"errors"
	"time"

	"github.com/appist/appy/support"
	"github.com/bxcodec/faker/v3"
)

type UserWithBeforeValidateError struct {
	Model     `masters:"primary" replicas:"primaryReplica" tableName:"callback_users" autoIncrement:"id" timezone:"local" faker:"-"`
	ID        int64         `db:"id" faker:"-"`
	Username  string        `db:"username" faker:"username,unique"`
	CreatedAt support.ZTime `db:"created_at" faker:"-"`
	DeletedAt *time.Time    `db:"deleted_at" faker:"-"`
	UpdatedAt *time.Time    `db:"updated_at" faker:"-"`
}

func (u *UserWithBeforeValidateError) BeforeValidate() error {
	return errors.New("before validate error")
}

type UserWithAfterValidateError struct {
	Model     `masters:"primary" replicas:"primaryReplica" tableName:"callback_users" autoIncrement:"id" timezone:"local" faker:"-"`
	ID        int64      `db:"id" faker:"-"`
	Username  string     `db:"username" faker:"username,unique"`
	CreatedAt *time.Time `db:"created_at" faker:"-"`
	DeletedAt *time.Time `db:"deleted_at" faker:"-"`
	UpdatedAt *time.Time `db:"updated_at" faker:"-"`
}

func (u *UserWithAfterValidateError) AfterValidate() error {
	return errors.New("after validate error")
}

type UserWithBeforeCreateError struct {
	Model     `masters:"primary" replicas:"primaryReplica" tableName:"callback_users" autoIncrement:"id" timezone:"local" faker:"-"`
	ID        int64      `db:"id" faker:"-"`
	Username  string     `db:"username" faker:"username,unique"`
	CreatedAt *time.Time `db:"created_at" faker:"-"`
	DeletedAt *time.Time `db:"deleted_at" faker:"-"`
	UpdatedAt *time.Time `db:"updated_at" faker:"-"`
}

func (u *UserWithBeforeCreateError) BeforeCreate() error {
	return errors.New("before create error")
}

type UserWithAfterCreateError struct {
	Model     `masters:"primary" replicas:"primaryReplica" tableName:"callback_users" autoIncrement:"id" timezone:"local" faker:"-"`
	ID        int64         `db:"id" faker:"-"`
	Username  string        `db:"username" faker:"username,unique"`
	CreatedAt support.NTime `db:"created_at" faker:"-"`
	DeletedAt *time.Time    `db:"deleted_at" faker:"-"`
	UpdatedAt *time.Time    `db:"updated_at" faker:"-"`
}

func (u *UserWithAfterCreateError) AfterCreate() error {
	return errors.New("after create error")
}

type UserWithBeforeUpdateError struct {
	Model     `masters:"primary" replicas:"primaryReplica" tableName:"callback_users" autoIncrement:"id" timezone:"local" faker:"-"`
	ID        int64      `db:"id" faker:"-"`
	Username  string     `db:"username" faker:"username,unique"`
	CreatedAt *time.Time `db:"created_at" faker:"-"`
	DeletedAt *time.Time `db:"deleted_at" faker:"-"`
	UpdatedAt *time.Time `db:"updated_at" faker:"-"`
}

func (u *UserWithBeforeUpdateError) BeforeUpdate() error {
	return errors.New("before update error")
}

type UserWithAfterUpdateError struct {
	Model     `masters:"primary" replicas:"primaryReplica" tableName:"callback_users" autoIncrement:"id" timezone:"local" faker:"-"`
	ID        int64      `db:"id" faker:"-"`
	Username  string     `db:"username" faker:"username,unique"`
	CreatedAt *time.Time `db:"created_at" faker:"-"`
	DeletedAt *time.Time `db:"deleted_at" faker:"-"`
	UpdatedAt time.Time  `db:"updated_at" faker:"-"`
}

func (u *UserWithAfterUpdateError) AfterUpdate() error {
	return errors.New("after update error")
}

type UserWithBeforeDeleteError struct {
	Model     `masters:"primary" replicas:"primaryReplica" tableName:"callback_users" autoIncrement:"id" timezone:"local" faker:"-"`
	ID        int64      `db:"id" faker:"-"`
	Username  string     `db:"username" faker:"username,unique"`
	CreatedAt *time.Time `db:"created_at" faker:"-"`
	DeletedAt *time.Time `db:"deleted_at" faker:"-"`
	UpdatedAt *time.Time `db:"updated_at" faker:"-"`
}

func (u *UserWithBeforeDeleteError) BeforeDelete() error {
	return errors.New("before delete error")
}

type UserWithAfterDeleteError struct {
	Model     `masters:"primary" replicas:"primaryReplica" tableName:"callback_users" autoIncrement:"id" timezone:"local" faker:"-"`
	ID        int64      `db:"id" faker:"-"`
	Username  string     `db:"username" faker:"username,unique"`
	CreatedAt *time.Time `db:"created_at" faker:"-"`
	DeletedAt *time.Time `db:"deleted_at" faker:"-"`
	UpdatedAt *time.Time `db:"updated_at" faker:"-"`
}

func (u *UserWithAfterDeleteError) AfterDelete() error {
	return errors.New("after delete error")
}

type UserWithAfterCreateCommitError struct {
	Model     `masters:"primary" replicas:"primaryReplica" tableName:"callback_users" autoIncrement:"id" timezone:"local" faker:"-"`
	ID        int64      `db:"id" faker:"-"`
	Message   string     `db:"-" faker:"-"`
	Username  string     `db:"username" faker:"username,unique"`
	CreatedAt *time.Time `db:"created_at" faker:"-"`
	DeletedAt *time.Time `db:"deleted_at" faker:"-"`
	UpdatedAt *time.Time `db:"updated_at" faker:"-"`
}

func (u *UserWithAfterCreateCommitError) AfterCreateCommit() error {
	return errors.New("after create commit error")
}

type UserWithAfterCreateCommit struct {
	Model     `masters:"primary" replicas:"primaryReplica" tableName:"callback_users" autoIncrement:"id" timezone:"local" faker:"-"`
	ID        int64      `db:"id" faker:"-"`
	Message   string     `db:"-" faker:"-"`
	Username  string     `db:"username" faker:"username,unique"`
	CreatedAt *time.Time `db:"created_at" faker:"-"`
	DeletedAt *time.Time `db:"deleted_at" faker:"-"`
	UpdatedAt *time.Time `db:"updated_at" faker:"-"`
}

func (u *UserWithAfterCreateCommit) AfterCreateCommit() error {
	u.Message = "after create commit"

	return nil
}

type UserWithAfterDeleteCommit struct {
	Model     `masters:"primary" replicas:"primaryReplica" tableName:"callback_users" autoIncrement:"id" timezone:"local" faker:"-"`
	ID        int64      `db:"id" faker:"-"`
	Message   string     `db:"-" faker:"-"`
	Username  string     `db:"username" faker:"username,unique"`
	CreatedAt *time.Time `db:"created_at" faker:"-"`
	DeletedAt *time.Time `db:"deleted_at" faker:"-"`
	UpdatedAt *time.Time `db:"updated_at" faker:"-"`
}

func (u *UserWithAfterDeleteCommit) AfterDeleteCommit() error {
	u.Message = "after delete commit"

	return nil
}

type UserWithAfterUpdateCommit struct {
	Model     `masters:"primary" replicas:"primaryReplica" tableName:"callback_users" autoIncrement:"id" timezone:"local" faker:"-"`
	ID        int64         `db:"id" faker:"-"`
	Message   string        `db:"-" faker:"-"`
	Username  string        `db:"username" faker:"username,unique"`
	CreatedAt *time.Time    `db:"created_at" faker:"-"`
	DeletedAt *time.Time    `db:"deleted_at" faker:"-"`
	UpdatedAt support.ZTime `db:"updated_at" faker:"-"`
}

func (u *UserWithAfterUpdateCommit) AfterUpdateCommit() error {
	u.Message = "after update commit"

	return nil
}

type UserWithAfterRollbackError struct {
	Model     `masters:"primary" replicas:"primaryReplica" tableName:"callback_users" autoIncrement:"id" timezone:"local" faker:"-"`
	ID        int64      `db:"id" faker:"-"`
	Message   string     `db:"-" faker:"-"`
	Username  string     `db:"username" faker:"username,unique"`
	CreatedAt *time.Time `db:"created_at" faker:"-"`
	DeletedAt *time.Time `db:"deleted_at" faker:"-"`
	UpdatedAt *time.Time `db:"updated_at" faker:"-"`
}

func (u *UserWithAfterRollbackError) AfterRollback() error {
	return errors.New("after rollback error")
}

type UserWithAfterRollback struct {
	Model     `masters:"primary" replicas:"primaryReplica" tableName:"callback_users" autoIncrement:"id" timezone:"local" faker:"-"`
	ID        int64      `db:"id" faker:"-"`
	Message   string     `db:"-" faker:"-"`
	Username  string     `db:"username" faker:"username,unique"`
	CreatedAt *time.Time `db:"created_at" faker:"-"`
	DeletedAt *time.Time `db:"deleted_at" faker:"-"`
	UpdatedAt *time.Time `db:"updated_at" faker:"-"`
}

func (u *UserWithAfterRollback) AfterRollback() error {
	u.Message = "after rollback"

	return nil
}

func (s *modelSuite) TestCallback() {
	for _, adapter := range support.SupportedDBAdapters {
		s.setupDB(adapter, "test_model_callback_with_"+adapter)

		{
			var user UserWithBeforeValidateError
			s.Nil(faker.FakeData(&user))

			count, err := s.model(&user).Create().Exec()
			s.Equal(int64(0), count)
			s.EqualError(err, "before validate error")
		}

		{
			var users []UserWithBeforeValidateError
			for i := 0; i < 10; i++ {
				user := UserWithBeforeValidateError{}
				s.Nil(faker.FakeData(&user))
				users = append(users, user)
			}

			count, err := s.model(&users).Create().Exec()
			s.Equal(int64(0), count)
			s.EqualError(err, "before validate error")
		}

		{
			var user UserWithAfterValidateError
			s.Nil(faker.FakeData(&user))

			count, err := s.model(&user).Create().Exec()
			s.Equal(int64(0), count)
			s.EqualError(err, "after validate error")
		}

		{
			var users []UserWithAfterValidateError
			for i := 0; i < 10; i++ {
				user := UserWithAfterValidateError{}
				s.Nil(faker.FakeData(&user))
				users = append(users, user)
			}

			count, err := s.model(&users).Create().Exec()
			s.Equal(int64(0), count)
			s.EqualError(err, "after validate error")
		}

		{
			var user UserWithBeforeCreateError
			s.Nil(faker.FakeData(&user))

			count, err := s.model(&user).Create().Exec()
			s.Equal(int64(0), count)
			s.EqualError(err, "before create error")
		}

		{
			var users []UserWithBeforeCreateError
			for i := 0; i < 10; i++ {
				user := UserWithBeforeCreateError{}
				s.Nil(faker.FakeData(&user))
				users = append(users, user)
			}

			count, err := s.model(&users).Create().Exec()
			s.Equal(int64(0), count)
			s.EqualError(err, "before create error")
		}

		{
			var user UserWithAfterCreateError
			s.Nil(faker.FakeData(&user))

			count, err := s.model(&user).Create().Exec()
			s.Equal(int64(1), count)
			s.EqualError(err, "after create error")
		}

		{
			var users []UserWithAfterCreateError
			for i := 0; i < 10; i++ {
				user := UserWithAfterCreateError{}
				s.Nil(faker.FakeData(&user))
				users = append(users, user)
			}

			count, err := s.model(&users).Create().Exec()
			s.Equal(int64(10), count)
			s.EqualError(err, "after create error")
		}

		{
			var user UserWithBeforeUpdateError
			s.Nil(faker.FakeData(&user))

			count, err := s.model(&user).Create().Exec()
			s.Equal(int64(1), count)
			s.Nil(err)

			user.Username = "foo"
			count, err = s.model(&user).Update().Exec()
			s.Equal(int64(0), count)
			s.EqualError(err, "before update error")
		}

		{
			var users []UserWithBeforeUpdateError
			for i := 0; i < 10; i++ {
				user := UserWithBeforeUpdateError{}
				s.Nil(faker.FakeData(&user))
				users = append(users, user)
			}

			count, err := s.model(&users).Create().Exec()
			s.Equal(int64(10), count)
			s.Nil(err)

			users[0].Username = "bar"
			count, err = s.model(&users).Update().Exec()
			s.Equal(int64(0), count)
			s.EqualError(err, "before update error")
		}

		{
			var user UserWithAfterUpdateError
			s.Nil(faker.FakeData(&user))
			user.UpdatedAt = time.Now()

			count, err := s.model(&user).Create().Exec()
			s.Equal(int64(1), count)
			s.Nil(err)

			user.Username = "foo"

			time.Sleep(1 * time.Second)
			count, err = s.model(&user).Update().Exec()
			s.Equal(int64(1), count)
			s.EqualError(err, "after update error")
		}

		{
			now := time.Now()
			var users []UserWithAfterUpdateError
			for i := 0; i < 10; i++ {
				user := UserWithAfterUpdateError{}
				s.Nil(faker.FakeData(&user))
				user.UpdatedAt = now
				users = append(users, user)
			}

			count, err := s.model(&users).Create().Exec()
			s.Equal(int64(10), count)
			s.Nil(err)

			users[0].Username = "bar"

			time.Sleep(1 * time.Second)
			count, err = s.model(&users).Update().Exec()
			s.Equal(int64(10), count)
			s.EqualError(err, "after update error")
		}

		{
			var user UserWithBeforeDeleteError
			s.Nil(faker.FakeData(&user))

			count, err := s.model(&user).Create().Exec()
			s.Equal(int64(1), count)
			s.Nil(err)

			count, err = s.model(&user).Delete().Exec()
			s.Equal(int64(0), count)
			s.EqualError(err, "before delete error")
		}

		{
			var users []UserWithBeforeDeleteError
			for i := 0; i < 10; i++ {
				user := UserWithBeforeDeleteError{}
				s.Nil(faker.FakeData(&user))
				users = append(users, user)
			}

			count, err := s.model(&users).Create().Exec()
			s.Equal(int64(10), count)

			count, err = s.model(&users).Delete().Exec()
			s.Equal(int64(0), count)
			s.EqualError(err, "before delete error")
		}

		{
			var user UserWithAfterDeleteError
			s.Nil(faker.FakeData(&user))

			count, err := s.model(&user).Create().Exec()
			s.Equal(int64(1), count)
			s.Nil(err)

			user.Username = "foo"
			count, err = s.model(&user).Delete().Exec()
			s.Equal(int64(1), count)
			s.EqualError(err, "after delete error")
		}

		{
			var users []UserWithAfterDeleteError
			for i := 0; i < 10; i++ {
				user := UserWithAfterDeleteError{}
				s.Nil(faker.FakeData(&user))
				users = append(users, user)
			}

			count, err := s.model(&users).Create().Exec()
			s.Equal(int64(10), count)
			s.Nil(err)

			users[0].Username = "bar"
			count, err = s.model(&users).Delete().Exec()
			s.Equal(int64(10), count)
			s.EqualError(err, "after delete error")
		}
	}
}

func (s *modelSuite) TestCallbackTx() {
	for _, adapter := range support.SupportedDBAdapters {
		s.setupDB(adapter, "test_model_callback_tx_with_"+adapter)

		{
			var user UserWithBeforeValidateError
			s.Nil(faker.FakeData(&user))

			model := s.model(&user)
			err := model.Begin()
			s.NotNil(model.Tx())
			s.Nil(err)

			count, err := model.Create().Exec()
			s.Equal(int64(0), count)
			s.EqualError(err, "before validate error")
			s.Nil(model.Tx())
		}

		{
			var users []UserWithBeforeValidateError
			for i := 0; i < 10; i++ {
				user := UserWithBeforeValidateError{}
				s.Nil(faker.FakeData(&user))
				users = append(users, user)
			}

			model := s.model(&users)
			err := model.Begin()
			s.NotNil(model.Tx())
			s.Nil(err)

			count, err := model.Create().Exec()
			s.Equal(int64(0), count)
			s.EqualError(err, "before validate error")
			s.Nil(model.Tx())
		}

		{
			var user UserWithAfterValidateError
			s.Nil(faker.FakeData(&user))

			model := s.model(&user)
			err := model.Begin()
			s.NotNil(model.Tx())
			s.Nil(err)

			count, err := model.Create().Exec()
			s.Equal(int64(0), count)
			s.EqualError(err, "after validate error")
			s.Nil(model.Tx())
		}

		{
			var users []UserWithAfterValidateError
			for i := 0; i < 10; i++ {
				user := UserWithAfterValidateError{}
				s.Nil(faker.FakeData(&user))
				users = append(users, user)
			}

			model := s.model(&users)
			err := model.Begin()
			s.NotNil(model.Tx())
			s.Nil(err)

			count, err := model.Create().Exec()
			s.Equal(int64(0), count)
			s.EqualError(err, "after validate error")
			s.Nil(model.Tx())
		}

		{
			var user UserWithBeforeCreateError
			s.Nil(faker.FakeData(&user))

			model := s.model(&user)
			err := model.Begin()
			s.NotNil(model.Tx())
			s.Nil(err)

			count, err := model.Create().Exec()
			s.Equal(int64(0), count)
			s.EqualError(err, "before create error")
			s.Nil(model.Tx())
		}

		{
			var users []UserWithBeforeCreateError
			for i := 0; i < 10; i++ {
				user := UserWithBeforeCreateError{}
				s.Nil(faker.FakeData(&user))
				users = append(users, user)
			}

			model := s.model(&users)
			err := model.Begin()
			s.NotNil(model.Tx())
			s.Nil(err)

			count, err := model.Create().Exec()
			s.Equal(int64(0), count)
			s.EqualError(err, "before create error")
			s.Nil(model.Tx())
		}

		{
			var user UserWithAfterCreateError
			s.Nil(faker.FakeData(&user))

			model := s.model(&user)
			err := model.Begin()
			s.NotNil(model.Tx())
			s.Nil(err)

			count, err := model.Create().Exec()
			s.Equal(int64(1), count)
			s.EqualError(err, "after create error")
			s.Nil(model.Tx())
		}

		{
			var users []UserWithAfterCreateError
			for i := 0; i < 10; i++ {
				user := UserWithAfterCreateError{}
				s.Nil(faker.FakeData(&user))
				users = append(users, user)
			}

			model := s.model(&users)
			err := model.Begin()
			s.NotNil(model.Tx())
			s.Nil(err)

			count, err := model.Create().Exec()
			s.Equal(int64(10), count)
			s.EqualError(err, "after create error")
			s.Nil(model.Tx())
		}

		{
			var user UserWithBeforeUpdateError
			s.Nil(faker.FakeData(&user))

			count, err := s.model(&user).Create().Exec()
			s.Equal(int64(1), count)
			s.Nil(err)

			model := s.model(&user)
			err = model.Begin()
			s.NotNil(model.Tx())
			s.Nil(err)

			user.Username = "foo"
			count, err = model.Update().Exec()
			s.Equal(int64(0), count)
			s.EqualError(err, "before update error")
			s.Nil(model.Tx())
		}

		{
			var users []UserWithBeforeUpdateError
			for i := 0; i < 10; i++ {
				user := UserWithBeforeUpdateError{}
				s.Nil(faker.FakeData(&user))
				users = append(users, user)
			}

			count, err := s.model(&users).Create().Exec()
			s.Equal(int64(10), count)
			s.Nil(err)

			model := s.model(&users)
			err = model.Begin()
			s.NotNil(model.Tx())
			s.Nil(err)

			users[0].Username = "bar"
			count, err = model.Update().Exec()
			s.Equal(int64(0), count)
			s.EqualError(err, "before update error")
			s.Nil(model.Tx())
		}

		{
			var user UserWithAfterUpdateError
			s.Nil(faker.FakeData(&user))
			user.UpdatedAt = time.Now()

			model := s.model(&user)
			err := model.Begin()
			s.NotNil(model.Tx())
			s.Nil(err)

			count, err := model.Create().Exec()
			s.Equal(int64(1), count)

			user.Username = "foo"

			time.Sleep(1 * time.Second)
			count, err = model.Update().Exec()
			s.Equal(int64(1), count)
			s.EqualError(err, "after update error")
			s.Nil(model.Tx())
		}

		{
			now := time.Now()
			var users []UserWithAfterUpdateError
			for i := 0; i < 10; i++ {
				user := UserWithAfterUpdateError{}
				s.Nil(faker.FakeData(&user))
				user.UpdatedAt = now
				users = append(users, user)
			}

			count, err := s.model(&users).Create().Exec()
			s.Equal(int64(10), count)
			s.Nil(err)

			model := s.model(&users)
			err = model.Begin()
			s.NotNil(model.Tx())
			s.Nil(err)

			users[0].Username = "bar"

			time.Sleep(1 * time.Second)
			count, err = model.Update().Exec()
			s.Equal(int64(10), count)
			s.EqualError(err, "after update error")
			s.Nil(model.Tx())
		}

		{
			var user UserWithBeforeDeleteError
			s.Nil(faker.FakeData(&user))

			count, err := s.model(&user).Create().Exec()
			s.Equal(int64(1), count)
			s.Nil(err)

			model := s.model(&user)
			err = model.Begin()
			s.NotNil(model.Tx())
			s.Nil(err)

			user.Username = "foo"
			count, err = model.Delete().Exec()
			s.Equal(int64(0), count)
			s.EqualError(err, "before delete error")
			s.Nil(model.Tx())
		}

		{
			var users []UserWithBeforeDeleteError
			for i := 0; i < 10; i++ {
				user := UserWithBeforeDeleteError{}
				s.Nil(faker.FakeData(&user))
				users = append(users, user)
			}

			count, err := s.model(&users).Create().Exec()
			s.Equal(int64(10), count)
			s.Nil(err)

			model := s.model(&users)
			err = model.Begin()
			s.NotNil(model.Tx())
			s.Nil(err)

			users[0].Username = "bar"
			count, err = model.Delete().Exec()
			s.Equal(int64(0), count)
			s.EqualError(err, "before delete error")
			s.Nil(model.Tx())
		}

		{
			var user UserWithAfterDeleteError
			s.Nil(faker.FakeData(&user))
			count, err := s.model(&user).Create().Exec()
			s.Equal(int64(1), count)
			s.Nil(err)

			model := s.model(&user)
			err = model.Begin()
			s.NotNil(model.Tx())
			s.Nil(err)

			count, err = model.Delete().Exec()
			s.Equal(int64(1), count)
			s.EqualError(err, "after delete error")
			s.Nil(model.Tx())
		}

		{
			var users []UserWithAfterDeleteError
			for i := 0; i < 10; i++ {
				user := UserWithAfterDeleteError{}
				s.Nil(faker.FakeData(&user))
				users = append(users, user)
			}

			count, err := s.model(&users).Create().Exec()
			s.Equal(int64(10), count)
			s.Nil(err)

			model := s.model(&users)
			err = model.Begin()
			s.NotNil(model.Tx())
			s.Nil(err)

			count, err = model.Delete().Exec()
			s.Equal(int64(10), count)
			s.EqualError(err, "after delete error")
			s.Nil(model.Tx())
		}

		{
			var user UserWithAfterCreateCommitError
			s.Nil(faker.FakeData(&user))

			model := s.model(&user)
			err := model.Begin()
			s.NotNil(model.Tx())
			s.Nil(err)

			count, err := model.Create().Exec()
			s.Equal(int64(1), count)
			s.Nil(err)

			err = model.Commit()
			s.Nil(model.Tx())
			s.EqualError(err, "after create commit error")
		}

		{
			var users []UserWithAfterCreateCommitError
			for i := 0; i < 10; i++ {
				user := UserWithAfterCreateCommitError{}
				s.Nil(faker.FakeData(&user))
				users = append(users, user)
			}

			model := s.model(&users)
			err := model.Begin()
			s.NotNil(model.Tx())
			s.Nil(err)

			count, err := model.Create().Exec()
			s.Equal(int64(10), count)
			s.Nil(err)

			err = model.Commit()
			s.Nil(model.Tx())
			s.EqualError(err, "after create commit error")
		}

		{
			var user UserWithAfterCreateCommit
			s.Nil(faker.FakeData(&user))

			model := s.model(&user)
			err := model.Begin()
			s.NotNil(model.Tx())
			s.Nil(err)

			count, err := model.Create().Exec()
			s.Equal(int64(1), count)
			s.Nil(err)

			err = model.Commit()
			s.Nil(model.Tx())
			s.Nil(err)
			s.Equal("after create commit", user.Message)
		}

		{
			var users []UserWithAfterCreateCommit
			for i := 0; i < 10; i++ {
				user := UserWithAfterCreateCommit{}
				s.Nil(faker.FakeData(&user))
				users = append(users, user)
			}

			model := s.model(&users)
			err := model.Begin()
			s.NotNil(model.Tx())
			s.Nil(err)

			count, err := model.Create().Exec()
			s.Equal(int64(10), count)
			s.Nil(err)

			err = model.Commit()
			s.Nil(model.Tx())
			s.Nil(err)

			for i := 0; i < 10; i++ {
				s.Equal("after create commit", users[i].Message)
			}
		}

		{
			var user UserWithAfterUpdateCommit
			s.Nil(faker.FakeData(&user))
			count, err := s.model(&user).Create().Exec()
			s.Equal(int64(1), count)
			s.Nil(err)

			model := s.model(&user)
			err = model.Begin()
			s.NotNil(model.Tx())
			s.Nil(err)

			user.Username = "foo"
			count, err = model.Update().Exec()
			s.Equal(int64(1), count)
			s.Nil(err)

			err = model.Commit()
			s.Nil(model.Tx())
			s.Nil(err)
			s.Equal("after update commit", user.Message)
		}

		{
			var users []UserWithAfterUpdateCommit
			for i := 0; i < 10; i++ {
				user := UserWithAfterUpdateCommit{}
				s.Nil(faker.FakeData(&user))
				users = append(users, user)
			}

			count, err := s.model(&users).Create().Exec()
			s.Equal(int64(10), count)
			s.Nil(err)

			model := s.model(&users)
			err = model.Begin()
			s.NotNil(model.Tx())
			s.Nil(err)

			users[0].Username = "bar"
			count, err = model.Update().Exec()
			s.Equal(int64(10), count)
			s.Nil(err)

			err = model.Commit()
			s.Nil(model.Tx())
			s.Nil(err)

			for i := 0; i < 10; i++ {
				s.Equal("after update commit", users[i].Message)
			}
		}

		{
			var user UserWithAfterDeleteCommit
			s.Nil(faker.FakeData(&user))
			count, err := s.model(&user).Create().Exec()
			s.Equal(int64(1), count)
			s.Nil(err)

			model := s.model(&user)
			err = model.Begin()
			s.NotNil(model.Tx())
			s.Nil(err)

			count, err = model.Delete().Exec()
			s.Equal(int64(1), count)
			s.Nil(err)

			err = model.Commit()
			s.Nil(model.Tx())
			s.Nil(err)
			s.Equal("after delete commit", user.Message)
		}

		{
			var users []UserWithAfterDeleteCommit
			for i := 0; i < 10; i++ {
				user := UserWithAfterDeleteCommit{}
				s.Nil(faker.FakeData(&user))
				users = append(users, user)
			}

			count, err := s.model(&users).Create().Exec()
			s.Equal(int64(10), count)
			s.Nil(err)

			model := s.model(&users)
			err = model.Begin()
			s.NotNil(model.Tx())
			s.Nil(err)

			count, err = model.Delete().Exec()
			s.Equal(int64(10), count)
			s.Nil(err)

			err = model.Commit()
			s.Nil(model.Tx())
			s.Nil(err)

			for i := 0; i < 10; i++ {
				s.Equal("after delete commit", users[i].Message)
			}
		}

		{
			var user UserWithAfterRollbackError
			s.Nil(faker.FakeData(&user))
			count, err := s.model(&user).Create().Exec()
			s.Equal(int64(1), count)
			s.Nil(err)

			model := s.model(&user)
			err = model.Begin()
			s.NotNil(model.Tx())
			s.Nil(err)

			count, err = model.Delete().Exec()
			s.Equal(int64(1), count)
			s.Nil(err)

			err = model.Rollback()
			s.Nil(model.Tx())
			s.EqualError(err, "after rollback error")
		}

		{
			var users []UserWithAfterRollbackError
			for i := 0; i < 10; i++ {
				user := UserWithAfterRollbackError{}
				s.Nil(faker.FakeData(&user))
				users = append(users, user)
			}

			count, err := s.model(&users).Create().Exec()
			s.Equal(int64(10), count)
			s.Nil(err)

			model := s.model(&users)
			err = model.Begin()
			s.NotNil(model.Tx())
			s.Nil(err)

			count, err = model.Delete().Exec()
			s.Equal(int64(10), count)
			s.Nil(err)

			err = model.Rollback()
			s.Nil(model.Tx())
			s.EqualError(err, "after rollback error")
		}

		{
			var user UserWithAfterRollback
			s.Nil(faker.FakeData(&user))
			count, err := s.model(&user).Create().Exec()
			s.Equal(int64(1), count)
			s.Nil(err)

			model := s.model(&user)
			err = model.Begin()
			s.NotNil(model.Tx())
			s.Nil(err)

			count, err = model.Delete().Exec()
			s.Equal(int64(1), count)
			s.Nil(err)

			err = model.Rollback()
			s.Nil(model.Tx())
			s.Nil(err)
			s.Equal("after rollback", user.Message)
		}

		{
			var users []UserWithAfterRollback
			for i := 0; i < 10; i++ {
				user := UserWithAfterRollback{}
				s.Nil(faker.FakeData(&user))
				users = append(users, user)
			}

			count, err := s.model(&users).Create().Exec()
			s.Equal(int64(10), count)
			s.Nil(err)

			model := s.model(&users)
			err = model.Begin()
			s.NotNil(model.Tx())
			s.Nil(err)

			count, err = model.Delete().Exec()
			s.Equal(int64(10), count)
			s.Nil(err)

			err = model.Rollback()
			s.Nil(model.Tx())
			s.Nil(err)

			for i := 0; i < 10; i++ {
				s.Equal("after rollback", users[i].Message)
			}
		}
	}
}