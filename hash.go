package main

import "math/rand"

type hashValue struct {
	score    int
	bestMove *move
}

// This is a pseudo-random array of numbers which contains a number for each piece at each square
var boardPieceNumbers [8][8][12]uint64
var boardFlagNumbers [numBoardFlags]uint64

// This is a global hash table which encodes the result of each searched position
var hashTable map[uint64]*hashValue

func initializeHash() {
	rand.Seed(30)

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			for k := 0; k < 12; k++ {
				boardPieceNumbers[i][j][k] = rand.Uint64()
			}
		}
	}

	for i := 0; i < numBoardFlags; i++ {
		boardFlagNumbers[i] = rand.Uint64()
	}
}

func hashPos(b *board) uint64 {
	var hash uint64

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if b.squares[i][j] == empty {
				continue
			}

			hash ^= boardPieceNumbers[i][j][b.squares[i][j]]
		}
	}

	// Hash board flags
	if hasFlag(b.flags, whiteRookKingSideMoved) {
		hash ^= boardFlagNumbers[0]
	}
	if hasFlag(b.flags, whiteRookQueenSideMoved) {
		hash ^= boardFlagNumbers[0]
	}
	if hasFlag(b.flags, whiteKingMoved) {
		hash ^= boardFlagNumbers[0]
	}
	if hasFlag(b.flags, blackRookKingSideMoved) {
		hash ^= boardFlagNumbers[0]
	}
	if hasFlag(b.flags, blackRookQueenSideMoved) {
		hash ^= boardFlagNumbers[0]
	}
	if hasFlag(b.flags, blackKingMoved) {
		hash ^= boardFlagNumbers[0]
	}

	return hash
}

// Will overwrite any existing entry at this hash
func addHashTableEntry(b *board, score int, bestMove *move) {
	hash := hashPos(b)
	hashTable[hash] = &hashValue{score: score, bestMove: bestMove}
}

func getHashTableEntry(b *board) *hashValue {
	hash := hashPos(b)
	if val, exists := hashTable[hash]; exists {
		return val
	}

	return nil
}
