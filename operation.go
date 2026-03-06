package tlog

type Op string

func (o Op) Extend(opt string) Op {
	return Op(string(o) + "-" + opt)
}
