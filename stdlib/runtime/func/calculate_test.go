/**
 * @Author       : Gainovel
 * @Organization : Copyright © 2024 gainovel.com All Rights Reserved.
 * @Date         : 2024-07-25 22:28:26 星期四
 * @ProductName  : GoLand
 * @ProjectName  : testcase
 * @File         : stdlib/runtime/func/calculate_test.go
 * @Description  : 开发中···
 **/

package funct

import "testing"

func TestName_2024_07_25_22_28_26(t *testing.T) {
	t.Run("case1", func(t *testing.T) {

	})
}

func Test_calculate(t *testing.T) {
	type args struct {
		x  int
		y  int
		op operate
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "case1",
			args: args{
				x:  8,
				y:  2,
				op: add,
			},
			want:    10,
			wantErr: true,
		},
		{
			name: "case2",
			args: args{
				x:  8,
				y:  2,
				op: sub,
			},
			want:    6,
			wantErr: true,
		},
		{
			name: "case3",
			args: args{
				x:  8,
				y:  2,
				op: mul,
			},
			want:    16,
			wantErr: true,
		},
		{
			name: "case4",
			args: args{
				x:  8,
				y:  2,
				op: div,
			},
			want:    4,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculate(tt.args.x, tt.args.y, tt.args.op)
			if (err == nil) != tt.wantErr {
				t.Errorf("calculate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("calculate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
