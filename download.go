package request

import (
	"io"
	"net/http"
	"os"
	"path"
)

//download file from network
func Download(url,filepath string)  error{
	//get []byte
	resp,err:=http.Get(url)
	if err!=nil{
		return err
	}
	defer resp.Body.Close()

	//check whether the Dir exists or not?
	dir := path.Dir(filepath)
	if !isExist(dir){
		err=os.MkdirAll(dir,0755)
		if err!=nil{
			return err
		}
	}
	//create the file
	out, err := os.Create(filepath)
	if err!=nil{
		return err
	}
	defer out.Close()

	//copy
	_,err=io.Copy(out,resp.Body)
	if err!=nil{
		return err
	}
	return nil
}

func isExist(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}


