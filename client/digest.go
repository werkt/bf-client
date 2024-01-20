package client

import (
  "fmt"
  reapi "github.com/bazelbuild/remote-apis/build/bazel/remote/execution/v2"
)

func DigestString(d *reapi.Digest) string {
  if d == nil {
    return "nil digest"
  }
  return fmt.Sprintf("%s/%d", d.Hash, d.SizeBytes)
}
