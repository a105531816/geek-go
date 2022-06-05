package main

type downloadFromNetDisk struct {
	secret   dymeNume
	filePath string
}

func (dd *downloadFromNetDisk) DownloadFile() {
	if err := dd.loginCheck(); err != nil {
		//todo 重新登录
	}
	dd.downloadFromAliYun(dd.filePath)
}

func (dd *downloadFromNetDisk) loginCheck() error {
	dd.checkSecret(dd.secret)
	return nil
}

func (dd *downloadFromNetDisk) downloadFromAliYun(file string) {

}

func (dd *downloadFromNetDisk) checkSecret(dymeNume) {
	//todo调用阿里云验证接口验证密码是否有效
}
