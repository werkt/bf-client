package client

import (
  "errors"
  reapi "github.com/bazelbuild/remote-apis/build/bazel/remote/execution/v2"
  bfpb "github.com/buildfarm/buildfarm/build/buildfarm/v1test"
  "github.com/golang/protobuf/ptypes"
  "github.com/golang/protobuf/proto"
  "google.golang.org/genproto/googleapis/longrunning"
)

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
  return nil, errors.New("Unexpected metadata: " + proto.MarshalTextString(m))
}
