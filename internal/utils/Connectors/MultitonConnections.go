package DbConnections

type MultiConnection struct {
	connections map[string]interface{}
}

func (data *MultiConnection) getPgConnect() {
	if data.connections["postgres"] != nil {
		return data.connections["postgres"].connector.getDbConnection()
	} else {
		connection, err := newPgConnection()
		if err != nil {
			fmt.Fatalf("Fatal error %s", err.Error())
		}
		data.connections["postgres"] = connection
	}

}
