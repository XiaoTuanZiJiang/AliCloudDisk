package entity

import (
	"bytes"
	"encoding/json"
	"reflect"
	"strings"
	"time"
)

type ResourcesMap map[string]Items

type TooLTT struct {
	Items []Items `json:"items"`
}

type Items struct {
	CreatedAt     time.Time    `json:"created_at"`
	DriveID       string       `json:"drive_id"`
	FileID        string       `json:"file_id"`
	Name          string       `json:"name"`
	ParentFileID  string       `json:"parent_file_id"`
	Starred       bool         `json:"starred"`
	Type          string       `json:"type"`
	UpdatedAt     time.Time    `json:"updated_at"`
	SyncMeta      string       `json:"sync_meta,omitempty"`
	Category      string       `json:"category,omitempty"`
	ContentHash   string       `json:"content_hash,omitempty"`
	FileExtension string       `json:"file_extension,omitempty"`
	MimeType      string       `json:"mime_type,omitempty"`
	PunishFlag    int          `json:"punish_flag,omitempty"`
	Size          int          `json:"size,omitempty"`
	Item          ResourcesMap `json:"item"`
}

func (i Items) String() string {
	bs, _ := json.MarshalIndent(i, "", " ")
	var out bytes.Buffer
	err := json.Indent(&out, bs, "", "\t")
	if err != nil {
		return ""
	}
	return out.String()
}

func (r ResourcesMap) String() string {
	bs, _ := json.MarshalIndent(r, "", " ")
	var out bytes.Buffer
	err := json.Indent(&out, bs, "", "\t")
	if err != nil {
		return ""
	}
	return out.String()
}

func (r ResourcesMap) FindResInfoByName(Name string) Items {
	return rangeResourcesMapOfName(&r, Name)
}

func rangeResourcesMapOfName(rm *ResourcesMap, Name string) Items {
	for k, v := range *rm {
		if strings.EqualFold(k, Name) {
			return v
		} else if !reflect.ValueOf(v.Item).IsNil() {
			item := rangeResourcesMapOfName(&v.Item, Name)
			if !reflect.DeepEqual(item, Items{}) {
				return item
			}
		}
	}
	return Items{}
}
