package setting

type ServerSettingS struct {
	HttpPort string
}

type LoggerSettingS struct {
	FileName  string
	MaxSize   int
	MaxBackup int
	MaxAge    int
	Compress  bool
	Type      string
	Level     string
}

type DatabaseSettingS struct {
	DbConnection       string
	DbHost             string
	DbPort             int
	DbDatabase         string
	Username           string
	Password           string
	Charset            string
	MaxIdleConnections int
	MaxOpenConnections int
	MaxLifeSeconds     int
}

/**
 *  viper 中的 UnmarshalKey 方法可分解配置到 struct 或者 map 等结构中
 */
func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
