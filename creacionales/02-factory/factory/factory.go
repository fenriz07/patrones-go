package factory

import "github.com/fenriz07/patrones-go/creacionales/02-factory/connection"

/*Factory devuelve un objeto connectio o postgresql segun sea el caso*/
func Factory(t int) connection.DBConnection {

	switch t {
	case 1:
		return &connection.MySQL{}
	case 2:
		return &connection.Postgres{}
	default:
		return nil
	}

}
