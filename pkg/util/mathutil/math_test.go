package mathutil

import (
	"testing"
)

func TestBinaryToHex(t *testing.T) {
	type args struct {
		binaryStr string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "binaryToHex1 success",
			args: args{
				binaryStr: "001010",
			},
			want:    "0A",
			wantErr: false,
		},
		{
			name: "binaryToHex2 success",
			args: args{
				binaryStr: "00001010",
			},
			want:    "0A",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BinaryToHex(tt.args.binaryStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("BinaryToHex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("BinaryToHex() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexToBinary(t *testing.T) {
	type args struct {
		hexStr string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "HexToBinary1 success",
			args: args{
				hexStr: "0A",
			},
			want:    "00001010",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HexToBinary(tt.args.hexStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("HexToBinary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HexToBinary() got = %v, want %v", got, tt.want)
			}
		})
	}
}
