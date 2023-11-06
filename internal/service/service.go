package service

import (
	"context"
	"ordering-system-backend/internal/domain"
	"ordering-system-backend/internal/repository"
)

// 不同層之間可能需要做資料轉換，所以delivery和service分開
type CreateOTPInput struct {
	Token string
	Email string
}

type VerifyOTPInput struct {
	Token    string
	Password string
}

type OTP interface {
	Create(ctx context.Context, input CreateOTPInput) error
	Verify(ctx context.Context, input VerifyOTPInput) error
}

type CreateUserInput struct {
	Email      string
	Password   string
	UserType   *int
	FirstName  string
	LastName   string
	LanguageId int
}

type UpdateUserInput struct {
	FirstName  string
	LastName   string
	LanguageId int
}

type ResetPasswordInput struct {
	UserId   string
	Password string
}

type Users interface {
	Create(ctx context.Context, input CreateUserInput) (string, error)
	Update(ctx context.Context, userId string, input UpdateUserInput) error
	Delete(ctx context.Context, userId string) error
	GetByEmail(ctx context.Context, email string, userType int) (string, error)
	GetByUid(ctx context.Context, uid string, userType int) (string, error)
	GetById(ctx context.Context, userId string) (domain.User, error)
	ResetPassword(ctx context.Context, input ResetPasswordInput) error
}

type Stores interface {
	Create(ctx context.Context, store domain.Store) (string, error)
	Update(ctx context.Context, store domain.Store) error
	Delete(ctx context.Context, userId string, storeId string) error
	GetAllByUserId(ctx context.Context, userId string) ([]domain.Store, error)
	GetByStoreId(ctx context.Context, userId string, storeId string) (domain.Store, error)
	// 不含UserId
	GetAll(ctx context.Context) ([]domain.Store, error)
}

type Seats interface {
	Create(ctx context.Context, seat domain.Seat) error
	Update(ctx context.Context, seat domain.Seat) error
	Delete(ctx context.Context, storeId string, seatId int) error
	GetAllByStoreId(ctx context.Context, storeId string) ([]domain.Seat, error)
	GetById(ctx context.Context, storeId string, seatId int) (domain.Seat, error)
}

type Categories interface {
	GetAllByUserId(ctx context.Context, userId string, languageId int) ([]domain.Category, error)
}

type CreateMenuInput struct {
	UserId      string
	Title       string
	Description string
	MenuItems   []MenuItemInput
}

type MenuItemInput struct {
	Title       string
	Description string
	Quantity    int
	Price       int
	CategoryId  int
	ImageBytes  []byte
}

type UpdateMenuInput struct {
	UserId      string
	MenuId      string
	Title       string
	Description string
	MenuItems   []MenuItemInput
}

type Menus interface {
	Create(ctx context.Context, input CreateMenuInput) (string, error)
	Update(ctx context.Context, input UpdateMenuInput) error
	Delete(ctx context.Context, userId string, menuId string) error
	GetAllByUserId(ctx context.Context, userId string, languageId int) ([]domain.Menu, error)
	GetById(ctx context.Context, userId string, menuId string, languageId int) (domain.Menu, error)
}

type StoreMenus interface {
	CreateMenuReference(ctx context.Context, userId string, storeId string, menuId string) error
	UpdateMenuReference(ctx context.Context, userId string, storeId string, menuId string) error
	DeleteMenuReference(ctx context.Context, userId string, storeId string) error
	GetMenuByStoreId(ctx context.Context, userId string, storeId string, languageId int) (domain.Menu, error)
}

type CreateOrderTicketInput struct {
	SeatId     int
	UserId     string
	TotalPrice float64
	OrderItems []OrderTicketItemInput
}

type OrderTicketItemInput struct {
	ProductId    int
	ProductName  string
	ProductPrice float64
	Quantity     int
}

type UpdateOrderTicketInput struct {
	OrderStatus string
}

type OrderTickets interface {
	Create(ctx context.Context, input CreateOrderTicketInput) error
	Update(ctx context.Context, storeId string, ticketId int, input UpdateOrderTicketInput) error
	GetAllByStoreId(ctx context.Context, storeId string) ([]domain.OrderTicket, error)
}

type Services struct {
	Users        Users
	OTP          OTP
	Stores       Stores
	Seats        Seats
	Categories   Categories
	Menus        Menus
	StoreMenus   StoreMenus
	OrderTickets OrderTickets
}

type Deps struct {
	Repos *repository.Repositories
}

func NewServices(deps Deps) *Services {
	usersService := NewUsersService(deps.Repos.Users)
	OTPService := NewOTPService(deps.Repos.OTP)
	categoriesService := NewCategoriesService(deps.Repos.Categories)
	menusService := NewMenusService(deps.Repos.Menus)
	storesService := NewStoresService(deps.Repos.Stores)
	seatsService := NewSeatsService(deps.Repos.Seats)
	storeMenusService := NewStoreMenusService(deps.Repos.StoreMenus)
	orderTicketsService := NewOrderTicketsService(deps.Repos.OrderTickets)

	return &Services{
		Users:        usersService,
		OTP:          OTPService,
		Stores:       storesService,
		Categories:   categoriesService,
		Menus:        menusService,
		Seats:        seatsService,
		StoreMenus:   storeMenusService,
		OrderTickets: orderTicketsService,
	}
}
