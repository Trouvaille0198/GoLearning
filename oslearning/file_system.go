package oslearning

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type command string

const (
	CREATE command = "create"
	DELETE command = "delete"
	OPEN   command = "open"
	CLOSE  command = "close"
	READ   command = "read"
	WRITE  command = "write"
	BYE    command = "bye"
)

type fileState string

const (
	READING fileState = "reading"
	WRITING fileState = "writing"
)

// MainFileDir 主文件目录
type MainFileDir struct {
	Owner        string         `json:"owner"`
	UserFileDirs []*UserFileDir `json:"userFileDirs"`
}

// UserFileDir 主文件目录
type UserFileDir struct {
	Name    string `json:"name"`
	ProCode string `json:"proCode"` // 保护码
	Len     int    `json:"len"`
}

// AccessFileDir 运行主目录
type AccessFileDir struct {
	fileDir *UserFileDir
	proCode [3]int
	state   fileState
}

type FileSystem struct {
	userName     string
	mainFileDirs []MainFileDir //TODO 应为单例
}

// NewFileSystem 新建文件管理系统 读取用户与用户文件
func NewFileSystem() (*FileSystem, error) {
	// 读取用户和文件信息
	cwd, _ := os.Getwd()
	fh, err := os.Open(cwd + "\\oslearning\\file_example.json")
	if err != nil {
		fmt.Println("Error while opening file: ", err)
		return &FileSystem{}, errors.New("can't open file_example")
	}
	jsonData, err := ioutil.ReadAll(fh)
	if err != nil {
		fmt.Println("Error while reading json: ", err)
		return &FileSystem{}, errors.New("wrong format with file_example")
	}
	// 解析json数据到MFD中
	mainFileDirs := make([]MainFileDir, 0)
	err = json.Unmarshal(jsonData, &mainFileDirs)
	if err != nil {
		fmt.Println("Error while converting json: ", err)
		return &FileSystem{}, errors.New("jsonData conflicts with struct MFD")
	}
	// 保存来自标准输入的用户名
	fmt.Println("Your name is?")
	var userName string
	_, err = fmt.Scan(&userName)
	if err != nil {
		fmt.Println("Error while inputting: ", err)
		os.Exit(1)
	}
	for {
		flag := false
		for _, mfd := range mainFileDirs {
			if userName == mfd.Owner {
				flag = true
			}
		}
		if flag {
			break
		} else {
			fmt.Printf("%s is not a valid user, please try again\n", userName)
		}
	}

	return &FileSystem{userName: userName, mainFileDirs: mainFileDirs}, nil
}

func (f *FileSystem) ShowUserName() {
	jsonData, _ := json.Marshal(f.mainFileDirs)
	fmt.Print(jsonData)
}
