package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-common/library/log"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

var ipfsHost = "http://127.0.0.1:5001"

type ipfsResp struct {
	Name string
	Hash string
	Size string
}

func cat(c *gin.Context) {
	hash := c.Param("hash")
	if hash == "" {
		c.String(200, "")
	}
	resp, err := http.Get(fmt.Sprintf("%s/api/v0/cat?arg=%s", ipfsHost, hash))
	if err != nil {
		c.String(400, "")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	}
	responseFile := string(body)
	c.String(200, responseFile)
}

func add(c *gin.Context) {
	fileData := c.PostForm("data")
	fileBuf := strings.NewReader(fileData)
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)
	formFile, err := writer.CreateFormFile("file", "test.txt")
	if err != nil {
	}

	_, err = io.Copy(formFile, fileBuf)
	if err != nil {
	}

	contentType := writer.FormDataContentType()

	resp, err := http.Post(fmt.Sprintf("%s/api/v0/add?pin=true", ipfsHost), contentType, buf)
	if err != nil {
		c.String(400, "")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	}

	jsonResp := &ipfsResp{}
	_ = json.Unmarshal([]byte(body), jsonResp)
	fmt.Printf("hash text: %s", jsonResp.Hash)
	c.Redirect(302, fmt.Sprintf("/e/%s", jsonResp.Hash))
}

func main() {
	g := gin.New()
	g.HTMLRender = gintemplate.Default()
	g.POST("/d", add)
	g.GET("/e/:hash", cat)
	if err := http.ListenAndServe(":7788", g); err != nil {
		log.Error("http server start error(%v)", err)
		panic(err)
	}
}
