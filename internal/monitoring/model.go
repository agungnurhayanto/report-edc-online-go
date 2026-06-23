package monitoring

import "time"

type Monitoring struct {
	ID         int
	Tgl        time.Time
	Kdcab      string
	Cabang     string
	Kdtk       string
	Nama       string
	Station    string
	Cek        string
	IP         string
	EDCBCA     string
	EDCMandiri string
	EDCMTI     string
	EDCMDRMTI  string
}
