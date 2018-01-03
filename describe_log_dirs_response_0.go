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

type DescribeLogDirsResponse0 struct {
    ThrottleTimeMs int32
    LogDirs []*DescribeLogDirsResponse0_LogDirs
}

func (that *DescribeLogDirsResponse0) Encode(enc *Encoder) error {
    enc.WriteInt32(that.ThrottleTimeMs)
    {
        arrayLength := len(that.LogDirs)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.LogDirs[i].Encode(enc)
        }
    }

    return nil
}

func (that *DescribeLogDirsResponse0) Decode(dec *Decoder) error {
    that.ThrottleTimeMs = dec.ReadInt32()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.LogDirs = nil
        } else {
            buf := make([]*DescribeLogDirsResponse0_LogDirs, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(DescribeLogDirsResponse0_LogDirs)
                item.Decode(dec)
                buf[i] = item
            }
            that.LogDirs = buf
        }
    }
    return nil
}

type DescribeLogDirsResponse0_Partitions struct {
    Partition int32
    Size int64
    OffsetLag int64
    IsFuture bool
}

func (that *DescribeLogDirsResponse0_Partitions) Encode(enc *Encoder) error {
    enc.WriteInt32(that.Partition)
    enc.WriteInt64(that.Size)
    enc.WriteInt64(that.OffsetLag)
    enc.WriteBool(that.IsFuture)
    return nil
}

func (that *DescribeLogDirsResponse0_Partitions) Decode(dec *Decoder) error {
    that.Partition = dec.ReadInt32()
    that.Size = dec.ReadInt64()
    that.OffsetLag = dec.ReadInt64()
    that.IsFuture = dec.ReadBool()
    return nil
}

type DescribeLogDirsResponse0_Topics struct {
    Topic string
    Partitions []*DescribeLogDirsResponse0_Partitions
}

func (that *DescribeLogDirsResponse0_Topics) Encode(enc *Encoder) error {
    enc.WriteString(that.Topic)
    {
        arrayLength := len(that.Partitions)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.Partitions[i].Encode(enc)
        }
    }

    return nil
}

func (that *DescribeLogDirsResponse0_Topics) Decode(dec *Decoder) error {
    that.Topic = dec.ReadString()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.Partitions = nil
        } else {
            buf := make([]*DescribeLogDirsResponse0_Partitions, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(DescribeLogDirsResponse0_Partitions)
                item.Decode(dec)
                buf[i] = item
            }
            that.Partitions = buf
        }
    }
    return nil
}

type DescribeLogDirsResponse0_LogDirs struct {
    ErrorCode int16
    LogDir string
    Topics []*DescribeLogDirsResponse0_Topics
}

func (that *DescribeLogDirsResponse0_LogDirs) Encode(enc *Encoder) error {
    enc.WriteInt16(that.ErrorCode)
    enc.WriteString(that.LogDir)
    {
        arrayLength := len(that.Topics)
        enc.WriteInt32(int32(arrayLength))
        for i := 0; i < arrayLength; i++ {
            that.Topics[i].Encode(enc)
        }
    }

    return nil
}

func (that *DescribeLogDirsResponse0_LogDirs) Decode(dec *Decoder) error {
    that.ErrorCode = dec.ReadInt16()
    that.LogDir = dec.ReadString()
    {
        arrayLength := dec.ReadInt32()
        if int(arrayLength) == -1 {
            that.Topics = nil
        } else {
            buf := make([]*DescribeLogDirsResponse0_Topics, arrayLength)
            var i int32
            for i = 0; i < arrayLength; i++ {
                item := new(DescribeLogDirsResponse0_Topics)
                item.Decode(dec)
                buf[i] = item
            }
            that.Topics = buf
        }
    }
    return nil
}

