package db

import "github.com/gocql/gocql"

func NewCassandraSession(host string) (*gocql.Session, error) {
	cluster := gocql.NewCluster(host)
	cluster.Keyspace = "online_bank"
	cluster.Consistency = gocql.Quorum
	return cluster.CreateSession()
}
