
# sql.Openは一度だけ
勝手にコネクションプールを確保してくれるので、初回(main関数など)で一回だけ初期化すれば良い。デフォルトでスレッドセーフ。

# xxxContextメソッドを使う
xxxメソッドは後方互換性のために残っているだけ。Contextで正しくキャンセル処理ができるようにする。

# transactionのロールバックはdefer文で呼ぶ
- defer文は関数スコープ終了時に確実に呼ばれる。
- transctionがcommitされた場合や、contextによりキャンセルされた場合はロールバック処理が呼ばれても実行されることはない。


# GoのRDBMS系のOSS
- github.com/jmoiron/sqlx
- https://entgo.io
- https://gorm.io
- https://sqlc.dev
- https://github.com/volatiletech/sqlboiler