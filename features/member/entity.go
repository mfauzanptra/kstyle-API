package member

import "github.com/labstack/echo/v4"

type Core struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Gender    string `json:"gender"`
	Skintype  string `json:"skintype"`
	Skincolor string `json:"skincolor"`
}

type MemberHandler interface {
	Insert() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	GetMembers() echo.HandlerFunc
}

type MemberService interface {
	Insert(newMember Core) (Core, error)
	Update(IdMember uint, updMember Core) (Core, error)
	Delete(id uint) error
	GetMembers() ([]Core, error)
}

type MemberData interface {
	Insert(newMember Core) (Core, error)
	Update(updMember Core) (Core, error)
	Delete(id uint) error
	GetMembers() ([]Core, error)
}
