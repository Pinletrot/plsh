package plsh

type MinHashValue struct {
	N      uint32
	values []uint32
	params hashParams
}

func NewMinHashValue(n uint32, params hashParams) *MinHashValue {
	msv := &MinHashValue{}
	msv.N = n
	msv.values = make([]uint32, msv.N)
	msv.params = params

	var i uint32
	for i = 0; i < msv.N; i++ {
		msv.values[i] = (uint32)(1<<31) - 1
	}

	return msv
}

func (msv *MinHashValue) Update(x string) {
	var i, tmp uint32
	md5value := MD5HashString(x)
	for i = 0; i < msv.N; i++ {
		tmp = msv.params[i].hash(md5value)
		if tmp < msv.values[i] {
			msv.values[i] = tmp
		}
	}
}

func (msv *MinHashValue) GetValues() []uint32 {
	return msv.values
}
