package conf

func GetEnv(key string, defaultValue ...string) string {
	if err := v.BindEnv(key); err != nil {
		return ""
	}

	return v.GetString(key)
}

type EnvVar string

func (e EnvVar) String() string {
	return string(e)
}

func (e EnvVar) Value(opts ...string) string {
	if len(opts) > 0 {
		for _, o := range opts {
			if o != "" {
				return o
			}
		}
	}

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
	PG_DSN EnvVar = "pg_dsn"
)
