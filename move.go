package main

type move struct {
	fromX int
	fromY int
	toX   int
	toY   int
	taken int
}

func generatePawnMoves(b *board, myCoords [2]int, kingCoords [2]int, colour int, moves []move) []move {
	// Enemy values to be cached
	enemyInLeftDirectionCached := false
	enemyInLeftDirection := false
	enemyInRightDirectionCached := false
	enemyInRightDirection := false
	enemyInBottomLeftDiagonalCached := false
	enemyInBottomLeftDiagonal := false
	enemyInTopLeftDiagonalCached := false
	enemyInTopLeftDiagonal := false
	enemyInTopRightDiagonalCached := false
	enemyInTopRightDiagonal := false
	enemyInBottomRightDiagonalCached := false
	enemyInBottomRightDiagonal := false

	type pawnMove struct {
		move move
		typ  string
	}
	// Generate list of possible moves, ignoring bounds or checks
	possibleMoves := [4]pawnMove{
		{move{myCoords[0], myCoords[1], myCoords[0], myCoords[1] + 1, empty}, "up1"},         // up 1 square
		{move{myCoords[0], myCoords[1], myCoords[0], myCoords[1] + 2, empty}, "up2"},         // up 2 squares
		{move{myCoords[0], myCoords[1], myCoords[0] - 1, myCoords[1] + 1, empty}, "upleft"},  // up and left
		{move{myCoords[0], myCoords[1], myCoords[0] + 1, myCoords[1] + 1, empty}, "upright"}, // up and right
	}
	// If black, reverse y axis
	if colour == black {
		possibleMoves[0].move.toY = myCoords[1] - 1
		possibleMoves[1].move.toY = myCoords[1] - 2
		possibleMoves[2].move.toY = myCoords[1] - 1
		possibleMoves[2].typ = "downleft"
		possibleMoves[3].move.toY = myCoords[1] - 1
		possibleMoves[3].typ = "downright"
	}

	for _, possibleMove := range possibleMoves {
		toX := possibleMove.move.toX
		toY := possibleMove.move.toY
		fromX := possibleMove.move.fromX
		fromY := possibleMove.move.fromY
		typ := possibleMove.typ

		// First check bounds
		if toX > 7 || toX < 0 || toY > 7 || toY < 0 {
			continue
		}

		// Next, check the target square. Forward moves require an empty square, diagonal moves require an enemy
		if toX == fromX {
			// 1 square
			if toY-fromY == 1 || toY-fromY == -1 {
				if b.squares[toX][toY] != empty {
					continue
				}
			}

			// 2 squares
			if toY-fromY == 2 || toY-fromY == -2 {
				if b.squares[toX][toY] != empty {
					continue
				}

				// Also need to make sure the square in between is empty
				if colour == white && b.squares[toX][toY-1] != empty ||
					colour == black && b.squares[toX][toY+1] != empty {
					continue
				}
			}
		} else {
			if colour == white && !isBlack(b.squares[toX][toY]) ||
				colour == black && !isWhite(b.squares[toX][toY]) {
				continue
			}
		}

		/*
		 * Next, check whether this move will lead to check
		 */
		enemyInWay := false

		/*
		 * Am I on the same row as my king? If so, will moving forward expose the king to a rook/queen?
		 */
		if myCoords[1] == kingCoords[1] {
			if myCoords[0] < kingCoords[0] {
				if enemyInLeftDirectionCached {
					enemyInWay = enemyInLeftDirection
				} else {
					enemyInWay = isEnemyInDirection(b, kingCoords, myCoords, colour, left)
					enemyInLeftDirection = enemyInWay
					enemyInLeftDirectionCached = true
				}
			} else {
				if enemyInRightDirectionCached {
					enemyInWay = enemyInRightDirection
				} else {
					enemyInWay = isEnemyInDirection(b, kingCoords, myCoords, colour, right)
					enemyInRightDirection = enemyInWay
					enemyInRightDirectionCached = true
				}
			}
		}

		if enemyInWay {
			continue
		}

		/*
		 * Am I on the same diagonal as my king? If so, will this move expose the king to a bishop/queen?
		 */
		xdiff := kingCoords[0] - myCoords[0]
		ydiff := kingCoords[1] - myCoords[1]

		if xdiff == ydiff {
			// Bottom left or top right
			if xdiff > 0 {
				// Bottom left
				if typ != "downleft" {
					if enemyInBottomLeftDiagonalCached {
						enemyInWay = enemyInBottomLeftDiagonal
					} else {
						enemyInWay = isEnemyOnDiagonal(b, kingCoords, myCoords, colour, bottomLeftDiagonal)
						enemyInBottomLeftDiagonal = enemyInWay
						enemyInBottomLeftDiagonalCached = true
					}
				}
			} else {
				// Top right
				if typ != "upright" {
					if enemyInTopRightDiagonalCached {
						enemyInWay = enemyInTopRightDiagonal
					} else {
						enemyInWay = isEnemyOnDiagonal(b, kingCoords, myCoords, colour, topRightDiagonal)
						enemyInTopRightDiagonal = enemyInWay
						enemyInTopRightDiagonalCached = true
					}
				}
			}
		} else if xdiff == -ydiff {
			// Bottom right or top left
			if xdiff > 0 {
				// Top left
				if typ != "upleft" {
					if enemyInTopLeftDiagonalCached {
						enemyInWay = enemyInTopLeftDiagonal
					} else {
						enemyInWay = isEnemyOnDiagonal(b, kingCoords, myCoords, colour, topLeftDiagonal)
						enemyInTopLeftDiagonal = enemyInWay
						enemyInTopLeftDiagonalCached = true
					}
				}
			} else {
				if typ != "downright" {
					// Bottom right
					if enemyInBottomRightDiagonalCached {
						enemyInWay = enemyInBottomRightDiagonal
					} else {
						enemyInWay = isEnemyOnDiagonal(b, kingCoords, myCoords, colour, bottomRightDiagonal)
						enemyInBottomRightDiagonal = enemyInWay
						enemyInBottomRightDiagonalCached = true
					}
				}
			}
		}

		if enemyInWay == false {
			moves = append(moves, possibleMove.move)
		}
	}
	return moves
}

func generateRookMoves(b *board, myCoords [2]int, kingCoords [2]int, colour int, moves []move) []move {
	// Can I move left or right?
	if myCoords[0] > 0 && (b.squares[myCoords[0]-1][myCoords[1]] == empty ||
		(colour == white && isBlack(b.squares[myCoords[0]-1][myCoords[1]])) ||
		(colour == black && isWhite(b.squares[myCoords[0]-1][myCoords[1]]))) ||
		myCoords[0] < 7 && (b.squares[myCoords[0]+1][myCoords[1]] == empty ||
			(colour == white && isBlack(b.squares[myCoords[0]+1][myCoords[1]])) ||
			(colour == black && isWhite(b.squares[myCoords[0]+1][myCoords[1]]))) {

		// First check whether I'll be moving into check
		if myCoords[0] != kingCoords[0] || (!isEnemyInDirection(b, kingCoords, myCoords, colour, up) &&
			!isEnemyInDirection(b, kingCoords, myCoords, colour, down)) {
			// Move left until find a piece or the edge
			for xmove := myCoords[0] - 1; xmove >= 0 && xmove < 8; xmove-- {
				// If this square is empty, it's a move and continue
				if b.squares[xmove][myCoords[1]] == empty {
					moves = append(moves, move{myCoords[0], myCoords[1], xmove, myCoords[1], empty})
				} else {
					// If this square is an enemy piece, it's a move
					if colour == white && isBlack(b.squares[xmove][myCoords[1]]) ||
						colour == black && isWhite(b.squares[xmove][myCoords[1]]) {
						moves = append(moves, move{myCoords[0], myCoords[1], xmove, myCoords[1], empty})
					}

					break
				}
			}

			// Move right until find a piece  or the edge
			for xmove := myCoords[0] + 1; xmove >= 0 && xmove < 8; xmove++ {
				// If this square is empty, it's a move and continue
				if b.squares[xmove][myCoords[1]] == empty {
					moves = append(moves, move{myCoords[0], myCoords[1], xmove, myCoords[1], empty})
				} else {
					// If this square is an enemy piece, it's a move
					if colour == white && isBlack(b.squares[xmove][myCoords[1]]) ||
						colour == black && isWhite(b.squares[xmove][myCoords[1]]) {
						moves = append(moves, move{myCoords[0], myCoords[1], xmove, myCoords[1], empty})
					}

					break
				}
			}
		}
	}

	// Can I move up or down?
	if myCoords[1] > 0 && (b.squares[myCoords[0]][myCoords[1]-1] == empty ||
		(colour == white && isBlack(b.squares[myCoords[0]][myCoords[1]-1])) ||
		(colour == black && isWhite(b.squares[myCoords[0]][myCoords[1]-1]))) ||
		myCoords[1] < 7 && (b.squares[myCoords[0]][myCoords[1]+1] == empty ||
			(colour == white && isBlack(b.squares[myCoords[0]][myCoords[1]+1])) ||
			(colour == black && isWhite(b.squares[myCoords[0]][myCoords[1]+1]))) {

		// First check whether I'll be moving into check
		if myCoords[1] != kingCoords[1] || (!isEnemyInDirection(b, kingCoords, myCoords, colour, left) &&
			!isEnemyInDirection(b, kingCoords, myCoords, colour, right)) {
			// Move left until find a piece of my own colour or the edge
			for ymove := myCoords[1] - 1; ymove >= 0 && ymove < 8; ymove-- {
				// If this square is empty, it's a move and continue
				if b.squares[myCoords[0]][ymove] == empty {
					moves = append(moves, move{myCoords[0], myCoords[1], myCoords[0], ymove, empty})
				} else {
					// If this square is an enemy piece, it's a move
					if colour == white && isBlack(b.squares[myCoords[0]][ymove]) ||
						colour == black && isWhite(b.squares[myCoords[0]][ymove]) {
						moves = append(moves, move{myCoords[0], myCoords[1], myCoords[0], ymove, empty})
					}

					break
				}
			}

			// Move right until find a piece of my own colour or the edge
			for ymove := myCoords[1] + 1; ymove >= 0 && ymove < 8; ymove++ {
				// If this square is empty, it's a move and continue
				if b.squares[myCoords[0]][ymove] == empty {
					moves = append(moves, move{myCoords[0], myCoords[1], myCoords[0], ymove, empty})
				} else {
					// If this square is an enemy piece, it's a move
					if colour == white && isBlack(b.squares[myCoords[0]][ymove]) ||
						colour == black && isWhite(b.squares[myCoords[0]][ymove]) {
						moves = append(moves, move{myCoords[0], myCoords[1], myCoords[0], ymove, empty})
					}

					break
				}
			}
		}
	}

	return moves
}

func generateKnightMoves(b *board, myCoords [2]int, kingCoords [2]int, colour int, moves []move) []move {
	/* First need to check whether I'm in between king and an enemy piece on row, col and diagonal. If so,
	 * I can't move anywhere
	 */

	// Check for enemy on col
	if myCoords[0] == kingCoords[0] {
		if myCoords[1] > kingCoords[1] {
			if isEnemyInDirection(b, kingCoords, myCoords, colour, up) {
				return moves
			}
		} else {
			if isEnemyInDirection(b, kingCoords, myCoords, colour, down) {
				return moves
			}
		}
	}

	// Check for enemy on row
	if myCoords[1] == kingCoords[1] {
		if myCoords[0] > kingCoords[0] {
			if isEnemyInDirection(b, kingCoords, myCoords, colour, right) {
				return moves
			}
		} else {
			if isEnemyInDirection(b, kingCoords, myCoords, colour, left) {
				return moves
			}
		}
	}

	// Check for enemy on diagonal
	xdiff := kingCoords[0] - myCoords[0]
	ydiff := kingCoords[1] - myCoords[1]

	if xdiff == ydiff {
		// Bottom left or top right
		if xdiff > 0 {
			// Bottom left
			if isEnemyOnDiagonal(b, kingCoords, myCoords, colour, bottomLeftDiagonal) {
				return moves
			}
		} else {
			// Top right
			if isEnemyOnDiagonal(b, kingCoords, myCoords, colour, topRightDiagonal) {
				return moves
			}
		}
	} else if xdiff == -ydiff {
		// Bottom right or top left
		if xdiff > 0 {
			// Top left
			if isEnemyOnDiagonal(b, kingCoords, myCoords, colour, topLeftDiagonal) {
				return moves
			}
		} else {
			// Bottom right
			if isEnemyOnDiagonal(b, kingCoords, myCoords, colour, bottomRightDiagonal) {
				return moves
			}
		}
	}

	// Now we've established no danger of moving into check, so can check possible moves

	possibleMoves := [8]move{
		{myCoords[0], myCoords[1], myCoords[0] - 1, myCoords[1] - 2, empty},
		{myCoords[0], myCoords[1], myCoords[0] - 2, myCoords[1] - 1, empty},
		{myCoords[0], myCoords[1], myCoords[0] - 2, myCoords[1] + 1, empty},
		{myCoords[0], myCoords[1], myCoords[0] - 1, myCoords[1] + 2, empty},
		{myCoords[0], myCoords[1], myCoords[0] + 1, myCoords[1] + 2, empty},
		{myCoords[0], myCoords[1], myCoords[0] + 2, myCoords[1] + 1, empty},
		{myCoords[0], myCoords[1], myCoords[0] + 2, myCoords[1] - 1, empty},
		{myCoords[0], myCoords[1], myCoords[0] + 1, myCoords[1] - 2, empty},
	}

	for _, v := range possibleMoves {
		// First check bounds
		if v.toX > 7 || v.toX < 0 || v.toY > 7 || v.toY < 0 {
			continue
		}

		// Is there a piece of my own colour there?
		if (colour == white && isWhite(b.squares[v.toX][v.toY])) || (colour == black && isBlack(b.squares[v.toX][v.toY])) {
			continue
		}

		moves = append(moves, v)
	}
	return moves
}

func generateBishopMoves(b *board, myCoords [2]int, kingCoords [2]int, colour int, moves []move) []move {
	/* First need to check whether I'm in between king and an enemy piece on row or col. If so,
	 * I can't move anywhere
	 */

	// Check for enemy on col
	if myCoords[0] == kingCoords[0] {
		if myCoords[1] > kingCoords[1] {
			if isEnemyInDirection(b, kingCoords, myCoords, colour, up) {
				return moves
			}
		} else {
			if isEnemyInDirection(b, kingCoords, myCoords, colour, down) {
				return moves
			}
		}
	}

	// Check for enemy on row
	if myCoords[1] == kingCoords[1] {
		if myCoords[0] > kingCoords[0] {
			if isEnemyInDirection(b, kingCoords, myCoords, colour, right) {
				return moves
			}
		} else {
			if isEnemyInDirection(b, kingCoords, myCoords, colour, left) {
				return moves
			}
		}
	}

	// Now, check whether I can move along each diagonal without moving into check
	canMoveBottomLeft := true
	canMoveTopRight := true
	canMoveBottomRight := true
	canMoveTopLeft := true

	xdiff := kingCoords[0] - myCoords[0]
	ydiff := kingCoords[1] - myCoords[1]

	if xdiff == ydiff {
		// Bottom left or top right
		if xdiff > 0 {
			// Bottom left
			if isEnemyOnDiagonal(b, kingCoords, myCoords, colour, bottomLeftDiagonal) {
				canMoveTopLeft = false
				canMoveBottomRight = false
			}
		} else {
			// Top right
			if isEnemyOnDiagonal(b, kingCoords, myCoords, colour, topRightDiagonal) {
				canMoveTopLeft = false
				canMoveBottomRight = false
			}
		}
	} else if xdiff == -ydiff {
		// Bottom right or top left
		if xdiff > 0 {
			// Top left
			if isEnemyOnDiagonal(b, kingCoords, myCoords, colour, topLeftDiagonal) {
				canMoveTopRight = false
				canMoveBottomLeft = false
			}
		} else {
			// Bottom right
			if isEnemyOnDiagonal(b, kingCoords, myCoords, colour, bottomRightDiagonal) {
				canMoveTopRight = false
				canMoveBottomLeft = false
			}
		}
	}

	// Now, find all moves along each safe diagonal
	if canMoveBottomLeft {
		// Move until find a piece or the edge
		for xcoord, ycoord := myCoords[0]-1, myCoords[1]-1; xcoord >= 0 && ycoord >= 0; xcoord, ycoord = xcoord-1, ycoord-1 {
			// If this square is empty, it's a move and continue
			if b.squares[xcoord][ycoord] == empty {
				moves = append(moves, move{myCoords[0], myCoords[1], xcoord, ycoord, empty})
			} else {
				// If this square is an enemy piece, it's a move
				if colour == white && isBlack(b.squares[xcoord][ycoord]) ||
					colour == black && isWhite(b.squares[xcoord][ycoord]) {
					moves = append(moves, move{myCoords[0], myCoords[1], xcoord, ycoord, empty})
				}

				break
			}
		}
	}

	if canMoveBottomRight {
		// Move until find a piece or the edge
		for xcoord, ycoord := myCoords[0]+1, myCoords[1]-1; xcoord < 8 && ycoord >= 0; xcoord, ycoord = xcoord+1, ycoord-1 {
			// If this square is empty, it's a move and continue
			if b.squares[xcoord][ycoord] == empty {
				moves = append(moves, move{myCoords[0], myCoords[1], xcoord, ycoord, empty})
			} else {
				// If this square is an enemy piece, it's a move
				if colour == white && isBlack(b.squares[xcoord][ycoord]) ||
					colour == black && isWhite(b.squares[xcoord][ycoord]) {
					moves = append(moves, move{myCoords[0], myCoords[1], xcoord, ycoord, empty})
				}

				break
			}
		}
	}

	if canMoveTopLeft {
		// Move until find a piece or the edge
		for xcoord, ycoord := myCoords[0]-1, myCoords[1]+1; xcoord >= 0 && ycoord < 8; xcoord, ycoord = xcoord-1, ycoord+1 {
			// If this square is empty, it's a move and continue
			if b.squares[xcoord][ycoord] == empty {
				moves = append(moves, move{myCoords[0], myCoords[1], xcoord, ycoord, empty})
			} else {
				// If this square is an enemy piece, it's a move
				if colour == white && isBlack(b.squares[xcoord][ycoord]) ||
					colour == black && isWhite(b.squares[xcoord][ycoord]) {
					moves = append(moves, move{myCoords[0], myCoords[1], xcoord, ycoord, empty})
				}

				break
			}
		}
	}

	if canMoveTopRight {
		// Move until find a piece or the edge
		for xcoord, ycoord := myCoords[0]+1, myCoords[1]+1; xcoord < 8 && ycoord < 8; xcoord, ycoord = xcoord+1, ycoord+1 {
			// If this square is empty, it's a move and continue
			if b.squares[xcoord][ycoord] == empty {
				moves = append(moves, move{myCoords[0], myCoords[1], xcoord, ycoord, empty})
			} else {
				// If this square is an enemy piece, it's a move
				if colour == white && isBlack(b.squares[xcoord][ycoord]) ||
					colour == black && isWhite(b.squares[xcoord][ycoord]) {
					moves = append(moves, move{myCoords[0], myCoords[1], xcoord, ycoord, empty})
				}

				break
			}
		}
	}

	return moves
}

func generateQueenMoves(b *board, myCoords [2]int, kingCoords [2]int, colour int, moves []move) []move {
	// Can I move left or right?
	if myCoords[0] > 0 && (b.squares[myCoords[0]-1][myCoords[1]] == empty ||
		(colour == white && isBlack(b.squares[myCoords[0]-1][myCoords[1]])) ||
		(colour == black && isWhite(b.squares[myCoords[0]-1][myCoords[1]]))) ||
		myCoords[0] < 7 && (b.squares[myCoords[0]+1][myCoords[1]] == empty ||
			(colour == white && isBlack(b.squares[myCoords[0]+1][myCoords[1]])) ||
			(colour == black && isWhite(b.squares[myCoords[0]+1][myCoords[1]]))) {

		// First check whether I'll be moving into check
		if myCoords[0] != kingCoords[0] || (!isEnemyInDirection(b, kingCoords, myCoords, colour, up) && !isEnemyInDirection(b, kingCoords, myCoords, colour, down)) {
			// Move left until find a piece or the edge
			for xmove := myCoords[0] - 1; xmove >= 0 && xmove < 8; xmove-- {
				// If this square is empty, it's a move and continue
				if b.squares[xmove][myCoords[1]] == empty {
					moves = append(moves, move{myCoords[0], myCoords[1], xmove, myCoords[1], empty})
				} else {
					// If this square is an enemy piece, it's a move
					if colour == white && isBlack(b.squares[xmove][myCoords[1]]) ||
						colour == black && isWhite(b.squares[xmove][myCoords[1]]) {
						moves = append(moves, move{myCoords[0], myCoords[1], xmove, myCoords[1], empty})
					}

					break
				}
			}

			// Move right until find a piece  or the edge
			for xmove := myCoords[0] + 1; xmove >= 0 && xmove < 8; xmove++ {
				// If this square is empty, it's a move and continue
				if b.squares[xmove][myCoords[1]] == empty {
					moves = append(moves, move{myCoords[0], myCoords[1], xmove, myCoords[1], empty})
				} else {
					// If this square is an enemy piece, it's a move
					if colour == white && isBlack(b.squares[xmove][myCoords[1]]) ||
						colour == black && isWhite(b.squares[xmove][myCoords[1]]) {
						moves = append(moves, move{myCoords[0], myCoords[1], xmove, myCoords[1], empty})
					}

					break
				}
			}
		}
	}

	// Can I move up or down?
	if myCoords[1] > 0 && (b.squares[myCoords[0]][myCoords[1]-1] == empty ||
		(colour == white && isBlack(b.squares[myCoords[0]][myCoords[1]-1])) ||
		(colour == black && isWhite(b.squares[myCoords[0]][myCoords[1]-1]))) ||
		myCoords[1] < 7 && (b.squares[myCoords[0]][myCoords[1]+1] == empty ||
			(colour == white && isBlack(b.squares[myCoords[0]][myCoords[1]+1])) ||
			(colour == black && isWhite(b.squares[myCoords[0]][myCoords[1]+1]))) {

		// First check whether I'll be moving into check
		// TODO cache directions
		if myCoords[1] != kingCoords[1] || (!isEnemyInDirection(b, kingCoords, myCoords, colour, left) &&
			!isEnemyInDirection(b, kingCoords, myCoords, colour, right)) {
			// Move left until find a piece of my own colour or the edge
			for ymove := myCoords[1] - 1; ymove >= 0 && ymove < 8; ymove-- {
				// If this square is empty, it's a move and continue
				if b.squares[myCoords[0]][ymove] == empty {
					moves = append(moves, move{myCoords[0], myCoords[1], myCoords[0], ymove, empty})
				} else {
					// If this square is an enemy piece, it's a move
					if colour == white && isBlack(b.squares[myCoords[0]][ymove]) ||
						colour == black && isWhite(b.squares[myCoords[0]][ymove]) {
						moves = append(moves, move{myCoords[0], myCoords[1], myCoords[0], ymove, empty})
					}

					break
				}
			}

			// Move right until find a piece of my own colour or the edge
			for ymove := myCoords[1] + 1; ymove >= 0 && ymove < 8; ymove++ {
				// If this square is empty, it's a move and continue
				if b.squares[myCoords[0]][ymove] == empty {
					moves = append(moves, move{myCoords[0], myCoords[1], myCoords[0], ymove, empty})
				} else {
					// If this square is an enemy piece, it's a move
					if colour == white && isBlack(b.squares[myCoords[0]][ymove]) ||
						colour == black && isWhite(b.squares[myCoords[0]][ymove]) {
						moves = append(moves, move{myCoords[0], myCoords[1], myCoords[0], ymove, empty})
					}

					break
				}
			}
		}
	}

	/* Now diagonals...first need to check whether I'm in between king and an enemy piece on row or col. If so,
	 * I can't move anywhere
	 */

	// Check for enemy on col
	if myCoords[0] == kingCoords[0] {
		if myCoords[1] > kingCoords[1] {
			if isEnemyInDirection(b, kingCoords, myCoords, colour, up) {
				return moves
			}
		} else {
			if isEnemyInDirection(b, kingCoords, myCoords, colour, down) {
				return moves
			}
		}
	}

	// Check for enemy on row
	if myCoords[1] == kingCoords[1] {
		if myCoords[0] > kingCoords[0] {
			if isEnemyInDirection(b, kingCoords, myCoords, colour, right) {
				return moves
			}
		} else {
			if isEnemyInDirection(b, kingCoords, myCoords, colour, left) {
				return moves
			}
		}
	}

	// Now, check whether I can move along each diagonal without moving into check
	canMoveBottomLeft := true
	canMoveTopRight := true
	canMoveBottomRight := true
	canMoveTopLeft := true

	xdiff := kingCoords[0] - myCoords[0]
	ydiff := kingCoords[1] - myCoords[1]

	if xdiff == ydiff {
		// Bottom left or top right
		if xdiff > 0 {
			// Bottom left
			if isEnemyOnDiagonal(b, kingCoords, myCoords, colour, bottomLeftDiagonal) {
				canMoveTopLeft = false
				canMoveBottomRight = false
			}
		} else {
			// Top right
			if isEnemyOnDiagonal(b, kingCoords, myCoords, colour, topRightDiagonal) {
				canMoveTopLeft = false
				canMoveBottomRight = false
			}
		}
	} else if xdiff == -ydiff {
		// Bottom right or top left
		if xdiff > 0 {
			// Top left
			if isEnemyOnDiagonal(b, kingCoords, myCoords, colour, topLeftDiagonal) {
				canMoveTopRight = false
				canMoveBottomLeft = false
			}
		} else {
			// Bottom right
			if isEnemyOnDiagonal(b, kingCoords, myCoords, colour, bottomRightDiagonal) {
				canMoveTopRight = false
				canMoveBottomLeft = false
			}
		}
	}

	// Now, find all moves along each safe diagonal
	if canMoveBottomLeft {
		// Move until find a piece or the edge
		for xcoord, ycoord := myCoords[0]-1, myCoords[1]-1; xcoord >= 0 && ycoord >= 0; xcoord, ycoord = xcoord-1, ycoord-1 {
			// If this square is empty, it's a move and continue
			if b.squares[xcoord][ycoord] == empty {
				moves = append(moves, move{myCoords[0], myCoords[1], xcoord, ycoord, empty})
			} else {
				// If this square is an enemy piece, it's a move
				if colour == white && isBlack(b.squares[xcoord][ycoord]) ||
					colour == black && isWhite(b.squares[xcoord][ycoord]) {
					moves = append(moves, move{myCoords[0], myCoords[1], xcoord, ycoord, empty})
				}

				break
			}
		}
	}

	if canMoveBottomRight {
		// Move until find a piece or the edge
		for xcoord, ycoord := myCoords[0]+1, myCoords[1]-1; xcoord < 8 && ycoord >= 0; xcoord, ycoord = xcoord+1, ycoord-1 {
			// If this square is empty, it's a move and continue
			if b.squares[xcoord][ycoord] == empty {
				moves = append(moves, move{myCoords[0], myCoords[1], xcoord, ycoord, empty})
			} else {
				// If this square is an enemy piece, it's a move
				if colour == white && isBlack(b.squares[xcoord][ycoord]) ||
					colour == black && isWhite(b.squares[xcoord][ycoord]) {
					moves = append(moves, move{myCoords[0], myCoords[1], xcoord, ycoord, empty})
				}

				break
			}
		}
	}

	if canMoveTopLeft {
		// Move until find a piece or the edge
		for xcoord, ycoord := myCoords[0]-1, myCoords[1]+1; xcoord >= 0 && ycoord < 8; xcoord, ycoord = xcoord-1, ycoord+1 {
			// If this square is empty, it's a move and continue
			if b.squares[xcoord][ycoord] == empty {
				moves = append(moves, move{myCoords[0], myCoords[1], xcoord, ycoord, empty})
			} else {
				// If this square is an enemy piece, it's a move
				if colour == white && isBlack(b.squares[xcoord][ycoord]) ||
					colour == black && isWhite(b.squares[xcoord][ycoord]) {
					moves = append(moves, move{myCoords[0], myCoords[1], xcoord, ycoord, empty})
				}

				break
			}
		}
	}

	if canMoveTopRight {
		// Move until find a piece or the edge
		for xcoord, ycoord := myCoords[0]+1, myCoords[1]+1; xcoord < 8 && ycoord < 8; xcoord, ycoord = xcoord+1, ycoord+1 {
			// If this square is empty, it's a move and continue
			if b.squares[xcoord][ycoord] == empty {
				moves = append(moves, move{myCoords[0], myCoords[1], xcoord, ycoord, empty})
			} else {
				// If this square is an enemy piece, it's a move
				if colour == white && isBlack(b.squares[xcoord][ycoord]) ||
					colour == black && isWhite(b.squares[xcoord][ycoord]) {
					moves = append(moves, move{myCoords[0], myCoords[1], xcoord, ycoord, empty})
				}

				break
			}
		}
	}

	return moves
}

func generateKingMoves(b *board, kingCoords [2]int, colour int, moves []move) []move {
	xcoord := kingCoords[0]
	ycoord := kingCoords[1]

	possibleMoves := [8]move{
		move{xcoord, ycoord, xcoord - 1, ycoord, empty},
		move{xcoord, ycoord, xcoord - 1, ycoord + 1, empty},
		move{xcoord, ycoord, xcoord, ycoord + 1, empty},
		move{xcoord, ycoord, xcoord + 1, ycoord + 1, empty},
		move{xcoord, ycoord, xcoord + 1, ycoord, empty},
		move{xcoord, ycoord, xcoord + 1, ycoord - 1, empty},
		move{xcoord, ycoord, xcoord, ycoord - 1, empty},
		move{xcoord, ycoord, xcoord - 1, ycoord - 1, empty},
	}

	for _, move := range possibleMoves {
		// First, check bounds
		if move.toX < 0 || move.toX > 7 || move.toY < 0 || move.toY > 7 {
			continue
		}

		// Next, check whether one of my pieces is there
		if colour == white && isWhite(b.squares[move.toX][move.toY]) {
			continue
		}
		if colour == black && isBlack(b.squares[move.toX][move.toY]) {
			continue
		}

		// Next, check for rook/queen on row/col
		if isEnemyInDirection(b, [2]int{move.toX, move.toY}, [2]int{-1, -1}, colour, left) || isEnemyInDirection(b, [2]int{move.toX, move.toY}, [2]int{-1, -1}, colour, right) ||
			isEnemyInDirection(b, [2]int{move.toX, move.toY}, [2]int{-1, -1}, colour, up) || isEnemyInDirection(b, [2]int{move.toX, move.toY}, [2]int{-1, -1}, colour, down) {
			continue
		}

		// Next, check for enemy in diagonal
		if isEnemyOnDiagonal(b, [2]int{move.toX, move.toY}, [2]int{-1, -1}, colour, bottomLeftDiagonal) || isEnemyOnDiagonal(b, [2]int{move.toX, move.toY}, [2]int{-1, -1}, colour, topLeftDiagonal) ||
			isEnemyOnDiagonal(b, [2]int{move.toX, move.toY}, [2]int{-1, -1}, colour, topRightDiagonal) || isEnemyOnDiagonal(b, [2]int{move.toX, move.toY}, [2]int{-1, -1}, colour, bottomRightDiagonal) {
			continue
		}

		// Next, check for knights
		knightFound := false
		if move.toX-2 >= 0 {
			if move.toY-1 >= 0 {
				if (colour == white && b.squares[move.toX-2][move.toY-1] == blackKnight) || (colour == black && b.squares[move.toX-2][move.toY-1] == whiteKnight) {
					knightFound = true
					goto KnightEnd
				}
			}
			if move.toY+1 < 8 {
				if (colour == white && b.squares[move.toX-2][move.toY+1] == blackKnight) || (colour == black && b.squares[move.toX-2][move.toY+1] == whiteKnight) {
					knightFound = true
					goto KnightEnd
				}
			}
		}

		if move.toX-1 >= 0 {
			if move.toY-2 >= 0 {
				if (colour == white && b.squares[move.toX-1][move.toY-2] == blackKnight) || (colour == black && b.squares[move.toX-1][move.toY-2] == whiteKnight) {
					knightFound = true
					goto KnightEnd
				}
			}
			if move.toY+2 < 8 {
				if (colour == white && b.squares[move.toX-1][move.toY+2] == blackKnight) || (colour == black && b.squares[move.toX-1][move.toY+2] == whiteKnight) {
					knightFound = true
					goto KnightEnd
				}
			}
		}

		if move.toX+2 < 8 {
			if move.toY-1 >= 0 {
				if (colour == white && b.squares[move.toX+2][move.toY-1] == blackKnight) || (colour == black && b.squares[move.toX+2][move.toY-1] == whiteKnight) {
					knightFound = true
					goto KnightEnd
				}
			}
			if move.toY+1 < 8 {
				if (colour == white && b.squares[move.toX+2][move.toY+1] == blackKnight) || (colour == black && b.squares[move.toX+2][move.toY+1] == whiteKnight) {
					knightFound = true
					goto KnightEnd
				}
			}
		}

		if move.toX+1 >= 0 {
			if move.toY-2 >= 0 {
				if (colour == white && b.squares[move.toX+1][move.toY-2] == blackKnight) || (colour == black && b.squares[move.toX+1][move.toY-2] == whiteKnight) {
					knightFound = true
					goto KnightEnd
				}
			}
			if move.toY+2 < 8 {
				if (colour == white && b.squares[move.toX+1][move.toY+2] == blackKnight) || (colour == black && b.squares[move.toX+1][move.toY+2] == whiteKnight) {
					knightFound = true
					goto KnightEnd
				}
			}
		}

	KnightEnd:
		if knightFound {
			continue
		}

		// Now pawns...
		pawnFound := false
		if colour == white {
			if (move.toX-1 > 0 && move.toY < 8 && b.squares[move.toX-1][move.toY+1] == blackPawn) || (move.toX+1 < 8 && move.toY < 8 && b.squares[move.toX+1][move.toY+1] == blackPawn) {
				pawnFound = true
			}
		} else {
			if (move.toX-1 >= 0 && move.toY-1 >= 0 && b.squares[move.toX-1][move.toY-1] == whitePawn) || (move.toX+1 >= 0 && move.toY-1 >= 0 && b.squares[move.toX+1][move.toY-1] == whitePawn) {
				pawnFound = true
			}
		}

		if pawnFound {
			continue
		}

		// Finally, the other king
		kingFound := false
		if move.toX-1 >= 0 {
			if colour == white && b.squares[move.toX-1][move.toY] == blackKing || colour == black && b.squares[move.toX-1][move.toY] == whiteKing {
				kingFound = true
				goto KingFound
			}

			if move.toY-1 >= 0 {
				if colour == white && b.squares[move.toX-1][move.toY-1] == blackKing || colour == black && b.squares[move.toX-1][move.toY-1] == whiteKing {
					kingFound = true
					goto KingFound
				}
			}

			if move.toY+1 < 8 {
				if colour == white && b.squares[move.toX-1][move.toY+1] == blackKing || colour == black && b.squares[move.toX-1][move.toY+1] == whiteKing {
					kingFound = true
					goto KingFound
				}
			}
		}

		if move.toX+1 < 8 {
			if colour == white && b.squares[move.toX+1][move.toY] == blackKing || colour == black && b.squares[move.toX+1][move.toY] == whiteKing {
				kingFound = true
				goto KingFound
			}

			if move.toY-1 >= 0 {
				if colour == white && b.squares[move.toX+1][move.toY-1] == blackKing || colour == black && b.squares[move.toX+1][move.toY-1] == whiteKing {
					kingFound = true
					goto KingFound
				}
			}

			if move.toY+1 < 8 {
				if colour == white && b.squares[move.toX+1][move.toY+1] == blackKing || colour == black && b.squares[move.toX+1][move.toY+1] == whiteKing {
					kingFound = true
					goto KingFound
				}
			}
		}

		if move.toY-1 >= 0 {
			if colour == white && b.squares[move.toX][move.toY-1] == blackKing || colour == black && b.squares[move.toX][move.toY-1] == whiteKing {
				kingFound = true
				goto KingFound
			}
		}

		if move.toY+1 < 8 {
			if colour == white && b.squares[move.toX][move.toY+1] == blackKing || colour == black && b.squares[move.toX][move.toY+1] == whiteKing {
				kingFound = true
				goto KingFound
			}
		}

	KingFound:
		if kingFound {
			continue
		}
		moves = append(moves, move)
	}

	return moves
}

/* This function will return a list of legal moves for both sides (that is, moves which will not
result in check). Assumptions:
- Not currently in check
- Kings exist
*/
func generateAllLegalMoves(b *board, colour int) []move {
	ret := make([]move, 0, 128)

	// Store location of both kings for reference
	whiteKingCoords := [2]int{0, 0}
	blackKingCoords := [2]int{0, 0}
	whiteKingFound := false
	blackKingFound := false
	for x, col := range b.squares {
		for y, sq := range col {
			if sq == whiteKing {
				whiteKingCoords[0] = x
				whiteKingCoords[1] = y
			} else if sq == blackKing {
				blackKingCoords[0] = x
				blackKingCoords[1] = y
			}
			if whiteKingFound && blackKingFound {
				break
			}
		}
		if whiteKingFound && blackKingFound {
			break
		}
	}

	// Iterate over all squares on board
	for x, col := range b.squares {
		for y, sq := range col {
			currentCoords := [2]int{x, y}

			// Is there a piece?
			switch square := sq; square {
			case empty:
				continue
			case whiteRook:
				if colour == white {
					ret = generateRookMoves(b, currentCoords, whiteKingCoords, white, ret)
				}
			case whiteKnight:
				if colour == white {
					ret = generateKnightMoves(b, currentCoords, whiteKingCoords, white, ret)
				}
			case whiteBishop:
				if colour == white {
					ret = generateBishopMoves(b, currentCoords, whiteKingCoords, white, ret)
				}
			case whiteQueen:
				if colour == white {
					ret = generateQueenMoves(b, currentCoords, whiteKingCoords, white, ret)
				}
			case whiteKing:
				if colour == white {
					ret = generateKingMoves(b, whiteKingCoords, white, ret)
				}
			case whitePawn:
				if colour == white {
					ret = generatePawnMoves(b, currentCoords, blackKingCoords, white, ret)
				}
			case blackPawn:
				if colour == black {
					ret = generatePawnMoves(b, currentCoords, blackKingCoords, black, ret)
				}
			case blackRook:
				if colour == black {
					ret = generateRookMoves(b, currentCoords, blackKingCoords, black, ret)
				}
			case blackKnight:
				if colour == black {
					ret = generateKnightMoves(b, currentCoords, blackKingCoords, black, ret)
				}
			case blackBishop:
				if colour == black {
					ret = generateBishopMoves(b, currentCoords, blackKingCoords, black, ret)
				}
			case blackKing:
				if colour == black {
					ret = generateKingMoves(b, blackKingCoords, black, ret)
				}
			case blackQueen:
				if colour == black {
					ret = generateQueenMoves(b, currentCoords, blackKingCoords, black, ret)
				}
			}
		}
	}

	// Is there a piece?

	return ret
}
