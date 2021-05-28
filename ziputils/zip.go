// Package ziputils
// Author: youngchan
// CreateDate: 2021/5/28 4:25 下午
// Copyright: ©2021 NEW CORE Technology Co. Ltd. All rights reserved.
// Description:
//
package ziputils

import (
	"archive/zip"
	"github.com/youngchan1988/gocommon/fileutils"
	"io"
	"os"
	"strings"
)

//CompressFiles 压缩文件
//files 文件数组，可以是不同dir下的文件或者文件夹
//dest 压缩文件存放地址
func CompressFiles(files []*os.File, dest string) error {
	//检查dest目录，如果不存在则创建
	dir := strings.Replace(dest, `\`, "/", -1)
	index := strings.LastIndex(dir, "/")
	dir = dir[:index]
	err := fileutils.CreateDir(dir)
	if err != nil {
		return err
	}
	d, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer d.Close()
	w := zip.NewWriter(d)
	defer w.Close()
	for _, file := range files {
		err := compress(file, "", w)
		if err != nil {
			return err
		}
	}
	return nil
}

//CompressFilesTemp 压缩文件并存储在temp目录
//files 文件数组，可以是不同dir下的文件或者文件夹
//dest 压缩文件名称，以`*.zip` 结尾
//返回zip 文件存放的路径
func CompressFilesTemp(files []*os.File, dest string) (string, error) {
	d, err := os.CreateTemp("", dest)
	if err != nil {
		return "", err
	}
	defer d.Close()
	w := zip.NewWriter(d)
	defer w.Close()
	for _, file := range files {
		err := compress(file, "", w)
		if err != nil {
			return "", err
		}
	}
	return d.Name(), nil
}

//CompressDir 压缩目录
//dirPath 目录
//dest 压缩文件存放地址
func CompressDir(dirPath string, dest string) error {
	f, err := os.Open(dirPath)
	if err != nil {
		return err
	}
	//检查dest目录，如果不存在则创建
	dir := strings.Replace(dest, `\`, "/", -1)
	index := strings.LastIndex(dir, "/")
	dir = dir[:index]
	err = fileutils.CreateDir(dir)
	if err != nil {
		return err
	}
	d, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer d.Close()
	w := zip.NewWriter(d)
	defer w.Close()

	err = compress(f, "", w)
	if err != nil {
		return err
	}
	return nil
}

//CompressDirTemp 压缩目录并存储在temp目录
//dirPath 目录
//dest 压缩文件名称，以`*.zip` 结尾
//返回zip 文件存放的路径
func CompressDirTemp(dirPath string, dest string) (string, error) {
	f, err := os.Open(dirPath)
	if err != nil {
		return "", err
	}

	d, err := os.CreateTemp("", dest)
	if err != nil {
		return "", err
	}
	defer d.Close()
	w := zip.NewWriter(d)
	defer w.Close()

	err = compress(f, "", w)
	if err != nil {
		return "", err
	}
	return d.Name(), nil
}

func compress(file *os.File, prefix string, zw *zip.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		prefix = prefix + "/" + info.Name()
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = compress(f, prefix, zw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := zip.FileInfoHeader(info)
		header.Name = prefix + "/" + header.Name
		if err != nil {
			return err
		}
		writer, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

//DeCompress 解压
func DeCompress(zipFile, dest string) error {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer reader.Close()
	for _, file := range reader.File {
		rc, err := file.Open()
		if err != nil {
			return err
		}
		defer rc.Close()
		filename := dest + file.Name
		err = os.MkdirAll(getDir(filename), 0755)
		if err != nil {
			return err
		}
		w, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer w.Close()
		_, err = io.Copy(w, rc)
		if err != nil {
			return err
		}
		w.Close()
		rc.Close()
	}
	return nil
}

func getDir(path string) string {
	return subString(path, 0, strings.LastIndex(path, "/"))
}

func subString(str string, start, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start is wrong")
	}

	if end < start || end > length {
		panic("end is wrong")
	}

	return string(rs[start:end])
}
