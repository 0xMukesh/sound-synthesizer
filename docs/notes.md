# notes

## wave file structure

### header

1. 0 - 4 - chunk id [must be equal to `RIFF`]
2. 4 - 8 - chunk size
3. 8 - 12 - format [must be equal to `WAVE`]

### fmt

1. 12 - 16 - sub chunk 1 id [must be equal to `FMT`]
2. 16 - 20 - sub chunk 1 size
3. 20 - 22 - audio format
4. 22 - 24 - num of channels
5. 24 - 28 - sample rate
6. 28 - 32 - byte rate
7. 32 - 34 - block align
8. 34 - 36 - bits per sample

### data

1. 36 - 40 - sub chunk 2 id
2. 40 - 44 - sub chunk 2 size
3. 44 - ... - raw data
