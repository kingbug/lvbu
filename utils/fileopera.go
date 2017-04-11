package utils

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/astaxie/beego"
)

type FileInfo struct {
	RelPath string
	Size    int64
	IsDir   bool
	DirMode os.FileMode
	Handle  *os.File
}

//复制文件数据
func ioCopy(srcHandle *os.File, dstPth string) (err error) {
	mode, err := srcHandle.Stat()
	if err != nil {
		beego.Debug(err)
	}

	dstHandle, err := os.OpenFile(dstPth, os.O_CREATE|os.O_WRONLY, mode.Mode())
	if err != nil {
		return err
	}
	defer srcHandle.Close()
	defer dstHandle.Close()
	_, err = io.Copy(dstHandle, srcHandle)
	return err
}

//遍历目录，将文件信息传入通道
func WalkFiles(srcDir, suffix string, c chan<- *FileInfo) {
	filepath.Walk(srcDir, func(f string, fi os.FileInfo, err error) error { //遍历目录
		if err != nil {
			beego.Error(err)
			return err
		}
		fileInfo := &FileInfo{}
		if suffix != "" {
			if fi.IsDir() && fi.Name() == suffix {
				return filepath.SkipDir
			}
		}

		if fh, err := os.OpenFile(f, os.O_RDONLY, os.ModePerm); err != nil {
			beego.Error(err)
		} else {
			fileInfo.Handle = fh
			fileInfo.RelPath, _ = filepath.Rel(srcDir, f) //相对路径
			fileInfo.Size = fi.Size()
			fileInfo.IsDir = fi.IsDir()
			fileInfo.DirMode = fi.Mode()
		}
		c <- fileInfo

		return nil
	})
	close(c) //遍历完成，关闭通道
}

//写目标文件
func WriteFiles(dstDir string, c <-chan *FileInfo) error {
	if err := os.Chdir(dstDir); err != nil { //切换工作路径
		return err
		beego.Error(err)
	}
	for f := range c {

		if fi, err := os.Stat(f.RelPath); os.IsNotExist(err) { //目标不存在
			if f.IsDir {
				if err := os.MkdirAll(f.RelPath, f.DirMode); err != nil {
					beego.Error(err)
					return err
				}
			} else {
				if err := ioCopy(f.Handle, f.RelPath); err != nil {
					beego.Error(err)
					return err
				} else {
					//beego.Info(" CP:", f.RelPath)
				}
			}
		} else if !f.IsDir { //目标存在，而且源不是一个目录
			if fi.IsDir() != f.IsDir { //检查文件名被目录名占用冲突
				beego.Error("filename conflict:", f.RelPath)
				return err
			} else if fi.Size() != f.Size { //源和目标的大小不一致时才重写
				if err := ioCopy(f.Handle, f.RelPath); err != nil {
					beego.Error(err)
					return err
				} else {
					//beego.Info(" CP:", f.RelPath)
				}
			}
		}
	}
	return nil
}

//suffix 忽略文件夹
func Copypath(src, des, suffix string) error {
	if isexists, _ := PathExists(des); isexists {
		return errors.New("目录:" + des + "is exists.千年不遇，万年不成的事，让撞到了，买彩票吧!")
	} else if err := os.MkdirAll(des, 0755); err != nil {
		return err
	}
	files_ch := make(chan *FileInfo, 100)
	go WalkFiles(src, suffix, files_ch) //在一个独立的 goroutine 中遍历文件
	if err := WriteFiles(des, files_ch); err != nil {
		return err
	}
	des = des[:strings.LastIndex(des, PD)]
	des = des[:strings.LastIndex(des, PD)]
	des = des[:strings.LastIndex(des, PD)]
	beego.Info("目录切换为：", des)
	os.Chdir(des)
	return nil
}

func Deployover(des string) error {
	beego.Debug("删除目录", des)
	if err := os.RemoveAll(des); err != nil {
		return err
	}
	return nil
}

func InitClear() {
	dirPth := EXECPATH + PD + ".code"
	//os.Chdir(dirPth)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		beego.Debug(err)
	}
	for _, fi := range dir {
		if fi.IsDir() {
			os.RemoveAll(dirPth + PD + fi.Name())
		}
	}
}
