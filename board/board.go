package board

type Board struct {
    board [8*8]bool
    colFilled [8]bool
    rowFilled [8]bool
}

/*
Replaces the column of the Board with one that has a queen on specified square
row and col must both be between 0 and 7
*/
func (b Board) PlaceQueen(row int, col int) Board{
    b.board[row * 8 + col] = true
    b.rowFilled[row] = true
    b.colFilled[col] = true
    return b
}

/*
Checks if a queen is placed on the board on the given square
row and col must both be between 0 and 7
*/
func (b Board) HasQueen(row int, col int) bool {
    return b.board[row * 8 + col]
}


func (b Board) queenOnDiags(row int, col int) bool {
	return b.queenOnPartialDiagonal(row, col, 1, 1) ||
		b.queenOnPartialDiagonal(row, col, 1, -1) ||
		b.queenOnPartialDiagonal(row, col, -1, 1) ||
		b.queenOnPartialDiagonal(row, col, -1, -1)
}

func (b Board) queenOnPartialDiagonal(row int, col int, deltaRow int, deltaCol int) bool {
    for r,c := row,col; r < 8 && r >= 0 && c < 8 && c >= 0; r,c  = r + deltaRow, c + deltaCol {
        if b.HasQueen(r, c) {
            return true
        }
    }
	return false
}

func (b Board) rowIsFilled(row int) bool{
    return b.rowFilled[row]
}
func (b Board) colIsFilled(col int) bool{
    return b.colFilled[col]
}

/*
Checks if a queen can be placed on the board on the given square so that it is not attacked by any of the queens already on the board
row and col must both be between 0 and 7
*/
func (b Board) CanPlaceQueen(row int, col int) bool {
    return !(b.rowIsFilled(row) || b.colIsFilled(col) || b.queenOnDiags(row, col))
}

/*
Check if this board has the same configuration as the given board
*/
func (a Board) Equals(b Board) bool{
    return a.board == b.board
}

/*
Same as CanPlaceQueen, except we do not check if there is already a queen in the same column. 
This is purely because for our algorithm in the 8-queens problem we incrementally fill the columns
row and col must both be between 0 and 7
*/
func (b Board) CanPlaceQueenNoVertCheck(row int, col int) bool {
    return !(b.queenOnDiags(row, col) || b.rowIsFilled(row))
}

/*
Gives a string representation of the board
*/
func (b Board) ToString() string {
    var s string
    for r := 0; r < 8; r++ {
        for c := 0; c < 8; c++ {
            if b.HasQueen(r, c) {
                s += "Q"
            } else {
                s += "X"
            }
            s += "|"
        }
        s += "\n----------------\n"
    }
    return s
}
