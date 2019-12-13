# ORM Benchmark

A benchmark to compare the performance of golang orm package.

## Results (2019-12-14)

see [here](https://kokizzu.blogspot.com/2019/12/go-orm-benchmark-on-memsql.html)

### Environment

* 32 GB RAM
* go version go1.13.4 linux/amd64
* [Go-MySQL-Driver Latest](https://github.com/go-sql-driver/mysql)
* MemSQL 6.7.16

### ORMs

All package run in no-cache mode.

* [Beego ORM](http://beego.me/docs/mvc/model/overview.md) latest in branch [develop](https://github.com/astaxie/beego/tree/develop)
* [xorm](https://github.com/lunny/xorm) latest
* [gorm](https://github.com/jinzhu/gorm) latest, [my fork](https://github.com/kokizzu/gorm) for memsql
* [gorp](https://github.com/coopernurse/gorp) latest
* [modl](https://github.com/jmoiron/modl) latest
* [Hood](https://github.com/eaigner/hood) latest, [my fork](https://github.com/kokizzu/hood) for memsql
* [Qbs](https://github.com/coocood/qbs) latest (Disabled stmt cache / [patch](https://gist.github.com/slene/8297019) / [full](https://gist.github.com/slene/8297565)), [my fork](https://github.com/kokizzu/qbs) for memsql
* [upper.io](https://upper.io/db) latest (updated to v3, can't fix the inefficient `UPDATE SET id = ? WHERE id = ?` that will always fail in MemSQL)

### Run

```go
go get github.com/kokizzu/orm-benchmark
orm-benchmark -multi=20 -orm=all
```

### Contact

Maintain by [slene](https://github.com/slene)
Updated by [kimaho](https://github.com/kihamo)
Updated by [kokizzu](https://github.com/kokizzu)
