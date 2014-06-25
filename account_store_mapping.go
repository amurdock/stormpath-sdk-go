package stormpath

type AccountStoreMapping struct {
	Href                  *string `json:"href,omitempty"`
	ListIndex             *int    `json:"listIndex,omitempty"`
	IsDefaultAccountStore *bool   `json:"isDefaultAccountStore,omitempty"`
	IsDefaultGroupStore   *bool   `json:"isDefaultGroupStore,omitempty"`
	Application           Link    `json:"application"`
	AccountStore          Link    `json:"accountStore"`
}

func NewAccountStoreMapping(applicationHref string, accountStoreHref string) *AccountStoreMapping {
	return &AccountStoreMapping{Application: Link{Href: applicationHref}, AccountStore: Link{Href: accountStoreHref}}
}