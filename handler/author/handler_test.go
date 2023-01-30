package author

//func TestHandler_GetByID(t *testing.T) {
//	testCases := []struct {
//		id                 int
//		method             string
//		expectedStatusCode int
//		expectedOutput     []byte
//	}{
//		{16, http.MethodGet, http.StatusOK, []byte(`{"id":16,"firstName":"Ganesh","lastName":"Manchi","penName":"Natraj","dateOfBirth":"2001-09-01","genre":"Fiction"}`)},
//		{1000, http.MethodGet, http.StatusNotFound, []byte(``)},
//		{1, http.MethodPost, http.StatusMethodNotAllowed, []byte(``)},
//	}
//	for i, tc := range testCases {
//		url := fmt.Sprintf("http://localhost:8000/author/%d", tc.id)
//		req := httptest.NewRequest(tc.method, url, nil)
//		w := httptest.NewRecorder()
//		authorStore := authorstore.MockStore{}
//		authorService := authorservice.New(authorStore)
//		authorHandler := New(authorService)
//		authorHandler.GetByID(w, req)
//		resp := w.Result()
//		body, err := io.ReadAll(resp.Body)
//		if err != nil {
//			t.Errorf("Error while reading the data")
//		}
//		if !reflect.DeepEqual(resp.StatusCode, tc.expectedStatusCode) {
//			t.Errorf("testcase %v failed, expected %v got %v", i, tc.expectedStatusCode, resp.StatusCode)
//		}
//		if !reflect.DeepEqual(body, tc.expectedOutput) {
//			t.Errorf("Task %v Failed", i)
//		}
//	}
//}
