package fakewordlib

import (
    "bufio"
    "strings"
    "os"
    "log"
    "fmt"
)

type FakeWordGenerator struct {
    letterCorrelationCounts map[rune]map[rune]uint
}

func (wg FakeWordGenerator) insertWord(word string) {
    log.Printf("inserting word: %q", word)
    word = strings.ToLower(word)

    var lastChar rune = rune(0)
    for _, char := range word {
        if lastChar != rune(0) {
            if _, ok := wg.letterCorrelationCounts[lastChar][char]; ok {
                log.Printf("incrementing [%q][%q]", lastChar, char)
                wg.letterCorrelationCounts[lastChar][char] += 1
            } else {
                log.Printf("initializisg [%q][%q] to 1", lastChar, char)
                wg.letterCorrelationCounts[lastChar][char] = 1
            }
        }
        if _, ok := wg.letterCorrelationCounts[char]; !ok {
            log.Printf("making letterCorrelationCounts[%q]", char)
            wg.letterCorrelationCounts[char] = make(map[rune]uint)
        }
        lastChar = char
    }
}

func NewFakeWordGenerator(wordlist string) (*FakeWordGenerator, error) {

    newFWG := FakeWordGenerator{letterCorrelationCounts: make(map[rune]map[rune]uint)}

    file, err := os.Open(wordlist)
    if err != nil {
        return nil, fmt.Errorf("error opening wordlist %q: %w", wordlist, err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        newFWG.insertWord(scanner.Text())
    }
    if scanner.Err() != nil {
        return nil, fmt.Errorf("error during wordlist scanning: %w", scanner.Err())
    }

    return &newFWG, nil
}
