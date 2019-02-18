package pub_ftp

import (
	"bankBigData/_public/log"
	"bankBigData/_public/util"
	"gitee.com/johng/gf/g"
	"github.com/jlaffaye/ftp"
	"io"
	"os"
)

var (
	ftpAddr   string = ""
	ftpUser   string = ""
	ftpPass   string = ""
	FtpFolder string = ""
)

func loginFtp() (*ftp.ServerConn, error) {
	if ftpAddr == "" {
		ftpAddr = g.Config().GetString("ftpAddr")
		ftpUser = g.Config().GetString("ftpUser")
		ftpPass = g.Config().GetString("ftpPass")
	}
	conn, e := ftp.Connect(ftpAddr)
	if e == nil {
		e = conn.Login(ftpUser, ftpPass)
	}
	return conn, e
}

func GetNameList(path string) []string {
	nameList := []string{}
	conn, e := loginFtp()
	if e == nil {
		nameList, e = conn.NameList(path)
	}
	if e == nil {
		e = conn.Logout()
	}
	if e == nil {
		_ = conn.Quit()
	}
	return nameList
}

func DownloadFile(path string) error {
	res := &ftp.Response{}
	file := &os.File{}
	conn, e := loginFtp()

	if e == nil {
		res, e = conn.Retr(path)
		defer conn.Logout()
	}
	defer conn.Quit()

	if e == nil {
		file, e = util.CreateFile(FtpFolder + path)
	}
	if e == nil {
		for {
			buf := make([]byte, 1024)
			n, e := res.Read(buf)
			if e == io.EOF {
				break
			}
			_, e = file.Write(buf[:n])
		}
		_ = res.Close()
	}
	if e == nil {
		log.Instance().Println("下载文件：", path)
	} else {
		log.Instance().Error("下载文件失败：", path, " 失败原因：", e)
	}
	defer file.Close()
	return e
}
