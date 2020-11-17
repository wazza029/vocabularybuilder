package newwords

func WordWriter(s string, e string, database map[string]string) map[string]string {

	//WordWriter registers new words in map, key in spanish, value in english

	database[s] = e
	return database

}
