package pttbbs

import (
	"github.com/PichuChen/go-bbs"

	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Connector struct {
	home string
}

func init() {
	bbs.Register("pttbbs", &Connector{})
}

// Open connect a file directory or SHMs, dataSourceName pointer to bbs home
// And it can append argument for SHM
// for example `file:///home/bbs/?UTMP=1993`
func (c *Connector) Open(dataSourceName string) error {

	if strings.HasPrefix(dataSourceName, "file://") {
		s := dataSourceName[len("file://"):]
		seg := strings.Split(s, "?")
		c.home = seg[0]
	} else {
		c.home = dataSourceName
	}
	return nil
}

func (c *Connector) GetUserRecordsPath() (string, error) {
	return GetPasswdsPath(c.home)
}

func (c *Connector) ReadUserRecordsFile(path string) ([]bbs.UserRecord, error) {
	rec, err := OpenUserecFile(path)
	ret := make([]bbs.UserRecord, len(rec))
	for i, v := range rec {
		ret[i] = v
	}
	return ret, err
}

func (c *Connector) GetBoardRecordsPath() (string, error) {
	return GetBoardPath(c.home)
}

func (c *Connector) ReadBoardRecordsFile(path string) ([]bbs.BoardRecord, error) {
	rec, err := OpenBoardHeaderFile(path)
	ret := make([]bbs.BoardRecord, len(rec))
	for i, v := range rec {
		ret[i] = v
	}
	return ret, err
}

func (c *Connector) GetBoardArticleRecordsPath(boardId string) (string, error) {
	return GetBoardArticlesDirectoryPath(c.home, boardId)
}

func (c *Connector) ReadBoardArticleRecordsFile(boardId string) ([]bbs.ArticleRecord, error) {
	filepath, err := GetBoardArticlesDirectoryPath(c.home, boardId)

	var fileHeaders []*FileHeader
	fileHeaders, err = OpenFileHeaderFile(filepath)
	if err != nil {
		return nil, err
	}
	ret := make([]bbs.ArticleRecord, len(fileHeaders))
	for i, v := range fileHeaders {
		ret[i] = v
	}
	return ret, err
}

func (c *Connector) GetBoardTreasureRecordsPath(boardId string, treasureId []string) (string, error) {
	return GetBoardTreasuresDirectoryPath(c.home, boardId, treasureId)
}

func (c *Connector) ReadBoardTreasureRecordsFile(boardId string, treasureId []string) ([]bbs.ArticleRecord, error) {
	filepath, err := GetBoardTreasuresDirectoryPath(c.home, boardId, treasureId)

	var fileHeaders []*FileHeader
	fileHeaders, err = OpenFileHeaderFile(filepath)
	if err != nil {
		return nil, err
	}
	ret := make([]bbs.ArticleRecord, len(fileHeaders))
	for i, v := range fileHeaders {
		ret[i] = v
	}
	return ret, err
}

// ReadBoardArticleFile returns raw file of specific boardId/filename article.
func (c *Connector) ReadBoardArticleFile(boardId, filename string) ([]byte, error) {

	path, err := GetBoardArticleFilePath(c.home, boardId, filename)
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("pttbbs: open file error: %v", err)
	}
	defer file.Close()
	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("pttbbs: readfile error: %v", err)
	}
	return buf, err
}

// ReadBoardTreasureFile returns raw file of specific boardId/treasureId/filename article.
func (c *Connector) ReadBoardTreasureFile(boardId string, treasureId []string, filename string) ([]byte, error) {

	path, err := GetBoardTreasureFilePath(c.home, boardId, treasureId, filename)
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("pttbbs: open file error: %v", err)
	}
	defer file.Close()
	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("pttbbs: readfile error: %v", err)
	}
	return buf, err
}
