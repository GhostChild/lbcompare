## Littlebakas's image compare

## http://littlebakas.moe/


#cross compile
Powershell:
```
$env:GOOS="linux"
$env:GOARCH="amd64"
go build lbcompare.go
```

PS:Linux 5.0以上内核无法运行 请使用4.x内核
可能是linux内核问题 也可能是go不兼容新的linux内核
有待验证