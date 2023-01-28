# 依存関係逆転の法則　DIPとは

Dependency inversion prinsipleの略

上位のモジュールは下位のモジュールに依存してはならない。
例) リポジトリレイヤーはUIレイヤーに依存してはならない。

実装の詳細が抽象に依存するべきである。
例) RepsitoryImplがRepositoryに依存する。

Goではデータベースを操作するとき、`databse/sql/driver`を利用する。
このライブラリは `MySQL`や`Postgresql`に関しての実装の詳細は含んでいない。
ただしそれぞれのDBドライバは`database/sql/driver`のインターフェースを実装しているので、
利用者は`database/sql/driver`のみを知っていれば、操作が可能である。

つまりもともと `database/sql/driver` → 各種RDBMSのドライバ という依存関係だったのに
インターフェースを解することで
`database/sql/driver` ← 各種RDBMSのドライバ という依存関係になったのである。
これを依存性の逆転と呼ぶ。

