# Casperin

Casperin is an open-source chess engine compliant with the UCI (Universal Chess Interface) standard, utilizing the [Alpha-beta pruning algorithm](https://en.wikipedia.org/wiki/Alpha%E2%80%93beta_pruning).

## UCI Configuration Options

### Hash
Defines the size of the transposition table in megabytes. Increasing this value can enhance performance.

### Threads
Specifies the number of threads to use during search operations. More threads generally lead to better performance.

### PawnHash
Determines the size of the Pawn Hash Table. The default value is usually sufficient due to the high hit-ratio in this table.

### Move Overhead
Sets a time buffer in milliseconds. Increase this value if you observe time losses during the game.

### Ponder
Activates the "thinking on opponent's time" mode, slightly adjusting the time management algorithm to use more time.

## Command-Line Interface (CLI) Options

### `casperin bench`
Runs a benchmark test to evaluate the engine's performance.

### `casperin tune`
Executes a tuning process that combines coordinate descent and gradient descent methods, where the gradient is calculated using symmetric derivatives.

### `casperin trace-tune`
Performs tuning based on gradient descent, using vectors that track the usage of evaluation constants in given positions. This requires compiling with the `tuning` constant set to `true` in the `evaluation/eval.go` file.

For tuning, games should be placed in the `games.fen` file.

The tuner is inspired by the [tuning methodology](https://github.com/AndyGrant/Ethereal/blob/master/Tuning.pdf) described by Andrew Grant.

## Static Exchange Evaluation

- Influenced by [Stockfish](https://github.com/official-stockfish/Stockfish/), developed by Tord Romstad, Marco Costalba, Joona Kiiski, and Gary Linscott.
- Incorporates various ideas and concepts from Stockfish in both the search procedure and evaluation.

## License

Casperin is licensed under the GNU General Public License version 3 (GPL v3).
