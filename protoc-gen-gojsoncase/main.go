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

package main

import (
	"github.com/gogo/protobuf/gogoproto"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	"github.com/gogo/protobuf/vanity"
	"github.com/gogo/protobuf/vanity/command"
)

func main() {
	req := command.Read()
	files := req.GetProtoFile()

	vanity.ForEachFieldInFiles(files, JsonTagCamelCase)

	resp := command.Generate(req)
	command.Write(resp)
}

func JsonTagCamelCase(field *descriptor.FieldDescriptorProto) {
	if field.IsRepeated() || field.IsMessage() {
		return
	}
	if field.DefaultValue != nil {
		return
	}
	SetStringFieldOption(gogoproto.E_Jsontag, generator.CamelCase(*field.Name))(field)
}

func SetStringFieldOption(extension *proto.ExtensionDesc, value string) func(field *descriptor.FieldDescriptorProto) {
	return func(field *descriptor.FieldDescriptorProto) {
		if FieldHasStringExtension(field, extension) {
			return
		}
		if field.Options == nil {
			field.Options = &descriptor.FieldOptions{}
		}
		if err := proto.SetExtension(field.Options, extension, &value); err != nil {
			panic(err)
		}
	}
}

func FieldHasStringExtension(field *descriptor.FieldDescriptorProto, extension *proto.ExtensionDesc) bool {
	if field.Options == nil {
		return false
	}
	value, err := proto.GetExtension(field.Options, extension)
	if err != nil {
		return false
	}
	if value == nil {
		return false
	}
	if value.(*string) == nil {
		return false
	}
	return true
}
