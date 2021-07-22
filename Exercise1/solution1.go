package main
import "fmt"
type Matrix struct{
	rows int
	columns int
	matrix [][] int
}
func (p Matrix) get_no_of_rows () int {
	return p.rows
}
func (p Matrix) get_no_of_columns () int {
	return p.columns
}
func (p *Matrix) set_element (i,j,element int) {
	p.matrix[i][j]= element
}

func (p Matrix) addition(v Matrix) *Matrix {
	rows :=v.rows
	columns :=v.columns
	sum:=Matrix{rows,columns,make([][]int,rows) }
	for iter:=0;iter<rows;iter++ {
		sum.matrix[iter]=make([]int,columns)
		for iter2:=0;iter2<columns; iter2++ {
			sum.matrix[iter][iter2]= p.matrix[iter][iter2] + v.matrix[iter][iter2]
		}
	}
	return &sum
}
func main(){
	firstmatrix:=Matrix{4,3,make([][]int,4)}
	secondmatrix:=Matrix{4,3,make([][]int,4)}
	for row:=0;row<4;row++ {
		firstmatrix.matrix[row]=make([]int,3)
		secondmatrix.matrix[row]=make([]int,3)
		for col:=0; col<3;col++{
			firstmatrix.set_element(row,col,row+col)
			secondmatrix.set_element(row,col,row-col)
		}
	}
	fmt.Println("Number of rows in Matrix is ",firstmatrix.get_no_of_rows())
	fmt.Println("Number of columns in Matrix is ", firstmatrix.get_no_of_columns())
	fmt.Println(firstmatrix.matrix)
	fmt.Println(secondmatrix.matrix)
	fmt.Print("Addition of matrices is ")
	result:= firstmatrix.addition(secondmatrix)
	fmt.Print(result.matrix)
}
