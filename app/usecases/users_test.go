package usecases

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"hoge.com/hoge/domain/entity"
	"hoge.com/hoge/domain/repository"
)

func Test_usersUsecase_FetchUsers(t *testing.T) {
	tests := []struct {
		name    string
		want    []entity.User
		wantErr bool
	}{
		{
			name: "success",
			want: []entity.User{
				{
					ID:   "123",
					Name: "456",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mock := repository.NewMockUser(ctrl)

			mock.EXPECT().Fetch().Return(tt.want, nil)

			u := &usersUsecase{
				usersRepository: mock,
			}
			got, err := u.FetchUsers()
			if (err != nil) != tt.wantErr {
				t.Errorf("usersUsecase.FetchUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("usersUsecase.FetchUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}
