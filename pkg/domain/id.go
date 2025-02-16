package domain

import (
	"time"

	"github.com/opoccomaxao/go-snowflake"
)

//nolint:wrapcheck,gosec,mnd
func NewGenerator() (*snowflake.Generator, error) {
	epoch := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()

	return snowflake.New(snowflake.Config{
		MachineID:             0,
		EpochStartUnixSeconds: uint64(epoch),
		MachineBits:           1,
		SequenceBits:          10,
	})
}
