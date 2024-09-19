package client

import (
  "crypto"
  "encoding/hex"
  "fmt"
  "github.com/golang/protobuf/proto"
  reapi "github.com/bazelbuild/remote-apis/build/bazel/remote/execution/v2"
)

func DigestString(d *reapi.Digest) string {
  if d == nil {
    return "nil digest"
  }
  return fmt.Sprintf("%s/%d", d.Hash, d.SizeBytes)
}

func DigestFromBlob(blob []byte, hashFn crypto.Hash) reapi.Digest {
  h := hashFn.New()
  h.Write(blob)
  arr := h.Sum(nil)
  return reapi.Digest{Hash: hex.EncodeToString(arr[:]), SizeBytes: int64(len(blob))}
}

func DigestFromMessage(msg proto.Message, hashFn crypto.Hash) (reapi.Digest, error) {
  blob, err := proto.Marshal(msg)
  if err != nil {
    Empty := DigestFromBlob([]byte{}, hashFn)
    return Empty, err
  }
  return DigestFromBlob(blob, hashFn), nil
}
