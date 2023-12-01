package main

import (
	"MyGoProject/project01_gin/Datebases"
	"database/sql"
)

func main() {
	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {

		}
	}(Datebases.DB)

}
