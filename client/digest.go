package client

import (
  "crypto"
  "encoding/hex"
  "fmt"
  "github.com/golang/protobuf/proto"
  reapi "github.com/bazelbuild/remote-apis/build/bazel/remote/execution/v2"
  bfpb "github.com/buildfarm/buildfarm/build/buildfarm/v1test"
)

func FromDigest(d bfpb.Digest) reapi.Digest {
  return reapi.Digest{Hash: d.Hash, SizeBytes: d.Size}
}

func ToDigest(d reapi.Digest, df reapi.DigestFunction_Value) bfpb.Digest {
  return bfpb.Digest{Hash: d.Hash, Size: d.SizeBytes, DigestFunction: df}
}

func DigestString(d bfpb.Digest) string {
  // optional prefix
  prefix := ""
  if d.DigestFunction == reapi.DigestFunction_BLAKE3 {
    prefix = "blake3/"
  }
  return fmt.Sprintf("%s%s/%d", prefix, d.Hash, d.Size)
}

func DigestFromBlob(blob []byte, hashFn crypto.Hash) bfpb.Digest {
  h := hashFn.New()
  h.Write(blob)
  arr := h.Sum(nil)

  dfs := map[crypto.Hash]reapi.DigestFunction_Value{
    crypto.MD5: reapi.DigestFunction_MD5,
    crypto.SHA1: reapi.DigestFunction_SHA1,
    crypto.SHA256: reapi.DigestFunction_SHA256,
    crypto.SHA384: reapi.DigestFunction_SHA384,
    crypto.SHA512: reapi.DigestFunction_SHA512,
    // crypto.BLAKE3: reapi.DigestFunction_BLAKE3,
  }

  df := dfs[hashFn]
  return bfpb.Digest{Hash: hex.EncodeToString(arr[:]), Size: int64(len(blob)), DigestFunction: df}
}

func DigestFromMessage(msg proto.Message, hashFn crypto.Hash) (bfpb.Digest, error) {
  blob, err := proto.Marshal(msg)
  if err != nil {
    Empty := DigestFromBlob([]byte{}, hashFn)
    return Empty, err
  }
  return DigestFromBlob(blob, hashFn), nil
}
