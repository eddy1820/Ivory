package setting

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Setting struct {
	vp       *viper.Viper
	sections map[string]interface{}
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("internal/infrastructure/configs/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	s := &Setting{vp, make(map[string]interface{})}
	return s, nil
}

func (s *Setting) WatchChange(onChange func()) {
	s.vp.WatchConfig()
	s.vp.OnConfigChange(func(in fsnotify.Event) {
		onChange()
		_ = s.ReloadAllSection()
	})
}

func (s *Setting) WatchSettingChange() {
	go func() {
		s.vp.WatchConfig()
		// 動態載入設定
		s.vp.OnConfigChange(func(in fsnotify.Event) {
			// log.Println("filePath: " + in.VersionName)
			// log.Println("op: " + in.Op.String())

			_ = s.ReloadAllSection()
		})
	}()
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	if _, ok := s.sections[k]; ok {
		s.sections[k] = v
	}

	return nil
}

func (s *Setting) ReloadAllSection() error {
	for k, v := range s.sections {
		err := s.ReadSection(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
