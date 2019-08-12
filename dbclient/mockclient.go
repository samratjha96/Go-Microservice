package dbclient
import (
    "github.com/stretchr/testify/mock"
    "github.com/samratjha96/Go-Microservice/model"
)

type MockBoltClient struct {
    mock.Mock
}

func (m *MockBoltClient) QueryAccount(accountId string) (model.Account, error) {
    args := m.Mock.Called(accountId)
    return args.Get(0).(model.Account), args.Error(1)
}
func (m *MockBoltClient) OpenBoltDb() { 
	// No-op 
}
func (m *MockBoltClient) Seed() { 
	// No-op 
}

func (m *MockBoltClient) Check() bool {
    args := m.Mock.Called()
    return args.Get(0).(bool)
}