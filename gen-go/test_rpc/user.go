// Code generated by Thrift Compiler (0.17.0). DO NOT EDIT.

package test_rpc

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"time"
	thrift "github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = errors.New
var _ = context.Background
var _ = time.Now
var _ = bytes.Equal

// Attributes:
//  - UserID
type TestReq struct {
  UserID *int64 `thrift:"user_id,1" db:"user_id" json:"user_id,omitempty"`
}

func NewTestReq() *TestReq {
  return &TestReq{}
}

var TestReq_UserID_DEFAULT int64
func (p *TestReq) GetUserID() int64 {
  if !p.IsSetUserID() {
    return TestReq_UserID_DEFAULT
  }
return *p.UserID
}
func (p *TestReq) IsSetUserID() bool {
  return p.UserID != nil
}

func (p *TestReq) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.I64 {
        if err := p.ReadField1(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *TestReq)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(ctx); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.UserID = &v
}
  return nil
}

func (p *TestReq) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "test_req"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *TestReq) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if p.IsSetUserID() {
    if err := oprot.WriteFieldBegin(ctx, "user_id", thrift.I64, 1); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:user_id: ", p), err) }
    if err := oprot.WriteI64(ctx, int64(*p.UserID)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.user_id (1) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(ctx); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 1:user_id: ", p), err) }
  }
  return err
}

func (p *TestReq) Equals(other *TestReq) bool {
  if p == other {
    return true
  } else if p == nil || other == nil {
    return false
  }
  if p.UserID != other.UserID {
    if p.UserID == nil || other.UserID == nil {
      return false
    }
    if (*p.UserID) != (*other.UserID) { return false }
  }
  return true
}

func (p *TestReq) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("TestReq(%+v)", *p)
}

// Attributes:
//  - UserID
//  - Age
//  - Name
type TestRes struct {
  UserID *int64 `thrift:"user_id,1" db:"user_id" json:"user_id,omitempty"`
  Age *int32 `thrift:"age,2" db:"age" json:"age,omitempty"`
  Name *string `thrift:"name,3" db:"name" json:"name,omitempty"`
}

func NewTestRes() *TestRes {
  return &TestRes{}
}

var TestRes_UserID_DEFAULT int64
func (p *TestRes) GetUserID() int64 {
  if !p.IsSetUserID() {
    return TestRes_UserID_DEFAULT
  }
return *p.UserID
}
var TestRes_Age_DEFAULT int32
func (p *TestRes) GetAge() int32 {
  if !p.IsSetAge() {
    return TestRes_Age_DEFAULT
  }
return *p.Age
}
var TestRes_Name_DEFAULT string
func (p *TestRes) GetName() string {
  if !p.IsSetName() {
    return TestRes_Name_DEFAULT
  }
return *p.Name
}
func (p *TestRes) IsSetUserID() bool {
  return p.UserID != nil
}

func (p *TestRes) IsSetAge() bool {
  return p.Age != nil
}

func (p *TestRes) IsSetName() bool {
  return p.Name != nil
}

func (p *TestRes) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.I64 {
        if err := p.ReadField1(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField2(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 3:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField3(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *TestRes)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(ctx); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.UserID = &v
}
  return nil
}

func (p *TestRes)  ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(ctx); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Age = &v
}
  return nil
}

func (p *TestRes)  ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(ctx); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  p.Name = &v
}
  return nil
}

func (p *TestRes) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "test_res"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
    if err := p.writeField2(ctx, oprot); err != nil { return err }
    if err := p.writeField3(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *TestRes) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if p.IsSetUserID() {
    if err := oprot.WriteFieldBegin(ctx, "user_id", thrift.I64, 1); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:user_id: ", p), err) }
    if err := oprot.WriteI64(ctx, int64(*p.UserID)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.user_id (1) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(ctx); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 1:user_id: ", p), err) }
  }
  return err
}

func (p *TestRes) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if p.IsSetAge() {
    if err := oprot.WriteFieldBegin(ctx, "age", thrift.I32, 2); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:age: ", p), err) }
    if err := oprot.WriteI32(ctx, int32(*p.Age)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.age (2) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(ctx); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 2:age: ", p), err) }
  }
  return err
}

func (p *TestRes) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if p.IsSetName() {
    if err := oprot.WriteFieldBegin(ctx, "name", thrift.STRING, 3); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:name: ", p), err) }
    if err := oprot.WriteString(ctx, string(*p.Name)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.name (3) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(ctx); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 3:name: ", p), err) }
  }
  return err
}

func (p *TestRes) Equals(other *TestRes) bool {
  if p == other {
    return true
  } else if p == nil || other == nil {
    return false
  }
  if p.UserID != other.UserID {
    if p.UserID == nil || other.UserID == nil {
      return false
    }
    if (*p.UserID) != (*other.UserID) { return false }
  }
  if p.Age != other.Age {
    if p.Age == nil || other.Age == nil {
      return false
    }
    if (*p.Age) != (*other.Age) { return false }
  }
  if p.Name != other.Name {
    if p.Name == nil || other.Name == nil {
      return false
    }
    if (*p.Name) != (*other.Name) { return false }
  }
  return true
}

func (p *TestRes) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("TestRes(%+v)", *p)
}

