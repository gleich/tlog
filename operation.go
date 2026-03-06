package oplog

type Operation string

func (o Operation) Inherit(opt string) Operation {
	return Operation(string(o) + "," + opt)
}
