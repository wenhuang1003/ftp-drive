package ftpdriver

import (
	//"errors"
	"fmt"
	"io"

	"github.com/dutchcoders/goftp"
	"github.com/goftp/server"
)

type FtpDriver struct {
	server string
	ftp    *goftp.FTP
}

func (driver *FtpDriver) Init(conn *server.Conn) {
	//driver.conn = conn
}

func (driver *FtpDriver) ChangeDir(path string) error {
	return nil
}

func (driver *FtpDriver) Stat(key string) (server.FileInfo, error) {

	return nil, nil
}

func (driver *FtpDriver) ListDir(prefix string, callback func(server.FileInfo) error) error {

	return nil
}

func (driver *FtpDriver) DeleteDir(key string) error {

	return nil
}

func (driver *FtpDriver) DeleteFile(key string) error {
	return nil
}

func (driver *FtpDriver) Rename(keySrc, keyDest string) error {
	fmt.Println("rename from", keySrc, keyDest)
	return nil
}

func (driver *FtpDriver) MakeDir(path string) error {
	return nil
}

func (driver *FtpDriver) GetFile(key string, offset int64) (int64, io.ReadCloser, error) {

	return 0, nil, nil
}

func (driver *FtpDriver) PutFile(key string, data io.Reader, appendData bool) (int64, error) {

	return 0, nil
}

type FtpDriverFactory struct {
	host string
	port string
	user string
	pass string
}

func NewFtpDriverFactory(host, port, user, pass string) server.DriverFactory {
	return &FtpDriverFactory{host, port, user, pass}
}

func (factory *FtpDriverFactory) NewDriver() (server.Driver, error) {
	var server = factory.host + ":" + factory.port
	var ftp *goftp.FTP
	var err error
	if ftp, err = goftp.Connect(server); err != nil {
		return nil, err
	}

	return &FtpDriver{server, ftp}, nil
}
