/**
 * @author Co2
 * @file_name static.go
 */

package main

import "fmt"

var (
	LuKong             = " "
	LuZhanWei          = "LNF"
	LuAnd              = "&&"
	LuOr               = "||"
	LuNon              = "!="
	LuLeftParenthesis  = "("
	LuRightParenthesis = ")"
	CeAndAdd           = "&+"
	CeAndProduct       = "&*"
	CeLuNonANotB       = "!&"
	CeLuNonAAddB       = "!+"
	FlagParserComplex  = fmt.Sprintf(`\&\&|\|\||\!\=|\&\+|\&\*|\!\&|\!\+`)
	AND                = "AND"
	OR                 = "OR"
	NOT                = "AND NOT"
	LuSlice            = []string{LuAnd, LuOr, LuNon, LuLeftParenthesis, LuRightParenthesis, CeAndAdd, CeAndProduct, CeLuNonAAddB, CeLuNonANotB}
	LuFMap             = map[string]string{LuAnd: AND, LuOr: OR, LuNon: NOT}
)
