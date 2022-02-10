package coolq

import (
	"encoding/json"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

func TestEnCode(t *testing.T) {
	var benchArray = gjson.Parse(`[{"type":"text","data":{"text":"asdfqwerqwerqwer"}},{"type":"face","data":{"id":"115","text":"111"}},{"type":"text","data":{"text":"asdfasdfasdfasdfasdfasdfasd"}},{"type":"face","data":{"id":"217"}},{"type":"text","data":{"text":"] "}},{"type":"text","data":{"text":"123"}},{"type":"text","data":{"text":" ["}}]`)
	var msg []MsgElem
	_ = json.Unmarshal([]byte(benchArray.String()), &msg)
	str := EnCodes(msg)
	log.Infof("code=%s", str)
}
