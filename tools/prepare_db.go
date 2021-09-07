package main

import "kasznar.hu/hello/app/repository"

func main() {
	repository.Connect()
	repository.MigrateOrm()

	/*DB().Exec(
		`DROP TABLE
				users,
				user_types,
				calendar_days,
				calendar_day_types,
				workdays,
				workday_types,
				teams`,
	)*/

	//db.Exec("DROP DATABASE IF EXISTS GO_DB")
	//db.Exec("CREATE DATABASE GO_DB")

	InsertMockData()
}
