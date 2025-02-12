package client

import (
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

func HasherFromDigestFunction(df reapi.DigestFunction_Value) Hasher {
  hashers := map[reapi.DigestFunction_Value]Hasher{
    reapi.DigestFunction_MD5: MD5,
    reapi.DigestFunction_SHA1: SHA1,
    reapi.DigestFunction_SHA256: SHA256,
    reapi.DigestFunction_SHA384: SHA384,
    reapi.DigestFunction_SHA512: SHA512,
    reapi.DigestFunction_BLAKE3: BLAKE3,
  }

  return hashers[df]
}

func DigestFromBlob(blob []byte, hasher Hasher) bfpb.Digest {
  h := hasher.New()
  h.Write(blob)
  arr := h.Sum(nil)

  dfs := map[Hasher]reapi.DigestFunction_Value{
    MD5: reapi.DigestFunction_MD5,
    SHA1: reapi.DigestFunction_SHA1,
    SHA256: reapi.DigestFunction_SHA256,
    SHA384: reapi.DigestFunction_SHA384,
    SHA512: reapi.DigestFunction_SHA512,
    BLAKE3: reapi.DigestFunction_BLAKE3,
  }

  df := dfs[hasher]
  return bfpb.Digest{Hash: hex.EncodeToString(arr[:]), Size: int64(len(blob)), DigestFunction: df}
}

func DigestFromMessage(msg proto.Message, hasher Hasher) (bfpb.Digest, error) {
  blob, err := proto.Marshal(msg)
  if err != nil {
    Empty := DigestFromBlob([]byte{}, hasher)
    return Empty, err
  }
  return DigestFromBlob(blob, hasher), nil
}
