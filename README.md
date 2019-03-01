# lean-search

Search LeanCloud documentation from the command line.

## Usage

```sh
; lean-search lastMessageAt
 即时通讯开发指南 · 进阶功能 > 对话查询 > 性能优化建议 > 对话的查询和存储 >> https://leancloud.cn/docs/realtime-guide-intermediate.html#hash-135440353
 即时通讯开发指南 · JavaScript > 对话 > 对话的查询 > 查询结果选项 > 对话的最后一条消息 >> https://leancloud.cn/docs/realtime_guide-js.html#hash1769400086
```

## Install

Download the binary file at GitHub's releases page according to your os,
rename it to `lean-search`,
grant it executable permission,
and put it under your `$PATH`.

Binary releases are available for the following operation systems on amd64:

1. darwin
2. freebsd
3. linux
4. windows

You can install it from source under other operation systems:

```sh
go get -u github.com/weakish/lean-search
```

## Limit

It will return **1000** results at most.

## Legal

**License** 0BSD

**Disclaimer** lean-search is not endorsed by LeanCloud in any way.



