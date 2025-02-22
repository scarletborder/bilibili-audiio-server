package model

type SongDetail struct {
	Aid     int    `json:"aid"`
	Bvid    string `json:"bvid"`
	Title   string `json:"title"`
	Desc    string `json:"desc,omitempty"`
	Cover   string `json:"cover"`
	Artist  Artist `json:"artist"`
	Songs   []Song `json:"songs,omitempty"`
	HasPart int   `json:"has_part,omitempty"`
}

type Song struct {
	Cid  int    `json:"cid"`
	Name string `json:"name"`
}
