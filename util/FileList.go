package util

import "encoding/json"

type FileRequestBody struct {
	All                   bool   `json:"all"`
	DriveID               string `json:"drive_id"`
	Fields                string `json:"fields"`
	ImageThumbnailProcess string `json:"image_thumbnail_process"`
	ImageURLProcess       string `json:"image_url_process"`
	Limit                 int64  `json:"limit"`
	OrderBy               string `json:"order_by"`
	OrderDirection        string `json:"order_direction"`
	ParentFileID          string `json:"parent_file_id"`
	URLExpireSEC          int64  `json:"url_expire_sec"`
	VideoThumbnailProcess string `json:"video_thumbnail_process"`
}

func (f *FileRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"drive_id":                f.DriveID,
		"parent_file_id":          f.ParentFileID,
		"limit":                   f.Limit,
		"all":                     f.All,
		"url_expire_sec":          f.URLExpireSEC,
		"image_thumbnail_process": f.ImageThumbnailProcess,
		"image_url_process":       f.ImageURLProcess,
		"video_thumbnail_process": f.VideoThumbnailProcess,
		"fields":                  f.Fields,
		"order_by":                f.OrderBy,
		"order_direction":         f.OrderDirection,
	})
}
