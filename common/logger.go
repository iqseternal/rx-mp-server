package common

type LoggerStruct struct {
}

var (
	Logger = LoggerStruct{}
)

func log(args ...string) {

}

func (*LoggerStruct) Info(args ...string) {
	log(args...)
}

func Error(args ...string) {
	log(args...)
}
