package m

import "encoding/json"

// Define the model resource structure
type TierModel struct { // 层聚合模型
	ModelID        string `json:"ModelID"`
	ModelType      int `json:"ModelType"`       // 0: 层聚合模型
	ModelTier      int `json:"ModelTier"`
	TierRound      int `json:"TierRound"`
	GlobalRound    int `json:"GlobalRound"`
	ModelEncryAddr string `json:ModelEncryAddr` // 模型的 IPFS 保存地址
}

type GlobalModel struct { // 全局聚合模型
	ModelID        string `json:"ModelID"`
	ModelType      int `json:"ModelType"`       // 1: 代表全局聚合模型
	GlobalRound    int `json:"GlobalRound"`
	ModelEncryAddr string `json:ModelEncryAddr` // 模型的 IPFS 保存地址
}

func (tm TierModel) ToBytes() []byte {
	bs, err := json.Marshal(tm)
	if err != nil {
		return nil
	}
	return bs
}

func (gm GlobalModel) ToBytes() []byte {
	bs, err := json.Marshal(gm)
	if err != nil {
		return nil
	}
	return bs
}


// func NewResource(b []byte) (Resource, error) {
// 	r := Resource{}
// 	err := json.Unmarshal(b, &r)
// 	return r, err
// }
