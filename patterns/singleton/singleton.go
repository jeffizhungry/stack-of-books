package singleton

var m_instance *SingleObject

type SingleObject struct{}

func Instance() *SingleObject {
	if m_instance == nil {
		m_instance = &SingleObject{}
	}
	return m_instance
}
