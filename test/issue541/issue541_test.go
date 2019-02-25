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

package issue541

import (
	"testing"
	"strings"
	"fmt"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/types"

)

func TestJsonpbMarshalUnmarshal(t *testing.T) {
	msg := &OuterMessage{
		Fieldone: "Hello",
		Inner: &InnerMessage{
			Fieldtwo: "Sailboat",
			Fieldthree: "Sailor",
		},
		OneOfFoo: &OuterMessage_FooOne{
			FooOne: int64(2),
		},
		BazListValue: &types.ListValue{
			Values:  []*types.Value{
				&types.Value{
					Kind: &types.Value_StringValue{
						StringValue: "ListValueValueOne",
					},
				},
			},
		},
	}
	marshaler := &jsonpb.Marshaler{}
	str, err := marshaler.MarshalToString(msg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", str)
	result := &OuterMessage{}
	err = jsonpb.Unmarshal(strings.NewReader(str), result)
	if err != nil {
		panic(err)
	}
}
