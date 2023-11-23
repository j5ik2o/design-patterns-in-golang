package singleton

type Singleton struct {
}

var instance *Singleton

func GetSingletonInstance() *Singleton {
	if instance == nil {
		instance = &Singleton{}
	}
	return instance
}
