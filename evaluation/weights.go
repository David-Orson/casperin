package evaluation

import (
	. "github.com/David-Orson/casperin/backend"
)

var PawnValue = S(85, 153)
var KnightValue = S(416, 546)
var BishopValue = S(389, 539)
var RookValue = S(561, 906)
var QueenValue = S(1252, 1734)

// Pawns Square scores
// Bishop Flag, Rank, Row
var PawnScores = [16][7][8]Score{
	{
		{},
		{S(-18, 16), S(-12, 5), S(-24, 12), S(-9, 13), S(-25, 22), S(15, 13), S(33, 6), S(-13, 11)},
		{S(-15, 7), S(-18, 5), S(-19, 6), S(-8, 0), S(-4, -4), S(-2, 4), S(6, 5), S(-19, 11)},
		{S(-15, 17), S(-21, 15), S(0, 0), S(6, -8), S(16, -16), S(16, -4), S(0, 13), S(-15, 18)},
		{S(-8, 29), S(-5, 17), S(-4, 3), S(4, -10), S(29, -17), S(31, -3), S(30, 10), S(0, 23)},
		{S(0, 51), S(9, 33), S(19, 23), S(45, -9), S(62, -17), S(63, 12), S(63, 28), S(2, 54)},
		{S(6, -4), S(-27, 26), S(-4, 21), S(12, 0), S(-20, -1), S(-14, 20), S(-127, 59), S(-149, 53)},
	},
	{
		{},
		{S(-12, 27), S(-7, 6), S(-21, 10), S(8, 18), S(-19, 17), S(27, 15), S(28, -2), S(-11, 13)},
		{S(-8, 12), S(-17, 3), S(-2, 7), S(-2, -4), S(10, 10), S(-4, -3), S(15, 11), S(-24, 11)},
		{S(-4, 21), S(-6, 12), S(7, -4), S(33, -4), S(26, -17), S(28, 2), S(0, 11), S(-12, 13)},
		{S(6, 34), S(15, 3), S(20, 7), S(19, -21), S(44, -12), S(47, -13), S(28, 14), S(-11, 31)},
		{S(-12, 82), S(34, 27), S(23, 27), S(67, -15), S(60, -13), S(102, -10), S(39, 23), S(-4, 52)},
		{S(-17, 67), S(11, 4), S(22, 10), S(21, -9), S(51, -5), S(19, -22), S(-85, 38), S(-111, 30)},
	},
	{
		{},
		{S(-18, 17), S(-13, 15), S(-16, 23), S(-20, 4), S(5, 29), S(8, 6), S(51, 4), S(-8, 20)},
		{S(-21, 12), S(-22, 12), S(-13, 0), S(-1, 10), S(4, -11), S(4, 16), S(-1, -3), S(-16, 19)},
		{S(-11, 7), S(-8, 12), S(10, 8), S(19, -12), S(42, -9), S(24, -7), S(8, 12), S(-12, 22)},
		{S(-7, 24), S(3, 16), S(2, 3), S(26, -5), S(37, -15), S(60, -2), S(22, 9), S(-3, 34)},
		{S(1, 45), S(10, 38), S(62, -5), S(45, -10), S(95, -25), S(71, 17), S(27, 25), S(7, 67)},
		{S(-25, 17), S(-14, 40), S(12, -19), S(23, 6), S(30, -30), S(29, -18), S(-90, 27), S(-111, 35)},
	},
	{
		{},
		{S(-19, 6), S(-13, -5), S(-4, -7), S(-1, -15), S(6, 1), S(9, 0), S(44, -8), S(-2, -16)},
		{S(-12, -4), S(-29, -6), S(-2, -7), S(-2, -12), S(9, -4), S(-8, -2), S(6, -11), S(-10, -21)},
		{S(-2, -10), S(-6, -12), S(12, -23), S(36, -18), S(33, -14), S(22, -7), S(22, -23), S(0, -22)},
		{S(-17, 17), S(0, -5), S(33, -41), S(19, -16), S(42, -41), S(45, -24), S(17, 6), S(-6, -11)},
		{S(-25, 71), S(12, 17), S(35, -18), S(27, -3), S(80, -28), S(48, 2), S(41, 12), S(-16, 32)},
		{S(-11, 27), S(-19, 25), S(-16, 5), S(19, -31), S(45, -6), S(38, -8), S(-90, 31), S(-122, 62)},
	},
	{
		{},
		{S(-23, 20), S(-11, 18), S(-7, 8), S(-3, 15), S(-23, 21), S(6, 17), S(35, 8), S(-11, 15)},
		{S(-18, 9), S(-21, 6), S(-9, -14), S(3, -6), S(-11, -10), S(-5, 11), S(-2, 2), S(-19, 18)},
		{S(-18, 27), S(-2, 11), S(12, -4), S(14, -29), S(33, -27), S(10, -9), S(6, 18), S(-17, 29)},
		{S(-5, 48), S(11, 18), S(5, -9), S(24, -13), S(33, -40), S(36, 8), S(28, 9), S(-7, 47)},
		{S(-23, 88), S(5, 42), S(8, 32), S(64, -23), S(71, -14), S(82, -7), S(57, 38), S(5, 73)},
		{S(30, 45), S(-20, 41), S(7, 24), S(42, 19), S(65, -31), S(48, 17), S(-77, 60), S(-121, 97)},
	},
	{
		{},
		{S(-2, 31), S(2, 19), S(-8, 21), S(-2, 20), S(-10, 24), S(16, 30), S(24, 9), S(-2, 15)},
		{S(-3, 19), S(-9, 5), S(1, 11), S(6, -5), S(5, 11), S(-1, 9), S(4, 19), S(-12, 15)},
		{S(-1, 34), S(9, 18), S(19, -1), S(31, -6), S(38, -13), S(24, 8), S(6, 23), S(-10, 24)},
		{S(17, 46), S(20, 14), S(11, 18), S(34, -14), S(50, -20), S(48, -6), S(37, 27), S(-6, 41)},
		{S(-4, 98), S(33, 45), S(39, 35), S(68, -16), S(78, -19), S(84, 7), S(30, 43), S(-2, 71)},
		{S(20, 71), S(-28, 16), S(17, 39), S(51, 18), S(27, -5), S(49, 24), S(-61, 26), S(-131, 45)},
	},
	{
		{},
		{S(-8, -16), S(0, -9), S(-5, -6), S(1, -25), S(1, 5), S(3, -18), S(46, 0), S(-2, -8)},
		{S(-6, -18), S(-9, -10), S(-14, -29), S(-1, -10), S(-11, -32), S(4, -6), S(-13, -22), S(-9, -1)},
		{S(-7, -8), S(10, -31), S(18, -16), S(13, -55), S(37, -29), S(8, -29), S(14, -4), S(-10, 7)},
		{S(7, -5), S(8, 2), S(9, -39), S(20, -15), S(33, -62), S(73, -25), S(27, -25), S(17, 10)},
		{S(-11, 25), S(42, -5), S(19, -4), S(43, -32), S(74, -50), S(91, -49), S(76, -8), S(45, 10)},
		{S(60, -73), S(2, 25), S(8, -78), S(27, -8), S(60, -99), S(26, 0), S(-60, -107), S(-118, 31)},
	},
	{
		{},
		{S(-6, 3), S(-4, 14), S(-5, 8), S(-4, 4), S(2, 20), S(5, 16), S(32, 15), S(-2, 12)},
		{S(-6, 0), S(-18, 6), S(-5, 0), S(0, 5), S(2, -2), S(-11, 29), S(-2, 0), S(-10, -4)},
		{S(-3, 5), S(-2, 5), S(20, -10), S(30, -21), S(34, -5), S(12, 11), S(5, 11), S(-12, 15)},
		{S(9, 26), S(16, 17), S(8, 7), S(26, -5), S(41, -29), S(53, -4), S(15, 11), S(-13, 21)},
		{S(-6, 56), S(26, 29), S(16, 21), S(45, 11), S(85, -32), S(88, 15), S(20, -9), S(-3, 74)},
		{S(9, -4), S(-30, 14), S(-16, -5), S(49, -7), S(-9, -49), S(31, -27), S(-70, 36), S(-133, 27)},
	},
	{
		{},
		{S(-16, 23), S(-9, 8), S(-17, 14), S(-9, 6), S(-15, 12), S(12, 10), S(23, 7), S(-11, 17)},
		{S(-13, 11), S(-17, 2), S(-10, -1), S(-5, -14), S(-1, -7), S(-15, 2), S(6, 9), S(-17, 13)},
		{S(-9, 27), S(-6, 21), S(1, -12), S(19, -23), S(22, -33), S(12, -1), S(-5, 15), S(-15, 23)},
		{S(2, 55), S(10, 11), S(-1, 4), S(7, -31), S(37, -13), S(23, -19), S(25, 19), S(-17, 40)},
		{S(17, 72), S(11, 44), S(64, 5), S(44, -17), S(69, -22), S(86, 13), S(98, 22), S(-8, 60)},
		{S(-26, 55), S(-8, 37), S(36, 21), S(45, -18), S(42, 20), S(39, -15), S(-88, 24), S(-123, 58)},
	},
	{
		{},
		{S(4, 0), S(3, -18), S(-10, -7), S(25, -2), S(-10, -6), S(36, -1), S(32, -24), S(5, -12)},
		{S(1, -8), S(-5, -27), S(14, -17), S(4, -31), S(19, -5), S(-6, -19), S(23, -5), S(-5, -22)},
		{S(5, 9), S(2, -12), S(17, -45), S(41, -22), S(33, -55), S(41, -16), S(11, -18), S(1, -7)},
		{S(22, 16), S(27, -31), S(20, -19), S(17, -65), S(56, -36), S(40, -48), S(35, 1), S(4, -3)},
		{S(42, 33), S(17, 10), S(57, -20), S(69, -34), S(81, -50), S(128, -41), S(64, -17), S(24, 11)},
		{S(-1, 25), S(27, -67), S(11, 43), S(46, -76), S(22, -28), S(47, -112), S(-98, 5), S(-131, -63)},
	},
	{
		{},
		{S(-8, 28), S(-1, 20), S(-9, 32), S(-9, 22), S(-16, 35), S(11, 14), S(34, 17), S(-2, 12)},
		{S(-10, 23), S(-6, 15), S(-7, 15), S(4, 7), S(1, -6), S(-4, 18), S(8, 9), S(-11, 12)},
		{S(-2, 26), S(4, 25), S(17, 7), S(21, 0), S(33, -4), S(23, 3), S(14, 17), S(-9, 22)},
		{S(2, 49), S(23, 26), S(11, 15), S(25, -7), S(42, -17), S(55, -2), S(31, 12), S(-2, 39)},
		{S(11, 76), S(19, 67), S(46, 11), S(35, 8), S(93, -22), S(79, 28), S(94, 7), S(4, 65)},
		{S(-5, 43), S(-2, 62), S(-6, 10), S(31, 18), S(73, -12), S(54, -41), S(-87, 23), S(-145, 70)},
	},
	{
		{},
		{S(-9, 22), S(-1, 8), S(-7, 13), S(7, 18), S(-16, 28), S(17, 14), S(32, -3), S(-3, 0)},
		{S(-6, 7), S(-13, 8), S(-3, 20), S(3, -2), S(5, 11), S(-9, 4), S(12, 2), S(-13, 2)},
		{S(2, 22), S(8, 7), S(17, -10), S(31, 7), S(35, -21), S(29, 4), S(10, 2), S(-16, 11)},
		{S(16, 22), S(22, 24), S(16, 20), S(30, -45), S(40, -20), S(53, -24), S(17, 21), S(-20, 39)},
		{S(3, 78), S(26, 52), S(34, 7), S(44, -19), S(63, -13), S(88, -11), S(42, 36), S(0, 34)},
		{S(-15, -5), S(5, 23), S(9, 4), S(24, 0), S(35, -11), S(36, -4), S(-88, 33), S(-118, 37)},
	},
	{
		{},
		{S(-27, 25), S(-15, 21), S(-7, 2), S(-6, 6), S(-28, -10), S(0, 5), S(20, -7), S(0, -10)},
		{S(-22, 24), S(-26, 12), S(-26, 4), S(-5, -9), S(-1, -28), S(-3, -10), S(-1, -5), S(-7, -3)},
		{S(-14, 36), S(-12, 21), S(6, -23), S(8, -33), S(24, -57), S(21, -28), S(-10, 17), S(-5, 8)},
		{S(-7, 41), S(3, 32), S(-1, 12), S(21, -30), S(28, -23), S(40, -49), S(20, 17), S(5, 40)},
		{S(-15, 122), S(5, 67), S(10, 40), S(50, 12), S(64, -6), S(83, 2), S(67, 30), S(-10, 96)},
		{S(-9, 40), S(-10, 29), S(-7, 26), S(9, -13), S(22, -13), S(29, -32), S(-65, 33), S(-107, 65)},
	},
	{
		{},
		{S(-8, 33), S(-1, 15), S(-11, 13), S(-1, 10), S(-19, 12), S(13, 27), S(14, 2), S(-3, 1)},
		{S(-7, 22), S(-19, 4), S(1, 9), S(4, -28), S(4, 9), S(-2, -7), S(3, 8), S(-13, -4)},
		{S(-1, 36), S(0, 20), S(17, -21), S(29, -5), S(34, -40), S(24, 2), S(4, 2), S(-9, 5)},
		{S(16, 56), S(16, 11), S(16, 14), S(30, -49), S(48, -13), S(32, -30), S(37, 22), S(-7, 35)},
		{S(13, 87), S(30, 72), S(31, 37), S(41, 13), S(34, 27), S(86, -11), S(37, 30), S(6, 64)},
		{S(18, 78), S(-18, 36), S(1, 22), S(32, -16), S(32, -10), S(-1, -41), S(-111, 5), S(-110, 70)},
	},
	{
		{},
		{S(-21, 23), S(-3, 14), S(-9, 28), S(-2, 4), S(-26, 27), S(-4, 4), S(23, 22), S(-5, 15)},
		{S(-17, 10), S(-10, 11), S(-9, -20), S(1, -1), S(-12, -16), S(-7, 10), S(-2, -7), S(-12, 14)},
		{S(-14, 26), S(0, 17), S(16, 2), S(17, -36), S(34, -19), S(12, -14), S(9, 6), S(-14, 16)},
		{S(-6, 51), S(27, 29), S(3, -1), S(24, 0), S(33, -45), S(57, -20), S(26, -3), S(-4, 43)},
		{S(2, 109), S(18, 41), S(49, 32), S(26, 2), S(90, -16), S(64, 9), S(49, 29), S(10, 59)},
		{S(-11, 11), S(-10, 54), S(-34, 13), S(18, 12), S(16, -26), S(30, 3), S(-77, 24), S(-131, 56)},
	},
	{
		{},
		{S(-12, 25), S(-3, 31), S(-16, 31), S(-8, 15), S(-22, 18), S(0, 30), S(12, 19), S(-8, 14)},
		{S(-10, 20), S(-16, 20), S(-8, 19), S(-3, 8), S(-4, 12), S(-16, 37), S(-9, 17), S(-18, 13)},
		{S(-3, 29), S(0, 30), S(14, 11), S(20, 5), S(28, -10), S(13, 12), S(-1, 23), S(-21, 25)},
		{S(8, 58), S(23, 28), S(4, 39), S(26, -20), S(35, -10), S(43, -9), S(18, 27), S(-15, 40)},
		{S(6, 97), S(21, 70), S(18, 70), S(46, 4), S(61, -19), S(64, -3), S(29, 23), S(-9, 68)},
		{S(-72, 78), S(-63, 11), S(-24, 31), S(12, -6), S(21, -35), S(33, -2), S(-63, 22), S(-155, 29)},
	},
}

// Piece Square Values
var PieceScores = [King + 1][8][8]Score{
	{},
	{ // knight
		{S(-87, -83), S(-24, -46), S(-46, -32), S(-21, -6), S(-23, -13), S(-26, -34), S(-29, -32), S(-90, -42)},
		{S(-32, -32), S(-19, -19), S(-16, -37), S(-15, -13), S(-25, -15), S(-17, -33), S(-26, -18), S(-21, -22)},
		{S(-13, -46), S(-7, -20), S(-7, -12), S(-1, 8), S(0, 6), S(-16, -14), S(-4, -24), S(-23, -46)},
		{S(-11, -4), S(-7, 6), S(9, 21), S(-3, 30), S(11, 27), S(2, 22), S(24, -6), S(-12, 1)},
		{S(-2, 9), S(-3, 7), S(11, 32), S(12, 49), S(10, 41), S(29, 25), S(9, 7), S(9, 5)},
		{S(-53, 8), S(-27, 15), S(-2, 40), S(9, 39), S(29, 28), S(29, 27), S(-13, 4), S(-20, -4)},
		{S(1, -19), S(-13, -5), S(26, -17), S(34, 17), S(41, 8), S(63, -39), S(-22, -8), S(15, -33)},
		{S(-203, -51), S(-96, -10), S(-151, 25), S(-27, -7), S(8, -7), S(-130, 33), S(-81, -9), S(-193, -63)},
	},
	{ // Bishop
		{S(24, -43), S(44, -12), S(6, 1), S(14, 1), S(4, 9), S(16, -9), S(15, -11), S(22, -41)},
		{S(35, -33), S(15, -36), S(42, -9), S(13, 4), S(19, -1), S(22, -15), S(33, -42), S(27, -65)},
		{S(14, -9), S(44, 2), S(5, -5), S(37, 14), S(21, 20), S(18, -15), S(36, -3), S(35, -8)},
		{S(5, -5), S(16, 9), S(29, 20), S(26, 27), S(47, 22), S(19, 20), S(30, 6), S(18, -8)},
		{S(-20, 17), S(23, 17), S(-1, 30), S(45, 33), S(29, 36), S(31, 22), S(23, 17), S(4, 15)},
		{S(-10, 14), S(-18, 33), S(9, 7), S(12, 33), S(30, 30), S(1, 16), S(12, 30), S(-5, 19)},
		{S(-37, 29), S(-15, 9), S(12, 25), S(-10, 39), S(-1, 35), S(18, 26), S(-19, 4), S(-32, 30)},
		{S(-33, -3), S(-64, 36), S(-128, 48), S(-104, 57), S(-105, 52), S(-98, 45), S(7, 20), S(-40, -1)},
	},
	{ // Rook
		{S(-23, -27), S(-18, -17), S(-11, -15), S(-2, -26), S(-4, -30), S(-5, -17), S(-4, -27), S(-13, -43)},
		{S(-63, -22), S(-26, -31), S(-22, -25), S(-21, -28), S(-22, -30), S(-1, -40), S(2, -44), S(-60, -19)},
		{S(-44, -16), S(-28, -3), S(-35, -5), S(-22, -13), S(-24, -12), S(-17, -13), S(7, -19), S(-28, -23)},
		{S(-41, 7), S(-33, 23), S(-28, 23), S(-15, 11), S(-22, 10), S(-15, 13), S(-7, 9), S(-16, 1)},
		{S(-32, 30), S(-3, 28), S(7, 27), S(36, 17), S(11, 20), S(25, 15), S(23, 9), S(3, 18)},
		{S(-35, 41), S(24, 24), S(19, 32), S(44, 17), S(45, 14), S(77, 19), S(93, -5), S(13, 25)},
		{S(-25, 49), S(-17, 50), S(15, 43), S(38, 42), S(17, 48), S(74, 11), S(22, 28), S(38, 22)},
		{S(12, 44), S(31, 46), S(0, 54), S(9, 47), S(12, 46), S(62, 40), S(85, 34), S(71, 38)},
	},
	{ // Queen
		{S(6, -108), S(2, -82), S(7, -89), S(10, -51), S(11, -87), S(-12, -80), S(3, -120), S(1, -112)},
		{S(-3, -82), S(5, -62), S(13, -85), S(4, -45), S(5, -55), S(17, -99), S(25, -110), S(17, -94)},
		{S(-9, -60), S(7, -25), S(1, 0), S(-6, -7), S(-4, -18), S(0, 6), S(18, -37), S(7, -57)},
		{S(-11, -48), S(-5, 8), S(-10, 16), S(-20, 69), S(-21, 53), S(6, 40), S(8, 11), S(13, 9)},
		{S(-27, -19), S(-14, 20), S(-31, 24), S(-26, 76), S(-31, 98), S(-4, 94), S(20, 69), S(13, 51)},
		{S(-46, -6), S(-28, -4), S(-30, 21), S(-10, 38), S(6, 74), S(50, 66), S(57, 51), S(26, 77)},
		{S(-48, 13), S(-84, 42), S(-28, 21), S(-61, 69), S(-36, 102), S(22, 82), S(-34, 80), S(14, 28)},
		{S(-45, 17), S(-6, 7), S(-12, 43), S(11, 42), S(15, 55), S(53, 51), S(85, 25), S(59, 13)},
	},
	{ // King
		{S(140, -8), S(129, 27), S(55, 65), S(52, 70), S(80, 50), S(51, 68), S(150, 24), S(167, -28)},
		{S(158, 32), S(100, 53), S(41, 89), S(40, 110), S(31, 120), S(51, 87), S(126, 48), S(160, 33)},
		{S(114, 25), S(154, 42), S(93, 81), S(92, 118), S(75, 126), S(102, 78), S(143, 42), S(104, 35)},
		{S(135, 12), S(270, 33), S(163, 90), S(84, 138), S(89, 140), S(149, 87), S(228, 38), S(83, 32)},
		{S(124, 51), S(322, 47), S(160, 108), S(108, 139), S(87, 145), S(192, 97), S(269, 48), S(91, 49)},
		{S(133, 47), S(323, 68), S(231, 101), S(178, 113), S(178, 117), S(253, 91), S(301, 62), S(127, 44)},
		{S(128, -3), S(219, 61), S(207, 78), S(166, 91), S(133, 98), S(178, 76), S(205, 67), S(120, -1)},
		{S(168, -120), S(301, -41), S(188, 1), S(167, 54), S(155, 53), S(169, 11), S(288, -25), S(150, -100)},
	},
}

var PawnsConnected = [7][4]Score{
	{S(0, 0), S(0, 0), S(0, 0), S(0, 0)},
	{S(-1, -15), S(7, 4), S(4, -1), S(15, 17)},
	{S(11, 5), S(22, 3), S(24, 10), S(27, 16)},
	{S(11, 5), S(20, 9), S(11, 10), S(17, 21)},
	{S(5, 17), S(16, 25), S(30, 29), S(28, 23)},
	{S(37, 34), S(25, 73), S(73, 74), S(80, 94)},
	{S(176, 37), S(285, 23), S(282, 53), S(328, 54)},
}

var MobilityBonus = [...][32]Score{
	{S(-62, -128), S(-44, -74), S(-28, -26), S(-19, -2), S(-12, 11), S(-7, 27), // Knights
		S(1, 29), S(10, 24), S(21, 9)},
	{S(-1, -133), S(7, -61), S(15, -20), S(23, 2), S(30, 18), S(35, 35), // Bishops
		S(37, 43), S(38, 46), S(37, 47), S(42, 47), S(45, 42), S(58, 31),
		S(75, 36), S(84, 9)},
	{S(-127, -153), S(-16, -39), S(-4, 15), S(-5, 43), S(-1, 56), S(1, 69), // Rooks
		S(3, 79), S(8, 85), S(12, 90), S(16, 96), S(19, 102), S(20, 107),
		S(26, 109), S(41, 99), S(99, 61)},
	{S(-413, -159), S(-127, -141), S(-48, -177), S(-26, -119), S(-11, -85), S(-10, -16), // Queens
		S(-6, 11), S(-2, 32), S(2, 49), S(5, 63), S(8, 71), S(12, 74),
		S(14, 77), S(15, 83), S(19, 76), S(16, 80), S(15, 75), S(13, 73),
		S(18, 62), S(26, 47), S(38, 27), S(41, 10), S(44, -9), S(58, -40),
		S(23, -29), S(-69, -9), S(135, -120), S(53, -80)},
}

var PassedFriendlyDistance = [8]Score{
	S(0, 0), S(-7, 39), S(-11, 23), S(-8, 8),
	S(-4, -5), S(1, -16), S(17, -27), S(5, -41),
}

var PassedEnemyDistance = [8]Score{
	S(0, 0), S(-100, -42), S(-4, -13), S(6, 5),
	S(13, 18), S(11, 26), S(6, 36), S(18, 42),
}

// PassedRank[Rank] contains a bonus according to the rank of a passed pawn, whether it can be pushed and whether the push would be safe
var PassedRank = [2][2][2][7]Score{
	{
		{
			{S(0, 0), S(-46, -22), S(-39, 8), S(-27, 13), S(29, 15), S(44, 20), S(174, 93)},
			{S(0, 0), S(-32, -50), S(-31, -18), S(-19, 6), S(29, 29), S(56, 43), S(168, 137)},
		},
		{
			{S(0, 0), S(-16, -27), S(-28, 0), S(-18, 19), S(14, 46), S(59, 81), S(251, 141)},
			{S(0, 0), S(-5, -50), S(-25, -6), S(-18, 24), S(13, 59), S(73, 87), S(230, 180)},
		},
	},
	{
		{
			{S(0, 0), S(-20, -9), S(-32, -3), S(-13, 22), S(28, 50), S(67, 98), S(278, 182)},
			{S(0, 0), S(-14, -24), S(-21, -7), S(-11, 35), S(37, 66), S(73, 109), S(238, 233)},
		},
		{
			{S(0, 0), S(-30, -14), S(-29, -11), S(-15, 32), S(19, 86), S(30, 230), S(96, 387)},
			{S(0, 0), S(-13, -35), S(-18, -17), S(-11, 32), S(23, 91), S(54, 217), S(163, 403)},
		},
	},
}

// PassedFile[File] contains a bonus according to the file of a passed pawn
var PassedFile = [8]Score{S(0, 23), S(-3, 27), S(-2, 11), S(-4, 0),
	S(-11, 4), S(-20, 13), S(-18, 25), S(11, 8),
}

var PassedStacked = [7]Score{S(0, 0), S(-17, -45), S(-25, -26), S(-33, -36), S(-14, -43), S(49, -112), S(0, 0)}
var PassedUncontested = [6]Score{S(0, 0), S(-90, 42), S(-84, 39), S(-95, 41), S(-89, 42), S(-69, 28)}
var PassedPushDefended = [6]Score{S(0, 0), S(-3, 20), S(-1, 8), S(6, 3), S(-1, 6), S(-7, 15)}
var PassedPushUncontestedDefended = [6]Score{S(0, 0), S(-41, 21), S(-27, 18), S(-54, 38), S(-63, 48), S(-46, 61)}

var Isolated = S(-7, -17)
var Doubled = S(-11, -31)
var Backward = S(8, -1)
var BackwardOpen = S(-4, -16)

var BishopPair = S(17, 74)
var BishopRammedPawns = S(-9, -23)

var BishopOutpostUndefendedBonus = S(24, -4)
var BishopOutpostDefendedBonus = S(55, 12)

var LongDiagonalBishop = S(23, 25)

var KnightOutpostUndefendedBonus = S(23, -22)
var KnightOutpostDefendedBonus = S(49, 12)

var DistantKnight = [4]Score{S(-15, -3), S(-17, -14), S(-28, -17), S(-54, -21)}

var MinorBehindPawn = S(8, 28)

var Tempo int16 = 15

// Rook on semiopen, open file
var RookOnFile = [2]Score{S(11, 12), S(41, 4)}
var RookOnQueenFile = S(6, 32)

var KingDefenders = [12]Score{
	S(-16, -9), S(-4, -11), S(-1, -10), S(3, -7),
	S(6, -3), S(9, 5), S(12, 12), S(16, 11),
	S(16, 13), S(17, -65), S(-10, -11), S(11, 0),
}

var KingShelter = [2][8][8]Score{
	{{S(-25, 8), S(10, -12), S(19, 2), S(24, 5),
		S(30, -13), S(21, -7), S(16, -38), S(-96, 38)},
		{S(15, 7), S(23, -6), S(-4, 10), S(-7, 7),
			S(3, -1), S(25, 3), S(40, -12), S(-46, 22)},
		{S(16, 3), S(4, -1), S(-26, 1), S(-22, -1),
			S(1, -17), S(-2, 5), S(19, -8), S(-31, -1)},
		{S(-18, 17), S(10, -2), S(-7, -9), S(-2, -8),
			S(5, -24), S(0, -9), S(22, 17), S(-15, -3)},
		{S(-32, 17), S(-15, 8), S(-32, 8), S(-20, 6),
			S(1, -8), S(-22, 0), S(7, -3), S(-27, 2)},
		{S(40, -18), S(23, -14), S(5, -12), S(5, -22),
			S(13, -28), S(26, -20), S(31, -19), S(-12, -4)},
		{S(16, -2), S(-8, -7), S(-28, 4), S(-18, 3),
			S(-8, -15), S(8, -3), S(14, -17), S(-24, 18)},
		{S(-40, 11), S(-26, -1), S(-19, 27), S(-20, 23),
			S(-21, 17), S(-10, 13), S(-12, -16), S(-71, 47)}},
	{{S(41, 53), S(-38, -10), S(-25, 5), S(-32, -3),
		S(-43, -14), S(-6, 8), S(-46, -13), S(-92, 36)},
		{S(143, 16), S(4, -2), S(-2, 5), S(-22, 18),
			S(-5, -11), S(17, 2), S(10, 2), S(-80, 24)},
		{S(0, 31), S(35, 4), S(12, 0), S(20, -5),
			S(28, -5), S(11, 2), S(38, 5), S(-24, 5)},
		{S(5, 35), S(-12, 21), S(-15, 15), S(-22, 21),
			S(0, 10), S(-5, 2), S(-4, 13), S(-42, -6)},
		{S(0, 24), S(-1, 9), S(-6, 4), S(-12, 1),
			S(-5, 1), S(1, -12), S(12, -2), S(-31, -5)},
		{S(21, 17), S(-6, 6), S(-13, 9), S(-7, 8),
			S(-3, 1), S(-38, 10), S(-9, 2), S(-39, 7)},
		{S(32, 4), S(-4, -15), S(-9, -12), S(-26, -6),
			S(-9, -20), S(1, -17), S(2, -20), S(-72, 19)},
		{S(-23, -5), S(-15, -30), S(-9, -16), S(-4, -23),
			S(-6, -34), S(0, -22), S(-5, -49), S(-66, 24)}},
}

var KingStorm = [2][4][8]Score{
	{{S(11, 8), S(3, 15), S(9, 8), S(4, 9),
		S(3, 7), S(11, 2), S(8, 9), S(-11, -5)},
		{S(23, 1), S(16, 6), S(19, 4), S(6, 12),
			S(20, 2), S(27, -9), S(18, -6), S(-7, -5)},
		{S(22, 9), S(3, 7), S(6, 10), S(0, 11),
			S(4, 8), S(11, 0), S(8, -9), S(-2, 4)},
		{S(12, 2), S(8, -2), S(11, -10), S(0, -13),
			S(-2, -7), S(16, -15), S(9, -12), S(-10, 2)}},
	{{S(0, 0), S(5, 4), S(-18, 11), S(10, -5),
		S(-2, 20), S(-13, 32), S(6, 51), S(-6, -4)},
		{S(0, 0), S(-51, -9), S(6, -3), S(47, -5),
			S(4, 0), S(-7, -4), S(14, 55), S(-7, 0)},
		{S(0, 0), S(-19, -10), S(-7, -5), S(17, -2),
			S(6, -6), S(-11, -18), S(39, -39), S(-7, 8)},
		{S(0, 0), S(-16, -24), S(-15, -24), S(-7, -12),
			S(1, -25), S(0, -56), S(3, 0), S(-11, -5)}},
}
var KingOnPawnlessFlank = S(5, -67)

var KingSafetyAttacksWeights = [King + 1]int16{0, -3, -7, -4, 4, 0}
var KingSafetyAttackValue int16 = 124
var KingSafetyWeakSquares int16 = 44
var KingSafetyFriendlyPawns int16 = -35
var KingSafetyNoEnemyQueens int16 = -176
var KingSafetySafeQueenCheck int16 = 90
var KingSafetySafeRookCheck int16 = 73
var KingSafetySafeBishopCheck int16 = 51
var KingSafetySafeKnightCheck int16 = 112
var KingSafetyAdjustment int16 = -12

var Hanging = S(38, 16)
var ThreatByKing = S(-9, 35)
var ThreatByMinor = [King + 1]Score{S(0, 0), S(20, 38), S(17, 38), S(33, 27), S(31, 29), S(1, 26)}
var ThreatByRook = [King + 1]Score{S(0, 0), S(-4, 13), S(-2, 17), S(-4, -11), S(33, 9), S(14, 9)}

// This weights are from black piece on black square perspective
var RookBishopExistence = [16]Score{
	S(24, -29), S(5, -12), S(6, -12), S(2, 2), S(-1, -10), S(-3, 36), S(-2, -2), S(-3, 18), S(2, -11), S(-3, -2), S(-3, 36), S(-3, 20), S(-8, -18), S(-4, 17), S(-1, 16), S(-18, 53),
}
var QueenBishopExistence = [16]Score{
	S(62, -58), S(-4, -6), S(-3, -20), S(-21, -7), S(-6, -8), S(9, 72), S(-16, 0), S(-1, -6), S(-1, -4), S(-11, 25), S(11, 68), S(3, 5), S(-31, -15), S(-6, 31), S(-11, 23), S(-12, 23),
}
var KingBishopExistence = [16]Score{
	S(0, 0), S(0, 3), S(1, 4), S(-15, -11), S(-2, -9), S(-2, -6), S(-2, -3), S(-1, 3), S(1, -1), S(2, 3), S(2, 6), S(2, 12), S(15, 11), S(0, -10), S(2, -7), S(0, 0),
}
