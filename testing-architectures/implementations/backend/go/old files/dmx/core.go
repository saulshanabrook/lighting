package core


const (
	maxChannels = 512
	)

type Data [maxChannels]uint8

type Output interface {
	Send(Data) error
}

