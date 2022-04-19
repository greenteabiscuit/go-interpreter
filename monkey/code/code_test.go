package code

import (
	"reflect"
	"testing"
)

func TestMake(t *testing.T) {
	type args struct {
		op       Opcode
		operands []int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "success",
			args: args{
				op:       OpConstant,
				operands: []int{65534},
			},
			want: []byte{byte(OpConstant), 255, 254},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Make(tt.args.op, tt.args.operands...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Make() = %v, want %v", got, tt.want)
			}
		})
	}
}
