package fpl

type Player struct {
	ID            int    `json:"id"`
	FirstName     string `json:"first_name"`
	SecondName    string `json:"second_name"`
	TeamID        int    `json:"team"`
	TotalPoints   int    `json:"total_points"`
	NowCost       int    `json:"now_cost"`
	Form          string `json:"form"`
	SelectedByPct string `json:"selected_by_percent"`
	ElementType   int    `json:"element_type"` // 1=GK, 2=DEF, 3=MID, 4=FWD
}

func (p Player) PositionLabel() string {
	switch p.ElementType {
	case 1:
		return "GK"
	case 2:
		return "DEF"
	case 3:
		return "MID"
	case 4:
		return "FWD"
	default:
		return "UNK"
	}
}

type Team struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
}

type Event struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	DeadlineTime string `json:"deadline_time"`
	IsNext       bool   `json:"is_next"`
	IsCurrent    bool   `json:"is_current"`
	IsFinished   bool   `json:"finished"`
}

type Bootstrap struct {
	Elements []Player `json:"elements"`
	Teams    []Team   `json:"teams"`
	Events   []Event  `json:"events"`
}