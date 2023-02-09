package service

import "testing"

func TestGenerateToken(t *testing.T) {
	type args struct {
		userID   uint64
		username string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateToken(tt.args.userID, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GenerateToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsUsernameExist(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsUsernameExist(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsUsernameExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsUsernameExist() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsUsernameLegal(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUsernameLegal(tt.args.username); got != tt.want {
				t.Errorf("IsUsernameLegal() = %v, want %v", got, tt.want)
			}
		})
	}
}
