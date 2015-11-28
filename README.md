govend [![GoDoc](http://godoc.org/github.com/gophersaurus/govend?status.png)](http://godoc.org/github.com/gophersaurus/govend) [![Build Status](https://travis-ci.org/gophersaurus/govend.svg?branch=master)](https://travis-ci.org/gophersaurus/govend) [![Go Report Card](http://goreportcard.com/badge/gophersaurus/govend?)](http://goreportcard.com/report/gophersaurus/govend)
============================================================================================================================

`govend` leverages the `GO15VENDOREXPERIMENT` to vendor external package dependencies.

**it does:**
* try to be compatible with any project
* take a both note and code from `go get`
* vendor the nested dependency tree to the `nth` degree
* utilize the `GO15VENDOREXPERIMENT` and `vendor` directory as specified in golang version 1.5

**it does not:**
* wrap the `go` command
* try to create a new project
* generate temporary directories or files
* alter any go environment variables, including `$GOPATH`

Install
=======

```bash
$ go get -u github.com/gophersaurus/govend
```

Vendor Dependencies
===================

```bash
$ cd project/root

$ govend -v
```

To vendor external package dependencies run `govend -v` while inside the project root directory.

```bash
→ govend -v
github.com/kr/fs
github.com/spf13/cobra
github.com/spf13/pflag
github.com/inconshreveable/mousetrap
github.com/cpuguy83/go-md2man
github.com/russross/blackfriday
github.com/shurcooL/sanitized_anchor_name
gopkg.in/yaml.v2
gopkg.in/check.v1
github.com/BurntSushi/toml
```

### `vendors.yml`

The `vendors.yml` file contains an array of import paths and commit revisions.

```yaml
vendors:
- path: github.com/BurntSushi/toml
  rev: 056c9bc7be7190eaa7715723883caffa5f8fa3e4
- path: github.com/cpuguy83/go-md2man
  rev: 71acacd42f85e5e82f70a55327789582a5200a90
- path: github.com/inconshreveable/mousetrap
  rev: 76626ae9c91c4f2a10f34cad8ce83ea42c93bb75
- path: github.com/kr/fs
  rev: 2788f0dbd16903de03cb8186e5c7d97b69ad387b
- path: github.com/russross/blackfriday
  rev: 300106c228d52c8941d4b3de6054a6062a86dda3
- path: github.com/shurcooL/sanitized_anchor_name
  rev: 10ef21a441db47d8b13ebcc5fd2310f636973c77
- path: github.com/spf13/cobra
  rev: 1c44ec8d3f1552cac48999f9306da23c4d8a288b
- path: github.com/spf13/pflag
  rev: 08b1a584251b5b62f458943640fc8ebd4d50aaa5
- path: gopkg.in/check.v1
  rev: 11d3bc7aa68e238947792f30573146a3231fc0f1
- path: gopkg.in/yaml.v2
  rev: 53feefa2559fb8dfa8d81baad31be332c97d6c77
```

### List Dependencies
If you want to scan your code to find out how many third party dependencies are
in your project run `govend list`. You can specify a path and output formats.

Here is an example of a path: `govend list some/project/path`
```bash
github.com/spf13/cobra
github.com/kr/fs
gopkg.in/yaml.v2
github.com/jackspirou/importsplus
golang.org/x/tools/go/vcs
```

**JSON**
`govend list -f json`
```bash
[
  "github.com/spf13/cobra",
  "github.com/kr/fs",
  "gopkg.in/yaml.v2",
  "github.com/jackspirou/importsplus",
  "golang.org/x/tools/go/vcs"
]%  
```

**YAML**
`govend list -f yml`
```bash
- github.com/spf13/cobra
- github.com/kr/fs
- gopkg.in/yaml.v2
- github.com/jackspirou/importsplus
- golang.org/x/tools/go/vcs
```
**XML**
`govend list -f xml`
```bash
<string>github.com/spf13/cobra</string>
<string>github.com/kr/fs</string>
<string>gopkg.in/yaml.v2</string>
<string>github.com/jackspirou/importsplus</string>
<string>golang.org/x/tools/go/vcs</string>%
```

Windows Support
===============

Does `govend` work on Windows platforms?

> It does, we have tested it, but some bugs may exist.

Contributing
============

### Can I Contribute?

> Please do! Simply fork the code and send a pull request.
