package m

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

type ABACRequest struct {
	AS AS
	AO AO
}

func (r *ABACRequest) ToBytes() []byte {
	b, err := json.Marshal(*r)
	if err != nil {
		return nil
	}
	return b
}

type Attrs struct {
	DeviceId  string
	UserId    string
	Timestamp int64
}

func (a Attrs) GetId() string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(a.UserId+a.DeviceId)))
}

func (r ABACRequest) GetAttrs() Attrs {
	return Attrs{DeviceId: r.AO.DeviceId, UserId: r.AS.UserId, Timestamp: time.Now().Unix()}
}
