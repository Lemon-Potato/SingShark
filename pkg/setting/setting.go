package setting

import "github.com/spf13/viper"

type Setting struct {
	// 设计为包内处理，对外输出参数对象
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	v := viper.New()
	v.AddConfigPath("configs/")
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Setting{v}, nil
}
