package factory

import "gitlab.com/chicho69-cesar/factory/connection"

func Factory(t int) connection.DBConnection {
	switch t {
		case 1:
			return &connection.MySQL {}
		case 2:
			return &connection.PostgreSQL {}
		default:
			return nil
	}
}
