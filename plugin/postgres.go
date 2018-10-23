package driver

import (
	"fmt"
)

func PGSource(cc *ConnectionConfig) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?%s",
		firstString(cc.Username, "root"),
		firstString(cc.Password, ""),
		firstString(cc.Host, "localhost"),
		firstInt(cc.Port, 5432),
		firstString(cc.Database, ""),
		firstString(cc.Params, "sslmode=disable"))
}
