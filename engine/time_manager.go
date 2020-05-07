package engine

import "time"
import . "github.com/David-Orson/casperin/backend"
import . "github.com/David-Orson/casperin/utils"

type timeManager interface {
	hardTimeout() time.Duration
	isSoftTimeout(depth, nodes int) bool
	updateTime(depth, score int)
	getStartedAt() time.Time
}

type depthMoveTimeManager struct {
	startedAt time.Time
	duration  int
	depth     int
}

func (manager *depthMoveTimeManager) hardTimeout() time.Duration {
	if manager.duration > 0 {
		return time.Duration(manager.duration) * time.Millisecond
	}
	return 0
}

func (manager *depthMoveTimeManager) isSoftTimeout(depth, nodes int) bool {
	return manager.depth > 0 && depth >= manager.depth
}

func (manager *depthMoveTimeManager) updateTime(int, int) {
}

func (manager *depthMoveTimeManager) getStartedAt() time.Time {
	return manager.startedAt
}

type tournamentTimeManager struct {
	startedAt time.Time
	hard      time.Duration
	ideal     time.Duration
	lastScore int
}

func (manager *tournamentTimeManager) getStartedAt() time.Time {
	return manager.startedAt
}

func (manager *tournamentTimeManager) hardTimeout() time.Duration {
	return manager.hard
}

func (manager *tournamentTimeManager) isSoftTimeout(int, int) bool {
	return time.Since(manager.startedAt) >= manager.ideal
}

func (manager *tournamentTimeManager) updateTime(depth, score int) {
	lastScore := manager.lastScore
	manager.lastScore = score
	if depth < 4 {
		return
	}

	if lastScore > score+16 {
		manager.ideal += manager.ideal / 20
	}
	if lastScore > score+21 {
		manager.ideal += manager.ideal / 20
	}
	if lastScore > score+61 {
		manager.ideal += manager.ideal / 20
	}

	if lastScore+23 < score {
		manager.ideal += manager.ideal / 40
	}
	if lastScore+46 < score {
		manager.ideal += manager.ideal / 20
	}
}

func newTournamentTimeManager(startedAt time.Time, limits LimitsType, overhead, sideToMove int) *tournamentTimeManager {
	res := &tournamentTimeManager{startedAt: startedAt}
	var limit, inc int
	if sideToMove == White {
		limit, inc = limits.WhiteTime, limits.WhiteIncrement
	} else {
		limit, inc = limits.BlackTime, limits.BlackIncrement
	}
	movesToGo := limits.MovesToGo
	var ideal, hard int

	if movesToGo > 0 {
		ideal = (((limit / (movesToGo + 5)) + inc) * 3) / 4
		hard = ((limit / (movesToGo + 7)) + inc) * 4
	} else {
		ideal = (limit + 25*inc) / 50
		hard = 5 * (limit + 25*inc) / 50
	}
	res.ideal = time.Duration(Min(ideal, limit-overhead)) * time.Millisecond
	res.hard = time.Duration(Min(hard, limit-overhead)) * time.Millisecond
	return res
}

func newTimeManager(limits LimitsType, overhead int, sideToMove int) timeManager {
	startedAt := time.Now()
	if limits.WhiteTime > 0 || limits.BlackTime > 0 {
		return newTournamentTimeManager(startedAt, limits, overhead, sideToMove)
	} else if limits.MoveTime > 0 {
		return &depthMoveTimeManager{startedAt: startedAt, duration: limits.MoveTime}
	} else if limits.Depth > 0 {
		return &depthMoveTimeManager{startedAt: startedAt, depth: limits.Depth}
	} else if limits.Infinite {
		return &depthMoveTimeManager{startedAt: startedAt}
	} else {
		return &depthMoveTimeManager{startedAt: startedAt, duration: 1000}
	}
}
