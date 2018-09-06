// Protocol Buffers for Go with Gadgets
//
// Copyright (c) 2018, The GoGo Authors. All rights reserved.
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

package issue466

import (
	"fmt"
	// "log"
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
)

func TestAny(t *testing.T) {
	em := &EventMessage{}
	em.Event = "aa"

	ff := &Foo{
		Bar: "bb",
		Baz: int64(1),
	}
	a, err := types.MarshalAny(ff)
	if err != nil {
		panic(err)
	}
	em.Myany = a

	b, err := proto.Marshal(em)
	if err != nil {
		panic(err)
	}
	fmt.Printf("buytes: %v\n", b)

	newEm := &EventMessage{}
	proto.Unmarshal(b, newEm)
	fmt.Printf("newEM: %#v\n", newEm)

	fmt.Printf("value: %v\n", newEm.Myany)
	var _ proto.Message = (*types.Any)(newEm.Myany)
	myProtoMessageFunc(newEm.Myany)
	mynewfoo := &Foo{}
	err = types.UnmarshalAny(newEm.Myany, mynewfoo)
	if err != nil {
		panic(err)
	}
	fmt.Printf("mynewfoo: %v\n", mynewfoo)
}

func myProtoMessageFunc(pb proto.Message) {
	fmt.Printf("Here is a pb message: %+v\n", pb)
}
