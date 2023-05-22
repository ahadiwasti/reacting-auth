//  Update by Emran Hamdan 14.5.2020
//  Applying Native gprc call to avoide Micro API handler
//  Fix the issue whre we call Micro API handler , now we are calling native Micro Client using JSON Metadata
//  DON NOT PLAY PLEASE
// [240 157 215 64 108 195 140 61 188 135 79 44 37 166 51 24]
// [179, 129, 235, 135, 122, 224, 248, 195, 134, 5, 84, 172, 28, 228, 117, 136]
// [127 79 93 119 190 52 233 199 55 241 205 221 109 207 59 225 253 156 219 150 186 223 125 124]
package controllers

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"io/ioutil"

	// "./pkg/api/cache"
	"./pkg/api/dto"
	"./pkg/api/logger"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"

	// "compress/zlib"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

//Play diffrent with native call
type Payload struct {
	Service string `json:"service"`
	Method  string `json:"method"`
	Request interface{}
}

type CRYPTO struct {
	Content string `json:"content" uri:"content" form:"content"`
	Iv      string `json:"iv" uri:"iv" form:"iv"`
}

type dynamic struct {
	value interface{}
}

type AmsController struct {
	BaseController
}
type StructData struct {
	Data []string `json:"data"`
}

func AESEncrypt(src string, key []byte, initVector []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("key error1", err)
	}
	if src == "" {
		fmt.Println("plain content empty")
	}
	ecb := cipher.NewCBCEncrypter(block, []byte(initVector))
	content := []byte(src)
	content = PKCS5Padding(content, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)
	return crypted
}
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
func AESDecrypt(crypt []byte, key []byte, initVector []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("key error1", err)
	}
	if len(crypt) == 0 {
		fmt.Println("plain content empty")
	}
	ecb := cipher.NewCBCDecrypter(block, []byte(initVector))
	decrypted := make([]byte, len(crypt))
	ecb.CryptBlocks(decrypted, crypt)
	return PKCS5Trimming(decrypted)
}
func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
func (a *AmsController) Amstech_rpc(c *gin.Context) {
	var encodedvalue CRYPTO
	var decryptedTextMoldel Payload
	pass, err1 := hex.DecodeString("C39BF77F8442B1C6C1D48808B1B35AE89CFDE6C373C693DBC99F2D16AC8408E7")
	if err1 != nil {
		fmt.Println("key error1", err1)
	}

	service := micro.NewService()
	service.Init()

	cl := service.Client()
	if a.BindAndValidate(c, &encodedvalue) {
		IV := encodedvalue.Iv
		initialVector, err := hex.DecodeString(IV)
		if err != nil {
			fmt.Println("key error1", err)
		}
		Content := encodedvalue.Content
		encryptedData, _ := base64.StdEncoding.DecodeString(Content)
		passphrase := []byte(pass)
		decryptedText := AESDecrypt(encryptedData, []byte(passphrase), initialVector)
		json.Unmarshal(decryptedText, &decryptedTextMoldel)

		var newpayload Payload
		newpayload.Service = decryptedTextMoldel.Service
		newpayload.Method = decryptedTextMoldel.Method
		newpayload.Request = decryptedTextMoldel.Request

		request := cl.NewRequest(newpayload.Service, newpayload.Method, newpayload.Request, client.WithContentType("application/json"))

		// var unpackedstr []string
		// unpackedstr = append(unpackedstr, string(newpayload.Request[:]))
		// fmt.Println("unpackedstr", unpackedstr)
		// epr := &StructData{unpackedstr}
		// fmt.Println("epr", epr)
		// r, _ := cache.Get(newpayload.Service)
		// fmt.Println("cache resp", r)
		var result map[string]interface{}
		var response interface{}
		if err := cl.Call(context.TODO(), request, &response); err != nil {
			rspbody, _ := json.Marshal(err)
			logger.Errorf("Request", err)
			if err := json.Unmarshal(rspbody, &result); err != nil {
				logger.Errorf("Request", err)

			}
			logger.Errorf("Request", err)
			resp(c, result)
		}

		rspbody, _ := json.Marshal(response)
		var unpacked []string
		unpacked = append(unpacked, string(rspbody[:]))
		pr := &StructData{unpacked}
		prAsBytes, err := json.Marshal(pr)
		if err != nil {
			fmt.Println("error :", err)
		}
		encryptedDataResp := AESEncrypt(string(prAsBytes[:]), passphrase, initialVector)
		encryptedString := base64.StdEncoding.EncodeToString(encryptedDataResp)
		resp(c, map[string]interface{}{
			"data": encryptedString,
		})
	}

}

// func (a *AmsController) Amstech_rpc(c *gin.Context) {

// 	var payloadDto dto.AmsPayload
// 	var newpayload Payload
// 	service := micro.NewService()
// 	service.Init()

// 	cl := service.Client()

// 	if a.BindAndValidate(c, &payloadDto) {
// 		newpayload.Service = payloadDto.Service
// 		newpayload.Method = payloadDto.Method

// 		//   fmt.Println("New Payload %q", newpayload)
// 		// convert to interface  to make Micro happy
// 		var reqdata = map[string]interface{}{}
// 		if err := json.Unmarshal([]byte(payloadDto.Request), &reqdata); err != nil {
// 			// keep panic for a while
// 			logger.Errorf("Request", err)
// 			panic(err)
// 		}

// 		newpayload.Request = reqdata

// 		// prepare request
// 		request := cl.NewRequest(newpayload.Service, newpayload.Method, newpayload.Request, client.WithContentType("application/json"))
// 		// prepare and call
// 		var result map[string]interface{}
// 		var response interface{}
// 		if err := cl.Call(context.TODO(), request, &response); err != nil {
// 			rspbody, _ := json.Marshal(err)
// 			logger.Errorf("Request", err)
// 			if err := json.Unmarshal(rspbody, &result); err != nil {
// 				logger.Errorf("Request", err)

// 			}
// 			logger.Errorf("Request", err)
// 			// fmt.Printlnln(err)
// 			resp(c, result)
// 		}

// 		rspbody, _ := json.Marshal(response)
// 		if err := json.Unmarshal(rspbody, &result); err != nil {
// 			logger.Errorf("Request", err)
// 			panic(err)
// 		}
// 		// what this do ? Ok getit return OK or not
// 		resp(c, result)
// 		// return result, nil
// 	}

// }

func (a *AmsController) Amstech_rpcevent(c *gin.Context) {

	var payloadDto dto.AmsPayload
	var newpayload Payload
	service := micro.NewService()
	service.Init()

	cl := service.Client()

	if a.BindAndValidate(c, &payloadDto) {
		newpayload.Service = payloadDto.Service
		newpayload.Method = payloadDto.Method

		// convert to interface  to make Micro happy
		var reqdata = map[string]interface{}{}
		if err := json.Unmarshal([]byte(payloadDto.Request), &reqdata); err != nil {
			// keep panic for a while
			panic(err)
		}

		newpayload.Request = reqdata

		// prepare request use publish to order Event , this is danger can create order

		p := micro.NewEvent("go.micro.api.srv.order", cl)
		p.Publish(context.TODO(), newpayload.Request)

		// fake resp
		result := map[string]interface{}{
			"result": "OK",
			"Code":   "000",
		}
		resp(c, result)
	}
}

func (a *AmsController) Amstech_pemfile(c *gin.Context) {
	// var w http.ResponseWriter = c.Writer
	var req *http.Request = c.Request
	req.ParseMultipartForm(10 << 20)
	file, handler, err := req.FormFile("file")
	if err != nil {
		fmt.Println("INVALID_FILE", err)
	}
	defer file.Close()
	tempFile, err := ioutil.TempFile("pkg/webui/dist/static/keys/", "key-*.asc")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("ReadFile_err", err)
	}
	tempFile.Write(fileBytes)
	mainFilePath := "pkg/webui/dist/static/keys/" + handler.Filename
	renamed := os.Rename(tempFile.Name(), mainFilePath)
	if renamed != nil {
		fmt.Println("Rename_err", err)
	}
	result := map[string]interface{}{
		"result": "successfully Saved",
		"Code":   "000",
	}
	resp(c, result)
	// fmt.Fprintf(w, "successfully saved")
}
func (a *AmsController) Amstech_removepemfile(c *gin.Context) {
	// var w http.ResponseWriter = c.Writer
	var delSt delete_st
	if a.BindAndValidate(c, &delSt) {
		path := "pkg/webui/dist/static/keys/" + delSt.Name
		errtorem := os.Remove(path)
		if errtorem != nil {
			fmt.Println(errtorem)
		}
		result := map[string]interface{}{
			"result": "successfully Deleted",
			"Code":   "000",
		}
		resp(c, result)
		// fmt.Fprintf(w, "successfully deleted")
	}
}

func (a *AmsController) Amstech_apk(c *gin.Context) {
	// var w http.ResponseWriter = c.Writer
	var req *http.Request = c.Request
	req.ParseMultipartForm(10 << 20)
	file, handler, err := req.FormFile("file")
	if err != nil {
		fmt.Println("INVALID_FILE", err)
	}
	defer file.Close()
	tempFile, err := ioutil.TempFile("pkg/webui/dist/static/apk/", "upload-*.apk")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("ReadFile_err", err)
	}
	tempFile.Write(fileBytes)
	mainFilePath := "pkg/webui/dist/static/apk/" + handler.Filename
	renamed := os.Rename(tempFile.Name(), mainFilePath)
	if renamed != nil {
		fmt.Println("Rename_err", err)
	}
	result := map[string]interface{}{
		"result": "successfully Saved",
		"Code":   "000",
	}
	resp(c, result)
	// fmt.Fprintf(w, "successfully saved")
}

func (a *AmsController) Amstech_upload(c *gin.Context) {
	// var w http.ResponseWriter = c.Writer
	var req *http.Request = c.Request
	req.ParseMultipartForm(10 << 20)
	file, handler, err := req.FormFile("file")
	if err != nil {
		fmt.Println("INVALID_FILE", err)
	}
	defer file.Close()
	tempFile, err := ioutil.TempFile("pkg/webui/dist/static/images/", "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("ReadFile_err", err)
	}
	tempFile.Write(fileBytes)
	mainFilePath := "pkg/webui/dist/static/images/" + handler.Filename
	renamed := os.Rename(tempFile.Name(), mainFilePath)
	if renamed != nil {
		fmt.Println("Rename_err", err)
	}
	result := map[string]interface{}{
		"result": "successfully Saved",
		"Code":   "000",
	}
	resp(c, result)
}

type logs_st struct {
	Name string `json:"name" uri:"name" form:"name"`
}

func (a *AmsController) AMSlogs(c *gin.Context) {
	var logSt logs_st
	if a.BindAndValidate(c, &logSt) {
		src := "logs/" + logSt.Name
		// Open compressed file
		gzipFile, err := os.Open(src)
		if err != nil {
			log.Fatal(err)
		}

		// Create a gzip reader on top of the file reader
		// Again, it could be any type reader though
		gzipReader, err := gzip.NewReader(gzipFile)
		if err != nil {
			log.Fatal(err)
		}
		defer gzipReader.Close()

		// Uncompress to a writer. We'll use a file writer
		outfileWriter, err := os.Create("test.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer outfileWriter.Close()

		// Copy contents of gzipped file to output file
		_, err = io.Copy(outfileWriter, gzipReader)
		if err != nil {
			log.Fatal(err)
		}
		content, err := ioutil.ReadFile("test.txt")
		if err != nil {
			log.Fatal(err)
		}

		// Convert []byte to string and print to screen
		// break the string and show in case needed
		text := string(content)
		fmt.Println(text)
		resp(c, map[string]interface{}{
			"result":   "hitting the api",
			"filename": logSt.Name,
		})
	}

}

type delete_st struct {
	Name string `json:"name" uri:"name" form:"name"`
}

func (a *AmsController) Amstech_remove(c *gin.Context) {
	// var w http.ResponseWriter = c.Writer
	var delSt delete_st
	if a.BindAndValidate(c, &delSt) {
		path := "pkg/webui/dist/static/images/" + delSt.Name
		errtorem := os.Remove(path)
		if errtorem != nil {
			fmt.Println(errtorem)
		}
		result := map[string]interface{}{
			"result": "successfully Removed",
			"Code":   "000",
		}
		resp(c, result)
	}
}

func (a *AmsController) Amstech_removeapk(c *gin.Context) {
	// var w http.ResponseWriter = c.Writer
	var delSt delete_st
	if a.BindAndValidate(c, &delSt) {
		path := "pkg/webui/dist/static/apk/" + delSt.Name
		errtorem := os.Remove(path)
		if errtorem != nil {
			fmt.Println(errtorem)
		}
		result := map[string]interface{}{
			"result": "successfully Removed",
			"Code":   "000",
		}
		resp(c, result)
	}
}

func (a *AmsController) Amstech_deliverynotes(c *gin.Context) {
	// var w http.ResponseWriter = c.Writer
	var req *http.Request = c.Request
	req.ParseMultipartForm(10 << 20)
	file, handler, err := req.FormFile("file")
	if err != nil {
		fmt.Println("INVALID_FILE", err)
	}
	defer file.Close()
	tempFile, err := ioutil.TempFile("pkg/webui/dist/static/deliverynotes/", "upload-*.pdf")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("ReadFile_err", err)
	}
	tempFile.Write(fileBytes)
	mainFilePath := "pkg/webui/dist/static/deliverynotes/" + handler.Filename
	renamed := os.Rename(tempFile.Name(), mainFilePath)
	if renamed != nil {
		fmt.Println("Rename_err", err)
	}
	result := map[string]interface{}{
		"result": "successfully Saved",
		"Code":   "000",
	}
	resp(c, result)
}
func (a *AmsController) Amstech_removedeliverynotes(c *gin.Context) {
	// var w http.ResponseWriter = c.Writer
	var delSt delete_st
	if a.BindAndValidate(c, &delSt) {
		path := "pkg/webui/dist/static/deliverynotes/" + delSt.Name
		errtorem := os.Remove(path)
		if errtorem != nil {
			fmt.Println(errtorem)
		}
		result := map[string]interface{}{
			"result": "successfully Deleted",
			"Code":   "000",
		}
		resp(c, result)
	}
}

func (a *AmsController) Amstech_uploadrefundfile(c *gin.Context) {
	// var w http.ResponseWriter = c.Writer
	var req *http.Request = c.Request
	req.ParseMultipartForm(10 << 20)
	file, handler, err := req.FormFile("file")
	if err != nil {
		fmt.Println("INVALID_FILE", err)
	}
	defer file.Close()
	tempFile, err := ioutil.TempFile("pkg/webui/dist/static/report/", "upload-*.pdf")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("ReadFile_err", err)
	}
	tempFile.Write(fileBytes)
	mainFilePath := "pkg/webui/dist/static/report/" + handler.Filename
	renamed := os.Rename(tempFile.Name(), mainFilePath)
	if renamed != nil {
		fmt.Println("Rename_err", err)
	}
	result := map[string]interface{}{
		"result": "successfully Saved",
		"Code":   "000",
	}
	resp(c, result)
}
func (a *AmsController) Amstech_removerefundfile(c *gin.Context) {
	// var w http.ResponseWriter = c.Writer
	var delSt delete_st
	if a.BindAndValidate(c, &delSt) {
		path := "pkg/webui/dist/static/report/" + delSt.Name
		errtorem := os.Remove(path)
		if errtorem != nil {
			fmt.Println(errtorem)
		}
		result := map[string]interface{}{
			"result": "successfully Deleted",
			"Code":   "000",
		}
		resp(c, result)
	}
}
