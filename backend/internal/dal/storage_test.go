package dal

import (
	"reflect"
	"testing"
)

func TestNewLocalVolume(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name string
		args args
		want *LocalVolume
	}{
		{
			name: "Desktop",
			args: args{p: "/Users/neo/Desktop"},
			want: &LocalVolume{
				volume{
					path:     "/Users/neo/Desktop/darkspace_volume",
					defaults: "/Users/neo/Desktop/darkspace_volume/defaults",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := NewLocalVolume(tt.args.p); !reflect.DeepEqual(
					got,
					tt.want,
				) {
					t.Errorf("NewLocalVolume() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
