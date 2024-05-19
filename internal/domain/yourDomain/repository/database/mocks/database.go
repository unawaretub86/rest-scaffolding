package mock

import "github.com/unawaretub86/restScaffolding/utils/mocker"

type mockContainer struct {
	mocker mocker.MockDataStore
}

func NewMockDatabase() *mockContainer {
	return &mockContainer{
		mocker: mocker.New(),
	}
}
