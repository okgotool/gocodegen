package biz

import (
	"testing"
)

func TestCreateAndComparePassword(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test_1",
			args: args{
				str: "111111",
			},
			want:    "111111",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreatePasswordHash(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("Md5() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !IsPasswordRight(got, tt.args.str) {
				t.Errorf("Md5() = %v, want %v", got, tt.args.str)
			}
			if IsPasswordRight(got, tt.args.str+"wrong") {
				t.Errorf("Md5() = %v", got)
			}
		})
	}
}

func TestCreateAndDecodeToken(t *testing.T) {
	type args struct {
		tokenPairs map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test_1",
			args: args{
				tokenPairs: map[string]interface{}{
					"name":  "admin",
					"roles": []string{"admin", "user"},
				},
			},
			want:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiYWRtaW4iLCJyb2xlcyI6WyJhZG1pbiIsInVzZXIiXX0.FcuF_BbJ83E7Oii_ch993yRiK6zV0XFTuCToGzDYPUs",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateToken(tt.args.tokenPairs)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateToken() = %v, want %v", got, tt.want)
			}

			decodeMap, err := DecodeToken(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if v, ok := decodeMap["name"]; !ok || v != "admin" {
				t.Errorf("DecodeToken() failed, decoded = %v, src = %v", decodeMap, tt.args.tokenPairs)
				return
			}
		})
	}
}
