package Test999

import (
	"testing"
)

func Test(t *testing.T) {
	//b:=[][]byte{[]byte{'.','.','.','.','.','.','.','.'},[]byte{'.','.','.','p','.','.','.','.'},[]byte{'.','.','.','R','.','.','.','p'},[]byte{'.','.','.','.','.','.','.','.'},[]byte{'.','.','.','.','.','.','.','.'},[]byte{'.','.','.','p','.','.','.','.'},[]byte{'.','.','.','.','.','.','.','.'},[]byte{'.','.','.','.','.','.','.','.'}}
	/*b:=[][]byte{
	[]byte{'.','.','.','.','.','.','.','.'},
	[]byte{'.','p','p','p','p','p','.','.'},
	[]byte{'.','p','p','B','p','p','.','.'},
	[]byte{'.','p','B','R','B','p','.','.'},
	[]byte{'.','p','p','B','p','p','.','.'},
	[]byte{'.','p','p','p','p','p','.','.'},
	[]byte{'.','.','.','.','.','.','.','.'},
	[]byte{'.','.','.','.','.','.','.','.'}}*/
	b := [][]byte{
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', 'p', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', 'p', '.', '.', '.', '.'},
		[]byte{'p', 'p', '.', 'R', '.', 'p', 'B', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', 'B', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', 'p', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.'}}
	t.Log(numRookCaptures(b))
}
func numRookCaptures(board [][]byte) int {
	var i, j int
	for ; i < 8; i++ {
		for j = i; j < 8; j++ {
			if board[i][j] == 'r' || board[i][j] == 'R' {
				goto start
			}
		}
	}
start:
	if i == 8 && j == 8 {
		return 0
	}
	var cnt int
	for k := i; k > 0; k-- { //上
		if board[k][j] == 'b' || board[k][j] == 'B' {
			break
		}
		if board[k][j] == (board[i][j]-('r'-'p'))^32 {
			cnt++
			break
		}
	}
	for k := i; k < 8; k++ { //下
		if board[k][j] == 'b' || board[k][j] == 'B' {
			break
		}
		if board[k][j] == (board[i][j]-('r'-'p'))^32 {
			cnt++
			break
		}
	}
	for k := j; k > 0; k-- { //左
		if board[i][k] == 'b' || board[i][k] == 'B' {
			break
		}
		if board[i][k] == (board[i][j]-('r'-'p'))^32 {
			cnt++
			break
		}
	}
	for k := j; k < 8; k++ { //右
		if board[i][k] == 'b' || board[i][k] == 'B' {
			break
		}
		if board[i][k] == (board[i][j]-('r'-'p'))^32 {
			cnt++
			break
		}
	}
	return cnt
}
