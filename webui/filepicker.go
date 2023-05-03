package main

import (
	"io/fs"
	"net/http"
	"path"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	selectType_File = iota
	selectType_Directory
	selectType_NewFile
	selectType_NewDirectory
)

type filesystemEntry struct {
	DisplayName string
	FullName    string
	Type        string
	LastModify  string
	Size        int64
}

type filesystemQuery struct {
	SelectType      string
	SourceDirectory string
	FilterSuffix    []string
}

type filesystemResponse struct {
	Result string
	Browse []filesystemEntry
}

func filePicker(c *gin.Context) {
	var (
		query    filesystemQuery
		response filesystemResponse
	)
	if c.ShouldBind(&query) != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"Result": "bad request",
		})
		return
	}

	selectType := selectType_File

	switch query.SelectType {
	case "file":
		selectType = selectType_File
	case "directory":
		selectType = selectType_Directory
	case "new_file":
		selectType = selectType_NewFile
	case "new_directory":
		selectType = selectType_NewDirectory
	}

	fullpath, err := filepath.Abs(query.SourceDirectory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"Result": err.Error(),
		})
		return
	}

	response.Browse = append(response.Browse, filesystemEntry{
		DisplayName: "..",
		FullName:    path.Join(fullpath, ".."),
		Type:        "dir",
	})

	filepath.Walk(fullpath, func(path string, info fs.FileInfo, err error) error {
		shouldAppend := false
		typeString := "file"

		if selectType == selectType_File {
			if info.IsDir() {
				shouldAppend = true
				typeString = "dir"
			}
			for _, suf := range query.FilterSuffix {
				if strings.HasSuffix(path, suf) {
					shouldAppend = true
					break
				}
			}
		} else if selectType == selectType_Directory {
			if info.IsDir() {
				shouldAppend = true
				typeString = "dir"
			}
		}

		if shouldAppend {
			response.Browse = append(response.Browse, filesystemEntry{
				DisplayName: info.Name(),
				FullName:    path,
				Type:        typeString,
				LastModify:  info.ModTime().String(),
				Size:        info.Size(),
			})
		}
		return nil
	})

	c.JSON(http.StatusOK, response)
}
