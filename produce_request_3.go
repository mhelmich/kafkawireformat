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

type ProduceRequest3_TopicData struct {
    Topic string
    Data []*ProduceRequest3_Data
}

func (that *ProduceRequest3_TopicData) Encode(enc *Encoder) error {
    enc.WriteString(that.Topic)
    {
        arrayLength := len(that.Data)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.Data[i].Encode(enc)
        }
    }

    return nil
}

func (that *ProduceRequest3_TopicData) Decode(dec *Decoder) error {
    that.Topic = dec.ReadString()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.Data = nil
        } else {
            buf := make([]*ProduceRequest3_Data, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(ProduceRequest3_Data)
                item.Decode(dec)
                buf[i] = item
            }
            that.Data = buf
        }
    }
    return nil
}

type ProduceRequest3_Data struct {
    Partition int32
    RecordSet *RecordBatch
}

func (that *ProduceRequest3_Data) Encode(enc *Encoder) error {
    enc.WriteInt32(that.Partition)
    that.RecordSet.Encode(enc)

    return nil
}

func (that *ProduceRequest3_Data) Decode(dec *Decoder) error {
    that.Partition = dec.ReadInt32()
    that.RecordSet = new(RecordBatch)
    that.RecordSet.Decode(dec)
    return nil
}

type ProduceRequest3 struct {
    TransactionalId string
    Acks int16
    Timeout int32
    TopicData []*ProduceRequest3_TopicData
}

func (that *ProduceRequest3) Encode(enc *Encoder) error {
    enc.WriteString(that.TransactionalId)
    enc.WriteInt16(that.Acks)
    enc.WriteInt32(that.Timeout)
    {
        arrayLength := len(that.TopicData)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.TopicData[i].Encode(enc)
        }
    }

    return nil
}

func (that *ProduceRequest3) Decode(dec *Decoder) error {
    that.TransactionalId = dec.ReadString()
    that.Acks = dec.ReadInt16()
    that.Timeout = dec.ReadInt32()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.TopicData = nil
        } else {
            buf := make([]*ProduceRequest3_TopicData, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(ProduceRequest3_TopicData)
                item.Decode(dec)
                buf[i] = item
            }
            that.TopicData = buf
        }
    }
    return nil
}

