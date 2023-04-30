# calconv-fun

時間割toカレンダーなgolangツール 未来大version.

## Usage

学校のある日のリストのcsvファイル (1行目は日付・2行目は曜日) を用意します。

これを -h オプションで表示されるヘルプを参考にしながらコマンドに与えてください

その後、Student (教務システム) に表示される時間割を標準入力に与えます。

その後、変換処理が行われます。

出力されたcsvをGoogle Calender等にインポートするといい感じになります。

※ミスがあったときに消すのが大変なので、新たなカレンダーを作ってインポートすることをおすすめします。

## 実行

go言語を使って実行できます。
必要な物はgo言語本体のみです。

``` bash
go run ./cmd/calfonv
```
