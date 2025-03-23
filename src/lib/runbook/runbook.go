package runbook

type DB interface {
	GetRunBooks() string
}

type DBClient struct{}

type RunBook struct {
	Id string `json:"id"`
}

func (dbClient *DBClient) GetRunBooks() string {
	return ""
}
