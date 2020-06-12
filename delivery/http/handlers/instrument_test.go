package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"

	entity "github.com/hobord/invst-portfolio-backend-golang/domain/entity"
	mocks "github.com/hobord/invst-portfolio-backend-golang/mocks"

	"github.com/gorilla/mux"
	"github.com/icrowley/fake"
	mock "github.com/stretchr/testify/mock"
)

func TestGetInstrumentByID(t *testing.T) {
	t.Parallel()

	type wantStruct struct {
		httpStatusCode int
		responseString string
	}

	fakeID := rand.Intn(1000) + 1
	fakeName := fake.Sentence()
	fakeSymbol := fake.Sentence()
	fakeType := fake.Sentence()
	var testCases = []struct {
		input *entity.Instrument
		want  wantStruct
	}{
		{
			input: &entity.Instrument{
				ID:     fakeID,
				Name:   fakeName,
				Symbol: fakeSymbol,
				Type:   fakeType,
			},
			want: wantStruct{
				httpStatusCode: http.StatusOK,
				responseString: fmt.Sprintf(`{"instrumentId":%d,"name":"%s","symbol":"%s","instrumentType":"%s"}`, fakeID, fakeName, fakeSymbol, fakeType),
			},
		},
		{
			input: nil,
			want: wantStruct{
				httpStatusCode: http.StatusBadRequest,
				responseString: "id should be integer\n",
			},
		},
	}

	for _, testCase := range testCases {

		// Create a mock uses case interactor and mock the results
		mockInteractor := &mocks.InstrumentInteractorInterface{}
		mockInteractor.On("GetByID", mock.Anything, mock.Anything).Return(testCase.input, nil)

		// Create a test HTTPApp with moc usecase
		app := CreateInstrumentRestHTTPModule(mockInteractor)

		var req *http.Request
		var err error
		// Create a test request
		if testCase.input == nil {
			req, err = http.NewRequest("GET", "/instruments/BAD_REQUEST", nil)
		} else {
			req, err = http.NewRequest("GET", fmt.Sprintf("/instruments/%d", testCase.input.ID), nil)
		}
		if err != nil {
			t.Fatal(err)
		}

		// Create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()

		// Create a router and assign our handler func
		router := mux.NewRouter()
		router.HandleFunc("/instruments/{id}", app.GetInstrumentByID)

		// Make a request into the router
		router.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != testCase.want.httpStatusCode {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, testCase.want.httpStatusCode)
		}

		// Check the response body is what we expect.
		expected := testCase.want.responseString
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	}
}
