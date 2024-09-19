package client

import (
  "context"
  "io"
  reapi "github.com/bazelbuild/remote-apis/build/bazel/remote/execution/v2"
  "github.com/golang/protobuf/proto"
  "google.golang.org/grpc"
  "google.golang.org/genproto/googleapis/bytestream"
)

func Expect(c *grpc.ClientConn, d *reapi.Digest, m proto.Message) error {
  bs := bytestream.NewByteStreamClient(c)

  bsrc, err := bs.Read(context.Background(), &bytestream.ReadRequest {
    ResourceName: "/blobs/" + DigestString(d),
  })
  if err != nil {
    return err
  }

  var b []byte
  for ;; {
    br, err := bsrc.Recv()
    if err == io.EOF {
      break
    }
    if int64(len(br.Data)) == d.SizeBytes {
      b = br.Data
    }
    err = proto.Unmarshal(b, m)
    if err != nil {
      return err
    }
    return nil
  }
  return io.EOF
}
