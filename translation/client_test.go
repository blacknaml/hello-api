package translation_test

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/blacknaml/hello-api/translation"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockService struct {
	mock.Mock
}

type HelloClientSuite struct {
	suite.Suite
	mockServerSuite *MockService
	server          *httptest.Server
	underTest       translation.HelloClient
}

func TestHelloClientSuite(t *testing.T) {
	suite.Run(t, new(HelloClientSuite))
}

func (m *MockService) Translate(word, language string) (string, error) {
	args := m.Called(word, language)
	return args.String(0), args.Error(1)
}

func (suite *HelloClientSuite) SetupSuite() {
	suite.mockServerSuite = new(MockService)
	handler := func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		defer func() {
			if closeErr := r.Body.Close(); closeErr != nil {
				// Log the error, but don't necessarily return it as it might overshadow
				// a more significant error from the main function logic.
				log.Printf("Error closing response body: %v", closeErr)
			}
		}()

		var m map[string]interface{}
		_ = json.Unmarshal(b, &m)

		word := m["word"].(string)
		language := m["language"].(string)

		resp, err := suite.mockServerSuite.Translate(word, language)
		if err != nil {
			http.Error(w, "error", 500)
		}
		if resp == "" {
			http.Error(w, "missing", 404)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, resp)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	suite.server = httptest.NewServer(mux)

	suite.underTest = translation.NewHelloClient(suite.server.URL)
}

func (suite *HelloClientSuite) TearDownSuite() {
	suite.server.Close()
}

func (suite *HelloClientSuite) TestCall() {
	// Arrange
	suite.mockServerSuite.On("Translate", "foo", "bar").Return(`{"translation":"baz"}`, nil)

	// Act
	resp, err := suite.underTest.Translate("foo", "bar")

	// Assert
	suite.NoError(err)
	suite.Equal(resp, "baz")
}

func (suite *HelloClientSuite) TestCall_APIError() {
	// Arrange
	suite.mockServerSuite.On("Translate", "foo", "bar").Return("", errors.New("this is a test"))

	// Act
	resp, err := suite.underTest.Translate("foo", "bar")

	// Assert
	suite.EqualError(err, "error in api")
	suite.Equal(resp, "")
}

func (suite *HelloClientSuite) TestCall_InvalidJSON() {
	// Arrange
	suite.mockServerSuite.On("Translate", "foo", "bar").Return(`invalid json`, nil)

	// Act
	resp, err := suite.underTest.Translate("foo", "bar")

	// Assert
	suite.EqualError(err, "unable to decode message")
	suite.Equal(resp, "")
}
