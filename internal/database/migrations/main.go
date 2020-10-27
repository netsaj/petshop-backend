package migrations

func Run() {

	// crear toda la estructura de tablas
	createTables()

	// creamos usuario admin si no existe
	createAdminIfNotExist()

	// agregamos datos por default
	agregarBarrios()
	agregarGruposVacunas()
	agregarVacunas()
	prefijoDefault()
	agregarGruposDesparasitantes()
	agregarDesparasitantes()
	agregarExamenes()
}
