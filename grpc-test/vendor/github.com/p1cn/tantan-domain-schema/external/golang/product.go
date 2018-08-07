package external

const ProductType = "product"

type Product struct {
	Id         string            `json:"id"`
	Name       string            `json:"name"`
	Privileges ProductPrivileges `json:"privileges"`
	Type       string            `json:"type"`
}

type ProductPrivileges struct {
	Membership *PrivilegeMembership `json:"membership"`
	Like       *PrivilegeLike       `json:"like"`
	Undo       *PrivilegeUndo       `json:"undo"`
	SuperLike  *PrivilegeSuperLike  `json:"superLike"`
	Boost      *PrivilegeBoost      `json:"boost"`
	Roaming    *PrivilegeRoaming    `json:"roaming"`
}

type PrivilegeMembership struct {
	Name  string `json:"name"`
	Level int    `json:"level"`
}

type PrivilegeLike struct {
	Total int `json:"total"`
	Reset int `json:"reset"`
}

type PrivilegeUndo struct {
	Total int `json:"total"`
	Reset int `json:"reset"`
}

type PrivilegeSuperLike struct {
	Quota int `json:"quota"`
	Reset int `json:"reset"`
}

type PrivilegeBoost struct {
	Multiplier int16 `json:"multiplier"`
	Duration   int   `json:"duration"`
	Quota      int   `json:"quota"`
	Reset      int   `json:"reset"`
}

type PrivilegeRoaming struct {
	Total    int `json:"total"`
	Duration int `json:"duration"`
	Reset    int `json:"reset"`
}
