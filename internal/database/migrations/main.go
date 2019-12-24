package migrations

func Run() {
	createTables()
	createAdminIfNotExist()
	agregarBarrios()
}
