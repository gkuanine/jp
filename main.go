package main

import (
	. "jp/irouter"
	"jp/jp"
)

func main() {
	jp.Ignite().Mount(NewUserRouter(), NewIndexRouter()).Launch()
}
