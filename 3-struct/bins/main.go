package bins

import "time"

type Bin struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

type BinList struct {
	Bins []Bin
}

func CreateBin(id string, private bool, createdAt time.Time, name string) (*Bin, string) {
	if id == "" {
		return nil, "Id is required"
	}

	if name == "" {
		return nil, "Name is required"
	}

	var bin = Bin{
		Id:        id,
		Private:   private,
		CreatedAt: createdAt,
		Name:      name,
	}

	return &bin, ""
}

func CreateBinList(bins []Bin) *BinList {
	var binList = BinList{
		Bins: bins,
	}

	return &binList
}
