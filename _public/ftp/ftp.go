package pub_ftp

import (
	"bankBigData/_public/util"
	"fmt"
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
	defer conn.Quit()
	if e == nil {
		nameList, e = conn.NameList(path)
		_ = conn.Logout()
	}
	return nameList
}

func DownloadFile(path string) error {
	res := &ftp.Response{}
	file := &os.File{}
	if FtpFolder == "" {
		FtpFolder = g.Config().GetString("ftpFolder")
	}
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
	defer file.Close()
	return e
}

func Test() {
	list := GetNameList("/")
	for _, v := range list {
		fmt.Println(v)
	}
	DownloadFile("/20180722/s_iccs_tbl_acm_accbsc_all_20180722.ddl")
}
