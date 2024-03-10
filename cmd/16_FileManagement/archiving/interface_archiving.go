package archiving

type IArchive interface {
	CompressFile(fileName string) error
}

var _ IArchive = &Zip{}
