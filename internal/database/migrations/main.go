package migrations

func Run() {
	createTables()
	createAdminIfNotExist()
	agregarBarrios()
	agregarGruposVacunas()
	agregarVacunas()
	prefijoDefault()
}
