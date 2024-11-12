package controller


import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"os"
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"github.com/labstack/echo/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProductUsecase struct {
	mock.Mock
}

func setMockJWTToken(c echo.Context, userID string) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
	})
	c.Set("user", token)
}

func (m *MockProductUsecase) GetProductList(organizationID uint) ([]model.ProductResponse, error){
	args := m.Called(organizationID)
	return nil, args.Error(0)
}

func (m *MockProductUsecase) SearchProduct(organizationID uint, term string) ([]model.ProductResponse, error){
	args := m.Called(organizationID, term)
	return nil, args.Error(0)
}
func (m *MockProductUsecase) CreateProduct(product model.Product) (model.ProductResponse, error) {
	args := m.Called(product)
	return model.ProductResponse{}, args.Error(0)
}

func (m *MockProductUsecase) UpdateProduct(product model.Product, id uint) (model.ProductResponse, error) {
	args := m.Called(product, id)
	return model.ProductResponse{}, args.Error(0)
}

func (m *MockProductUsecase) DeleteProduct(product model.Product, id uint) error {
	args := m.Called(product, id)
	return args.Error(0)
}
func TestCreateProduct(t *testing.T) {
    e := echo.New()
	e.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
	}))
    req := httptest.NewRequest(http.MethodPost, "/api/product", strings.NewReader(`{
    "image": "choco_cake.jpg",
    "name": "Chocolate Cake",
    "sell_price": 15.25
}`))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

	setMockJWTToken(c, "1")  // Set mock JWT token

    mockProductUsecase := new(MockProductUsecase)
    mockProductUsecase.On("CreateProduct", mock.AnythingOfType("*models.Product")).Return(nil)

    h := &productController{mockProductUsecase}

    if assert.NoError(t, h.CreateProduct(c)) {
        assert.Equal(t, http.StatusCreated, rec.Code)
        assert.Contains(t, rec.Body.String(), "Product 1")
    }

    mockProductUsecase.AssertExpectations(t)
	// mockProductUsecase = New(mockProductUsecase)
}
