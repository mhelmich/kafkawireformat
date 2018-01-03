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
import ()

type CreateTopicsResponse2 struct {
    ThrottleTimeMs int32
    TopicErrors []*CreateTopicsResponse2_TopicErrors
}

func (that *CreateTopicsResponse2) Encode(enc *Encoder) error {
    enc.WriteInt32(that.ThrottleTimeMs)
    {
        arrayLength := len(that.TopicErrors)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.TopicErrors[i].Encode(enc)
        }
    }

    return nil
}

func (that *CreateTopicsResponse2) Decode(dec *Decoder) error {
    that.ThrottleTimeMs = dec.ReadInt32()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.TopicErrors = nil
        } else {
            buf := make([]*CreateTopicsResponse2_TopicErrors, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(CreateTopicsResponse2_TopicErrors)
                item.Decode(dec)
                buf[i] = item
            }
            that.TopicErrors = buf
        }
    }
    return nil
}

type CreateTopicsResponse2_TopicErrors struct {
    Topic string
    ErrorCode int16
    ErrorMessage string
}

func (that *CreateTopicsResponse2_TopicErrors) Encode(enc *Encoder) error {
    enc.WriteString(that.Topic)
    enc.WriteInt16(that.ErrorCode)
    enc.WriteString(that.ErrorMessage)
    return nil
}

func (that *CreateTopicsResponse2_TopicErrors) Decode(dec *Decoder) error {
    that.Topic = dec.ReadString()
    that.ErrorCode = dec.ReadInt16()
    that.ErrorMessage = dec.ReadString()
    return nil
}

