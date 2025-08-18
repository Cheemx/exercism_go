package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	arr, ok := cb[file]
    if !ok {
        return 0
    }
    cnt := 0
    for _, val := range arr {
        if val {
            cnt++
        }
    }
    return cnt
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
    if rank < 1 || rank > 8 {
        return 0
    }
    cnt := 0
	for _, v := range cb {
        if v[rank - 1] {
            cnt++
        }
    }
    return cnt
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
    cnt := 0
	for _, v := range cb {
        cnt += len(v)
    }
    return cnt
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	cnt := 0
	for _, v := range cb {
        for _, val := range v {
            if val {
            cnt++
        	}
        }
    }
    return cnt
}
