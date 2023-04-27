/**
 * @author Co2
 * @file_name relation_test.go
 */

package main

import (
	"log"
	"testing"
)

func TestRelationalAnalysis(t *testing.T) {

	rule1 := "rule1 && (rule2 && (rule3 || rule4) && (rule5 && rule6)) &&  (rule7 != (rule8 || rule9)) || rule10"
	priority, record, keyRules, _ := RelationalAnalysis(rule1)
	log.Printf("测试优先级 = %#v \n", priority)
	log.Printf("测试映射结果为 = %#v \n", record)
	log.Printf("生成key值 = %#v \n", keyRules)
	MakeRuleTree(record)
}
