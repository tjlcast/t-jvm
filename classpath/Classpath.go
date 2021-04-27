package classpath

type Classpath struct {
	bootClasspath Entry
	extClasspath Entry
	userClasspath Entry
}

func Parse(jreOption, cpOption string) *Classpath {
	// todo
	return nil
}

func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	// todo
	return nil, nil, nil
}

func (self *Classpath) String() string {
	// todo
	return ""
}
