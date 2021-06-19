package evaluation

import (
	. "github.com/David-Orson/casperin/backend"
)

var PawnValue = S(79, 166)
var KnightValue = S(410, 555)
var BishopValue = S(386, 543)
var RookValue = S(541, 926)
var QueenValue = S(1229, 1760)

// Pawns Square scores
// Bishop Flag, Rank, Row
var PawnScores = [16][7][8]Score{
	{
		{},
		{S(-16, 17), S(-10, 8), S(-19, 15), S(-5, 16), S(-19, 25), S(18, 21), S(33, 13), S(-11, 17)},
		{S(-15, 8), S(-15, 7), S(-16, 7), S(-4, 2), S(-1, 2), S(0, 11), S(7, 12), S(-17, 16)},
		{S(-15, 19), S(-18, 17), S(2, 1), S(8, -8), S(16, -11), S(18, 5), S(3, 22), S(-15, 26)},
		{S(-8, 32), S(-1, 19), S(-3, 5), S(8, -10), S(29, -14), S(33, 5), S(34, 21), S(1, 32)},
		{S(2, 55), S(13, 37), S(20, 27), S(47, -7), S(64, -15), S(67, 19), S(67, 35), S(6, 60)},
		{S(12, -2), S(-21, 29), S(2, 22), S(17, -1), S(-20, 4), S(-17, 21), S(-126, 62), S(-142, 54)},
	},
	{
		{},
		{S(-12, 25), S(-7, 6), S(-19, 9), S(8, 17), S(-17, 18), S(27, 15), S(27, -2), S(-11, 13)},
		{S(-9, 9), S(-17, 2), S(-3, 5), S(-1, -6), S(9, 10), S(-4, -3), S(14, 10), S(-23, 10)},
		{S(-6, 18), S(-8, 12), S(5, -7), S(29, -6), S(23, -18), S(29, 0), S(2, 11), S(-13, 13)},
		{S(4, 32), S(15, 3), S(17, 5), S(19, -23), S(42, -14), S(46, -12), S(29, 15), S(-12, 31)},
		{S(-13, 81), S(35, 24), S(23, 25), S(68, -19), S(60, -16), S(103, -12), S(39, 25), S(-4, 50)},
		{S(-21, 67), S(10, 5), S(28, 6), S(17, -8), S(57, -11), S(16, -24), S(-83, 34), S(-108, 28)},
	},
	{
		{},
		{S(-19, 15), S(-15, 14), S(-16, 22), S(-17, 3), S(5, 28), S(9, 6), S(48, 3), S(-9, 19)},
		{S(-22, 9), S(-22, 10), S(-14, -1), S(-1, 8), S(4, -12), S(3, 15), S(-1, -4), S(-18, 17)},
		{S(-13, 6), S(-9, 10), S(8, 5), S(17, -14), S(39, -12), S(23, -8), S(7, 13), S(-13, 21)},
		{S(-9, 23), S(2, 15), S(0, 0), S(24, -8), S(35, -18), S(58, -2), S(22, 10), S(-5, 32)},
		{S(0, 45), S(9, 38), S(63, -8), S(44, -12), S(97, -28), S(68, 18), S(25, 26), S(6, 65)},
		{S(-32, 19), S(-18, 42), S(12, -21), S(21, 6), S(27, -29), S(30, -22), S(-91, 27), S(-101, 29)},
	},
	{
		{},
		{S(-19, 6), S(-12, -5), S(-3, -6), S(3, -18), S(8, -1), S(12, -1), S(43, -9), S(-1, -16)},
		{S(-12, -4), S(-27, -4), S(-1, -8), S(-1, -12), S(10, -6), S(-6, -3), S(7, -13), S(-9, -22)},
		{S(-3, -9), S(-6, -12), S(12, -24), S(35, -19), S(32, -16), S(23, -10), S(23, -24), S(3, -26)},
		{S(-17, 18), S(0, -4), S(34, -45), S(19, -16), S(41, -42), S(45, -24), S(18, 4), S(-5, -12)},
		{S(-27, 74), S(13, 16), S(40, -25), S(24, 1), S(83, -32), S(40, 7), S(38, 13), S(-16, 31)},
		{S(-15, 32), S(-28, 33), S(-27, 9), S(17, -36), S(52, 0), S(44, -8), S(-91, 32), S(-120, 71)},
	},
	{
		{},
		{S(-22, 19), S(-11, 18), S(-6, 7), S(-3, 14), S(-21, 21), S(7, 18), S(34, 8), S(-11, 17)},
		{S(-18, 8), S(-20, 5), S(-8, -15), S(4, -7), S(-9, -10), S(-5, 12), S(1, 2), S(-19, 19)},
		{S(-18, 26), S(-4, 11), S(11, -7), S(13, -29), S(31, -28), S(11, -8), S(6, 19), S(-16, 29)},
		{S(-6, 48), S(11, 18), S(4, -10), S(23, -14), S(34, -39), S(36, 7), S(29, 10), S(-9, 48)},
		{S(-24, 88), S(5, 43), S(8, 32), S(63, -22), S(71, -16), S(82, -7), S(57, 39), S(4, 71)},
		{S(37, 40), S(-27, 43), S(5, 24), S(42, 18), S(73, -36), S(52, 17), S(-78, 61), S(-125, 98)},
	},
	{
		{},
		{S(-3, 30), S(1, 18), S(-8, 20), S(-2, 20), S(-9, 24), S(15, 31), S(25, 9), S(-2, 16)},
		{S(-4, 18), S(-9, 5), S(0, 10), S(6, -4), S(6, 10), S(-1, 8), S(6, 18), S(-12, 14)},
		{S(-2, 32), S(7, 18), S(17, -3), S(28, -7), S(35, -14), S(24, 9), S(7, 23), S(-10, 25)},
		{S(14, 44), S(20, 15), S(10, 15), S(33, -14), S(48, -20), S(48, -6), S(37, 29), S(-7, 43)},
		{S(-7, 98), S(31, 44), S(36, 34), S(67, -17), S(76, -20), S(84, 6), S(28, 46), S(-3, 73)},
		{S(17, 70), S(-37, 21), S(12, 41), S(52, 19), S(22, -2), S(51, 31), S(-52, 21), S(-132, 47)},
	},
	{
		{},
		{S(-6, -17), S(0, -11), S(-4, -8), S(4, -26), S(3, 4), S(6, -18), S(46, -1), S(0, -10)},
		{S(-6, -20), S(-8, -12), S(-12, -31), S(2, -13), S(-7, -32), S(4, -6), S(-9, -23), S(-8, -3)},
		{S(-6, -11), S(9, -31), S(17, -19), S(13, -57), S(36, -31), S(11, -29), S(15, -4), S(-9, 7)},
		{S(6, -5), S(9, -1), S(9, -41), S(21, -18), S(35, -63), S(73, -27), S(29, -22), S(17, 8)},
		{S(-7, 23), S(45, -8), S(23, -8), S(45, -33), S(80, -54), S(94, -51), S(85, -15), S(46, 10)},
		{S(86, -84), S(5, 23), S(19, -83), S(27, -11), S(76, -106), S(21, 3), S(-37, -120), S(-111, 21)},
	},
	{
		{},
		{S(-6, 4), S(-5, 15), S(-5, 7), S(-2, 5), S(2, 20), S(6, 16), S(30, 17), S(-3, 12)},
		{S(-7, 0), S(-18, 7), S(-6, 0), S(1, 6), S(3, -4), S(-10, 28), S(-1, 1), S(-10, -4)},
		{S(-3, 7), S(-2, 7), S(17, -9), S(27, -22), S(31, -6), S(12, 9), S(6, 11), S(-11, 15)},
		{S(8, 27), S(15, 17), S(7, 7), S(26, -4), S(39, -28), S(52, -6), S(16, 13), S(-12, 21)},
		{S(-6, 56), S(25, 30), S(15, 22), S(44, 14), S(84, -33), S(87, 17), S(21, -10), S(-4, 78)},
		{S(12, -6), S(-41, 20), S(-21, -2), S(58, -8), S(-26, -55), S(33, -39), S(-64, 39), S(-135, 18)},
	},
	{
		{},
		{S(-15, 22), S(-7, 8), S(-16, 14), S(-7, 6), S(-14, 14), S(13, 11), S(22, 8), S(-11, 17)},
		{S(-12, 10), S(-15, 1), S(-10, -2), S(-2, -14), S(1, -7), S(-15, 2), S(6, 10), S(-18, 13)},
		{S(-10, 26), S(-7, 21), S(2, -14), S(19, -26), S(21, -34), S(14, 1), S(-5, 17), S(-17, 25)},
		{S(-1, 55), S(11, 12), S(-2, 3), S(10, -32), S(36, -15), S(24, -16), S(25, 21), S(-17, 40)},
		{S(15, 73), S(11, 45), S(62, 4), S(45, -19), S(67, -21), S(85, 13), S(98, 21), S(-8, 61)},
		{S(-31, 58), S(-12, 36), S(44, 17), S(44, -19), S(38, 22), S(42, -16), S(-86, 21), S(-123, 60)},
	},
	{
		{},
		{S(5, -3), S(6, -21), S(-7, -9), S(27, -4), S(-6, -8), S(37, -1), S(32, -24), S(8, -15)},
		{S(2, -11), S(-2, -30), S(15, -20), S(7, -32), S(20, -8), S(-4, -20), S(22, -5), S(-3, -23)},
		{S(5, 7), S(3, -14), S(18, -47), S(40, -27), S(32, -57), S(42, -17), S(12, -18), S(1, -7)},
		{S(21, 13), S(29, -31), S(20, -22), S(20, -66), S(56, -40), S(41, -46), S(37, 0), S(5, -3)},
		{S(41, 31), S(21, 4), S(57, -21), S(74, -41), S(82, -53), S(136, -47), S(68, -19), S(28, 6)},
		{S(2, 21), S(36, -71), S(7, 46), S(53, -79), S(17, -31), S(65, -119), S(-100, -9), S(-117, -70)},
	},
	{
		{},
		{S(-8, 26), S(-2, 19), S(-9, 31), S(-7, 21), S(-15, 36), S(12, 15), S(30, 18), S(-3, 12)},
		{S(-10, 20), S(-7, 14), S(-7, 13), S(4, 8), S(1, -6), S(-4, 18), S(6, 9), S(-12, 12)},
		{S(-3, 25), S(3, 24), S(15, 6), S(19, -3), S(30, -6), S(24, 3), S(12, 18), S(-11, 23)},
		{S(1, 48), S(21, 25), S(9, 12), S(24, -8), S(40, -16), S(52, -2), S(31, 16), S(-4, 40)},
		{S(9, 76), S(17, 68), S(45, 11), S(34, 9), S(92, -22), S(76, 28), S(94, 7), S(2, 67)},
		{S(-12, 47), S(-10, 67), S(-13, 15), S(27, 20), S(80, -15), S(66, -50), S(-89, 25), S(-155, 76)},
	},
	{
		{},
		{S(-8, 22), S(-1, 10), S(-6, 14), S(9, 18), S(-13, 29), S(17, 13), S(30, -3), S(-3, 2)},
		{S(-7, 9), S(-12, 8), S(-3, 20), S(3, -1), S(6, 11), S(-7, 4), S(11, 1), S(-13, 2)},
		{S(0, 22), S(7, 8), S(15, -10), S(29, 5), S(32, -22), S(29, 4), S(9, 3), S(-15, 11)},
		{S(14, 23), S(21, 23), S(15, 19), S(30, -44), S(40, -20), S(52, -25), S(18, 21), S(-19, 39)},
		{S(1, 80), S(24, 54), S(34, 8), S(44, -17), S(62, -11), S(89, -13), S(40, 42), S(0, 34)},
		{S(-16, -8), S(3, 27), S(9, 5), S(21, 5), S(34, -7), S(37, 0), S(-88, 35), S(-117, 35)},
	},
	{
		{},
		{S(-28, 25), S(-16, 19), S(-7, 1), S(-5, 4), S(-25, -12), S(1, 5), S(20, -7), S(0, -10)},
		{S(-24, 23), S(-26, 12), S(-26, 2), S(-4, -10), S(0, -29), S(-3, -12), S(0, -5), S(-7, -4)},
		{S(-15, 35), S(-14, 20), S(5, -24), S(8, -35), S(23, -59), S(22, -28), S(-9, 17), S(-5, 9)},
		{S(-8, 40), S(2, 31), S(-1, 11), S(20, -30), S(28, -25), S(41, -50), S(21, 18), S(4, 39)},
		{S(-23, 131), S(-2, 74), S(4, 47), S(48, 15), S(63, -2), S(84, 2), S(70, 31), S(-18, 106)},
		{S(-15, 49), S(-13, 39), S(-16, 36), S(-1, -8), S(15, -9), S(30, -45), S(-52, 33), S(-93, 78)},
	},
	{
		{},
		{S(-9, 32), S(-1, 15), S(-10, 13), S(-1, 9), S(-16, 13), S(13, 27), S(14, 2), S(-2, 1)},
		{S(-8, 22), S(-17, 4), S(1, 7), S(5, -27), S(4, 7), S(-3, -7), S(3, 7), S(-12, -4)},
		{S(-1, 35), S(0, 19), S(15, -20), S(26, -7), S(31, -40), S(24, 2), S(5, 3), S(-8, 5)},
		{S(13, 55), S(16, 13), S(15, 12), S(30, -49), S(46, -15), S(34, -27), S(37, 22), S(-7, 35)},
		{S(10, 89), S(27, 75), S(29, 39), S(40, 16), S(32, 33), S(89, -14), S(35, 35), S(5, 66)},
		{S(12, 99), S(-29, 48), S(-3, 30), S(33, -15), S(30, -5), S(-11, -50), S(-120, -9), S(-106, 83)},
	},
	{
		{},
		{S(-19, 24), S(-5, 14), S(-8, 26), S(-2, 4), S(-23, 27), S(-2, 5), S(20, 22), S(-6, 16)},
		{S(-17, 9), S(-11, 10), S(-8, -20), S(1, -1), S(-10, -16), S(-7, 10), S(-3, -5), S(-13, 16)},
		{S(-14, 27), S(-2, 17), S(13, 1), S(15, -36), S(31, -20), S(13, -14), S(8, 7), S(-15, 18)},
		{S(-7, 51), S(23, 29), S(2, -2), S(23, -1), S(34, -45), S(55, -21), S(27, -2), S(-6, 44)},
		{S(-1, 110), S(16, 45), S(47, 33), S(25, 6), S(91, -17), S(61, 12), S(48, 31), S(7, 61)},
		{S(-13, 14), S(-17, 76), S(-51, 27), S(10, 26), S(8, -24), S(28, 10), S(-76, 21), S(-134, 64)},
	},
	{
		{},
		{S(-11, 26), S(-4, 31), S(-15, 30), S(-7, 16), S(-20, 19), S(0, 31), S(11, 20), S(-9, 17)},
		{S(-10, 22), S(-15, 20), S(-8, 20), S(-3, 10), S(-3, 12), S(-16, 36), S(-9, 18), S(-19, 15)},
		{S(-3, 28), S(0, 29), S(12, 11), S(18, 5), S(25, -9), S(13, 14), S(-1, 24), S(-20, 25)},
		{S(8, 57), S(22, 28), S(4, 37), S(25, -16), S(34, -7), S(41, -7), S(19, 28), S(-14, 41)},
		{S(5, 98), S(19, 72), S(15, 75), S(44, 9), S(58, -15), S(63, -1), S(28, 27), S(-9, 70)},
		{S(-89, 109), S(-74, 21), S(-36, 47), S(5, 1), S(15, -41), S(32, 3), S(-60, 16), S(-162, 25)},
	},
}

// Piece Square Values
var PieceScores = [King + 1][8][8]Score{
	{},
	{ // knight
		{S(-79, -88), S(-23, -45), S(-42, -31), S(-21, -3), S(-23, -10), S(-25, -33), S(-28, -32), S(-87, -44)},
		{S(-29, -31), S(-18, -20), S(-15, -36), S(-16, -10), S(-24, -12), S(-17, -32), S(-25, -16), S(-21, -22)},
		{S(-12, -45), S(-7, -19), S(-8, -11), S(-3, 10), S(-2, 9), S(-17, -11), S(-5, -23), S(-22, -46)},
		{S(-10, -5), S(-8, 6), S(7, 21), S(-6, 34), S(8, 30), S(0, 25), S(21, -4), S(-13, 0)},
		{S(-4, 9), S(-4, 6), S(8, 33), S(8, 52), S(7, 44), S(24, 29), S(6, 9), S(7, 5)},
		{S(-52, 7), S(-27, 16), S(-4, 40), S(5, 41), S(26, 28), S(26, 30), S(-14, 7), S(-21, -4)},
		{S(-1, -21), S(-14, -5), S(22, -18), S(32, 17), S(38, 8), S(59, -40), S(-23, -7), S(13, -35)},
		{S(-203, -53), S(-98, -9), S(-153, 27), S(-30, -7), S(8, -9), S(-127, 32), S(-75, -10), S(-187, -69)},
	},
	{ // Bishop
		{S(33, -41), S(50, -14), S(6, 1), S(13, 2), S(0, 10), S(15, -8), S(13, -13), S(22, -40)},
		{S(45, -34), S(20, -35), S(44, -10), S(8, 6), S(18, -1), S(17, -15), S(32, -42), S(23, -65)},
		{S(21, -7), S(50, 0), S(8, -5), S(33, 15), S(16, 21), S(17, -14), S(30, -1), S(33, -8)},
		{S(16, -5), S(20, 8), S(32, 18), S(21, 28), S(42, 22), S(14, 21), S(26, 5), S(13, -7)},
		{S(-12, 20), S(28, 15), S(0, 31), S(40, 34), S(22, 37), S(27, 22), S(18, 17), S(2, 15)},
		{S(1, 13), S(-13, 33), S(12, 6), S(6, 34), S(26, 31), S(-3, 18), S(11, 29), S(-9, 19)},
		{S(-29, 32), S(-5, 8), S(11, 26), S(-12, 40), S(-7, 38), S(17, 28), S(-18, 7), S(-31, 31)},
		{S(-22, -1), S(-64, 40), S(-126, 51), S(-107, 60), S(-105, 54), S(-99, 47), S(23, 17), S(-40, 2)},
	},
	{ // Rook
		{S(-23, -28), S(-18, -18), S(-12, -16), S(-3, -27), S(-4, -30), S(-6, -17), S(-4, -27), S(-12, -46)},
		{S(-59, -22), S(-24, -33), S(-21, -26), S(-19, -30), S(-19, -31), S(-1, -40), S(2, -46), S(-57, -19)},
		{S(-43, -16), S(-28, -2), S(-34, -4), S(-20, -14), S(-22, -13), S(-16, -11), S(6, -20), S(-27, -22)},
		{S(-39, 9), S(-34, 25), S(-28, 24), S(-15, 11), S(-23, 12), S(-16, 15), S(-7, 10), S(-18, 2)},
		{S(-33, 35), S(-7, 31), S(3, 31), S(30, 20), S(8, 22), S(21, 19), S(18, 14), S(0, 21)},
		{S(-36, 47), S(19, 27), S(13, 38), S(39, 19), S(41, 18), S(69, 25), S(86, -1), S(9, 30)},
		{S(-26, 55), S(-19, 55), S(10, 48), S(32, 48), S(15, 52), S(68, 16), S(19, 33), S(31, 27)},
		{S(15, 44), S(29, 46), S(0, 56), S(10, 47), S(13, 49), S(61, 41), S(81, 34), S(66, 40)},
	},
	{ // Queen
		{S(7, -108), S(2, -83), S(6, -86), S(8, -50), S(10, -85), S(-11, -79), S(4, -122), S(3, -112)},
		{S(-3, -81), S(5, -62), S(11, -82), S(3, -44), S(4, -52), S(17, -99), S(24, -110), S(17, -95)},
		{S(-8, -59), S(6, -26), S(-1, -1), S(-6, -9), S(-4, -17), S(-1, 6), S(17, -37), S(8, -55)},
		{S(-11, -48), S(-6, 5), S(-11, 14), S(-20, 68), S(-21, 55), S(3, 42), S(8, 9), S(12, 11)},
		{S(-26, -21), S(-15, 20), S(-30, 23), S(-26, 77), S(-30, 100), S(-5, 94), S(18, 72), S(13, 52)},
		{S(-44, -4), S(-28, -3), S(-30, 20), S(-11, 39), S(6, 75), S(49, 68), S(57, 53), S(25, 80)},
		{S(-46, 16), S(-78, 43), S(-28, 23), S(-60, 74), S(-35, 108), S(25, 85), S(-33, 81), S(15, 29)},
		{S(-44, 22), S(-7, 11), S(-12, 47), S(13, 44), S(16, 59), S(61, 49), S(90, 19), S(63, 10)},
	},
	{ // King
		{S(142, -9), S(128, 26), S(55, 62), S(54, 67), S(81, 44), S(54, 64), S(150, 18), S(167, -34)},
		{S(155, 34), S(101, 56), S(40, 90), S(40, 111), S(30, 121), S(52, 88), S(127, 47), S(158, 30)},
		{S(113, 28), S(153, 45), S(92, 82), S(90, 120), S(72, 129), S(100, 81), S(139, 45), S(105, 35)},
		{S(142, 11), S(268, 31), S(160, 91), S(85, 140), S(88, 144), S(146, 91), S(221, 42), S(88, 31)},
		{S(132, 51), S(324, 44), S(157, 108), S(109, 139), S(88, 147), S(192, 102), S(266, 52), S(96, 53)},
		{S(140, 47), S(324, 66), S(224, 103), S(173, 112), S(175, 119), S(253, 94), S(301, 65), S(133, 46)},
		{S(136, -4), S(222, 63), S(212, 76), S(177, 89), S(127, 98), S(174, 78), S(203, 70), S(124, -3)},
		{S(180, -123), S(312, -41), S(201, 1), S(173, 54), S(154, 53), S(167, 11), S(287, -26), S(145, -100)},
	},
}

var PawnsConnected = [7][4]Score{
	{S(0, 0), S(0, 0), S(0, 0), S(0, 0)},
	{S(0, -15), S(6, 3), S(3, -1), S(14, 15)},
	{S(11, 5), S(20, 3), S(22, 11), S(24, 16)},
	{S(9, 5), S(18, 10), S(10, 11), S(15, 20)},
	{S(3, 18), S(13, 26), S(25, 30), S(25, 25)},
	{S(34, 30), S(21, 71), S(65, 75), S(74, 92)},
	{S(176, 37), S(280, 22), S(275, 55), S(317, 56)},
}

var MobilityBonus = [...][32]Score{
	{S(-58, -125), S(-42, -69), S(-28, -22), S(-21, 1), S(-14, 14), S(-10, 30), // Knights
		S(-2, 32), S(6, 28), S(16, 12)},
	{S(4, -128), S(9, -58), S(16, -17), S(23, 5), S(29, 19), S(34, 34), // Bishops
		S(36, 43), S(35, 46), S(36, 49), S(39, 49), S(43, 43), S(54, 34),
		S(74, 35), S(81, 8)},
	{S(-126, -156), S(-14, -35), S(-4, 19), S(-6, 47), S(-2, 59), S(0, 72), // Rooks
		S(1, 81), S(5, 89), S(10, 94), S(13, 100), S(16, 107), S(17, 113),
		S(22, 114), S(38, 101), S(95, 57)},
	{S(-413, -159), S(-130, -143), S(-46, -178), S(-23, -117), S(-10, -83), S(-10, -13), // Queens
		S(-6, 14), S(-2, 34), S(1, 50), S(4, 63), S(7, 71), S(11, 74),
		S(13, 77), S(14, 82), S(17, 77), S(14, 82), S(14, 78), S(12, 77),
		S(16, 66), S(23, 52), S(35, 30), S(38, 13), S(40, -5), S(53, -40),
		S(24, -29), S(-61, -10), S(140, -118), S(56, -79)},
}

var PassedFriendlyDistance = [8]Score{
	S(0, 0), S(-1, 36), S(-7, 21), S(-7, 7),
	S(-4, -7), S(-1, -17), S(14, -29), S(1, -40),
}

var PassedEnemyDistance = [8]Score{
	S(0, 0), S(-87, -43), S(2, -14), S(5, 5),
	S(10, 17), S(9, 24), S(5, 32), S(15, 37),
}

// PassedRank[Rank] contains a bonus according to the rank of a passed pawn, whether it can be pushed and whether the push would be safe
var PassedRank = [2][2][2][7]Score{
	{
		{
			{S(0, 0), S(-47, -24), S(-39, 8), S(-25, 13), S(29, 13), S(44, 16), S(173, 87)},
			{S(0, 0), S(-32, -55), S(-31, -21), S(-18, 3), S(30, 25), S(60, 37), S(166, 129)},
		},
		{
			{S(0, 0), S(-17, -29), S(-27, -1), S(-17, 18), S(15, 44), S(58, 78), S(252, 140)},
			{S(0, 0), S(-6, -54), S(-25, -6), S(-19, 24), S(14, 58), S(75, 85), S(226, 175)},
		},
	},
	{
		{
			{S(0, 0), S(-19, -9), S(-30, -4), S(-10, 21), S(29, 48), S(64, 97), S(276, 179)},
			{S(0, 0), S(-13, -24), S(-19, -5), S(-9, 35), S(36, 65), S(74, 107), S(235, 229)},
		},
		{
			{S(0, 0), S(-28, -14), S(-26, -13), S(-13, 31), S(20, 86), S(32, 233), S(77, 410)},
			{S(0, 0), S(-12, -35), S(-17, -16), S(-10, 33), S(21, 93), S(56, 221), S(138, 422)},
		},
	},
}

// PassedFile[File] contains a bonus according to the file of a passed pawn
var PassedFile = [8]Score{S(0, 23), S(-2, 27), S(-2, 10), S(-5, -4),
	S(-10, 2), S(-18, 10), S(-16, 23), S(8, 8),
}

var PassedStacked = [7]Score{S(0, 0), S(-22, -58), S(-31, -40), S(-44, -51), S(-23, -52), S(42, -101), S(0, 0)}
var PassedUncontested = [6]Score{S(0, 0), S(-88, 45), S(-77, 43), S(-86, 46), S(-80, 47), S(-66, 37)}
var PassedPushDefended = [6]Score{S(0, 0), S(-3, 20), S(0, 7), S(7, 1), S(-3, 6), S(-15, 16)}
var PassedPushUncontestedDefended = [6]Score{S(0, 0), S(-57, 31), S(-36, 24), S(-64, 48), S(-72, 57), S(-58, 70)}

var Isolated = S(-5, -15)

// Is doubled, can be traded
var StackedPawns = [2][2][8]Score{
	{
		{S(-14, -37), S(-11, -28), S(-5, -27), S(-5, -25), S(-10, -26), S(-4, -26), S(-4, -30), S(-1, -34)},
		{S(-17, -39), S(-10, -26), S(-3, -23), S(-10, -26), S(-11, -23), S(-7, -26), S(-8, -31), S(-4, -38)},
	},
	{
		{S(-17, -40), S(-11, -30), S(-14, -35), S(-18, -37), S(-13, -32), S(-3, -32), S(-8, -34), S(-13, -41)},
		{S(-20, -42), S(-7, -34), S(-18, -38), S(-18, -34), S(-17, -34), S(-7, -40), S(-1, -39), S(-19, -43)},
	},
}
var Backward = S(8, -2)
var BackwardOpen = S(-4, -19)

var BishopPair = S(15, 79)
var BishopRammedPawns = S(-7, -23)

var BishopOutpostUndefendedBonus = S(22, -3)
var BishopOutpostDefendedBonus = S(51, 13)

var LongDiagonalBishop = S(21, 25)
var DistantBishop = [4]Score{S(-7, 1), S(-11, -1), S(-14, -4), S(-18, -21)}

var KnightOutpostUndefendedBonus = S(21, -22)
var KnightOutpostDefendedBonus = S(43, 16)

var DistantKnight = [4]Score{S(-15, -2), S(-16, -14), S(-28, -16), S(-53, -21)}

var MinorBehindPawn = S(7, 29)

var Tempo int16 = 15

// Rook on semiopen, open file
var RookOnFile = [2]Score{S(12, 12), S(38, 4)}
var RookOnQueenFile = S(6, 29)

var KingDefenders = [12]Score{
	S(-18, -15), S(-2, -11), S(1, -7), S(4, -4),
	S(6, -1), S(9, 5), S(12, 10), S(15, 10),
	S(14, 14), S(16, -62), S(-16, -12), S(11, 0),
}

var KingShelter = [2][8][8]Score{
	{{S(-25, 10), S(10, -13), S(17, 0), S(22, 4),
		S(28, -14), S(19, -6), S(17, -39), S(-90, 40)},
		{S(17, 7), S(24, -7), S(-5, 10), S(-7, 7),
			S(2, 0), S(23, 3), S(41, -15), S(-45, 24)},
		{S(15, 3), S(4, -1), S(-24, -1), S(-20, -2),
			S(1, -17), S(0, 4), S(19, -7), S(-30, 1)},
		{S(-18, 16), S(8, -2), S(-5, -10), S(-2, -10),
			S(6, -24), S(1, -9), S(17, 20), S(-15, -2)},
		{S(-32, 21), S(-15, 9), S(-31, 9), S(-22, 9),
			S(0, -8), S(-23, 1), S(8, -3), S(-26, 3)},
		{S(39, -17), S(23, -15), S(5, -13), S(4, -23),
			S(12, -29), S(26, -19), S(31, -19), S(-11, -3)},
		{S(16, -2), S(-7, -8), S(-28, 3), S(-19, 2),
			S(-7, -16), S(8, -4), S(15, -17), S(-25, 20)},
		{S(-40, 10), S(-26, -2), S(-20, 27), S(-21, 23),
			S(-22, 15), S(-12, 10), S(-11, -16), S(-70, 50)}},
	{{S(34, 56), S(-35, -13), S(-25, 4), S(-30, -5),
		S(-41, -15), S(-6, 6), S(-42, -16), S(-90, 41)},
		{S(138, 18), S(7, -6), S(-5, 4), S(-23, 18),
			S(-6, -11), S(12, 3), S(13, 0), S(-79, 27)},
		{S(-4, 33), S(34, 3), S(13, -1), S(19, -5),
			S(26, -6), S(12, 3), S(36, 5), S(-23, 7)},
		{S(6, 34), S(-13, 21), S(-14, 14), S(-23, 22),
			S(-1, 10), S(-5, 4), S(-5, 11), S(-42, -3)},
		{S(4, 23), S(-1, 8), S(-7, 5), S(-11, 0),
			S(-4, 0), S(0, -11), S(12, 0), S(-29, -4)},
		{S(17, 20), S(-6, 8), S(-15, 10), S(-10, 8),
			S(-4, 2), S(-37, 12), S(-9, 3), S(-38, 10)},
		{S(30, 4), S(-5, -15), S(-9, -13), S(-26, -8),
			S(-10, -18), S(1, -16), S(2, -19), S(-70, 22)},
		{S(-27, -3), S(-14, -34), S(-9, -20), S(-4, -24),
			S(-4, -37), S(-1, -24), S(-5, -52), S(-64, 28)}},
}

var KingStorm = [2][4][8]Score{
	{{S(11, 8), S(3, 16), S(8, 7), S(5, 8),
		S(4, 5), S(11, 1), S(8, 9), S(-11, -6)},
		{S(24, 0), S(16, 6), S(19, 3), S(6, 10),
			S(20, 2), S(27, -9), S(19, -6), S(-7, -4)},
		{S(22, 10), S(4, 7), S(5, 10), S(0, 11),
			S(3, 8), S(11, 2), S(8, -7), S(-3, 6)},
		{S(11, 2), S(8, -2), S(11, -11), S(0, -13),
			S(-2, -7), S(16, -15), S(9, -13), S(-11, 4)}},
	{{S(0, 0), S(3, 3), S(-15, 8), S(11, -7),
		S(-2, 21), S(-12, 32), S(4, 51), S(-7, -5)},
		{S(0, 0), S(-53, -11), S(7, -7), S(46, -8),
			S(5, 0), S(-8, -4), S(15, 54), S(-6, 0)},
		{S(0, 0), S(-20, -13), S(-6, -7), S(16, -2),
			S(6, -5), S(-10, -17), S(36, -39), S(-7, 7)},
		{S(0, 0), S(-18, -26), S(-15, -26), S(-6, -11),
			S(2, -24), S(0, -54), S(4, -2), S(-11, -4)}},
}
var KingOnPawnlessFlank = S(-13, -74)

var Hanging = S(33, 19)
var ThreatByKing = S(-11, 37)
var ThreatByMinor = [King + 1]Score{S(0, 0), S(18, 41), S(16, 39), S(32, 30), S(30, 30), S(2, 28)}
var ThreatByRook = [King + 1]Score{S(0, 0), S(-3, 13), S(-2, 18), S(-5, -10), S(33, 9), S(14, 10)}

// This weights are from black piece on black square perspective
var RookBishopExistence = [16]Score{
	S(18, -11), S(3, -9), S(3, -9), S(1, 0), S(0, -6), S(-6, 36), S(-2, -1), S(-6, 19), S(2, -6), S(-3, -2), S(-6, 37), S(-6, 21), S(-8, -19), S(-4, 18), S(-2, 18), S(-21, 56),
}
var QueenBishopExistence = [16]Score{
	S(85, -57), S(-5, -9), S(-3, -23), S(-24, -8), S(-5, -9), S(10, 78), S(-19, 0), S(-1, -3), S(-1, -4), S(-15, 25), S(12, 74), S(3, 6), S(-38, -14), S(-8, 38), S(-12, 30), S(-15, 35),
}
var KingBishopExistence = [16]Score{
	S(0, 0), S(-2, 5), S(-1, 6), S(-19, -6), S(1, -11), S(-2, -6), S(-1, -4), S(-1, 6), S(3, -3), S(1, 4), S(2, 6), S(1, 15), S(19, 6), S(1, -13), S(3, -10), S(0, 0),
}

//
// King safety
//

var KingSafetyAttacksWeights = [Queen + 1]Score{S(0, 26), S(-2, -3), S(-6, 19), S(-3, -9), S(1, 29)}
var KingSafetyWeakSquares = S(41, 32)
var KingSafetyFriendlyPawns = S(-34, -14)
var KingSafetyNoEnemyQueens = S(-170, -173)
var KingSafetySafeQueenCheck = S(89, 95)
var KingSafetySafeRookCheck = S(77, 81)
var KingSafetySafeBishopCheck = S(50, 51)
var KingSafetySafeKnightCheck = S(114, 107)
var KingSafetyAdjustment = S(-20, -3)

// Attack value is special as it is scaled by a fraction
var KingSafetyAttackValue = S(118, 137)

//
// Complexity
//

var ComplexityTotalPawns = S(0, 6)
var ComplexityPawnEndgame = S(0, 83)
var ComplexityPawnBothFlanks = S(0, 106)
var ComplexityInfiltration = S(0, 6)
var ComplexityAdjustment = S(0, -173)
