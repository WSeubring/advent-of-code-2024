package day5

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/wseubring/aoc2024/utils"
)

type PageOrderRule struct {
	PageBefore int
	PageAfter  int
}

type SafetyManualUpdate struct {
	pages []int
}

func (s *SafetyManualUpdate) GetMiddlePage() int {
	return s.pages[len(s.pages)/2]
}

func (s *SafetyManualUpdate) IsRuleSatisfied(rule PageOrderRule) bool {
	foundAfter := false
	for _, page := range s.pages {
		if page == rule.PageBefore {
			return !foundAfter
		}
		if page == rule.PageAfter {
			foundAfter = true
		}
	}

	// If both pages are not included in the manual, the rule is satisfied
	return true
}

func (s *SafetyManualUpdate) IsRuleSetSatisfied(ruleSet []PageOrderRule) bool {
	for _, rule := range ruleSet {
		if !s.IsRuleSatisfied(rule) {
			return false
		}
	}
	return true
}

//go:embed input.txt
var input string

func SolveA() {
	var rules []PageOrderRule
	var manuals []SafetyManualUpdate

	// Parse
	for _, line := range strings.Split(input, "\n") {
		if IsRule(line) {
			rules = append(rules, ParseRule(line))
		}
		if IsSafetyManualUpdate(line) {
			manuals = append(manuals, ParseSafetyManualUpdate(line))
		}
	}

	// Solve
	sum := 0
	for _, manual := range manuals {
		if manual.IsRuleSetSatisfied(rules) {
			sum += manual.GetMiddlePage()
		}
	}
	fmt.Println(sum)
}

func (s *SafetyManualUpdate) FixRule(rule PageOrderRule) {
	beforeIndex := -1
	afterIndex := -1

	for i, page := range s.pages {
		if page == rule.PageBefore {
			beforeIndex = i
		}
		if page == rule.PageAfter {
			afterIndex = i
		}
	}

	s.pages = utils.RemoveIndexFromArr(s.pages, beforeIndex)
	s.pages = utils.InsertAtIndex(s.pages, afterIndex, rule.PageBefore)
}

func (s *SafetyManualUpdate) FixOrderToSatisfyRules(rules []PageOrderRule) {
	for allSatisfied := false; !allSatisfied; {
		allSatisfied = true
		for _, rule := range rules {
			if !s.IsRuleSatisfied(rule) {
				s.FixRule(rule)
				allSatisfied = false
			}
		}
	}
}

func SolveB() {
	var rules []PageOrderRule
	var manuals []SafetyManualUpdate

	// Parse
	for _, line := range strings.Split(input, "\n") {
		if IsRule(line) {
			rules = append(rules, ParseRule(line))
		}
		if IsSafetyManualUpdate(line) {
			manuals = append(manuals, ParseSafetyManualUpdate(line))
		}
	}

	// Solve
	sum := 0
	for _, manual := range manuals {
		if !manual.IsRuleSetSatisfied(rules) {
			manual.FixOrderToSatisfyRules(rules)
			sum += manual.GetMiddlePage()
		}
	}
	fmt.Println(sum)
}

// PARSING
func IsRule(line string) bool {
	return strings.Contains(line, "|")
}

func ParseRule(line string) PageOrderRule {
	rule := strings.Split(line, "|")
	beforePage, err := utils.StringToInt(strings.TrimSpace(rule[0]))
	if err != nil {
		panic(err)
	}

	afterPage, err := utils.StringToInt(strings.TrimSpace(rule[1]))
	if err != nil {
		panic(err)
	}
	return PageOrderRule{PageBefore: beforePage, PageAfter: afterPage}
}

func IsSafetyManualUpdate(line string) bool {
	return strings.Contains(line, ",")
}

func ParseSafetyManualUpdate(line string) SafetyManualUpdate {
	pages := strings.Split(line, ",")
	pageNumbers, err := utils.StringsToInts(pages)

	if err != nil {
		panic(err)
	}

	return SafetyManualUpdate{pages: pageNumbers}
}
