package issue580

import (
	"bytes"
	"fmt"
	testing "testing"

	"github.com/gogo/protobuf/jsonpb"
)

func TestMyMessageJSONProto(t *testing.T) {
	msg := `
{
    "e": "CASE1"
}
`
	plan := &MyMessage{}
	err := jsonpb.Unmarshal(bytes.NewBufferString(msg), plan)
	if err != nil {
		fmt.Println("Error " + err.Error())
		return
	}
	fmt.Println("msg " + plan.String())
	return
}
