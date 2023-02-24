package main

import (
    "flag"
    "fmt"
)

func main() {
    // コマンドライン引数をパースする
    str := flag.String("string", "", "description")
    flag.Parse()
    fmt.Println(*str)

    //command := exec.Command("git", "add", "-A")
    //command.Stdout = os.Stdout
    //err := command.Run()
    //if err != nil {
    //    log.Fatal("git add の実行に失敗しました")
    //}
    //
    //command = exec.Command("git", "commit", "-m", "Update")
    //command.Stdout = os.Stdout
    //err = command.Run()
    //if err != nil {
    //    log.Fatal("git commit の実行に失敗しました")
    //}
    //
    //command = exec.Command("git", "push", "origin", "main")
    //command.Stdout = os.Stdout
    //err = command.Run()
    //if err != nil {
    //    log.Fatal("git push の実行に失敗しました")
    //}
}
