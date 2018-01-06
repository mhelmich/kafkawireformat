/*
 * Copyright 2018 Marco Helmich
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package kafkawireformat

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestVarIntSerialization(t *testing.T) {
    i := int64(11)
    enc := NewEncoder()
    enc.WriteVarInt(i)
    assert.Equal(t, 1, enc.Len(), "Length isn't quite right")
    dec := NewDecoderFromBytes(enc.Bytes())
    i2, err := dec.ReadVarInt()
    t.Logf("Got original: %d loaded: %d error: %s", i, i2, err)
    assert.Nil(t, err, err)
    assert.Equal(t, i, i2, "They should be equal")
}

func TestVarIntArraySerialization(t *testing.T) {
    bites := []byte("elevenchars")
    enc := NewEncoder()
    enc.WriteVarIntByteArray(bites)
    assert.Equal(t, len(bites)+1, enc.Len(), "Length isn't quite right")
    dec := NewDecoderFromBytes(enc.Bytes())
    bites2 := dec.ReadVarIntByteArray()
    t.Logf("bites: %x, bites2: %x", bites, bites2)
    assert.Equal(t, bites, bites2, "bytes aren't the same")

    dec2 := NewDecoderFromBytes(enc.Bytes())
    i, err := dec2.ReadVarInt()
    t.Logf("varint array size: %d %x", i, err)
    assert.Equal(t, int64(len(bites)), i, "varint byte array size wasn't right")
}
