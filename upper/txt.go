package upper

import (
	"github.com/sirupsen/logrus"
)

type txtservice struct {
}

func (s *txtservice) Handle(srcport uint16, destport uint16, data *[]byte) {
	logrus.Infof("[upper.txt] recv %s.", *data)
}

func init() {
	ts := Service(new(txtservice))
	Register(1, &ts)
	logrus.Debugln("[upper.txt] service registered.")
}
