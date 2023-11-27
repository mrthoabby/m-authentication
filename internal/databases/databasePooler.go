package databases

import (
	"sync"

	"com.github/mrthoabby/m-authentication/globalConfig"
	"com.github/mrthoabby/m-authentication/util"
)

type DatabasePooler struct {
	connectionFactory iDatabaseConnectionFactory
	databasePool      chan IDatabaseConnectionRepository
	mutex             sync.Mutex
}

func NewDatabasePooler(connectionFactory iDatabaseConnectionFactory) *DatabasePooler {
	pool := make(chan IDatabaseConnectionRepository, globalConfig.POOL_SIZE_DATABASE_CONNECTION)

	for index := 0; index < globalConfig.POOL_SIZE_DATABASE_CONNECTION; index++ {
		connection, errorCreatingConnection := connectionFactory.CreateConnection()
		if errorCreatingConnection != nil {
			util.LoggerHandler().Error("Error creating database connection", "error", errorCreatingConnection.Error())
			return nil
		}
		pool <- connection
	}

	return &DatabasePooler{
		connectionFactory: connectionFactory,
		databasePool:      pool,
	}
}

func (dp *DatabasePooler) GetConnection() IDatabaseConnectionRepository {
	dp.mutex.Lock()
	defer dp.mutex.Unlock()
	return <-dp.databasePool
}

func (dp *DatabasePooler) ReleaseConnection(connection IDatabaseConnectionRepository) {
	dp.mutex.Lock()
	defer dp.mutex.Unlock()
	dp.databasePool <- connection
}
