package pb

import (
	"reflect"
	"testing"

	"google.golang.org/protobuf/runtime/protoimpl"
)

func TestNewOrderCreated(t *testing.T) {
	type args struct {
		id      string
		version int64
	}
	tests := []struct {
		name string
		args args
		want *OrderCreated
	}{
		{
			name: "75ca6ff5-df06-414c-a8b7-ffc38e167858",
			args: args{
				id:      "1",
				version: 0,
			},
			want: &OrderCreated{
				state:   protoimpl.MessageState{},
				OrderId: "1",
				At:      nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOrderCreated(tt.args.id, tt.args.version); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOrderCreated() = %v, want %v", got, tt.want)
			}
		})
	}
}
