package databases

import (
	"strconv"
	"sync"

	"com.github/mrthoabby/m-authentication/globalConfig"
	"com.github/mrthoabby/m-authentication/types"
	"com.github/mrthoabby/m-authentication/util"
)

type DatabasePooler struct {
	connectionFactory       iDatabaseConnectionFactory
	databasePool            chan IDatabaseConnectionRepository
	mutex                   sync.Mutex
	currentConnetionsOpened int
}

func NewDatabasePooler(connectionFactory iDatabaseConnectionFactory, tableMapper *types.TableMapper) *DatabasePooler {
	pool := make(chan IDatabaseConnectionRepository, globalConfig.POOL_SIZE_DATABASE_CONNECTION)
	succesConnections := 0
	for index := 0; index < globalConfig.POOL_SIZE_DATABASE_CONNECTION; index++ {
		connection, errorCreatingConnection := connectionFactory.CreateConnection(tableMapper)
		if errorCreatingConnection != nil {
			util.LoggerHandler().Error("Error creating database connection number: "+strconv.Itoa(index+1), "error", errorCreatingConnection.Error())
			continue
		}
		succesConnections++
		pool <- connection
	}

	return &DatabasePooler{
		currentConnetionsOpened: succesConnections,
		connectionFactory:       connectionFactory,
		databasePool:            pool,
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

func (dp *DatabasePooler) GetConnectionsOpened() int {
	return dp.currentConnetionsOpened
}
