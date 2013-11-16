package boggle

// Specifies functions pertaining to a set of dice.
type Dice interface {
	Shake() []byte
	Count() int
}

// An individual dice where the string is the letter on each face.
type Die []string

// A set of dice for a single game.
type DieSet []Die

// Returns a random letter from each die.
func (d DieSet) Shake() []byte {
	return []byte{}
}

// Returns the number of dice in the set.
func (d DieSet) Count() int {
	return len(d)
}

var (
	// 2008 Boggle dice. Slight changes when compared to the 1976 version.
	NewBoggleDice = DieSet{
		{"AAEEGN"},
		{"ELRTTY"},
		{"AOOTTW"},
		{"ABBJOO"},
		{"EHRTVW"},
		{"CIMOTU"},
		{"DISTTY"},
		{"EIOSST"},
		{"DELRVY"},
		{"ACHOPS"},
		{"HIMNQU"},
		{"EEINSU"},
		{"EEGHNW"},
		{"AFFKPS"},
		{"HLNNRZ"},
		{"DEILRX"},
	}

	// 1976
	OldBoggleDice = DieSet{
		{"AACIOT"},
		{"AHMORS"},
		{"EGKLUY"},
		{"ABILTY"},
		{"ACDEMP"},
		{"EGINTV"},
		{"GILRUW"},
		{"ELPSTU"},
		{"DENOSW"},
		{"ACELRS"},
		{"ABJMOQ"},
		{"EEFHIY"},
		{"EHINPS"},
		{"DKNOTU"},
		{"ADENVZ"},
		{"BIFORX"},
	}
)

/*
Boggle

1976          2008
New Version | Old Version
AAEEGN | AACIOT
ELRTTY | AHMORS
AOOTTW | EGKLUY
ABBJOO | ABILTY
EHRTVW | ACDEMP
CIMOTU | EGINTV
DISTTY | GILRUW
EIOSST | ELPSTU
DELRVY | DENOSW
ACHOPS | ACELRS
HIMNQU | ABJMOQ
EEINSU | EEFHIY
EEGHNW | EHINPS
AFFKPS | DKNOTU
HLNNRZ | ADENVZ
DEILRX | BIFORX


Big Boggle
aaafrs
aaeeee
aafirs
adennn
aeeeem

aeegmu
aegmnn
afirsy
bjkqxz
ccenst

ceiilt
ceilpt
ceipst
ddhnot
dhhlor

dhlnor
dhlnor (duplicate)
eiiitt
emottt
ensssu

fiprsy
gorrvw
iprrry
nootuw
ooottu

Challenge Cube (1983 edition, perhaps others)
iklmqu

Boggle Deluxe
AAAFRS
AAEEEE
AAFIRS
ADENNN
AEEEEM
AEEGMU
AEGMNN
AFIRSY
BJKQXZ
CCNSTW *
CEIILT
CEILPT
CEIPST
DHHNOT *
DHHLOR
DHLNOR
DDLNOR *
EIIITT
EMOTTT
ENSSSU
FIPRSY
GORRVW
HIPRRY *
NOOTUW
OOOTTU

* indicates a different die from Big Boggle
*/
