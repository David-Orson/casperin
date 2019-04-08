package engine

import (
	"context"
	"fmt"
	"math/rand"
	"sync"

	. "github.com/David-Orson/casperin/backend"
)

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1
const ValueWin = Mate - 150
const SMPCycles = 16

var SkipSize = []int{1, 1, 1, 2, 2, 2, 1, 3, 2, 2, 1, 3, 3, 2, 2, 1}
var SkipDepths = []int{1, 2, 2, 4, 4, 3, 2, 5, 4, 3, 2, 6, 5, 4, 3, 2}

func lossIn(height int) int {
	return -Mate + height
}

func depthToMate(val int) int {
	if val >= ValueWin {
		return Mate - val
	}
	return val - Mate
}

func (t *thread) EvaluateMoves(moves []EvaledMove, fromTrans Move, height int) {
	var pos *Position = &t.stack[height].position
	var counter Move
	if pos.LastMove != NullMove {
		counter = t.CounterMoves[pos.LastMove.From()][pos.LastMove.To()]
	}
	for i := range moves {
		if moves[i].Move == fromTrans {
			moves[i].Value = 100000
		} else if moves[i].Move.IsPromotion() {
			moves[i].Value = 70000 + SEEValues[moves[i].Move.CapturedPiece()]
		} else if moves[i].Move.IsCapture() {
			moves[i].Value = SEEValues[moves[i].Move.CapturedPiece()]*8 - SEEValues[moves[i].Move.MovedPiece()] + 10000
		} else {
			if moves[i].Move == t.KillerMoves[height][0] {
				moves[i].Value = 9000
			} else if moves[i].Move == t.KillerMoves[height][1] {
				moves[i].Value = 8000
			} else if moves[i].Move == counter {
				moves[i].Value = 7999
			} else {
				moves[i].Value = t.EvalHistory[moves[i].Move.From()][moves[i].Move.To()]
			}
		}
	}
}

func (t *thread) EvaluateQsMoves(moves []EvaledMove) {
	for i := range moves {
		moves[i].Value = PieceValues[moves[i].Move.CapturedPiece()] - PieceValues[moves[i].Move.MovedPiece()]
	}
}

func (t *thread) quiescence(alpha, beta, height int) int {
	t.incNodes()
	t.stack[height].PV.clear()
	pos := &t.stack[height].position

	if height >= MAX_HEIGHT {
		return contempt(pos)
	}

	child := &t.stack[height+1].position
	moveCount := 0

	val := Evaluate(pos)

	if val >= beta {
		return beta
	}
	if alpha < val {
		alpha = val
	}

	evaled := pos.GenerateAllCaptures(t.stack[height].moves[:])
	t.EvaluateQsMoves(evaled)

	for i := range evaled {
		maxMoveToFirst(evaled[i:])
		if !pos.MakeMove(evaled[i].Move, child) {
			continue
		}
		moveCount++
		val = -t.quiescence(-beta, -alpha, height+1)
		if val > alpha {
			alpha = val
			if val >= beta {
				return beta
			}
			t.stack[height].PV.assign(evaled[i].Move, &t.stack[height+1].PV)
		}
	}

	if moveCount == 0 {
		return val
	}

	return alpha
}

func contempt(pos *Position) int {
	return 0
}

func maxMoveToFirst(moves []EvaledMove) {
	maxIdx := 0
	for i := 1; i < len(moves); i++ {
		if moves[i].Value > moves[maxIdx].Value {
			maxIdx = i
		}
	}
	moves[0], moves[maxIdx] = moves[maxIdx], moves[0]
}

func (t *thread) alphaBeta(depth, alpha, beta, height int) int {
	t.incNodes()
	t.stack[height].PV.clear()

	var pos = &t.stack[height].position

	if t.isDraw(height) {
		return contempt(pos)
	}

	var tmpVal int

	alphaOrig := alpha
	hashOk, hashValue, hashDepth, hashMove, hashFlag := t.engine.TransTable.Get(pos.Key, height)
	if hashOk {
		tmpVal = int(hashValue)
		if hashDepth >= uint8(depth) {
			if hashFlag == TransExact {
				if !hashMove.IsCapture() {
					t.EvalHistory[uint(hashMove.From())][uint(hashMove.To())] += depth
				}
				return tmpVal
			}
			if hashFlag == TransAlpha && tmpVal <= alpha {
				return alpha
			}
			if hashFlag == TransBeta && tmpVal >= beta {
				return beta
			}
		}
	}

	inCheck := pos.IsInCheck()

	// check extension
	if depth == 0 {
		if !inCheck {
			return t.quiescence(alpha, beta, height)
		}
		depth = 1
	}

	var child = &t.stack[height+1].position

	if pos.LastMove != NullMove && depth >= 4 && !inCheck && !isLateEndGame(pos) {
		pos.MakeNullMove(child)
		tmpVal = -t.alphaBeta(depth-3, -beta, -beta+1, height+1)
		if tmpVal >= beta {
			return beta
		}
	}

	val := MinInt

	evaled := pos.GenerateAllMoves(t.stack[height].moves[:])
	t.EvaluateMoves(evaled, hashMove, height)
	bestMove := NullMove
	moveCount := 0
	for i := range evaled {
		maxMoveToFirst(evaled[i:])
		if !pos.MakeMove(evaled[i].Move, child) {
			continue
		}
		tmpVal = -t.alphaBeta(depth-1, -beta, -alpha, height+1)
		moveCount++

		if tmpVal > val {
			val = tmpVal
			bestMove = evaled[i].Move
			if val > alpha {
				alpha = val

				if !evaled[i].Move.IsCapture() {
					t.EvalHistory[uint(evaled[i].Move.From())][uint(evaled[i].Move.To())] += depth
				}

				if alpha >= beta {
					if !evaled[i].Move.IsCapture() && pos.LastMove != NullMove {
						t.KillerMoves[height][0], t.KillerMoves[height][1] = evaled[i].Move, t.KillerMoves[height][0]
						t.CounterMoves[pos.LastMove.From()][pos.LastMove.To()] = evaled[i].Move
					}
					t.engine.TransTable.Set(pos.Key, beta, depth, evaled[i].Move, TransBeta, height)
					return beta
				}
				t.stack[height].PV.assign(evaled[i].Move, &t.stack[height+1].PV)
			}
		}
	}

	if moveCount == 0 {
		if inCheck {
			return lossIn(height)
		}
		return contempt(pos)
	}

	if alpha == alphaOrig {
		t.engine.TransTable.Set(pos.Key, alpha, depth, bestMove, TransAlpha, height)
	} else {
		t.engine.TransTable.Set(pos.Key, alpha, depth, bestMove, TransExact, height)
	}
	return alpha
}

func (t *thread) isDraw(height int) bool {
	var pos *Position = &t.stack[height].position
	if pos.FiftyMove > 100 {
		return true
	}

	if (pos.Pawns|pos.Rooks|pos.Queens) == 0 && !MoreThanOne(pos.Knights|pos.Bishops) {
		return true
	}

	for i := height - 1; i >= 0; i-- {
		descendant := &t.stack[i].position
		if descendant.Key == pos.Key {
			return true
		}
		if descendant.FiftyMove == 0 || descendant.LastMove == NullMove {
			return false
		}
	}

	if t.engine.MoveHistory[pos.Key] >= 2 {
		return true
	}

	return false
}

type result struct {
	Move
	value int
	depth int
	moves []Move
}

func (t *thread) depSearch(depth int, moves []EvaledMove, resultChan chan result, mainThread bool) {
	var pos *Position = &t.stack[0].position
	var child *Position = &t.stack[1].position
	var bestMove Move = NullMove
	if mainThread {
		hashOk, _, _, hashMove, _ := t.engine.TransTable.Get(pos.Key, 0)
		lastBestMove := NullMove
		if hashOk {
			lastBestMove = hashMove
		}
		t.EvaluateMoves(moves, lastBestMove, 0)
	}
	alpha := -MaxInt
	moveCount := 0
	t.stack[0].PV.clear()
	bestMoveIdx := -1
	for i := range moves {
		if mainThread {
			maxMoveToFirst(moves[i:])
		}
		if !pos.MakeMove(moves[i].Move, child) {
			continue
		}
		moveCount++
		val := -t.alphaBeta(depth-1, -MaxInt, -alpha, 1)
		if val > alpha {
			bestMoveIdx = i
			alpha = val
			bestMove = moves[i].Move
			t.stack[0].PV.assign(moves[i].Move, &t.stack[1].PV)
		}
	}
	if moveCount == 0 {
		if pos.IsInCheck() {
			alpha = lossIn(0)
		} else {
			alpha = contempt(pos)
		}
	}
	if !mainThread && bestMoveIdx != -1 {
		moveToFirst(moves, bestMoveIdx)
	}
	fmt.Println(depth, mainThread)
	t.engine.TransTable.Set(pos.Key, alpha, depth, bestMove, TransExact, 0)
	resultChan <- result{bestMove, alpha, depth, cloneMoves(t.stack[0].PV.items[:t.stack[0].PV.size])}
}

func moveToFirst(moves []EvaledMove, idx int) {
	if idx == 0 {
		return
	}
	move := moves[idx]
	for i := idx; idx > 0; idx-- {
		moves[i] = moves[i-1]
	}
	moves[0] = move
}

func (e *Engine) singleThreadBestMove(ctx context.Context, pos *Position) Move {
	var lastBestMove Move
	thread := e.threads[0]
	thread.stack[0].position = *pos
	moves := pos.GenerateAllMoves(e.threads[0].stack[0].moves[:])
	for i := 1; ; i++ {
		resultChan := make(chan result, 1)
		go func(d int) {
			defer recoverFromTimeout()
			thread.depSearch(d, moves, resultChan, true)
		}(i)
		select {
		case <-ctx.Done():
			return lastBestMove
		case res := <-resultChan:
			if i >= 3 {
				e.callUpdate(SearchInfo{res.value, i, thread.nodes, res.moves})
			}
			if res.value >= ValueWin && depthToMate(res.value) <= i {
				return res.Move
			}
			if res.Move == 0 {
				return lastBestMove
			}
			if i >= MAX_HEIGHT {
				return res.Move
			}
			if e.isSoftTimeout(i, thread.nodes) {
				return res.Move
			}
			lastBestMove = res.Move
		}
	}
}

func (t *thread) iterativeDeepening(resultChan chan result, idx int) {
	mainThread := idx == 0
	moves := t.stack[0].position.GenerateAllMoves(t.stack[0].moves[:])
	if !mainThread {
		rand.Shuffle(len(moves), func(i, j int) {
			moves[i], moves[j] = moves[j], moves[i]
		})
	}
	cycle := idx % SMPCycles
	for depth := 1; depth < MAX_HEIGHT; depth++ {
		t.depSearch(depth, moves, resultChan, mainThread)
		if !mainThread && (depth+cycle)%SkipDepths[cycle] == 0 {
			depth += SkipSize[cycle]
		}
	}
}

func (e *Engine) bestMove(ctx context.Context, pos *Position) Move {
	if e.Threads.Val == 1 {
		return e.singleThreadBestMove(ctx, pos)
	}
	for i := range e.threads {
		e.threads[i].stack[0].position = *pos
	}

	var wg = &sync.WaitGroup{}
	resultChan := make(chan result)
	for i := range e.threads {
		wg.Add(1)
		go func(idx int) {
			defer recoverFromTimeout()
			e.threads[idx].iterativeDeepening(resultChan, idx)
			wg.Done()
		}(i)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	prevDepth := 0
	var lastBestMove Move
	for {
		select {
		case <-e.done:
			return lastBestMove
		case res := <-resultChan:
			if res.depth <= prevDepth {
				continue
			}
			nodes := e.nodes()
			e.callUpdate(SearchInfo{res.value, res.depth, nodes, res.moves})
			if res.value >= ValueWin && depthToMate(res.value) <= res.depth {
				return res.Move
			}
			if res.Move == 0 {
				return lastBestMove
			}
			if res.depth >= MAX_HEIGHT {
				return res.Move
			}
			if e.isSoftTimeout(res.depth, nodes) {
				return res.Move
			}
			lastBestMove = res.Move
			prevDepth = res.depth
		}
	}
}

func cloneMoves(src []Move) []Move {
	dst := make([]Move, len(src))
	copy(dst, src)
	return dst
}

func recoverFromTimeout() {
	err := recover()
	if err != nil && err != errTimeout {
		panic(err)
	}
}
