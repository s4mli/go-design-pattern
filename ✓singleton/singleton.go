package singleton

type _singleton struct{} // package name vs struct name

var instance *_singleton

func GetInstance() *_singleton {
	if instance == nil {
		instance = &_singleton{}
	}
	return instance
}
