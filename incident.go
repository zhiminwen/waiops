package waiops

type Incident struct {
	State              string      `json:"state,omitempty"`
	AlertIDs           []string    `json:"alertIds,omitempty"`
	ContextualAlertIDs []string    `json:"contextualAlertIds,omitempty"`
	Insights           []EvInsight `json:"insights,omitempty"`
	ID                 string      `json:"id,omitempty"`
	CreatedTime        EvTime      `json:"createdTime,omitempty"`
	CreatedBy          string      `json:"createdBy,omitempty"`
	Title              string      `json:"title,omitempty"`
	Description        string      `json:"description,omitempty"`
	LangID             string      `json:"langId,omitempty"`
	Priority           int         `json:"priority,omitempty"`
	ExternalEndPoint   []string    `json:"externalEndPoint,omitempty"`
	LastChangedTime    EvTime      `json:"lastChangedTime,omitempty"`
	Owner              string      `json:"owner,omitempty"`
	Team               string      `json:"team,omitempty"`
}
