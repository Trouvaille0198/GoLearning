package oslearning

// 一些情况
// 文件以ufd结构体形式保存
// 文件以文件名作为唯一索引
// afd中的number没有实际意义
// 所有文件数据 都在程序退出后正式写入外存

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type Command string

const (
	SHOW         Command = "show"
	SHOW_OPENING Command = "showo"
	CREATE       Command = "create"
	DELETE       Command = "delete"
	OPEN         Command = "open"
	CLOSE        Command = "close"
	READ         Command = "read"
	WRITE        Command = "write"
	BYE          Command = "bye"
)

type FileState string

const (
	OPENING FileState = "opening"
	READING FileState = "reading"
	WRITING FileState = "writing"
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

// Create 创建文件
func (mfd *MainFileDir) Create(name, proCode string) {
	mfd.UserFileDirs = append(mfd.UserFileDirs, &UserFileDir{Name: name, ProCode: proCode, Len: 0})
}

// Delete 根据文件名删除文件
func (mfd *MainFileDir) Delete(name string) (*UserFileDir, error) {
	for i, ufd := range mfd.UserFileDirs {
		if ufd.Name == name {
			mfd.UserFileDirs = append(mfd.UserFileDirs[:i], mfd.UserFileDirs[i+1:]...)
			return ufd, nil
		}
	}
	return &UserFileDir{}, errors.New("file not found")
}

// GetFileByName 根据文件名获取文件
func (mfd *MainFileDir) GetFileByName(name string) (*UserFileDir, error) {
	for _, ufd := range mfd.UserFileDirs {
		if ufd.Name == name {
			return ufd, nil
		}
	}
	return &UserFileDir{}, errors.New("file not found")
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
	ProCode string
	State   FileState
}

type FileSystem struct {
	UserName         string           // 用户名
	OwnerMainFileDir *MainFileDir     // 用户的mfd
	AccessFileDirs   []*AccessFileDir // 正在运行的文件目录

	MainFileDirs []*MainFileDir // 所有mfd 用于程序退出时更新外存内容
	FilePath     string         // 外存地址
}

// NewFileSystem 新建文件管理系统 读取用户与用户文件
func NewFileSystem(filePath string) (*FileSystem, error) {
	// 读取用户和文件信息
	fh, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error while opening file: ", err)
		return &FileSystem{}, errors.New("can't open file_example")
	}
	defer fh.Close()

	jsonData, err := ioutil.ReadAll(fh)
	if err != nil {
		fmt.Println("Error while reading json: ", err)
		return &FileSystem{}, errors.New("wrong format with file_example")
	}
	// 解析json数据到MFD中
	mainFileDirs := make([]*MainFileDir, 0)
	err = json.Unmarshal(jsonData, &mainFileDirs)
	if err != nil {
		fmt.Println("Error while converting json: ", err)
		return &FileSystem{}, errors.New("jsonData conflicts with struct MFD")
	}

	// 判断用户名是否注册
	var ownerMainFileDir *MainFileDir // 用户的mfd
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

	return &FileSystem{
		UserName: userName, MainFileDirs: mainFileDirs,
		OwnerMainFileDir: ownerMainFileDir, FilePath: filePath}, nil
}

// ShowFiles 打印文件列表
func (f *FileSystem) ShowFiles() {
	fmt.Println("-----------------------------")

	if len(f.OwnerMainFileDir.UserFileDirs) != 0 {
		fmt.Printf("%-10s %-10s %-10s \n", "name", "pro code", "length")
		for _, ufd := range f.OwnerMainFileDir.UserFileDirs {
			fmt.Printf("%-10s %-10s %-10d \n", ufd.Name, ufd.ProCode, ufd.Len)
		}
	} else {
		fmt.Println("No files.")
	}

	fmt.Println("-----------------------------")
}

// ShowOpeningFiles 打印打开的文件列表
func (f *FileSystem) ShowOpeningFiles() {
	fmt.Println("-----------------------------")
	if len(f.AccessFileDirs) != 0 {
		fmt.Printf("%-10%s %-10s %-10s %-10s \n", "number", "name", "pro code", "length")
		for _, afd := range f.AccessFileDirs {
			ufd := afd.FileDir
			fmt.Printf("%-10d %-10s %-10s %-10d \n", afd.Number, ufd.Name, afd.ProCode, ufd.Len)
		}
	} else {
		fmt.Println("No files.")
	}
	fmt.Println("-----------------------------")
}

// Create 创建文件 TODO 不能同名
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

// Delete 删除文件
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
	fmt.Printf("File %s is deleted \n", name)
	return nil
}

// Open 打开文件 即将目标文件链接到afd中
func (f *FileSystem) Open() error {
	if len(f.AccessFileDirs) > 5 {
		return errors.New("the amount of opening files is up to 5")
	}
	fmt.Println("Enter the name of file to be opened: ")
	var name string
	_, err := fmt.Scan(&name)
	if err != nil {
		return err
	}

	ufd, err := f.OwnerMainFileDir.GetFileByName(name)
	if err != nil {
		return err
	}

	fmt.Println("Enter the open mode: ")
	var proCode string
	_, err = fmt.Scan(&proCode)
	if err != nil {
		return err
	}

	fileNumber := len(f.AccessFileDirs) + 1
	afd := &AccessFileDir{Number: fileNumber, FileDir: ufd, ProCode: proCode}
	f.AccessFileDirs = append(f.AccessFileDirs, afd)

	fmt.Printf("File %s opened, number %d \n", name, fileNumber)
	return nil
}

// Close 关闭文件 即将目标文件从afds中删除
func (f *FileSystem) Close() error {
	fmt.Println("Enter the name of file to be closed: ")
	var name string
	_, err := fmt.Scan(&name)
	if err != nil {
		return err
	}

	for i, afd := range f.AccessFileDirs {
		if afd.FileDir.Name == name {
			f.AccessFileDirs = append(f.AccessFileDirs[:i], f.AccessFileDirs[i+1:]...)
			return nil
		}
	}
	return errors.New("file is not in opening file list")
}

// Read 读文件
func (f *FileSystem) Read() error {
	fmt.Println("Enter the name of file to read: ")
	var name string
	_, err := fmt.Scan(&name)
	if err != nil {
		return err
	}

	for _, afd := range f.AccessFileDirs {
		if afd.FileDir.Name == name {
			if afd.ProCode[0] == '0' {
				return errors.New("file is not allowed to read")
			}
			afd.State = READING
			fmt.Printf("File %s is being read \n", name)
			return nil
		}
	}
	return errors.New("file is not in opening file list")
}

// Write 读文件
func (f *FileSystem) Write() error {
	fmt.Println("Enter the name of file to write: ")
	var name string
	_, err := fmt.Scan(&name)
	if err != nil {
		return err
	}

	for _, afd := range f.AccessFileDirs {
		if afd.FileDir.Name == name {
			if afd.ProCode[1] == '0' {
				return errors.New("file is not allowed to write")
			}
			afd.State = WRITING
			fmt.Println("How many characters to be written into the file? ")
			var length int
			_, err = fmt.Scan(&length)
			if err != nil {
				return err
			}
			afd.FileDir.Len += length
			fmt.Printf("file %s is written", name)
			return nil
		}
	}
	return errors.New("file is not in opening file list")
}

// Update 将文件更新至外存
func (f *FileSystem) Update() error {
	fh, err := os.OpenFile(f.FilePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error while opening file: ", err)
		return errors.New("can't open file_example")
	}
	defer fh.Close()

	jsonData, err := json.MarshalIndent(f.MainFileDirs, "", "  ")
	if err != nil {
		return err
	}

	_, err = fh.Write(jsonData)
	if err != nil {
		return err
	}
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
		case SHOW_OPENING:
			f.ShowOpeningFiles()
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
			err = f.Open()
			if err != nil {
				fmt.Println(err)
			}
		case CLOSE:
			err = f.Close()
			if err != nil {
				fmt.Println(err)
			}
		case READ:
			err = f.Read()
			if err != nil {
				fmt.Println(err)
			}
		case WRITE:
			err = f.Write()
			if err != nil {
				fmt.Println(err)
			}
		case BYE:
			err = f.Update()
			if err != nil {
				fmt.Println(err)
			}
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
