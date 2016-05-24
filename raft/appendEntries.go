package raft

type AppendEntries string

func (e *AppendEntries) HeartBeat(s string, r *string) error {
	*r = s + "--" + s
	return nil
}
