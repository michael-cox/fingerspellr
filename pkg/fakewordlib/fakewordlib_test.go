package fakewordlib

import "testing"

func TestNewFakeWordGenerator(t *testing.T) {
    res, err := NewFakeWordGenerator("../../test/data/testwordlist.txt")
    if err != nil {
        t.Errorf("error during NewFakeWordGenerator: %v", err)
    }
    if res.letterCorrelationCounts['t']['e'] != 2 {
        t.Errorf("letterCorrelationCounts[%q][%q] should be 2, got %d", 't', 'e',
        res.letterCorrelationCounts['t']['e'])
    } else if res.letterCorrelationCounts['e']['a'] != 2 {
        t.Errorf("letterCorrelationCounts[%q][%q] should be 2, got %d", 'e', 'a',
        res.letterCorrelationCounts['e']['a'])
    } else if res.letterCorrelationCounts['e']['s'] != 1 {
        t.Errorf("letterCorrelationCounts[%q][%q] should be 1, got %d", 'e', 's',
        res.letterCorrelationCounts['e']['s'])
    } else if res.letterCorrelationCounts['e']['l'] != 1 {
        t.Errorf("letterCorrelationCounts[%q][%q] should be 1, got %d", 'e', 'l',
        res.letterCorrelationCounts['e']['l'])
    } else if res.letterCorrelationCounts['s']['t'] != 1 {
        t.Errorf("letterCorrelationCounts[%q][%q] should be 1, got %d", 's', 't',
        res.letterCorrelationCounts['s']['t'])
    } else if res.letterCorrelationCounts['h']['e'] != 1 {
        t.Errorf("letterCorrelationCounts[%q][%q] should be 1, got %d", 'h', 'e',
        res.letterCorrelationCounts['h']['e'])
    } else if res.letterCorrelationCounts['l']['l'] != 1 {
        t.Errorf("letterCorrelationCounts[%q][%q] should be 1, got %d", 'l', 'l',
        res.letterCorrelationCounts['l']['l'])
    } else if res.letterCorrelationCounts['l']['o'] != 1 {
        t.Errorf("letterCorrelationCounts[%q][%q] should be 1, got %d", 'l', 'o',
        res.letterCorrelationCounts['l']['o'])
    } else if res.letterCorrelationCounts['l']['d'] != 1 {
        t.Errorf("letterCorrelationCounts[%q][%q] should be 1, got %d", 'l', 'd',
        res.letterCorrelationCounts['l']['d'])
    } else if res.letterCorrelationCounts['w']['o'] != 1 {
        t.Errorf("letterCorrelationCounts[%q][%q] should be 1, got %d", 'w', 'o',
        res.letterCorrelationCounts['w']['o'])
    } else if res.letterCorrelationCounts['o']['r'] != 1 {
        t.Errorf("letterCorrelationCounts[%q][%q] should be 1, got %d", 'o', 'r',
        res.letterCorrelationCounts['o']['r'])
    } else if res.letterCorrelationCounts['r']['l'] != 1 {
        t.Errorf("letterCorrelationCounts[%q][%q] should be 1, got %d", 'r', 'l',
        res.letterCorrelationCounts['r']['l'])
    } else if res.letterCorrelationCounts['a']['s'] != 1 {
        t.Errorf("letterCorrelationCounts[%q][%q] should be 1, got %d", 'a', 's',
        res.letterCorrelationCounts['a']['s'])
    } else if res.letterCorrelationCounts['s']['e'] != 1 {
        t.Errorf("letterCorrelationCounts[%q][%q] should be 1, got %d", 's', 'e',
        res.letterCorrelationCounts['s']['e'])
    } else if res.letterCorrelationCounts['m']['e'] != 1 {
        t.Errorf("letterCorrelationCounts[%q][%q] should be 1, got %d", 'm', 'e',
        res.letterCorrelationCounts['m']['e'])
    } else if res.letterCorrelationCounts['a']['t'] != 1 {
        t.Errorf("letterCorrelationCounts[%q][%q] should be 1, got %d", 'a', 't',
        res.letterCorrelationCounts['a']['t'])
    }
}
