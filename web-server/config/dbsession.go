package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

type ConfDbSession struct {
	Host    string `toml:"host"`    // ホスト名
	Port    int    `toml:"port"`    // ポート番号
	DbName  string `toml:"db-name"` // 接続先DB名
	Charset string `toml:"charset"` // 文字コード
	User    string `toml:"user"`    // 接続ユーザ名
	Pass    string // 接続パスワード
}

type ConfigSession struct {
	ConfDbSession `toml:"database1"`
}

// URL設定の構造体
func ReadConfDbSession() (*ConfigSession, error) {

	// 設定ファイル名
	confpath := "config/db.toml"

	// 構造体を初期化
	conf := new(ConfigSession)

	// 読み込んだjson文字列をデコードし構造体にセット
	_, err := toml.DecodeFile(confpath, &conf)
	if err != nil {
		return conf, err
	}
	conf.ConfDbSession.Pass = os.Getenv("PASSWORD1")
	return conf, nil
}
