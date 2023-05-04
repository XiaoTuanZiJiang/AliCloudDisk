package AliCloudDisk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/XiaoTuanZiJiang/AliCloudDisk/entity"
	"github.com/XiaoTuanZiJiang/AliCloudDisk/util"
	"io"
	"net/http"
	"strings"
)

type AliCloudDisk struct {
	Authorization string //Token
	UserInfo      *entity.UserInfo
	ResourcesMap  entity.ResourcesMap
}

type Config struct {
	GetResourcesMap bool
}

func NewCloudDiskConnection(Authorization string, config *Config) *AliCloudDisk {
	c := &AliCloudDisk{}
	c.Authorization = Authorization
	return initCloudDisk(c, config)
}

func initCloudDisk(c *AliCloudDisk, config *Config) *AliCloudDisk {
	c.getUserInfo()
	if config.GetResourcesMap {
		c.ResourcesMap = c.getResourcesMap("root")
	}
	return c
}

// GetUserInfo  获取用户信息
func (c *AliCloudDisk) getUserInfo() {
	url := "https://api.aliyundrive.com/adrive/v2/user/get"
	method := "POST"
	payload := strings.NewReader(`{}`)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		panic("Request Error :" + err.Error())
	}

	requestDefault(c.Authorization, req)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		panic("clientDo Error :" + err.Error())
	}
	defer resp.Body.Close()

	var byt []byte
	buffer := bytes.NewBuffer(byt)
	_, err = io.Copy(buffer, resp.Body)

	if err != nil {
		panic("ReadAllBody Error :" + err.Error())
	}
	if resp.StatusCode != 200 {
		panic("clientDo Error :" + buffer.String())
	}

	err = json.Unmarshal(buffer.Bytes(), &c.UserInfo)
	if err != nil {
		errors.New("jsonUnmarshal Error")
		return
	}
}

func (c *AliCloudDisk) getResourcesMap(parentFileId string) entity.ResourcesMap {
	url := "https://api.aliyundrive.com/adrive/v3/file/list?jsonmask=next_marker,items(name,file_id,drive_id,type,size,created_at,updated_at,category,file_extension,parent_file_id,mime_type,starred,thumbnail,url,streams_info,content_hash,user_tags,user_meta,trashed,video_media_metadata,video_preview_metadata,sync_meta,sync_device_flag,sync_flag,punish_flag)"

	method := "POST"

	frb := &util.FileRequestBody{
		DriveID:               c.UserInfo.DefaultDriveId,
		ParentFileID:          parentFileId,
		Limit:                 20,
		All:                   false,
		URLExpireSEC:          14400,
		ImageThumbnailProcess: "image/resize,w_256/format,jpeg",
		ImageURLProcess:       "image/resize,w_1920/format,jpeg/interlace,1",
		VideoThumbnailProcess: "video/snapshot,t_1000,f_jpg,ar_auto,w_256",
		Fields:                "*",
		OrderBy:               "updated_at",
		OrderDirection:        "DESC",
	}
	frbJSON, _ := frb.MarshalJSON()
	payload := strings.NewReader(string(frbJSON))
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	requestDefault(c.Authorization, req)

	res, err := client.Do(req)
	if res != nil {
		defer res.Body.Close()
	}
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var byt []byte
	buffer := bytes.NewBuffer(byt)
	_, err = io.Copy(buffer, res.Body)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	toolTT := &entity.TooLTT{}
	resMap := make(entity.ResourcesMap)
	err = json.Unmarshal(buffer.Bytes(), &toolTT)
	if err != nil {
		fmt.Println("err = ", err)
		return nil
	}
	for _, k := range toolTT.Items {
		if k.Type == "folder" {
			k.Item = c.getResourcesMap(k.FileID)
		}
		name := k.Name
		resMap[name] = k
	}
	return resMap
}

func requestDefault(Authorization string, req *http.Request) {
	req.Header.Add("Authorization", Authorization)
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://www.apifox.cn)")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Host", "api.aliyundrive.com")
	req.Header.Add("Connection", "keep-alive")
}
