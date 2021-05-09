package main

/**

Hard:自分のマシンのIPアドレスを取得し画面表示
自分のマシンのIPアドレス”のみ”を表示する
ただし、ループバックインタフェースのIPアドレス(127.0.0.1)は除外して表示する
※もし余裕があれば、WindowsとMacOSXとLinuxに対応したソースコードを書いてみてください。余裕があれば。

**/

// MacOSX, Linuxで動作するソースコード

import (
  "net"
  "os"
  "fmt"
)
 
func main() {
  addrs, err := net.InterfaceAddrs()
//  fmt.Println(addrs)

//  fmt.Println("-------------------------------------------------")

//  for _, addrs := range addrs {
//    fmt.Println(addrs.String())
//  }  

//  fmt.Println("-------------------------------------------------")
  if err != nil {
    os.Stderr.WriteString("Oops: " + err.Error() + "\n")
  os.Exit(1)
  }
 
  for _, a := range addrs {
    if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
      if ipnet.IP.To4() != nil {
        fmt.Println(ipnet.IP.String())
      }
    }
  }

}
