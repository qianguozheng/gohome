package admin

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

var fileDirectory string = "./files/"

func (admin AdminCtx) HandleUpload(c echo.Context) error {
	fmt.Println("admin upload page")

	//for _, cookie := range c.Cookies() {
	//	fmt.Printf("Name:[%s][%s]\n", cookie.Name, cookie.Value)
	//}

	cookie, err := c.Cookie("sessionId")
	if err != nil {
		return err
	}
	fmt.Printf("sessionId:%s\n", cookie.Value)
	//TODO: judge the cookie Name and Value exist in db

	//return c.File("html/admin.html")

	if c.Request().Method == "POST" {
		//获取文件句柄，然后对文件进行存储等处理
		file, err := c.FormFile("uploadfile")
		if err != nil {
			fmt.Println("form file err: ", err)
			return err
		}

		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		//创建上传的目的文件
		exist, _ := PathExists(fileDirectory)
		if !exist {
			os.Mkdir(fileDirectory, os.ModePerm)
		}
		dst, err := os.OpenFile(fileDirectory+file.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println("open file err: ", err)
			return err
		}
		defer dst.Close()
		//拷贝文件
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}
		return c.Render(http.StatusOK, "upload_done.html", map[string]interface{}{
			"name": file.Filename,
		})
	}

	return c.Render(http.StatusOK, "upload.html", map[string]interface{}{})
}

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
