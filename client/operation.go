package client

import (
  "errors"
  reapi "github.com/bazelbuild/remote-apis/build/bazel/remote/execution/v2"
  bfpb "github.com/buildfarm/buildfarm/build/buildfarm/v1test"
  "github.com/golang/protobuf/ptypes"
  "github.com/golang/protobuf/proto"
  "google.golang.org/genproto/googleapis/longrunning"
)

type Operation struct {
  Name string
  Metadata *reapi.RequestMetadata
}

func ExecuteOperationMetadata(op *longrunning.Operation) (*reapi.ExecuteOperationMetadata, error) {
  m := op.Metadata
  em := &reapi.ExecuteOperationMetadata{}
  qm := &bfpb.QueuedOperationMetadata{}
  if ptypes.Is(m, em) {
    if err := ptypes.UnmarshalAny(m, em); err != nil {
      return nil, err
    }
    return em, nil
  } else if ptypes.Is(m, qm) {
    if err := ptypes.UnmarshalAny(m, qm); err != nil {
      return nil, err
    }
    return qm.ExecuteOperationMetadata, nil
  }
  return nil, errors.New("Unexpected metadata: " + proto.MarshalTextString(op))
}

func RequestMetadata(o *longrunning.Operation) *reapi.RequestMetadata {
  m := o.Metadata
  em := &reapi.ExecuteOperationMetadata{}
  qm := &bfpb.QueuedOperationMetadata{}
  if ptypes.Is(m, em) {
    return nil
  } else if ptypes.Is(m, qm) {
    if err := ptypes.UnmarshalAny(m, qm); err != nil {
      return nil
    } else {
      return qm.RequestMetadata
    }
  } else {
    return nil
  }
}

func ExecutedActionMetadata(o *longrunning.Operation) (*reapi.ExecutedActionMetadata, error) {
  switch r := o.Result.(type) {
  case *longrunning.Operation_Response:
    er := &reapi.ExecuteResponse{}
    if ptypes.Is(r.Response, er) {
      err := ptypes.UnmarshalAny(r.Response, er)
      if err == nil && er.Result != nil {
        if er.Result.ExecutionMetadata == nil {
          panic(er.Result)
        }
        return er.Result.ExecutionMetadata, nil
      }
      if err == nil {
        err = errors.New("ExecuteResponse.Result was nil for: " + proto.MarshalTextString(o))
      }
      return nil, err
    }
  }

  em, err := ExecuteOperationMetadata(o)
  if err != nil {
    return nil, err
  }
  if em == nil || em.PartialExecutionMetadata == nil {
    return &reapi.ExecutedActionMetadata{}, nil
  }
  return em.PartialExecutionMetadata, nil
}
