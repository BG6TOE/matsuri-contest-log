package guiserver

import (
	"io"
	"os"
	"sort"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
)

type match struct {
	callsign       string
	formatCallsign string
	score          int
}

type trieNode struct {
	callsign string
	next     map[byte]*trieNode
}

type PartialChecker interface {
	insert(call string)
	match(call string) []match
	name() string
}

type PartialCheckerBase struct {
	checkerName string
	count       int
	trie        *trieNode
	mut         sync.RWMutex
}

type matchList []match

func (l matchList) Len() int {
	return len(l)
}

func (l matchList) Less(i, j int) bool {
	return l[i].score < l[j].score
}
func (l matchList) Swap(i, j int) {
	t := l[i]
	l[i] = l[j]
	l[j] = t
}

func (n *trieNode) insert(index int, call string) {
	if index == 0 {
		n.callsign = call
		return
	}
	char := call[index-1]
	child, ok := n.next[char]
	if !ok {
		child = &trieNode{
			next: make(map[byte]*trieNode),
		}
		n.next[char] = child
	}
	child.insert(index-1, call)
}

func (n *trieNode) match(maxError int, index int, call string, score int, matchCall string, res *[]match) {
	if index == 0 {
		if len(n.callsign) != 0 {
			*res = append(*res, match{callsign: n.callsign, formatCallsign: matchCall, score: score})
		}
	}
	var char byte
	if index > 0 {
		char = call[index-1]
	}
	for k, nxt := range n.next {
		if k == char {
			nextCall := string([]byte{char}) + matchCall
			nxt.match(maxError, index-1, call, score, nextCall, res)
			continue
		}
		if maxError >= 2 || (maxError >= 1 && index == 0) || (index == 0 && len(call) >= 3) {
			nextCall := string([]byte{'!', k}) + matchCall
			// allow missing
			if index != 0 {
				nxt.match(maxError-2, index, call, score+2, nextCall, res)
			} else if len(call) < 3 {
				// prefix is cheap
				nxt.match(maxError-1, index, call, score+2, nextCall, res)
			} else {
				// region prefix is free
				nxt.match(maxError, index, call, score+2, nextCall, res)
			}
		}
		if maxError >= 1 && index > 0 {
			// allow wrong
			nextCall := string([]byte{'@', k}) + matchCall
			nxt.match(maxError-1, index-1, call, score+2, nextCall, res)
		}
	}
}

func (c *PartialCheckerBase) insert(call string) {
	c.mut.Lock()
	defer c.mut.Unlock()

	c.trie.insert(len(call), call)
	c.count++
}

func (c *PartialCheckerBase) match(call string) []match {
	c.mut.RLock()
	defer c.mut.RUnlock()

	result := make([]match, 0)
	if len(call) <= 3 {
		c.trie.match(1, len(call), call, 0, "", &result)
	} else {
		c.trie.match(2, len(call), call, 0, "", &result)
	}

	var tr matchList = result
	sort.Sort(tr)

	if len(result) > 10 {
		result = result[:10]
	}

	return result
}

func (c *PartialCheckerBase) name() string {
	return c.checkerName
}

func newPartialChecker(name string) *PartialCheckerBase {
	return &PartialCheckerBase{
		checkerName: name,
		trie: &trieNode{
			next: make(map[byte]*trieNode),
		},
	}
}

func scpChecker(file string) *PartialCheckerBase {
	ret := newPartialChecker("Database")

	fp, err := os.Open(file)
	if err != nil {
		return ret
	}
	defer fp.Close()

	data, err := io.ReadAll(fp)
	if err != nil {
		return ret
	}

	lines := strings.Split(string(data), "\n")
	for _, l := range lines {
		l = strings.TrimSpace(l)
		if len(l) == 0 {
			continue
		}
		if l[0] == '#' {
			continue
		}
		ret.insert(l)
	}

	logrus.Infof("Loaded %d callsigns from SCP", ret.count)

	return ret
}
