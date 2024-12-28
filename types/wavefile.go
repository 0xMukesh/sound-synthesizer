package types

type Sample float64

type WaveHeader struct {
	ChunkId   []byte
	ChunkSize int
}

type WaveFmt struct {
	SubChunk1Id   []byte // must be equal to "FMT"
	SubChunk1Size int    // 16 for PCM
	AudioFormat   int    // 1 for PCM (pulse code modulation)
	NumOfChannels int    // mono = 1, stereo = 2
	SampleRate    int    // generally, it is 44100
	ByteRate      int    // byte rate = (sample rate) * (num of channels) * (bits per sample)/8
	BlockAlign    int    // block align = (num of channels) * (bits per sample)/8
	BitsPerSample int
}
