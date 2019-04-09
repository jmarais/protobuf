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

package test

import (
	"testing"
)

var result []byte

func generateWide(width int) *Test {
	test := Test{
		Tests: map[string]*Test{},
	}
	for i := 0; i < width; i++ {
		test.Tests[string(int('a')+i)] = nil
	}
	return &test
}

func generateDeep(depth int) *Test {
	test := Test{
		Tests: nil,
	}
	current := &test
	for i := 0; i < depth; i++ {
		current.Tests = map[string]*Test{
			string(int('a') + i): &Test{
				Tests: nil,
			},
		}
		current = current.Tests[string(int('a')+i)]
	}
	return &test
}

func BenchmarkWideProtobuf1(b *testing.B) {
	test := generateWide(1)
	var data []byte
	for n := 0; n < b.N; n++ {
		data, _ = test.Marshal()
	}
	result = data
}

func BenchmarkWideProtobuf2(b *testing.B) {
	test := generateWide(2)
	var data []byte
	for n := 0; n < b.N; n++ {
		data, _ = test.Marshal()
	}
	result = data
}

func BenchmarkWideProtobuf5(b *testing.B) {
	test := generateWide(5)
	var data []byte
	for n := 0; n < b.N; n++ {
		data, _ = test.Marshal()
	}
	result = data
}

func BenchmarkWideProtobuf10(b *testing.B) {
	test := generateWide(10)
	var data []byte
	for n := 0; n < b.N; n++ {
		data, _ = test.Marshal()
	}
	result = data
}

func BenchmarkWideProtobuf20(b *testing.B) {
	test := generateWide(20)
	var data []byte
	for n := 0; n < b.N; n++ {
		data, _ = test.Marshal()
	}
	result = data
}

func BenchmarkWideProtobuf50(b *testing.B) {
	test := generateWide(50)
	var data []byte
	for n := 0; n < b.N; n++ {
		data, _ = test.Marshal()
	}
	result = data
}

func BenchmarkWideProtobuf100(b *testing.B) {
	test := generateWide(100)
	var data []byte
	for n := 0; n < b.N; n++ {
		data, _ = test.Marshal()
	}
	result = data
}

func BenchmarkDeepProtobuf1(b *testing.B) {
	test := generateDeep(1)
	var data []byte
	for n := 0; n < b.N; n++ {
		data, _ = test.Marshal()
	}
	result = data
}

func BenchmarkDeepProtobuf2(b *testing.B) {
	test := generateDeep(2)
	var data []byte
	for n := 0; n < b.N; n++ {
		data, _ = test.Marshal()
	}
	result = data
}

func BenchmarkDeepProtobuf5(b *testing.B) {
	test := generateDeep(5)
	var data []byte
	for n := 0; n < b.N; n++ {
		data, _ = test.Marshal()
	}
	result = data
}

func BenchmarkDeepProtobuf10(b *testing.B) {
	test := generateDeep(10)
	var data []byte
	for n := 0; n < b.N; n++ {
		data, _ = test.Marshal()
	}
	result = data
}

func BenchmarkDeepProtobuf20(b *testing.B) {
	test := generateDeep(20)
	var data []byte
	for n := 0; n < b.N; n++ {
		data, _ = test.Marshal()
	}
	result = data
}

func BenchmarkDeepProtobuf50(b *testing.B) {
	test := generateDeep(50)
	var data []byte
	for n := 0; n < b.N; n++ {
		data, _ = test.Marshal()
	}
	result = data
}

func BenchmarkDeepProtobuf100(b *testing.B) {
	test := generateDeep(100)
	var data []byte
	for n := 0; n < b.N; n++ {
		data, _ = test.Marshal()
	}
	result = data
}
