package util

import (
	"archive/tar"
	"bankBigData/_public/log"
	"compress/gzip"
	"io"
	"os"
	"strings"
)

func Compress(files []*os.File, dest string) error {
	d, _ := os.Create(dest)
	defer d.Close()
	gw := gzip.NewWriter(d)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()
	for _, file := range files {
		err := compress(file, "", tw)
		if err != nil {
			return err
		}
	}
	return nil
}

func compress(file *os.File, prefix string, tw *tar.Writer) error {
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
			err = compress(f, prefix, tw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := tar.FileInfoHeader(info, "")
		header.Name = prefix + "/" + header.Name
		if err != nil {
			return err
		}
		err = tw.WriteHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(tw, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

//解压 tar.gz
func DeCompress(tarFile, dest string) error {
	srcFile, err := os.Open(tarFile)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	gr, err := gzip.NewReader(srcFile)
	if err != nil {
		return err
	}
	defer gr.Close()
	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		filename := dest + hdr.Name
		file, err := CreateFile(filename)
		if err != nil {
			return err
		}
		_, _ = io.Copy(file, tr)
	}
	return nil
}

func CreateFile(name string) (*os.File, error) {
	err := os.MkdirAll(string([]rune(name)[0:strings.LastIndex(name, "/")]), 0755)
	if err != nil {
		return nil, err
	}
	return os.Create(name)
}

// 解压文件到指定目录
func Gzip(gzFilePath, exportPath string) (bool, error) {
	basePath := strings.Split(gzFilePath, "/")
	baseExportPath := exportPath
	fileName := basePath[len(basePath)-1]
	fileName = strings.Replace(fileName, ".gz", "", -1)
	if exportPath == "" {
		baseExportPath = strings.Join(basePath[:len(basePath)-1], "/")
	}
	log.Instance().Println("gz路径：", gzFilePath, " 解压存放路径：", baseExportPath)
	// file read
	fileRes, err := os.Open(gzFilePath)
	if err != nil {
		return false, err
	}
	defer fileRes.Close()
	// gzip read
	gzipRes, err := gzip.NewReader(fileRes)
	if err != nil {
		return false, err
	}
	defer gzipRes.Close()

	// created file to export path
	fileTemp, err := os.OpenFile(baseExportPath+"/"+fileName, os.O_CREATE|os.O_WRONLY, 0644)

	defer fileTemp.Close()
	// write file
	_, err = io.Copy(fileTemp, gzipRes)
	//if err != nil {
	//	return false, err
	//}

	// tar read
	//tarRes := tar.NewReader(gzipRes)
	// read file
	//for {
	//	h, err := tarRes.Next()
	//	if err == io.EOF {
	//		break
	//	}
	//	if err != nil {
	//		return false, err
	//	}
	//	// created file to export path
	//	fileTemp, err := os.OpenFile(baseExportPath+"/"+h.Name, os.O_CREATE|os.O_WRONLY, 0644)
	//	logs.Info("解压文件：%s", baseExportPath+"/"+h.Name)
	//	if err != nil {
	//		return false, err
	//	}
	//	defer fileTemp.Close()
	//	// write file
	//	_, err = io.Copy(fileTemp, tarRes)
	//	if err != nil {
	//		return false, err
	//	}
	//}
	return err == nil, err
}

// 根据路径判断文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}