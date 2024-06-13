// Code generated from Cvlt.g4 by ANTLR 4.13.1. DO NOT EDIT.

package cvlt

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type CvltLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var CvltLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func cvltlexerLexerInit() {
	staticData := &CvltLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'.'", "'List'", "'<'", "'>'", "'Interval'", "'Tuple'", "'{'", "','",
		"'}'", "'Choice'", "':'", "'true'", "'false'", "'null'", "'['", "'('",
		"']'", "')'", "'display'", "'Code'", "'from'", "'Concept'", "'year'",
		"'month'", "'week'", "'day'", "'hour'", "'minute'", "'second'", "'millisecond'",
		"'years'", "'months'", "'weeks'", "'days'", "'hours'", "'minutes'",
		"'seconds'", "'milliseconds'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "DATE", "DATETIME", "TIME", "IDENTIFIER", "DELIMITEDIDENTIFIER",
		"QUOTEDIDENTIFIER", "STRING", "NUMBER", "LONGNUMBER", "WS", "COMMENT",
		"LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
		"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32",
		"T__33", "T__34", "T__35", "T__36", "T__37", "DATE", "DATETIME", "TIME",
		"DATEFORMAT", "TIMEFORMAT", "TIMEZONEOFFSETFORMAT", "IDENTIFIER", "DELIMITEDIDENTIFIER",
		"QUOTEDIDENTIFIER", "STRING", "NUMBER", "LONGNUMBER", "WS", "COMMENT",
		"LINE_COMMENT", "ESC", "UNICODE", "HEX",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 50, 483, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 1, 0, 1, 0, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6,
		1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15,
		1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27,
		1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30,
		1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33,
		1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1,
		35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36,
		1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1,
		37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39,
		1, 39, 3, 39, 322, 8, 39, 1, 39, 3, 39, 325, 8, 39, 1, 40, 1, 40, 1, 40,
		1, 40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1,
		41, 3, 41, 341, 8, 41, 3, 41, 343, 8, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1,
		42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 4, 42, 355, 8, 42, 11, 42, 12, 42,
		356, 3, 42, 359, 8, 42, 3, 42, 361, 8, 42, 3, 42, 363, 8, 42, 1, 43, 1,
		43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 3, 43, 372, 8, 43, 1, 44, 3, 44,
		375, 8, 44, 1, 44, 5, 44, 378, 8, 44, 10, 44, 12, 44, 381, 9, 44, 1, 45,
		1, 45, 1, 45, 5, 45, 386, 8, 45, 10, 45, 12, 45, 389, 9, 45, 1, 45, 1,
		45, 1, 46, 1, 46, 1, 46, 5, 46, 396, 8, 46, 10, 46, 12, 46, 399, 9, 46,
		1, 46, 1, 46, 1, 47, 1, 47, 1, 47, 5, 47, 406, 8, 47, 10, 47, 12, 47, 409,
		9, 47, 1, 47, 1, 47, 1, 48, 3, 48, 414, 8, 48, 1, 48, 4, 48, 417, 8, 48,
		11, 48, 12, 48, 418, 1, 48, 1, 48, 4, 48, 423, 8, 48, 11, 48, 12, 48, 424,
		3, 48, 427, 8, 48, 1, 49, 3, 49, 430, 8, 49, 1, 49, 4, 49, 433, 8, 49,
		11, 49, 12, 49, 434, 1, 49, 1, 49, 1, 50, 4, 50, 440, 8, 50, 11, 50, 12,
		50, 441, 1, 50, 1, 50, 1, 51, 1, 51, 1, 51, 1, 51, 5, 51, 450, 8, 51, 10,
		51, 12, 51, 453, 9, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 52, 1, 52,
		1, 52, 1, 52, 5, 52, 464, 8, 52, 10, 52, 12, 52, 467, 9, 52, 1, 52, 1,
		52, 1, 53, 1, 53, 1, 53, 3, 53, 474, 8, 53, 1, 54, 1, 54, 1, 54, 1, 54,
		1, 54, 1, 54, 1, 55, 1, 55, 4, 387, 397, 407, 451, 0, 56, 1, 1, 3, 2, 5,
		3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25,
		13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43,
		22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61,
		31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79,
		40, 81, 41, 83, 0, 85, 0, 87, 0, 89, 42, 91, 43, 93, 44, 95, 45, 97, 46,
		99, 47, 101, 48, 103, 49, 105, 50, 107, 0, 109, 0, 111, 0, 1, 0, 8, 1,
		0, 48, 57, 2, 0, 43, 43, 45, 45, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48,
		57, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 10, 10,
		13, 13, 9, 0, 34, 34, 39, 39, 47, 47, 92, 92, 96, 96, 102, 102, 110, 110,
		114, 114, 116, 116, 3, 0, 48, 57, 65, 70, 97, 102, 502, 0, 1, 1, 0, 0,
		0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0,
		0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0,
		0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1,
		0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33,
		1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0,
		41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0,
		0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0,
		0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0,
		0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1,
		0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79,
		1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0,
		93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0,
		0, 101, 1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105, 1, 0, 0, 0, 1, 113, 1,
		0, 0, 0, 3, 115, 1, 0, 0, 0, 5, 120, 1, 0, 0, 0, 7, 122, 1, 0, 0, 0, 9,
		124, 1, 0, 0, 0, 11, 133, 1, 0, 0, 0, 13, 139, 1, 0, 0, 0, 15, 141, 1,
		0, 0, 0, 17, 143, 1, 0, 0, 0, 19, 145, 1, 0, 0, 0, 21, 152, 1, 0, 0, 0,
		23, 154, 1, 0, 0, 0, 25, 159, 1, 0, 0, 0, 27, 165, 1, 0, 0, 0, 29, 170,
		1, 0, 0, 0, 31, 172, 1, 0, 0, 0, 33, 174, 1, 0, 0, 0, 35, 176, 1, 0, 0,
		0, 37, 178, 1, 0, 0, 0, 39, 186, 1, 0, 0, 0, 41, 191, 1, 0, 0, 0, 43, 196,
		1, 0, 0, 0, 45, 204, 1, 0, 0, 0, 47, 209, 1, 0, 0, 0, 49, 215, 1, 0, 0,
		0, 51, 220, 1, 0, 0, 0, 53, 224, 1, 0, 0, 0, 55, 229, 1, 0, 0, 0, 57, 236,
		1, 0, 0, 0, 59, 243, 1, 0, 0, 0, 61, 255, 1, 0, 0, 0, 63, 261, 1, 0, 0,
		0, 65, 268, 1, 0, 0, 0, 67, 274, 1, 0, 0, 0, 69, 279, 1, 0, 0, 0, 71, 285,
		1, 0, 0, 0, 73, 293, 1, 0, 0, 0, 75, 301, 1, 0, 0, 0, 77, 314, 1, 0, 0,
		0, 79, 317, 1, 0, 0, 0, 81, 326, 1, 0, 0, 0, 83, 330, 1, 0, 0, 0, 85, 344,
		1, 0, 0, 0, 87, 371, 1, 0, 0, 0, 89, 374, 1, 0, 0, 0, 91, 382, 1, 0, 0,
		0, 93, 392, 1, 0, 0, 0, 95, 402, 1, 0, 0, 0, 97, 413, 1, 0, 0, 0, 99, 429,
		1, 0, 0, 0, 101, 439, 1, 0, 0, 0, 103, 445, 1, 0, 0, 0, 105, 459, 1, 0,
		0, 0, 107, 470, 1, 0, 0, 0, 109, 475, 1, 0, 0, 0, 111, 481, 1, 0, 0, 0,
		113, 114, 5, 46, 0, 0, 114, 2, 1, 0, 0, 0, 115, 116, 5, 76, 0, 0, 116,
		117, 5, 105, 0, 0, 117, 118, 5, 115, 0, 0, 118, 119, 5, 116, 0, 0, 119,
		4, 1, 0, 0, 0, 120, 121, 5, 60, 0, 0, 121, 6, 1, 0, 0, 0, 122, 123, 5,
		62, 0, 0, 123, 8, 1, 0, 0, 0, 124, 125, 5, 73, 0, 0, 125, 126, 5, 110,
		0, 0, 126, 127, 5, 116, 0, 0, 127, 128, 5, 101, 0, 0, 128, 129, 5, 114,
		0, 0, 129, 130, 5, 118, 0, 0, 130, 131, 5, 97, 0, 0, 131, 132, 5, 108,
		0, 0, 132, 10, 1, 0, 0, 0, 133, 134, 5, 84, 0, 0, 134, 135, 5, 117, 0,
		0, 135, 136, 5, 112, 0, 0, 136, 137, 5, 108, 0, 0, 137, 138, 5, 101, 0,
		0, 138, 12, 1, 0, 0, 0, 139, 140, 5, 123, 0, 0, 140, 14, 1, 0, 0, 0, 141,
		142, 5, 44, 0, 0, 142, 16, 1, 0, 0, 0, 143, 144, 5, 125, 0, 0, 144, 18,
		1, 0, 0, 0, 145, 146, 5, 67, 0, 0, 146, 147, 5, 104, 0, 0, 147, 148, 5,
		111, 0, 0, 148, 149, 5, 105, 0, 0, 149, 150, 5, 99, 0, 0, 150, 151, 5,
		101, 0, 0, 151, 20, 1, 0, 0, 0, 152, 153, 5, 58, 0, 0, 153, 22, 1, 0, 0,
		0, 154, 155, 5, 116, 0, 0, 155, 156, 5, 114, 0, 0, 156, 157, 5, 117, 0,
		0, 157, 158, 5, 101, 0, 0, 158, 24, 1, 0, 0, 0, 159, 160, 5, 102, 0, 0,
		160, 161, 5, 97, 0, 0, 161, 162, 5, 108, 0, 0, 162, 163, 5, 115, 0, 0,
		163, 164, 5, 101, 0, 0, 164, 26, 1, 0, 0, 0, 165, 166, 5, 110, 0, 0, 166,
		167, 5, 117, 0, 0, 167, 168, 5, 108, 0, 0, 168, 169, 5, 108, 0, 0, 169,
		28, 1, 0, 0, 0, 170, 171, 5, 91, 0, 0, 171, 30, 1, 0, 0, 0, 172, 173, 5,
		40, 0, 0, 173, 32, 1, 0, 0, 0, 174, 175, 5, 93, 0, 0, 175, 34, 1, 0, 0,
		0, 176, 177, 5, 41, 0, 0, 177, 36, 1, 0, 0, 0, 178, 179, 5, 100, 0, 0,
		179, 180, 5, 105, 0, 0, 180, 181, 5, 115, 0, 0, 181, 182, 5, 112, 0, 0,
		182, 183, 5, 108, 0, 0, 183, 184, 5, 97, 0, 0, 184, 185, 5, 121, 0, 0,
		185, 38, 1, 0, 0, 0, 186, 187, 5, 67, 0, 0, 187, 188, 5, 111, 0, 0, 188,
		189, 5, 100, 0, 0, 189, 190, 5, 101, 0, 0, 190, 40, 1, 0, 0, 0, 191, 192,
		5, 102, 0, 0, 192, 193, 5, 114, 0, 0, 193, 194, 5, 111, 0, 0, 194, 195,
		5, 109, 0, 0, 195, 42, 1, 0, 0, 0, 196, 197, 5, 67, 0, 0, 197, 198, 5,
		111, 0, 0, 198, 199, 5, 110, 0, 0, 199, 200, 5, 99, 0, 0, 200, 201, 5,
		101, 0, 0, 201, 202, 5, 112, 0, 0, 202, 203, 5, 116, 0, 0, 203, 44, 1,
		0, 0, 0, 204, 205, 5, 121, 0, 0, 205, 206, 5, 101, 0, 0, 206, 207, 5, 97,
		0, 0, 207, 208, 5, 114, 0, 0, 208, 46, 1, 0, 0, 0, 209, 210, 5, 109, 0,
		0, 210, 211, 5, 111, 0, 0, 211, 212, 5, 110, 0, 0, 212, 213, 5, 116, 0,
		0, 213, 214, 5, 104, 0, 0, 214, 48, 1, 0, 0, 0, 215, 216, 5, 119, 0, 0,
		216, 217, 5, 101, 0, 0, 217, 218, 5, 101, 0, 0, 218, 219, 5, 107, 0, 0,
		219, 50, 1, 0, 0, 0, 220, 221, 5, 100, 0, 0, 221, 222, 5, 97, 0, 0, 222,
		223, 5, 121, 0, 0, 223, 52, 1, 0, 0, 0, 224, 225, 5, 104, 0, 0, 225, 226,
		5, 111, 0, 0, 226, 227, 5, 117, 0, 0, 227, 228, 5, 114, 0, 0, 228, 54,
		1, 0, 0, 0, 229, 230, 5, 109, 0, 0, 230, 231, 5, 105, 0, 0, 231, 232, 5,
		110, 0, 0, 232, 233, 5, 117, 0, 0, 233, 234, 5, 116, 0, 0, 234, 235, 5,
		101, 0, 0, 235, 56, 1, 0, 0, 0, 236, 237, 5, 115, 0, 0, 237, 238, 5, 101,
		0, 0, 238, 239, 5, 99, 0, 0, 239, 240, 5, 111, 0, 0, 240, 241, 5, 110,
		0, 0, 241, 242, 5, 100, 0, 0, 242, 58, 1, 0, 0, 0, 243, 244, 5, 109, 0,
		0, 244, 245, 5, 105, 0, 0, 245, 246, 5, 108, 0, 0, 246, 247, 5, 108, 0,
		0, 247, 248, 5, 105, 0, 0, 248, 249, 5, 115, 0, 0, 249, 250, 5, 101, 0,
		0, 250, 251, 5, 99, 0, 0, 251, 252, 5, 111, 0, 0, 252, 253, 5, 110, 0,
		0, 253, 254, 5, 100, 0, 0, 254, 60, 1, 0, 0, 0, 255, 256, 5, 121, 0, 0,
		256, 257, 5, 101, 0, 0, 257, 258, 5, 97, 0, 0, 258, 259, 5, 114, 0, 0,
		259, 260, 5, 115, 0, 0, 260, 62, 1, 0, 0, 0, 261, 262, 5, 109, 0, 0, 262,
		263, 5, 111, 0, 0, 263, 264, 5, 110, 0, 0, 264, 265, 5, 116, 0, 0, 265,
		266, 5, 104, 0, 0, 266, 267, 5, 115, 0, 0, 267, 64, 1, 0, 0, 0, 268, 269,
		5, 119, 0, 0, 269, 270, 5, 101, 0, 0, 270, 271, 5, 101, 0, 0, 271, 272,
		5, 107, 0, 0, 272, 273, 5, 115, 0, 0, 273, 66, 1, 0, 0, 0, 274, 275, 5,
		100, 0, 0, 275, 276, 5, 97, 0, 0, 276, 277, 5, 121, 0, 0, 277, 278, 5,
		115, 0, 0, 278, 68, 1, 0, 0, 0, 279, 280, 5, 104, 0, 0, 280, 281, 5, 111,
		0, 0, 281, 282, 5, 117, 0, 0, 282, 283, 5, 114, 0, 0, 283, 284, 5, 115,
		0, 0, 284, 70, 1, 0, 0, 0, 285, 286, 5, 109, 0, 0, 286, 287, 5, 105, 0,
		0, 287, 288, 5, 110, 0, 0, 288, 289, 5, 117, 0, 0, 289, 290, 5, 116, 0,
		0, 290, 291, 5, 101, 0, 0, 291, 292, 5, 115, 0, 0, 292, 72, 1, 0, 0, 0,
		293, 294, 5, 115, 0, 0, 294, 295, 5, 101, 0, 0, 295, 296, 5, 99, 0, 0,
		296, 297, 5, 111, 0, 0, 297, 298, 5, 110, 0, 0, 298, 299, 5, 100, 0, 0,
		299, 300, 5, 115, 0, 0, 300, 74, 1, 0, 0, 0, 301, 302, 5, 109, 0, 0, 302,
		303, 5, 105, 0, 0, 303, 304, 5, 108, 0, 0, 304, 305, 5, 108, 0, 0, 305,
		306, 5, 105, 0, 0, 306, 307, 5, 115, 0, 0, 307, 308, 5, 101, 0, 0, 308,
		309, 5, 99, 0, 0, 309, 310, 5, 111, 0, 0, 310, 311, 5, 110, 0, 0, 311,
		312, 5, 100, 0, 0, 312, 313, 5, 115, 0, 0, 313, 76, 1, 0, 0, 0, 314, 315,
		5, 64, 0, 0, 315, 316, 3, 83, 41, 0, 316, 78, 1, 0, 0, 0, 317, 318, 5,
		64, 0, 0, 318, 319, 3, 83, 41, 0, 319, 321, 5, 84, 0, 0, 320, 322, 3, 85,
		42, 0, 321, 320, 1, 0, 0, 0, 321, 322, 1, 0, 0, 0, 322, 324, 1, 0, 0, 0,
		323, 325, 3, 87, 43, 0, 324, 323, 1, 0, 0, 0, 324, 325, 1, 0, 0, 0, 325,
		80, 1, 0, 0, 0, 326, 327, 5, 64, 0, 0, 327, 328, 5, 84, 0, 0, 328, 329,
		3, 85, 42, 0, 329, 82, 1, 0, 0, 0, 330, 331, 7, 0, 0, 0, 331, 332, 7, 0,
		0, 0, 332, 333, 7, 0, 0, 0, 333, 342, 7, 0, 0, 0, 334, 335, 5, 45, 0, 0,
		335, 336, 7, 0, 0, 0, 336, 340, 7, 0, 0, 0, 337, 338, 5, 45, 0, 0, 338,
		339, 7, 0, 0, 0, 339, 341, 7, 0, 0, 0, 340, 337, 1, 0, 0, 0, 340, 341,
		1, 0, 0, 0, 341, 343, 1, 0, 0, 0, 342, 334, 1, 0, 0, 0, 342, 343, 1, 0,
		0, 0, 343, 84, 1, 0, 0, 0, 344, 345, 7, 0, 0, 0, 345, 362, 7, 0, 0, 0,
		346, 347, 5, 58, 0, 0, 347, 348, 7, 0, 0, 0, 348, 360, 7, 0, 0, 0, 349,
		350, 5, 58, 0, 0, 350, 351, 7, 0, 0, 0, 351, 358, 7, 0, 0, 0, 352, 354,
		5, 46, 0, 0, 353, 355, 7, 0, 0, 0, 354, 353, 1, 0, 0, 0, 355, 356, 1, 0,
		0, 0, 356, 354, 1, 0, 0, 0, 356, 357, 1, 0, 0, 0, 357, 359, 1, 0, 0, 0,
		358, 352, 1, 0, 0, 0, 358, 359, 1, 0, 0, 0, 359, 361, 1, 0, 0, 0, 360,
		349, 1, 0, 0, 0, 360, 361, 1, 0, 0, 0, 361, 363, 1, 0, 0, 0, 362, 346,
		1, 0, 0, 0, 362, 363, 1, 0, 0, 0, 363, 86, 1, 0, 0, 0, 364, 372, 5, 90,
		0, 0, 365, 366, 7, 1, 0, 0, 366, 367, 7, 0, 0, 0, 367, 368, 7, 0, 0, 0,
		368, 369, 5, 58, 0, 0, 369, 370, 7, 0, 0, 0, 370, 372, 7, 0, 0, 0, 371,
		364, 1, 0, 0, 0, 371, 365, 1, 0, 0, 0, 372, 88, 1, 0, 0, 0, 373, 375, 7,
		2, 0, 0, 374, 373, 1, 0, 0, 0, 375, 379, 1, 0, 0, 0, 376, 378, 7, 3, 0,
		0, 377, 376, 1, 0, 0, 0, 378, 381, 1, 0, 0, 0, 379, 377, 1, 0, 0, 0, 379,
		380, 1, 0, 0, 0, 380, 90, 1, 0, 0, 0, 381, 379, 1, 0, 0, 0, 382, 387, 5,
		96, 0, 0, 383, 386, 3, 107, 53, 0, 384, 386, 9, 0, 0, 0, 385, 383, 1, 0,
		0, 0, 385, 384, 1, 0, 0, 0, 386, 389, 1, 0, 0, 0, 387, 388, 1, 0, 0, 0,
		387, 385, 1, 0, 0, 0, 388, 390, 1, 0, 0, 0, 389, 387, 1, 0, 0, 0, 390,
		391, 5, 96, 0, 0, 391, 92, 1, 0, 0, 0, 392, 397, 5, 34, 0, 0, 393, 396,
		3, 107, 53, 0, 394, 396, 9, 0, 0, 0, 395, 393, 1, 0, 0, 0, 395, 394, 1,
		0, 0, 0, 396, 399, 1, 0, 0, 0, 397, 398, 1, 0, 0, 0, 397, 395, 1, 0, 0,
		0, 398, 400, 1, 0, 0, 0, 399, 397, 1, 0, 0, 0, 400, 401, 5, 34, 0, 0, 401,
		94, 1, 0, 0, 0, 402, 407, 5, 39, 0, 0, 403, 406, 3, 107, 53, 0, 404, 406,
		9, 0, 0, 0, 405, 403, 1, 0, 0, 0, 405, 404, 1, 0, 0, 0, 406, 409, 1, 0,
		0, 0, 407, 408, 1, 0, 0, 0, 407, 405, 1, 0, 0, 0, 408, 410, 1, 0, 0, 0,
		409, 407, 1, 0, 0, 0, 410, 411, 5, 39, 0, 0, 411, 96, 1, 0, 0, 0, 412,
		414, 7, 1, 0, 0, 413, 412, 1, 0, 0, 0, 413, 414, 1, 0, 0, 0, 414, 416,
		1, 0, 0, 0, 415, 417, 7, 0, 0, 0, 416, 415, 1, 0, 0, 0, 417, 418, 1, 0,
		0, 0, 418, 416, 1, 0, 0, 0, 418, 419, 1, 0, 0, 0, 419, 426, 1, 0, 0, 0,
		420, 422, 5, 46, 0, 0, 421, 423, 7, 0, 0, 0, 422, 421, 1, 0, 0, 0, 423,
		424, 1, 0, 0, 0, 424, 422, 1, 0, 0, 0, 424, 425, 1, 0, 0, 0, 425, 427,
		1, 0, 0, 0, 426, 420, 1, 0, 0, 0, 426, 427, 1, 0, 0, 0, 427, 98, 1, 0,
		0, 0, 428, 430, 7, 1, 0, 0, 429, 428, 1, 0, 0, 0, 429, 430, 1, 0, 0, 0,
		430, 432, 1, 0, 0, 0, 431, 433, 7, 0, 0, 0, 432, 431, 1, 0, 0, 0, 433,
		434, 1, 0, 0, 0, 434, 432, 1, 0, 0, 0, 434, 435, 1, 0, 0, 0, 435, 436,
		1, 0, 0, 0, 436, 437, 5, 76, 0, 0, 437, 100, 1, 0, 0, 0, 438, 440, 7, 4,
		0, 0, 439, 438, 1, 0, 0, 0, 440, 441, 1, 0, 0, 0, 441, 439, 1, 0, 0, 0,
		441, 442, 1, 0, 0, 0, 442, 443, 1, 0, 0, 0, 443, 444, 6, 50, 0, 0, 444,
		102, 1, 0, 0, 0, 445, 446, 5, 47, 0, 0, 446, 447, 5, 42, 0, 0, 447, 451,
		1, 0, 0, 0, 448, 450, 9, 0, 0, 0, 449, 448, 1, 0, 0, 0, 450, 453, 1, 0,
		0, 0, 451, 452, 1, 0, 0, 0, 451, 449, 1, 0, 0, 0, 452, 454, 1, 0, 0, 0,
		453, 451, 1, 0, 0, 0, 454, 455, 5, 42, 0, 0, 455, 456, 5, 47, 0, 0, 456,
		457, 1, 0, 0, 0, 457, 458, 6, 51, 0, 0, 458, 104, 1, 0, 0, 0, 459, 460,
		5, 47, 0, 0, 460, 461, 5, 47, 0, 0, 461, 465, 1, 0, 0, 0, 462, 464, 8,
		5, 0, 0, 463, 462, 1, 0, 0, 0, 464, 467, 1, 0, 0, 0, 465, 463, 1, 0, 0,
		0, 465, 466, 1, 0, 0, 0, 466, 468, 1, 0, 0, 0, 467, 465, 1, 0, 0, 0, 468,
		469, 6, 52, 0, 0, 469, 106, 1, 0, 0, 0, 470, 473, 5, 92, 0, 0, 471, 474,
		7, 6, 0, 0, 472, 474, 3, 109, 54, 0, 473, 471, 1, 0, 0, 0, 473, 472, 1,
		0, 0, 0, 474, 108, 1, 0, 0, 0, 475, 476, 5, 117, 0, 0, 476, 477, 3, 111,
		55, 0, 477, 478, 3, 111, 55, 0, 478, 479, 3, 111, 55, 0, 479, 480, 3, 111,
		55, 0, 480, 110, 1, 0, 0, 0, 481, 482, 7, 7, 0, 0, 482, 112, 1, 0, 0, 0,
		29, 0, 321, 324, 340, 342, 356, 358, 360, 362, 371, 374, 377, 379, 385,
		387, 395, 397, 405, 407, 413, 418, 424, 426, 429, 434, 441, 451, 465, 473,
		1, 0, 1, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// CvltLexerInit initializes any static state used to implement CvltLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewCvltLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func CvltLexerInit() {
	staticData := &CvltLexerLexerStaticData
	staticData.once.Do(cvltlexerLexerInit)
}

// NewCvltLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewCvltLexer(input antlr.CharStream) *CvltLexer {
	CvltLexerInit()
	l := new(CvltLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &CvltLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Cvlt.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CvltLexer tokens.
const (
	CvltLexerT__0                = 1
	CvltLexerT__1                = 2
	CvltLexerT__2                = 3
	CvltLexerT__3                = 4
	CvltLexerT__4                = 5
	CvltLexerT__5                = 6
	CvltLexerT__6                = 7
	CvltLexerT__7                = 8
	CvltLexerT__8                = 9
	CvltLexerT__9                = 10
	CvltLexerT__10               = 11
	CvltLexerT__11               = 12
	CvltLexerT__12               = 13
	CvltLexerT__13               = 14
	CvltLexerT__14               = 15
	CvltLexerT__15               = 16
	CvltLexerT__16               = 17
	CvltLexerT__17               = 18
	CvltLexerT__18               = 19
	CvltLexerT__19               = 20
	CvltLexerT__20               = 21
	CvltLexerT__21               = 22
	CvltLexerT__22               = 23
	CvltLexerT__23               = 24
	CvltLexerT__24               = 25
	CvltLexerT__25               = 26
	CvltLexerT__26               = 27
	CvltLexerT__27               = 28
	CvltLexerT__28               = 29
	CvltLexerT__29               = 30
	CvltLexerT__30               = 31
	CvltLexerT__31               = 32
	CvltLexerT__32               = 33
	CvltLexerT__33               = 34
	CvltLexerT__34               = 35
	CvltLexerT__35               = 36
	CvltLexerT__36               = 37
	CvltLexerT__37               = 38
	CvltLexerDATE                = 39
	CvltLexerDATETIME            = 40
	CvltLexerTIME                = 41
	CvltLexerIDENTIFIER          = 42
	CvltLexerDELIMITEDIDENTIFIER = 43
	CvltLexerQUOTEDIDENTIFIER    = 44
	CvltLexerSTRING              = 45
	CvltLexerNUMBER              = 46
	CvltLexerLONGNUMBER          = 47
	CvltLexerWS                  = 48
	CvltLexerCOMMENT             = 49
	CvltLexerLINE_COMMENT        = 50
)
