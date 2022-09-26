package chapter20

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDoubles(t *testing.T) {
	result := FindMultiSportsPlayers(basketballPlayers, footballPlayers)
	sort.Strings(result)
	fmt.Println(result)
	assert.Equal(t, []string{"Jill Huang", "Wanda Vakulsak"}, result)
}
