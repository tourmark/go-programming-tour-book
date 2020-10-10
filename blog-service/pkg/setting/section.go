package setting


type DatabaseSettingS struct {
	DBType   string
	Username string
	Pwd      string
	Host     string
	Database   string
	Charset  string
}


func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}