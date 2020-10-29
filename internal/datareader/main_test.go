package datareader

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToDomainPlayer(t *testing.T) {
	player := testFactory.NewTestPlayer()
	t.Parallel()
	t.Run("Valid conversion", func(t *testing.T) {
		t.Parallel()
		domainPlayer, err := toDomainPlayer(player)
		assert.NoError(t, err)
		assert.Equal(t, 1034, domainPlayer.TotalRushingYards())
	})
	t.Run("Invalid conversion", func(t *testing.T) {
		t.Parallel()
		player.TotalRushingYards = "Fff"
		_, err := toDomainPlayer(player)
		assert.Error(t, err)
		assert.Equal(t, "strconv.ParseInt: parsing \"Fff\": invalid syntax",err.Error())
	})
}
