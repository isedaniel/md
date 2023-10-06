package md_test

import (
    "testing"
    "github.com/approvals/go-approval-tests"
    "github.com/isedaniel/md"
)

func TestMdToHtml(t *testing.T) {
    const m = `# This is a test

This is the body of the test.

Here is a [link](github.com/isedaniel)`

    html := md.MdToHtml([]byte(m))
    approvals.VerifyString(t, string(html))
}
