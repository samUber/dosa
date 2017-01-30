// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package dosa

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCQL(t *testing.T) {
	data := []struct {
		Instance  DomainObject
		Statement string
	}{
		{
			Instance:  &SinglePrimaryKey{},
			Statement: `create table "singleprimarykey" ("primarykey" bigint, "data" text, primary key (primarykey))`,
		},
		{
			Instance:  &AllTypes{},
			Statement: `create table "alltypes" ("booltype" boolean, "int32type" int, "int64type" bigint, "doubletype" double, "stringtype" text, "blobtype" blob, "timetype" timestamp, "uuidtype" uuid, primary key (booltype))`,
		},
		// TODO: Add more test cases
	}

	for _, d := range data {
		table, err := TableFromInstance(d.Instance)
		assert.Nil(t, err) // this code does not test TableFromInstance
		statement := table.EntityDefinition.CqlCreateTable()
		assert.Equal(t, statement, d.Statement, fmt.Sprintf("Instance: %T", d.Instance))
	}
}

func BenchmarkCQL(b *testing.B) {
	table, _ := TableFromInstance(&AllTypes{})
	for i := 0; i < b.N; i++ {
		table.EntityDefinition.CqlCreateTable()
	}
}

func TestTypemapUnknown(t *testing.T) {
	assert.Equal(t, "unknown", typemap(Invalid))
}
