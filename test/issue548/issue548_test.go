// Protocol Buffers for Go with Gadgets
//
// Copyright (c) 2019, The GoGo Authors. All rights reserved.
// http://github.com/gogo/protobuf
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package issue548

import (
	"fmt"
	"bytes"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
)

// Original test case submitted by issue creator
func TestMarshalUnMarshal(t *testing.T) {
	msg1 := &Request{
		RequestId: "testReqId",
		Kind: &Request_SessionStatus{
			SessionStatus: &RequestSessionStatus{
				SessionId: "testSessionID",
			},
		},
	}

	m := jsonpb.Marshaler{}
	var packedMsg bytes.Buffer
	err1 := m.Marshal(&packedMsg, msg1)
	if err1 != nil {
		t.Fatalf("err1: %v", err1)
	}
	fmt.Printf("%s\n", msg1)

	msg2 := &Request{}
	err2 := jsonpb.Unmarshal(bytes.NewReader(packedMsg.Bytes()), msg2)
	if err2 != nil {
		t.Fatalf("err2: %v", err2)
	}
	fmt.Printf("%s\n", msg2)
}
