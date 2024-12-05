package day5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMiddlePage(t *testing.T) {
	s := SafetyManualUpdate{pages: []int{1, 2, 3, 4, 5}}
	assert.Equal(t, 3, s.GetMiddlePage())
}

func TestIsRuleSatisfiedTrue(t *testing.T) {
	s := SafetyManualUpdate{pages: []int{1, 2, 3, 4, 5}}
	rule := PageOrderRule{PageBefore: 3, PageAfter: 4}
	assert.True(t, s.IsRuleSatisfied(rule))
}

func TestIsRuleSatisfiedFalse(t *testing.T) {
	s := SafetyManualUpdate{pages: []int{1, 2, 3, 4, 5}}
	rule := PageOrderRule{PageBefore: 4, PageAfter: 3}
	assert.False(t, s.IsRuleSatisfied(rule))
}

func TestIsRuleSatisfiedBothMissing(t *testing.T) {
	s := SafetyManualUpdate{pages: []int{1, 2, 3, 5}}
	rule := PageOrderRule{PageBefore: 4, PageAfter: 6}
	assert.True(t, s.IsRuleSatisfied(rule))
}

func TestIsRuleSatisfiedOnlyAfterIncluded(t *testing.T) {
	s := SafetyManualUpdate{pages: []int{1, 2, 3, 4, 5}}
	rule := PageOrderRule{PageBefore: 6, PageAfter: 4}
	assert.True(t, s.IsRuleSatisfied(rule))
}

func TestIsRuleSetSatisfiedTrue(t *testing.T) {
	s := SafetyManualUpdate{pages: []int{1, 2, 3, 4, 5}}
	ruleSet := []PageOrderRule{
		{PageBefore: 3, PageAfter: 4},
		{PageBefore: 4, PageAfter: 6},
	}
	assert.True(t, s.IsRuleSetSatisfied(ruleSet))
}

func TestIsRuleSetSatisfiedFalse(t *testing.T) {
	s := SafetyManualUpdate{pages: []int{1, 2, 3, 4, 5}}
	ruleSet := []PageOrderRule{
		{PageBefore: 3, PageAfter: 4},
		{PageBefore: 4, PageAfter: 3},
	}

	assert.False(t, s.IsRuleSetSatisfied(ruleSet))
}

func TestSolveA(t *testing.T) {
	SolveA()
}

func TestFixOrderToSatisfyRules(t *testing.T) {
	s := SafetyManualUpdate{pages: []int{1, 2}}
	ruleSet := []PageOrderRule{
		{PageBefore: 2, PageAfter: 1},
	}

	s.FixOrderToSatisfyRules(ruleSet)
	assert.Equal(t, []int{2, 1}, s.pages)
}

func TestFixRule(t *testing.T) {
	s := SafetyManualUpdate{pages: []int{1, 2}}
	rule := PageOrderRule{PageBefore: 2, PageAfter: 1}

	s.FixRule(rule)
	assert.Equal(t, []int{2, 1}, s.pages)
}

func TestFixOrderToSatisfyRulesMultipleRules(t *testing.T) {
	s := SafetyManualUpdate{pages: []int{1, 2, 3}}
	ruleSet := []PageOrderRule{
		{PageBefore: 2, PageAfter: 1},
		{PageBefore: 3, PageAfter: 2},
	}

	s.FixOrderToSatisfyRules(ruleSet)
	assert.Equal(t, []int{3, 2, 1}, s.pages)
}

func TestSolveB(t *testing.T) {
	SolveB()
}
