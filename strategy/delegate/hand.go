package delegate

import "fmt"

type Hand int

const (
	GUU Hand = iota
	CHO
	PAA
)

func (h *Hand) Name() string {
	var result string
	switch *h {
	case GUU:
		result = "グー"
	case CHO:
		result = "チョキ"
	case PAA:
		result = "パー"
	}
	return result
}

func GetHand(value int) Hand {
	var result Hand
	switch value {
	case 0:
		result = GUU
	case 1:
		result = CHO
	case 2:
		result = PAA
	default:
		panic(fmt.Sprintf("not found: value = %d", value))
	}
	return result
}

func (h *Hand) HandValue() int {
	return (int)(*h)
}

func (h *Hand) Fight(other Hand) int {
	if h.HandValue() == other.HandValue() {
		return 0
	} else if ((h.HandValue() + 1) % 3) == other.HandValue() {
		return 1
	} else {
		return -1
	}
}

func (h *Hand) IsStrongerThan(other Hand) bool {
	return h.Fight(other) == 1
}

func (h *Hand) IsWeakerThan(other Hand) bool {
	return h.Fight(other) == -1
}

func (h *Hand) ToString() string {
	return h.Name()
}
