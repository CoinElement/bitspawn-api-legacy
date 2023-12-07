package enum

type FeeType string

const (
	Credit FeeType = "CREDIT"
	Spwn   FeeType = "SPWN"
	Usdc   FeeType = "USDC"
)

func (ft FeeType) IsValid() bool {
	if ft != Credit && ft != Spwn && ft != Usdc {
		return false
	}
	return true
}

func (ft FeeType) ToString() string {
	return string(ft)
}
