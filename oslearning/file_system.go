package oslearning

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type Command string

const (
	SHOW   Command = "show"
	CREATE Command = "create"
	DELETE Command = "delete"
	OPEN   Command = "open"
	CLOSE  Command = "close"
	READ   Command = "read"
	WRITE  Command = "write"
	BYE    Command = "bye"
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

// isZeroOrOne 判断一个字符串是否只包含0或1
func isZeroOrOne(str string) bool {
	for _, char := range str {
		if char != '0' && char != '1' {
			return false
		}
	}
	return true
}

// Create 创建文件 TODO 写入json文件
func (mfd *MainFileDir) Create(name, proCode string) {
	mfd.UserFileDirs = append(mfd.UserFileDirs, &UserFileDir{Name: name, ProCode: proCode, Len: 0})
}

// Delete 根据文件名删除文件 TODO 写入json文件
func (mfd *MainFileDir) Delete(name string) error {
	for i, ufd := range mfd.UserFileDirs {
		if ufd.Name == name {
			mfd.UserFileDirs = append(mfd.UserFileDirs[:i], mfd.UserFileDirs[i+1:]...)
			return nil
		}
	}
	return errors.New("file not found")
}

// UserFileDir 用户文件目录
type UserFileDir struct {
	Name    string `json:"name"`
	ProCode string `json:"proCode"` // 保护码
	Len     int    `json:"len"`
}

// AccessFileDir 运行主目录
type AccessFileDir struct {
	Number  int
	FileDir *UserFileDir
	ProCode [3]int
	State   fileState
}

type FileSystem struct {
	UserName         string           // 用户名
	OwnerMainFileDir *MainFileDir     // 用户的mfd
	AccessFileDirs   []*AccessFileDir // 正在运行的文件目录
}

// NewFileSystem 新建文件管理系统 读取用户与用户文件
func NewFileSystem(filePath string) (*FileSystem, error) {
	// 读取用户和文件信息
	fh, err := os.Open(filePath)
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

	// 判断用户名是否注册
	var ownerMainFileDir MainFileDir // 用户的mfd
	var userName string
	for {
		// 获取来自标准输入的用户名
		fmt.Println("Your name is?")
		_, err = fmt.Scan(&userName)
		if err != nil {
			fmt.Println("Error while inputting username: ", err)
			os.Exit(1)
		}

		flag := false
		for _, mfd := range mainFileDirs {
			if userName == mfd.Owner {
				flag = true
				ownerMainFileDir = mfd
			}
		}
		if flag {
			break
		} else {
			fmt.Printf("%s is not a valid user, please try again\n", userName)
		}
	}

	return &FileSystem{UserName: userName, OwnerMainFileDir: &ownerMainFileDir}, nil
}

// ShowFiles 打印文件列表
func (f *FileSystem) ShowFiles() {
	fmt.Println("-----------------------------")
	fmt.Printf("%-10s %-10s %-10s \n", "name", "pro code", "length")
	for _, ufd := range f.OwnerMainFileDir.UserFileDirs {
		fmt.Printf("%-10s %-10s %-10d \n", ufd.Name, ufd.ProCode, ufd.Len)
	}
	fmt.Println("-----------------------------")
}

func (f *FileSystem) Create() error {
	fmt.Println("The new file's name: ")
	var name string
	_, err := fmt.Scan(&name)
	if err != nil {
		return err
	}

	fmt.Println("The new file's protection code: ")
	var proCode string
	_, err = fmt.Scan(&proCode)
	if err != nil {
		return err
	}
	// 检查proCode是否合法
	if len(proCode) != 3 || !isZeroOrOne(proCode) {
		return errors.New("protection code is invalid")
	}
	f.OwnerMainFileDir.Create(name, proCode)
	fmt.Println("File is created")
	return nil
}

func (f *FileSystem) Delete() error {
	fmt.Println("Enter the name of file to be deleted: ")
	var name string
	_, err := fmt.Scan(&name)
	if err != nil {
		return err
	}
	err = f.OwnerMainFileDir.Delete(name)
	if err != nil {
		return err
	}
	fmt.Printf("File %s is deleted", name)
	return nil
}

// Run 入口函数
func (f *FileSystem) Run() {
	f.ShowFiles()
	for {
		flag := true
		fmt.Println("What would you like to do?")
		var command Command
		_, err := fmt.Scan(&command)
		if err != nil {
			fmt.Println("Error while inputting Command: ", err)
			os.Exit(1)
		}
		switch command {
		case SHOW:
			f.ShowFiles()
		case CREATE:
			err = f.Create()
			if err != nil {
				fmt.Println(err)
			}
		case DELETE:
			err = f.Delete()
			if err != nil {
				fmt.Println(err)
			}
		case OPEN:
			fmt.Println("open")
		case CLOSE:
			fmt.Println("close")
		case READ:
			fmt.Println("read")
		case WRITE:
			fmt.Println("write")
		case BYE:
			fmt.Println("See u.")
			flag = false
		default:
			fmt.Println("Seems like you've typed a wrong command...")
			continue
		}
		if !flag {
			break
		}
	}
}
