package logging

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"time"

	"github.com/canghai908/ms-agent/setting"
)

//getLogFilePath func
func getLogFilePath() string {
	return fmt.Sprintf("%s", setting.AppSetting.LogSavePath)
}

//getLogFileName a
func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		"ms-agent_",
		time.Now().Format("20060102"),
		"log",
	)
}

//GetSize a
func GetSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)

	return len(content), err
}

//GetExt a
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

//CheckNotExist a
func CheckNotExist(src string) bool {
	_, err := os.Stat(src)

	return os.IsNotExist(err)
}

//CheckPermission a
func CheckPermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}

//IsNotExistMkDir a
func IsNotExistMkDir(src string) error {
	if notExist := CheckNotExist(src); notExist == true {
		if err := MkDir(src); err != nil {
			return err
		}
	}

	return nil
}

//MkDir a
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

//Open a
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}

//MustOpen a
func MustOpen(fileName, filePath string) (*os.File, error) {

	src := filePath + "/"
	perm := CheckPermission(src)
	if perm == true {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	err := IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	f, err := Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("Fail to OpenFile :%v", err)
	}

	return f, nil
}
