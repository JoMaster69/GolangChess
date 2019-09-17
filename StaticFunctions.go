package main

import (
	"math"
)

type MoveSortingMode byte
const (
	NoSorting        MoveSortingMode = 0
	BestMovesFirst   MoveSortingMode = 1
	OnlyHittingMoves MoveSortingMode = 2
)

// BoardPosition functions

func eval(boardPos BoardPosition) float32{
	return -0.3
}

func kingAtColor(boardPos BoardPosition, point Point, color bool) bool {
	if color == true {
		for _, element := range boardPos.WhitePieces {
			if element.getPosition() == point { if element.getValue() == 10 {return true} }
		}
	}
	if color == false {
		for _, element := range boardPos.BlackPieces {
			if element.getPosition() == point { if element.getValue() == 10 {return true} }
		}
	}
	return false
}

// returns true, if piece exists at point
func pieceAt(boardPos BoardPosition, point Point) bool{
	for _, element := range boardPos.WhitePieces {
		if element.getPosition() == point { return true }
	}
	for _, element := range boardPos.BlackPieces {
		if element.getPosition() == point { return true }
	}
	return false
}

// returns true, if piece of given color exists at point
func pieceAtColor(boardPos BoardPosition, point Point, color bool) bool{
	if color == true {
		for _, element := range boardPos.WhitePieces {
			if element.getPosition() == point { return true }
		}
	}
	if color == false {
		for _, element := range boardPos.BlackPieces {
			if element.getPosition() == point { return true }
		}
	}
	return false
}

// returns piece at point, else emptyPiece
func getPiece(boardPos BoardPosition, point Point) *Pawn{
	a := Pawn{}
	return &a
}

// returns all valid moves for given position
func allValidMoves(boardPos BoardPosition, sortingMode MoveSortingMode) []Move{
	var retMoveList []Move
	// All white moves
	if boardPos.nextMove {
		for _, element := range boardPos.WhitePieces {
			retMoveList = append(retMoveList, element.allMoves(boardPos)...)
		}
	}
	// All black moves
	if !boardPos.nextMove {
		for _, element := range boardPos.BlackPieces {
			retMoveList = append(retMoveList, element.allMoves(boardPos)...)
		}
	}
	return retMoveList
}

// returns a clone of boardPos
func clone(boardPos BoardPosition) BoardPosition {
	var newPos BoardPosition
	for _, piece := range boardPos.WhitePieces {
		newPos.WhitePieces = append(newPos.WhitePieces, piece.clone())
	}
	for _, piece := range boardPos.BlackPieces {
		newPos.BlackPieces = append(newPos.BlackPieces, piece.clone())
	}
	newPos.nextMove = boardPos.nextMove
	newPos.prevMove.start.x = boardPos.prevMove.start.x
	newPos.prevMove.start.y = boardPos.prevMove.start.y
	newPos.prevMove.end.x = boardPos.prevMove.end.x
	newPos.prevMove.end.y = boardPos.prevMove.end.y
	return newPos
}

// returns false if way is blocked
func freeWay(boardPos BoardPosition, move Move) bool {
	// horizontal
	if move.start.y == move.end.y && move.start.x != move.end.x {
		horizontalDiff := int(move.start.x)-int(move.end.x)
		if horizontalDiff < 0 {horizontalDiff *= -1}
		if move.start.x < move.end.x {
			for i := 1; i < horizontalDiff; i++ {
				if pieceAt(boardPos, Point{move.start.x+byte(i), move.start.y}){ return false }
			}
		}
		if move.start.x > move.end.x {
			for i := 1; i < horizontalDiff; i++ {
				if pieceAt(boardPos, Point{move.start.x-byte(i), move.start.y}){ return false }
			}
		}
	}

	// vertical
	if move.start.y != move.end.y && move.start.x == move.end.x {
		verticalDiff := int(move.start.y)-int(move.end.y)
		if verticalDiff < 0 {verticalDiff *= -1}
		if move.start.y < move.end.y {
			for i := 1; i < verticalDiff; i++ {
				if pieceAt(boardPos, Point{move.start.x, move.start.y+byte(i)}){ return false }
			}
		}
		if move.start.y > move.end.y {
			for i := 1; i < verticalDiff; i++ {
				if pieceAt(boardPos, Point{move.start.x, move.start.y-byte(i)}){ return false }
			}
		}
	}

	//diagonal
	pointX := int(move.end.x) - int(move.start.x)
	pointY := int(move.end.y) - int(move.start.y)
	if math.Abs(float64(pointX)) == math.Abs(float64(pointY)) {
		var i byte
		// topRight
		if pointX > 0 && pointY > 0 {
			for i = 1; i < byte(pointY); i++ {
				if pieceAt(boardPos, Point{move.start.x+i, move.start.y+i}) { return false }
			}
		}
		// topLeft
		if pointX < 0 && pointY > 0 {
			for i = 1; i < byte(pointY); i++ {
				if pieceAt(boardPos, Point{move.start.x-i, move.start.y+i}) { return false }
			}
		}
		// bottomRight
		if pointX > 0 && pointY < 0 {
			for i = 1; i < byte(pointX); i++ {
				if pieceAt(boardPos, Point{move.start.x+i, move.start.y-i}) { return false }
			}
		}
		//bottomLeft
		if pointX < 0 && pointY < 0 {
			for i = 1; i < byte(math.Abs(float64(pointX))); i++ {
				if pieceAt(boardPos, Point{move.start.x-i, move.start.y-i}) { return false }
			}
		}
	}
	return true
}

func KingBlockingKing(boardPos BoardPosition, point Point) bool{
	if kingAtColor(boardPos, Point{point.x+1, point.y+1}, !boardPos.nextMove) {return true}
	if kingAtColor(boardPos, Point{point.x+1, point.y}, !boardPos.nextMove) {return true}
	if kingAtColor(boardPos, Point{point.x+1, point.y-1}, !boardPos.nextMove) {return true}
	if kingAtColor(boardPos, Point{point.x-1, point.y+1}, !boardPos.nextMove) {return true}
	if kingAtColor(boardPos, Point{point.x-1, point.y}, !boardPos.nextMove) {return true}
	if kingAtColor(boardPos, Point{point.x-1, point.y-1}, !boardPos.nextMove) {return true}
	if kingAtColor(boardPos, Point{point.x, point.y+1}, !boardPos.nextMove) {return true}
	if kingAtColor(boardPos, Point{point.x, point.y-1}, !boardPos.nextMove) {return true}
	return false
}

// returns board coordinate from 1-768 coordinates
func posIntToPoint(posInt int) Point{
	if posInt == 0 {
		return Point{8,8}
	}
	var helpY float32 = (float32)(posInt) / 8.0
	var posY byte
	if helpY <= 1 {
		posY = 1
	}
	if helpY > 1 && helpY <= 2 {
		posY = 2
	}
	if helpY > 2 && helpY <= 3 {
		posY = 3
	}
	if helpY > 3 && helpY <= 4 {
		posY = 4
	}
	if helpY > 4 && helpY <= 5 {
		posY = 5
	}
	if helpY > 5 && helpY <= 6 {
		posY = 6
	}
	if helpY > 6 && helpY <= 7 {
		posY = 7
	}
	if helpY > 7 {
		posY = 8
	}
	var posX byte = byte(posInt) - (posY-1) * 8
	return Point{posX, posY}
}