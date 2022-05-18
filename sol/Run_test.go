package sol

import (
	"reflect"
	"strconv"
	"testing"
)

func CreateBinaryTree(input *[]string) *TreeNode {
	var tree *TreeNode
	arr := *input
	result := make([]*TreeNode, len(arr))
	for idx, val := range arr {
		if val != "null" {
			num, _ := strconv.Atoi(val)
			result[idx] = &TreeNode{Val: num}
		}
		if idx == 0 {
			tree = result[idx]
		}
	}
	for idx := range result {
		if result[idx] != nil {
			if 2*idx+1 < len(result) {
				result[idx].Left = result[2*idx+1]
			}
			if 2*idx+2 < len(result) {
				result[idx].Right = result[2*idx+2]
			}
		}
	}
	return tree
}
func TestRun(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "[1,2,3,null,null,4,5]",
			args: args{root: CreateBinaryTree(&[]string{"1", "2", "3", "null", "null", "4", "5"})},
			want: CreateBinaryTree(&[]string{"1", "2", "3", "null", "null", "4", "5"}),
		},
		{
			name: "[1,2,3,null,null,4,5,6,7]",
			args: args{root: CreateBinaryTree(&[]string{"1", "2", "3", "null", "null", "4", "5", "null", "null", "null", "null", "6", "7"})},
			want: CreateBinaryTree(&[]string{"1", "2", "3", "null", "null", "4", "5", "null", "null", "null", "null", "6", "7"}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Run(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
