package easy

// In MATLAB, there is a handy function called reshape which can reshape an
// m x n matrix into a new one with a different size r x c keeping its original data.
//
// You are given an m x n matrix mat and two integers r and c representing the number
// of rows and the number of columns of the wanted reshaped matrix.
//
// The reshaped matrix should be filled with all the elements of the original matrix in
// the same row-traversing order as they were.
//
// If the reshape operation with given parameters is possible and legal, output the new
// reshaped matrix; Otherwise, output the original matrix.
//
//
//
// Example 1:
// Input: mat = [[1,2],[3,4]], r = 1, c = 4
// Output: [[1,2,3,4]]
//
// Example 2:
// Input: mat = [[1,2],[3,4]], r = 2, c = 4
// Output: [[1,2],[3,4]]
func matrixReshape(mat [][]int, r int, c int) [][]int {
	rows := len(mat)
	columns := len(mat[0])

	if r*c != rows*columns {
		return mat
	}

	cn := 0
	rn := 0

	res := make([][]int, r)
	res[rn] = make([]int, c)

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			res[rn][cn] = mat[i][j]
			cn++

			if cn == c {
				cn = 0
				rn++
				if rn < r {
					res[rn] = make([]int, c)
				}
			}
		}
	}

	return res
}

// better solution...
// vector<vector<int>> matrixReshape(vector<vector<int>>& mat, int r, int c) {
//        if (r * c != mat.size() * mat[0].size())return mat;
//
//        vector<vector<int>>ans(r , vector<int>(c));
//        for(int i = 0 , length = r * c; i < length ; ++ i)
//            ans[i / c][i % c] = mat[i / mat[0].size()][i % mat[0].size()];
//        return ans;
//    }
