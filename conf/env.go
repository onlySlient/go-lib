package conf

func GetEnv(key string) string {
	if err := v.BindEnv(key); err != nil {
		return ""
	}

	return v.GetString(key)
}

type EnvVar string

func (e EnvVar) String() string {
	return string(e)
}

func (e EnvVar) Value() string {
	err := v.BindEnv(string(e))
	if err != nil {
		return ""
	}

	return v.GetString(string(e))
}

const (
	// pg
	PG_HOST     EnvVar = "PG_HOST"
	PG_PORT     EnvVar = "PG_PORT"
	PG_USER     EnvVar = "PG_USER"
	PG_PASSWORD EnvVar = "PG_PASSWORD"
	PG_DBNAME   EnvVar = "PG_DBNAME"

	// redis
	REDIS_HOST     EnvVar = "REDIS_HOST"
	REDIS_PORT     EnvVar = "REDIS_PORT"
	REDIS_PASSWORD EnvVar = "REDIS_PASSWORD"
)
