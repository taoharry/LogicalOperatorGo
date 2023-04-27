/**
 * @author Co2
 * @file_name relational_analysis.go
 */

package LogicalOperatorGo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	logging "log"
)

func iterNum(id int, str string) string {
	if id-1 > 0 {
		str = str + fmt.Sprintf("%v", id-1)
		return iterNum(id-1, str)
	} else {
		return str
	}
}

func MakeRuleTree(record map[int]string) {
	keys := []int{}
	traversalTrees := make([]TraversalTree, 0)
	nowDept := ""
	fatherDept := ""

	for k, _ := range record {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		strK := strconv.Itoa(k)
		lenK := len(strK)
		childTree := TraversalTree{}
		if lenK == 1 {
			nowDept = strK
			fatherDept = ""
			if k > 1 {
				nowDept = strK
				fatherDept = iterNum(k, "")
			}
		} else {
			orgDept := strK[:1]
			orgInt, _ := strconv.Atoi(orgDept)
			if orgInt == 1 {
				nowDept = strK
				fatherDept = strK[:lenK-1]
			} else {
				if lenK > 2 {
					//nowDept = iterNum(orgInt, "") + strK[1:]
					nowDept = strK
					fatherDept = strK[:lenK-1] + orgDept + iterNum(orgInt, "")
				} else {
					//nowDept = iterNum(orgInt, "") + strK[1:]
					nowDept = strK
					fatherDept = orgDept + iterNum(orgInt, "")
				}

			}
		}
		childTree.NowId = nowDept
		childTree.FatherId = strings.Trim(fatherDept, " ")
		childTree.Rule = record[k]
		traversalTrees = append(traversalTrees, childTree)
	}

	logging.Printf("result : %#v\n", traversalTrees)
	pTraversalTrees := make([]*TraversalTree, 0)
	for i, _ := range traversalTrees {
		var a *TraversalTree
		a = &traversalTrees[i]
		pTraversalTrees = append(pTraversalTrees, a)
	}

	var node *TraversalTree
	node = &traversalTrees[0]
	MakeTree(pTraversalTrees, node)

	data, _ := json.Marshal(node)
	var out bytes.Buffer
	json.Indent(&out, data, "", "\t")
	fmt.Printf("Tree=%v\n", out.String())
}

func FlagParser(rule string, dept int, priority []string, record map[int]string) ([]string, map[int]string) {
	find := regexp.MustCompile(FlagParserComplex)
	results := find.FindAllString(rule, -1)

	record[dept] = rule

	if len(results) == 1 {
		priority = append(priority, rule)
		return priority, record
	}
	ruleSilce := strings.Split(rule, LuKong)
	rules := RemoveKong(ruleSilce)
	luFlag := 0
	next := 0
	splitRule := ""
	perchRule := ""
	for n, key := range rules {
		if Contains(LuSlice, key) {
			luFlag += 1
			if luFlag == 1 {
				continue
			}
			if perchRule == "" {
				splitRule = strings.Join(rules[next:n], " ")
			} else {
				splitRule = perchRule + " " + strings.Join(rules[next:n], " ")
			}
			perchRule = fmt.Sprintf("%v%v", dept, luFlag-1)

			next = n
			priority = append(priority, splitRule)
			tmpCreateRuleInt, _ := strconv.Atoi(perchRule)
			record[tmpCreateRuleInt] = splitRule
			rule = strings.Replace(rule, splitRule, perchRule, 1)
			//logging.Println("rule解析后: ", rule)

			if n+2 == len(rules) {
				rule = perchRule + " " + strings.Join(rules[n:], LuKong)
				priority = append(priority, rule)
				record[tmpCreateRuleInt+1] = rule
			}
		}
	}
	return priority, record
}

func ParserRule(rule string, dept int, priority []string, record map[int]string) ([]string, map[int]string) {
	record[dept] = rule
	complex := fmt.Sprintf(`\([\&\|\!\=\+\*a-z0-9\ ]*?\)`)
	find := regexp.MustCompile(complex)
	results := find.FindAllString(rule, -1)
	//logging.Printf("运行第", dept, "层 result = ", results)
	if len(results) > 0 {
		for n, k := range results {

			newDept := fmt.Sprintf("%v%v", dept, n+1)
			trimK := strings.Trim(strings.Trim(k, LuLeftParenthesis), LuRightParenthesis)
			newDeptInt, _ := strconv.Atoi(newDept)
			priority, record = FlagParser(trimK, newDeptInt, priority, record)
			rule = strings.Replace(rule, k, newDept, 1)
		}
		priority, record = ParserRule(rule, dept+1, priority, record)
	} else {
		priority, record = FlagParser(rule, dept, priority, record)
	}
	return priority, record
}

// 拆分与, 或, 非和有限级逻辑
func RelationalAnalysis(rule string) (priority []string, record map[int]string, keyRules []string, err error) {
	err = nil
	rule, keyRules = FormatRules(rule)
	priority = []string{}
	record = map[int]string{}
	priority, record = ParserRule(rule, 1, priority, record)
	return
}

// 规则组合结构, 标准化为 rule1 && rule2
func FormatRules(rule string) (string, []string) {
	var (
		keyRules = []string{}
	)
	for _, i := range LuSlice {
		rule = strings.ReplaceAll(rule, i, fmt.Sprintf(" %v ", i))
	}

	rule2 := rule
	for _, i := range LuSlice {
		rule2 = strings.ReplaceAll(rule2, i, " ")
	}

	ruleSilce := strings.Split(rule2, LuKong)
	keyRulesNoTrims := RemoveKong(ruleSilce)
	for _, k := range keyRulesNoTrims {
		keyRules = append(keyRules, strings.ReplaceAll(k, LuKong, ""))
	}
	return rule, keyRules
}
