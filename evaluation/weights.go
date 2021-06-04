package evaluation

import (
	. "github.com/David-Orson/casperin/backend"
)

var PawnValue = S(78, 160)
var KnightValue = S(415, 551)
var BishopValue = S(388, 541)
var RookValue = S(551, 918)
var QueenValue = S(1238, 1752)

// Pawns Square scores
// Bishop Flag, Rank, Row
var PawnScores = [16][7][8]Score{
	{
		{},
		{S(-17, 17), S(-10, 6), S(-20, 14), S(-7, 14), S(-21, 23), S(17, 18), S(34, 9), S(-11, 15)},
		{S(-15, 8), S(-16, 6), S(-17, 7), S(-5, 1), S(-1, 0), S(1, 9), S(7, 9), S(-17, 14)},
		{S(-14, 19), S(-19, 16), S(2, 1), S(9, -8), S(18, -11), S(18, 1), S(3, 19), S(-14, 24)},
		{S(-7, 31), S(-3, 18), S(-3, 4), S(6, -10), S(29, -15), S(32, 1), S(32, 16), S(2, 29)},
		{S(2, 54), S(11, 35), S(20, 25), S(46, -8), S(64, -16), S(65, 15), S(65, 32), S(5, 58)},
		{S(12, -4), S(-22, 27), S(1, 20), S(16, -2), S(-21, 2), S(-18, 20), S(-128, 59), S(-144, 52)},
	},
	{
		{},
		{S(-12, 26), S(-7, 6), S(-20, 9), S(8, 17), S(-18, 18), S(27, 14), S(27, -2), S(-12, 13)},
		{S(-8, 10), S(-17, 2), S(-3, 6), S(-1, -5), S(9, 10), S(-3, -3), S(15, 9), S(-24, 11)},
		{S(-5, 20), S(-7, 12), S(6, -5), S(31, -5), S(25, -17), S(28, -1), S(0, 11), S(-13, 13)},
		{S(5, 33), S(14, 1), S(18, 6), S(18, -22), S(43, -14), S(46, -14), S(28, 14), S(-12, 31)},
		{S(-13, 82), S(35, 25), S(23, 26), S(68, -18), S(60, -15), S(103, -12), S(38, 24), S(-4, 51)},
		{S(-21, 68), S(10, 5), S(28, 6), S(17, -8), S(57, -10), S(16, -23), S(-83, 35), S(-108, 28)},
	},
	{
		{},
		{S(-19, 16), S(-14, 15), S(-15, 22), S(-18, 3), S(5, 28), S(9, 5), S(48, 2), S(-9, 20)},
		{S(-21, 10), S(-22, 11), S(-13, 0), S(-1, 9), S(5, -12), S(3, 15), S(0, -4), S(-17, 18)},
		{S(-12, 7), S(-9, 11), S(9, 6), S(19, -13), S(41, -10), S(23, -8), S(7, 12), S(-13, 21)},
		{S(-9, 24), S(2, 15), S(1, 1), S(25, -7), S(35, -17), S(59, -3), S(21, 8), S(-5, 33)},
		{S(0, 45), S(9, 38), S(64, -7), S(44, -11), S(97, -27), S(68, 17), S(25, 26), S(6, 66)},
		{S(-32, 20), S(-18, 42), S(12, -20), S(21, 6), S(27, -29), S(30, -22), S(-91, 27), S(-101, 29)},
	},
	{
		{},
		{S(-19, 6), S(-12, -5), S(-3, -7), S(2, -18), S(8, -1), S(11, -1), S(44, -9), S(-2, -17)},
		{S(-12, -4), S(-28, -5), S(-1, -8), S(-1, -12), S(10, -6), S(-7, -3), S(7, -12), S(-9, -22)},
		{S(-2, -9), S(-7, -12), S(13, -24), S(36, -19), S(33, -15), S(22, -9), S(22, -24), S(2, -25)},
		{S(-17, 18), S(-1, -5), S(34, -45), S(19, -16), S(42, -42), S(45, -25), S(17, 4), S(-5, -12)},
		{S(-27, 73), S(13, 16), S(40, -25), S(24, 1), S(83, -32), S(40, 7), S(38, 13), S(-16, 31)},
		{S(-15, 32), S(-28, 33), S(-27, 9), S(17, -36), S(52, 0), S(44, -8), S(-91, 32), S(-120, 71)},
	},
	{
		{},
		{S(-23, 19), S(-11, 18), S(-7, 7), S(-3, 14), S(-22, 21), S(6, 16), S(34, 7), S(-11, 17)},
		{S(-18, 8), S(-21, 6), S(-8, -14), S(4, -7), S(-10, -11), S(-5, 12), S(0, 1), S(-19, 19)},
		{S(-18, 27), S(-3, 11), S(11, -6), S(14, -28), S(32, -27), S(10, -9), S(5, 19), S(-17, 29)},
		{S(-6, 47), S(10, 18), S(4, -10), S(23, -14), S(33, -40), S(35, 6), S(27, 8), S(-9, 47)},
		{S(-23, 88), S(5, 42), S(8, 32), S(63, -23), S(71, -16), S(82, -8), S(57, 38), S(4, 72)},
		{S(37, 41), S(-27, 43), S(5, 24), S(42, 18), S(73, -35), S(52, 17), S(-78, 60), S(-125, 98)},
	},
	{
		{},
		{S(-3, 30), S(1, 18), S(-8, 20), S(-2, 20), S(-9, 24), S(15, 30), S(24, 8), S(-2, 16)},
		{S(-4, 19), S(-9, 5), S(0, 11), S(6, -5), S(6, 10), S(0, 9), S(5, 18), S(-12, 15)},
		{S(-1, 34), S(8, 18), S(18, -1), S(30, -6), S(36, -13), S(24, 8), S(6, 23), S(-9, 25)},
		{S(15, 45), S(18, 14), S(10, 17), S(32, -14), S(49, -20), S(47, -7), S(37, 28), S(-6, 43)},
		{S(-6, 99), S(32, 44), S(37, 34), S(68, -17), S(77, -19), S(84, 6), S(28, 45), S(-3, 72)},
		{S(18, 71), S(-37, 21), S(13, 41), S(52, 19), S(22, -2), S(51, 31), S(-52, 21), S(-132, 47)},
	},
	{
		{},
		{S(-7, -17), S(1, -11), S(-4, -8), S(3, -26), S(2, 4), S(5, -19), S(46, -1), S(0, -10)},
		{S(-6, -19), S(-8, -11), S(-12, -30), S(1, -12), S(-8, -33), S(4, -6), S(-9, -23), S(-8, -2)},
		{S(-6, -10), S(10, -31), S(18, -18), S(15, -55), S(38, -30), S(10, -29), S(14, -4), S(-9, 7)},
		{S(7, -5), S(9, 0), S(9, -40), S(22, -17), S(34, -63), S(74, -27), S(28, -25), S(17, 9)},
		{S(-7, 23), S(45, -7), S(23, -8), S(45, -33), S(80, -54), S(93, -51), S(85, -14), S(47, 10)},
		{S(86, -83), S(5, 23), S(19, -83), S(27, -11), S(76, -106), S(21, 3), S(-38, -120), S(-111, 21)},
	},
	{
		{},
		{S(-6, 4), S(-5, 15), S(-4, 8), S(-3, 4), S(2, 20), S(5, 16), S(30, 16), S(-3, 11)},
		{S(-7, 0), S(-18, 7), S(-5, 0), S(1, 6), S(3, -3), S(-10, 28), S(-1, 1), S(-10, -3)},
		{S(-3, 6), S(-2, 7), S(18, -9), S(29, -21), S(32, -5), S(12, 10), S(5, 11), S(-11, 15)},
		{S(8, 27), S(14, 17), S(7, 7), S(25, -4), S(40, -29), S(52, -5), S(15, 12), S(-12, 21)},
		{S(-6, 56), S(25, 30), S(15, 22), S(44, 14), S(85, -33), S(87, 17), S(21, -10), S(-4, 78)},
		{S(12, -6), S(-41, 20), S(-21, -2), S(58, -8), S(-26, -55), S(33, -39), S(-64, 39), S(-135, 18)},
	},
	{
		{},
		{S(-15, 23), S(-8, 8), S(-16, 13), S(-7, 6), S(-14, 13), S(12, 10), S(23, 7), S(-12, 17)},
		{S(-12, 11), S(-16, 1), S(-10, -1), S(-3, -14), S(1, -7), S(-14, 2), S(6, 10), S(-18, 13)},
		{S(-9, 27), S(-7, 21), S(2, -13), S(20, -25), S(22, -33), S(12, -1), S(-5, 16), S(-17, 24)},
		{S(1, 55), S(9, 11), S(-1, 3), S(8, -32), S(36, -14), S(22, -19), S(24, 19), S(-18, 39)},
		{S(16, 73), S(11, 44), S(63, 4), S(45, -19), S(68, -22), S(85, 12), S(98, 20), S(-8, 60)},
		{S(-30, 58), S(-12, 37), S(44, 17), S(44, -19), S(38, 22), S(42, -16), S(-86, 20), S(-123, 58)},
	},
	{
		{},
		{S(5, -2), S(5, -20), S(-8, -9), S(27, -4), S(-7, -7), S(37, -2), S(33, -25), S(7, -14)},
		{S(2, -10), S(-3, -29), S(15, -19), S(6, -32), S(20, -7), S(-3, -19), S(23, -6), S(-3, -22)},
		{S(6, 8), S(3, -13), S(19, -46), S(41, -25), S(34, -55), S(42, -18), S(12, -18), S(1, -7)},
		{S(22, 14), S(28, -32), S(21, -21), S(19, -66), S(57, -39), S(40, -49), S(36, 0), S(5, -3)},
		{S(42, 32), S(21, 5), S(58, -21), S(74, -40), S(82, -52), S(136, -47), S(67, -19), S(28, 7)},
		{S(2, 22), S(36, -71), S(7, 46), S(53, -79), S(17, -31), S(64, -119), S(-100, -9), S(-118, -69)},
	},
	{
		{},
		{S(-8, 27), S(-2, 19), S(-9, 32), S(-7, 21), S(-15, 36), S(11, 14), S(32, 17), S(-3, 12)},
		{S(-10, 22), S(-7, 15), S(-7, 15), S(5, 7), S(1, -6), S(-5, 18), S(7, 9), S(-12, 12)},
		{S(-2, 26), S(3, 25), S(16, 8), S(20, -1), S(32, -4), S(23, 2), S(13, 17), S(-11, 23)},
		{S(1, 49), S(21, 26), S(10, 13), S(24, -8), S(39, -17), S(53, -3), S(29, 13), S(-4, 40)},
		{S(10, 77), S(17, 67), S(46, 11), S(34, 9), S(93, -22), S(76, 28), S(94, 6), S(3, 67)},
		{S(-11, 48), S(-9, 67), S(-13, 14), S(28, 20), S(80, -14), S(66, -50), S(-89, 24), S(-155, 76)},
	},
	{
		{},
		{S(-9, 23), S(-1, 9), S(-6, 14), S(9, 18), S(-14, 29), S(17, 13), S(31, -3), S(-4, 1)},
		{S(-7, 9), S(-12, 8), S(-3, 20), S(3, -1), S(6, 11), S(-8, 4), S(12, 2), S(-14, 2)},
		{S(1, 23), S(7, 8), S(17, -9), S(30, 6), S(34, -21), S(28, 4), S(10, 3), S(-16, 10)},
		{S(15, 23), S(21, 23), S(15, 20), S(30, -45), S(40, -20), S(52, -25), S(17, 21), S(-20, 39)},
		{S(1, 80), S(24, 54), S(34, 8), S(44, -17), S(62, -11), S(89, -13), S(40, 42), S(0, 34)},
		{S(-16, -8), S(3, 27), S(9, 5), S(21, 5), S(34, -7), S(37, 0), S(-88, 35), S(-117, 35)},
	},
	{
		{},
		{S(-28, 25), S(-16, 20), S(-8, 1), S(-5, 4), S(-26, -12), S(0, 5), S(20, -7), S(0, -10)},
		{S(-24, 24), S(-27, 12), S(-26, 3), S(-5, -10), S(-1, -29), S(-2, -11), S(0, -5), S(-7, -3)},
		{S(-15, 36), S(-13, 21), S(5, -23), S(9, -34), S(24, -58), S(21, -29), S(-10, 17), S(-5, 8)},
		{S(-7, 41), S(2, 31), S(0, 12), S(20, -30), S(27, -25), S(40, -51), S(20, 17), S(4, 40)},
		{S(-22, 131), S(-2, 74), S(4, 47), S(48, 15), S(63, -2), S(84, 2), S(70, 31), S(-17, 106)},
		{S(-15, 49), S(-13, 39), S(-16, 36), S(-1, -8), S(15, -9), S(30, -45), S(-52, 33), S(-93, 78)},
	},
	{
		{},
		{S(-9, 33), S(-1, 15), S(-10, 13), S(-1, 9), S(-17, 13), S(12, 27), S(14, 2), S(-2, 1)},
		{S(-8, 22), S(-18, 5), S(1, 8), S(5, -28), S(4, 8), S(0, -6), S(3, 8), S(-12, -3)},
		{S(-1, 36), S(0, 20), S(16, -19), S(27, -6), S(32, -39), S(24, 1), S(5, 3), S(-8, 6)},
		{S(14, 56), S(14, 12), S(15, 13), S(29, -49), S(46, -14), S(32, -29), S(36, 22), S(-7, 35)},
		{S(11, 89), S(28, 75), S(29, 39), S(40, 16), S(31, 33), S(89, -14), S(35, 35), S(5, 66)},
		{S(13, 99), S(-29, 48), S(-3, 30), S(33, -15), S(30, -5), S(-11, -50), S(-120, -9), S(-106, 83)},
	},
	{
		{},
		{S(-21, 24), S(-4, 15), S(-9, 27), S(-2, 4), S(-24, 28), S(-4, 4), S(21, 22), S(-6, 16)},
		{S(-18, 10), S(-10, 11), S(-9, -19), S(1, -1), S(-11, -17), S(-7, 10), S(-2, -5), S(-13, 16)},
		{S(-14, 27), S(-1, 17), S(14, 2), S(16, -35), S(32, -19), S(12, -14), S(8, 7), S(-15, 17)},
		{S(-7, 52), S(24, 29), S(2, -2), S(23, -1), S(32, -45), S(55, -21), S(25, -3), S(-6, 44)},
		{S(0, 110), S(16, 44), S(48, 33), S(24, 6), S(92, -17), S(61, 11), S(48, 31), S(8, 61)},
		{S(-13, 14), S(-17, 76), S(-50, 26), S(10, 26), S(8, -24), S(28, 10), S(-76, 21), S(-134, 64)},
	},
	{
		{},
		{S(-12, 27), S(-4, 32), S(-15, 31), S(-7, 16), S(-21, 19), S(0, 31), S(11, 20), S(-9, 16)},
		{S(-11, 22), S(-16, 21), S(-8, 21), S(-3, 9), S(-4, 12), S(-16, 37), S(-9, 18), S(-19, 15)},
		{S(-4, 29), S(0, 30), S(12, 13), S(19, 6), S(26, -8), S(13, 14), S(-1, 24), S(-20, 25)},
		{S(8, 58), S(22, 28), S(4, 38), S(24, -18), S(33, -9), S(41, -8), S(19, 27), S(-15, 41)},
		{S(5, 98), S(19, 72), S(16, 75), S(44, 8), S(59, -15), S(63, -1), S(28, 27), S(-9, 70)},
		{S(-88, 108), S(-74, 21), S(-36, 47), S(5, 1), S(15, -41), S(32, 3), S(-60, 16), S(-162, 25)},
	},
}

// Piece Square Values
var PieceScores = [King + 1][8][8]Score{
	{},
	{ // knight
		{S(-80, -87), S(-22, -46), S(-43, -31), S(-21, -5), S(-23, -12), S(-25, -34), S(-28, -33), S(-88, -44)},
		{S(-29, -31), S(-18, -20), S(-14, -37), S(-15, -11), S(-24, -13), S(-17, -32), S(-25, -17), S(-21, -22)},
		{S(-12, -45), S(-7, -20), S(-7, -12), S(-3, 9), S(-1, 7), S(-16, -12), S(-4, -24), S(-22, -46)},
		{S(-10, -5), S(-7, 6), S(8, 21), S(-4, 33), S(9, 29), S(1, 24), S(23, -5), S(-12, 0)},
		{S(-3, 8), S(-3, 6), S(9, 33), S(9, 51), S(8, 44), S(26, 27), S(7, 8), S(8, 5)},
		{S(-53, 7), S(-26, 17), S(-3, 40), S(7, 40), S(27, 28), S(28, 29), S(-13, 7), S(-20, -4)},
		{S(0, -20), S(-13, -5), S(24, -18), S(33, 17), S(40, 8), S(61, -40), S(-22, -7), S(14, -34)},
		{S(-202, -52), S(-98, -9), S(-153, 26), S(-30, -7), S(9, -8), S(-127, 32), S(-75, -10), S(-187, -68)},
	},
	{ // Bishop
		{S(33, -40), S(51, -13), S(7, 1), S(14, 2), S(0, 10), S(15, -9), S(13, -12), S(23, -40)},
		{S(45, -33), S(20, -35), S(45, -10), S(8, 5), S(18, -1), S(17, -15), S(32, -42), S(22, -65)},
		{S(22, -7), S(51, 0), S(7, -5), S(34, 14), S(16, 20), S(17, -14), S(31, -2), S(33, -8)},
		{S(16, -5), S(20, 8), S(32, 18), S(21, 27), S(42, 22), S(14, 20), S(27, 5), S(14, -8)},
		{S(-12, 20), S(29, 15), S(0, 31), S(41, 34), S(22, 37), S(28, 21), S(18, 17), S(3, 16)},
		{S(1, 14), S(-13, 33), S(13, 7), S(6, 34), S(27, 31), S(-2, 18), S(11, 30), S(-9, 20)},
		{S(-29, 32), S(-5, 9), S(12, 25), S(-11, 39), S(-6, 37), S(17, 27), S(-18, 6), S(-31, 31)},
		{S(-22, -1), S(-64, 39), S(-126, 50), S(-107, 59), S(-106, 53), S(-99, 46), S(22, 16), S(-40, 1)},
	},
	{ // Rook
		{S(-23, -28), S(-18, -17), S(-12, -15), S(-3, -26), S(-4, -29), S(-6, -16), S(-4, -27), S(-13, -45)},
		{S(-60, -21), S(-25, -32), S(-21, -26), S(-20, -29), S(-20, -31), S(-1, -40), S(3, -45), S(-58, -19)},
		{S(-43, -16), S(-28, -2), S(-34, -5), S(-21, -14), S(-23, -12), S(-16, -12), S(6, -19), S(-28, -23)},
		{S(-39, 9), S(-33, 25), S(-28, 24), S(-15, 11), S(-22, 12), S(-15, 15), S(-7, 10), S(-18, 2)},
		{S(-32, 34), S(-5, 30), S(4, 30), S(32, 18), S(9, 21), S(22, 17), S(20, 12), S(1, 20)},
		{S(-35, 46), S(21, 25), S(16, 36), S(41, 18), S(43, 16), S(73, 23), S(89, -3), S(11, 28)},
		{S(-25, 54), S(-18, 54), S(12, 46), S(35, 46), S(17, 51), S(71, 14), S(21, 32), S(35, 26)},
		{S(13, 44), S(29, 46), S(0, 55), S(9, 48), S(13, 48), S(62, 40), S(83, 33), S(68, 38)},
	},
	{ // Queen
		{S(7, -107), S(2, -82), S(6, -87), S(8, -51), S(10, -86), S(-11, -79), S(4, -122), S(2, -112)},
		{S(-3, -81), S(5, -62), S(12, -83), S(3, -44), S(4, -53), S(17, -99), S(24, -110), S(17, -95)},
		{S(-8, -60), S(6, -26), S(-1, 0), S(-6, -8), S(-4, -18), S(-1, 6), S(17, -37), S(8, -55)},
		{S(-11, -47), S(-6, 6), S(-11, 15), S(-20, 69), S(-21, 55), S(4, 41), S(8, 10), S(12, 10)},
		{S(-27, -20), S(-15, 20), S(-30, 24), S(-26, 77), S(-30, 100), S(-4, 94), S(19, 70), S(13, 52)},
		{S(-45, -4), S(-28, -3), S(-30, 21), S(-11, 39), S(7, 75), S(50, 67), S(58, 52), S(26, 79)},
		{S(-46, 16), S(-80, 44), S(-28, 22), S(-60, 72), S(-35, 106), S(24, 84), S(-33, 81), S(15, 28)},
		{S(-45, 20), S(-7, 10), S(-13, 46), S(12, 43), S(15, 58), S(60, 48), S(91, 19), S(64, 9)},
	},
	{ // King
		{S(142, -9), S(128, 26), S(54, 63), S(52, 69), S(80, 46), S(53, 65), S(150, 19), S(168, -33)},
		{S(156, 33), S(101, 55), S(40, 90), S(39, 111), S(30, 121), S(51, 87), S(126, 47), S(159, 30)},
		{S(113, 27), S(154, 44), S(93, 82), S(91, 119), S(74, 128), S(101, 80), S(141, 44), S(105, 34)},
		{S(142, 11), S(269, 31), S(161, 90), S(84, 140), S(89, 143), S(148, 91), S(225, 41), S(86, 32)},
		{S(131, 51), S(325, 45), S(156, 108), S(108, 139), S(87, 147), S(192, 101), S(268, 52), S(94, 53)},
		{S(139, 47), S(324, 67), S(225, 103), S(173, 113), S(176, 119), S(253, 93), S(301, 65), S(132, 46)},
		{S(136, -4), S(222, 62), S(212, 77), S(177, 89), S(127, 98), S(174, 77), S(203, 69), S(124, -3)},
		{S(180, -123), S(312, -42), S(201, 0), S(173, 53), S(154, 53), S(167, 11), S(287, -25), S(145, -100)},
	},
}

var PawnsConnected = [7][4]Score{
	{S(0, 0), S(0, 0), S(0, 0), S(0, 0)},
	{S(-1, -14), S(6, 4), S(3, 0), S(15, 16)},
	{S(11, 6), S(21, 4), S(23, 11), S(25, 17)},
	{S(9, 6), S(19, 10), S(10, 11), S(15, 20)},
	{S(3, 18), S(14, 25), S(26, 30), S(25, 24)},
	{S(35, 32), S(22, 70), S(68, 73), S(76, 92)},
	{S(176, 37), S(280, 23), S(275, 55), S(318, 56)},
}

var MobilityBonus = [...][32]Score{
	{S(-58, -126), S(-41, -71), S(-27, -23), S(-19, -1), S(-13, 13), S(-9, 28), // Knights
		S(0, 31), S(8, 26), S(18, 10)},
	{S(3, -130), S(9, -60), S(16, -18), S(23, 4), S(30, 19), S(34, 34), // Bishops
		S(36, 42), S(36, 46), S(36, 49), S(40, 48), S(44, 43), S(56, 33),
		S(75, 35), S(83, 9)},
	{S(-127, -156), S(-14, -36), S(-4, 18), S(-5, 46), S(-1, 58), S(1, 71), // Rooks
		S(2, 80), S(6, 88), S(10, 93), S(14, 99), S(17, 106), S(18, 112),
		S(24, 113), S(39, 99), S(96, 57)},
	{S(-413, -159), S(-130, -143), S(-47, -178), S(-24, -118), S(-9, -84), S(-9, -14), // Queens
		S(-5, 13), S(-2, 33), S(1, 50), S(4, 63), S(7, 71), S(11, 74),
		S(12, 77), S(14, 82), S(18, 77), S(15, 81), S(14, 77), S(13, 76),
		S(17, 65), S(24, 50), S(36, 29), S(39, 12), S(41, -7), S(54, -40),
		S(24, -29), S(-62, -11), S(140, -118), S(56, -79)},
}

var PassedFriendlyDistance = [8]Score{
	S(0, 0), S(-4, 39), S(-8, 24), S(-7, 9),
	S(-4, -5), S(-1, -15), S(15, -27), S(3, -39),
}

var PassedEnemyDistance = [8]Score{
	S(0, 0), S(-94, -41), S(-1, -12), S(5, 7),
	S(11, 19), S(11, 26), S(7, 35), S(17, 40),
}

// PassedRank[Rank] contains a bonus according to the rank of a passed pawn, whether it can be pushed and whether the push would be safe
var PassedRank = [2][2][2][7]Score{
	{
		{
			{S(0, 0), S(-47, -22), S(-39, 9), S(-26, 13), S(28, 13), S(43, 16), S(174, 88)},
			{S(0, 0), S(-31, -51), S(-30, -18), S(-18, 6), S(29, 27), S(59, 38), S(168, 131)},
		},
		{
			{S(0, 0), S(-16, -27), S(-27, 1), S(-17, 19), S(14, 45), S(58, 79), S(253, 139)},
			{S(0, 0), S(-5, -50), S(-24, -4), S(-18, 26), S(14, 60), S(76, 85), S(228, 176)},
		},
	},
	{
		{
			{S(0, 0), S(-20, -8), S(-31, -3), S(-12, 21), S(28, 47), S(64, 96), S(276, 179)},
			{S(0, 0), S(-13, -24), S(-19, -5), S(-10, 35), S(36, 65), S(74, 107), S(237, 230)},
		},
		{
			{S(0, 0), S(-29, -14), S(-27, -12), S(-14, 31), S(19, 86), S(33, 234), S(75, 404)},
			{S(0, 0), S(-12, -35), S(-17, -16), S(-10, 33), S(22, 94), S(58, 221), S(138, 418)},
		},
	},
}

// PassedFile[File] contains a bonus according to the file of a passed pawn
var PassedFile = [8]Score{S(0, 24), S(-3, 29), S(-1, 11), S(-4, -3),
	S(-11, 3), S(-19, 13), S(-17, 25), S(9, 10),
}

var PassedStacked = [7]Score{S(0, 0), S(-17, -50), S(-25, -31), S(-35, -40), S(-17, -44), S(43, -99), S(0, 0)}
var PassedUncontested = [6]Score{S(0, 0), S(-90, 46), S(-80, 43), S(-90, 47), S(-85, 48), S(-68, 36)}
var PassedPushDefended = [6]Score{S(0, 0), S(-3, 20), S(0, 7), S(6, 1), S(-3, 6), S(-14, 15)}
var PassedPushUncontestedDefended = [6]Score{S(0, 0), S(-57, 30), S(-36, 23), S(-65, 46), S(-72, 56), S(-57, 69)}

var Isolated = S(-7, -17)
var Doubled = S(-10, -32)
var Backward = S(8, -3)
var BackwardOpen = S(-3, -16)

var BishopPair = S(16, 77)
var BishopRammedPawns = S(-8, -23)

var BishopOutpostUndefendedBonus = S(22, -4)
var BishopOutpostDefendedBonus = S(52, 13)

var LongDiagonalBishop = S(21, 25)
var DistantBishop = [4]Score{S(-6, 2), S(-11, -1), S(-13, -4), S(-18, -21)}

var KnightOutpostUndefendedBonus = S(21, -22)
var KnightOutpostDefendedBonus = S(44, 15)

var DistantKnight = [4]Score{S(-15, -2), S(-16, -14), S(-28, -17), S(-54, -21)}

var MinorBehindPawn = S(7, 29)

var Tempo int16 = 15

// Rook on semiopen, open file
var RookOnFile = [2]Score{S(11, 10), S(37, 5)}
var RookOnQueenFile = S(5, 30)

var KingDefenders = [12]Score{
	S(-18, -15), S(-2, -11), S(1, -7), S(4, -4),
	S(6, -1), S(9, 5), S(12, 11), S(15, 10),
	S(14, 14), S(15, -62), S(-16, -12), S(11, 0),
}

var KingShelter = [2][8][8]Score{
	{{S(-25, 9), S(11, -13), S(18, 0), S(23, 5),
		S(29, -14), S(20, -6), S(17, -38), S(-92, 39)},
		{S(17, 6), S(25, -8), S(-5, 10), S(-7, 8),
			S(2, 0), S(23, 3), S(41, -14), S(-46, 23)},
		{S(15, 3), S(4, -1), S(-25, 0), S(-21, -1),
			S(1, -17), S(0, 4), S(19, -8), S(-31, 0)},
		{S(-19, 16), S(8, -2), S(-6, -9), S(-1, -9),
			S(6, -24), S(1, -9), S(17, 20), S(-15, -3)},
		{S(-32, 20), S(-16, 9), S(-31, 9), S(-21, 9),
			S(0, -8), S(-23, 1), S(8, -3), S(-27, 2)},
		{S(38, -18), S(23, -14), S(5, -12), S(5, -21),
			S(13, -28), S(25, -20), S(31, -20), S(-12, -4)},
		{S(16, -2), S(-7, -7), S(-27, 4), S(-18, 4),
			S(-8, -15), S(8, -4), S(15, -18), S(-24, 19)},
		{S(-41, 11), S(-26, -2), S(-20, 27), S(-20, 24),
			S(-22, 15), S(-11, 11), S(-11, -16), S(-71, 48)}},
	{{S(34, 56), S(-35, -12), S(-24, 4), S(-30, -4),
		S(-41, -15), S(-5, 6), S(-42, -16), S(-90, 39)},
		{S(138, 18), S(7, -5), S(-5, 5), S(-23, 19),
			S(-6, -11), S(12, 3), S(13, 0), S(-80, 25)},
		{S(-4, 32), S(34, 3), S(13, -1), S(20, -5),
			S(27, -5), S(12, 2), S(37, 4), S(-24, 6)},
		{S(6, 34), S(-13, 21), S(-15, 15), S(-23, 22),
			S(-1, 10), S(-5, 3), S(-5, 11), S(-43, -4)},
		{S(4, 23), S(-1, 9), S(-7, 5), S(-11, 1),
			S(-4, 0), S(-1, -12), S(12, -1), S(-30, -4)},
		{S(17, 19), S(-6, 7), S(-14, 10), S(-8, 9),
			S(-3, 2), S(-38, 11), S(-9, 2), S(-39, 9)},
		{S(31, 4), S(-4, -15), S(-9, -12), S(-25, -6),
			S(-9, -19), S(0, -17), S(2, -20), S(-70, 21)},
		{S(-27, -3), S(-14, -33), S(-8, -19), S(-3, -23),
			S(-4, -36), S(-1, -23), S(-5, -52), S(-65, 27)}},
}

var KingStorm = [2][4][8]Score{
	{{S(11, 8), S(2, 15), S(8, 7), S(5, 8),
		S(4, 5), S(11, 0), S(8, 8), S(-11, -4)},
		{S(23, 0), S(16, 6), S(19, 3), S(6, 11),
			S(20, 2), S(27, -9), S(19, -7), S(-7, -3)},
		{S(21, 9), S(3, 6), S(5, 10), S(0, 11),
			S(4, 7), S(11, 1), S(8, -9), S(-2, 7)},
		{S(10, 1), S(7, -2), S(10, -11), S(0, -13),
			S(-2, -8), S(15, -16), S(9, -13), S(-11, 5)}},
	{{S(0, 0), S(3, 4), S(-16, 8), S(10, -7),
		S(-2, 20), S(-12, 32), S(5, 50), S(-6, -4)},
		{S(0, 0), S(-53, -10), S(7, -6), S(46, -7),
			S(5, 1), S(-7, -4), S(15, 54), S(-7, 1)},
		{S(0, 0), S(-19, -12), S(-5, -6), S(17, -2),
			S(6, -5), S(-11, -17), S(37, -39), S(-6, 8)},
		{S(0, 0), S(-18, -25), S(-15, -25), S(-7, -11),
			S(2, -24), S(0, -54), S(4, -2), S(-11, -4)}},
}
var KingOnPawnlessFlank = S(-5, -74)

var Hanging = S(34, 17)
var ThreatByKing = S(-10, 38)
var ThreatByMinor = [King + 1]Score{S(0, 0), S(18, 40), S(16, 39), S(32, 29), S(30, 29), S(1, 27)}
var ThreatByRook = [King + 1]Score{S(0, 0), S(-4, 14), S(-2, 19), S(-4, -11), S(33, 9), S(13, 10)}

// This weights are from black piece on black square perspective
var RookBishopExistence = [16]Score{
	S(25, -19), S(4, -10), S(5, -10), S(2, 1), S(0, -8), S(-5, 37), S(-1, -2), S(-5, 20), S(2, -8), S(-2, -3), S(-4, 37), S(-5, 21), S(-9, -19), S(-4, 18), S(-2, 18), S(-21, 56),
}
var QueenBishopExistence = [16]Score{
	S(85, -66), S(-5, -9), S(-3, -23), S(-23, -6), S(-5, -10), S(10, 78), S(-18, 0), S(0, -2), S(-1, -6), S(-14, 25), S(12, 74), S(4, 8), S(-38, -15), S(-8, 37), S(-12, 29), S(-13, 36),
}
var KingBishopExistence = [16]Score{
	S(0, 0), S(0, 4), S(1, 5), S(-18, -8), S(-1, -10), S(-2, -6), S(0, -3), S(0, 5), S(1, -2), S(0, 3), S(2, 6), S(2, 14), S(18, 8), S(0, -12), S(2, -9), S(0, 0),
}

//
// King safety
//

var KingSafetyAttacksWeights = [Queen + 1]Score{S(0, 12), S(-2, 3), S(-6, 5), S(-4, 3), S(1, 15)}
var KingSafetyWeakSquares = S(41, 40)
var KingSafetyFriendlyPawns = S(-34, -24)
var KingSafetyNoEnemyQueens = S(-173, -172)
var KingSafetySafeQueenCheck = S(89, 92)
var KingSafetySafeRookCheck = S(76, 77)
var KingSafetySafeBishopCheck = S(50, 51)
var KingSafetySafeKnightCheck = S(113, 110)
var KingSafetyAdjustment = S(-17, -5)

// Attack value is special as it is scaled by a fraction
var KingSafetyAttackValue = S(120, 132)

//
// Complexity
//

var ComplexityTotalPawns = S(0, 6)
var ComplexityPawnEndgame = S(0, 83)
var ComplexityPawnBothFlanks = S(0, 106)
var ComplexityInfiltration = S(0, 6)
var ComplexityAdjustment = S(0, -173)
