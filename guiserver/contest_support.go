package guiserver

import (
	"encoding/csv"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/sirupsen/logrus"
	"matsu.dev/matsuri-contest-log/config"
)

type ContestSupport struct {
	PartialCheckerBase

	callsignInfo map[string][]string
}

func newContestSupport(name string) *ContestSupport {
	var (
		err error
		row []string
	)
	ret := &ContestSupport{
		PartialCheckerBase: PartialCheckerBase{
			checkerName: "Contest",
			count:       0,
			trie: &trieNode{
				callsign: "",
				next:     make(map[byte]*trieNode),
			},
		},
		callsignInfo: make(map[string][]string),
	}
	contestSupportFile := config.GetConfigFilePath(path.Join("contest", fmt.Sprintf("%s.csv", name)))
	logrus.Infof("trying to load contest support file: %s", contestSupportFile)
	csvFile, err := os.Open(contestSupportFile)
	if err != nil {
		return ret
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	for row, err = reader.Read(); err == nil; row, err = reader.Read() {
		ret.callsignInfo[row[0]] = row[1:]
		ret.insert(row[0])
	}

	fmt.Printf("Failed to read: %v", err)

	return ret
}

func (c *ContestSupport) match(str string) []match {
	ret := c.PartialCheckerBase.match(str)

	for i, _ := range ret {
		extra, ok := c.callsignInfo[ret[i].callsign]
		if ok && len(extra) != 0 {
			ret[i].formatCallsign = ret[i].formatCallsign + fmt.Sprintf("(%s)", strings.Join(extra, ","))
		}
	}
	return ret
}
