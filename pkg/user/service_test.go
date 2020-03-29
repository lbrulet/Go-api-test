package user

import (
	"reflect"
	"testing"
	"time"

	"github.com/lbrulet/Go-api-test/pkg/models"
)

func TestNewService(t *testing.T) {
	type args struct {
		r Repository
	}
	r := NewInMemRepository()
	tests := []struct {
		name string
		args args
		want *Service
	}{
		{
			name: "test_new_service",
			args: args{r: r},
			want: &Service{r: r},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_DeleteUserByID(t *testing.T) {
	type fields struct {
		r Repository
	}
	type args struct {
		ID int
	}
	r := NewInMemRepository()
	r.m[1] = &models.User{
		ID:        1,
		FirstName: "test",
		LastName:  "test",
		Password:  "test",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "test_delete_user_by_id",
			fields:  fields{r: r},
			args:    args{ID: 1},
			wantErr: false,
		},
		{
			name:    "test_delete_user_by_id_with_error",
			fields:  fields{r: r},
			args:    args{ID: 2},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				r: tt.fields.r,
			}
			if err := s.DeleteUserByID(tt.args.ID); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUserByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_GetAllUsers(t *testing.T) {
	type fields struct {
		r Repository
	}
	r := NewInMemRepository()
	users := []*models.User{{
		ID:        1,
		FirstName: "test",
		LastName:  "test",
		Password:  "test",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}}
	r.m[users[0].ID] = users[0]
	tests := []struct {
		name    string
		fields  fields
		want    []*models.User
		wantErr bool
	}{
		{
			name:    "test_get_all_users",
			fields:  fields{r: r},
			want:    users,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				r: tt.fields.r,
			}
			got, err := s.GetAllUsers()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllUsers() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetUserByID(t *testing.T) {
	type fields struct {
		r Repository
	}
	type args struct {
		ID int
	}
	r := NewInMemRepository()
	user := &models.User{
		ID:        1,
		FirstName: "test",
		LastName:  "test",
		Password:  "test",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}
	r.m[user.ID] = user
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.User
		wantErr bool
	}{
		{
			name:    "test_get_user_by_id",
			fields:  fields{r},
			args:    args{ID: 1},
			want:    user,
			wantErr: false,
		},
		{
			name:    "test_get_user_by_id_with_error",
			fields:  fields{r},
			args:    args{ID: 0},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				r: tt.fields.r,
			}
			got, err := s.GetUserByID(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_InsertUser(t *testing.T) {
	type fields struct {
		r Repository
	}
	type args struct {
		user *models.User
	}
	r := NewInMemRepository()
	user := &models.User{
		FirstName: "test",
		LastName:  "test",
		Password:  "test",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "test_insert_user",
			fields:  fields{r: r},
			args:    args{user: user},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				r: tt.fields.r,
			}
			if err := s.InsertUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("InsertUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_Migrate(t *testing.T) {
	type fields struct {
		r Repository
	}
	r := NewInMemRepository()
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name:   "test_migrate",
			fields: fields{r: r},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				r: tt.fields.r,
			}
			s.Migrate()
		})
	}
}

func TestService_UpdateUserByID(t *testing.T) {
	type fields struct {
		r Repository
	}
	type args struct {
		user *models.User
	}
	r := NewInMemRepository()
	user := &models.User{
		ID:        1,
		FirstName: "test",
		LastName:  "test",
		Password:  "test",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}
	r.m[user.ID] = user
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "test_update_user_by_id",
			fields: fields{r: r},
			args: args{user: &models.User{
				ID:        1,
				FirstName: "test",
				LastName:  "test",
				Password:  "test",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: nil,
			}},
			wantErr: false,
		},
		{
			name:   "test_update_user_by_id_with_error",
			fields: fields{r: r},
			args: args{user: &models.User{
				ID:        0,
				FirstName: "test",
				LastName:  "test",
				Password:  "test",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: nil,
			}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				r: tt.fields.r,
			}
			if err := s.UpdateUserByID(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUserByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
