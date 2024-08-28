package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	kr "github.com/99designs/keyring"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockKeyring struct {
	mock.Mock
}

type K8SHelperClientMock struct {
	mock.Mock
}

func (m *MockKeyring) Get(key string) (kr.Item, error) {
	args := m.Called(key)
	return args.Get(0).(kr.Item), args.Error(1)
}

func (m *MockKeyring) GetMetadata(key string) (kr.Metadata, error) {
	args := m.Called(key)
	return args.Get(0).(kr.Metadata), args.Error(1)
}

func (m *MockKeyring) Set(item kr.Item) error {
	args := m.Called(item)
	return args.Error(0)
}

func (m *MockKeyring) Remove(key string) error {
	args := m.Called(key)
	return args.Error(0)
}

func (m *MockKeyring) Keys() ([]string, error) {
	args := m.Called()
	return args.Get(0).([]string), args.Error(1)
}

// // CreateKeypair mocks the CreateKeypair method of the K8SHelperClient interface
// func (m *K8SHelperClientMock) CreateKeypair(ctx context.Context, in *dia.K8SHelperRequest, opts ...grpc.CallOption) (*proto.KeyPair, error) {
// 	args := m.Called(ctx, in, opts)
// 	return args.Get(0).(*proto.KeyPair), args.Error(1)
// }

// // GetKey mocks the GetKey method of the K8SHelperClient interface
// func (m *K8SHelperClientMock) GetKey(ctx context.Context, in *proto.K8SHelperRequest, opts ...grpc.CallOption) (*proto.KeyPair, error) {
// 	args := m.Called(ctx, in, opts)
// 	return args.Get(0).(*proto.KeyPair), args.Error(1)
// }

// // CreatePod mocks the CreatePod method of the K8SHelperClient interface
// func (m *K8SHelperClientMock) CreatePod(ctx context.Context, in *proto.FeederConfig, opts ...grpc.CallOption) (*proto.CreatePodResult, error) {
// 	args := m.Called(ctx, in, opts)
// 	return args.Get(0).(*proto.CreatePodResult), args.Error(1)
// }

// // RestartPod mocks the RestartPod method of the K8SHelperClient interface
// func (m *K8SHelperClientMock) RestartPod(ctx context.Context, in *proto.FeederConfig, opts ...grpc.CallOption) (*proto.RestartPodResult, error) {
// 	args := m.Called(ctx, in, opts)
// 	return args.Get(0).(*proto.RestartPodResult), args.Error(1)
// }

// // DeletePod mocks the DeletePod method of the K8SHelperClient interface
// func (m *K8SHelperClientMock) DeletePod(ctx context.Context, in *proto.FeederConfig, opts ...grpc.CallOption) (*proto.DeletePodResult, error) {
// 	args := m.Called(ctx, in, opts)
// 	return args.Get(0).(*proto.DeletePodResult), args.Error(1)
// }

func TestCreate(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockKeyring := new(MockKeyring)
	mockRelDB := new(RelDatastoreMock)
	// mockK8sbridgeClient := new(MockK8sbridgeClient)

	ob := &Env{
		Keyring: mockKeyring,
		RelDB:   mockRelDB,
		// k8sbridgeClient: mockK8sbridgeClient,
	}

	keyringItem := kr.Item{Data: []byte(`{"publicKey":"mockPublicKey"}`)}

	mockKeyring.On("Set", mock.Anything).Return(nil)
	mockKeyring.On("Get", "feeder-feeder-").Return(keyringItem, nil)

	mockRelDB.On("GetFeederByID", mock.Anything).Return("")
	mockRelDB.On("SetOracleConfig", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	mockRelDB.On("GetOracleConfig", mock.Anything, mock.Anything).Return(dia.OracleConfig{}, nil)

	// mockK8sbridgeClient.On("CreatePod", mock.Anything, mock.Anything).Return(nil, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Setup POST form data

	form := url.Values{}
	form.Add("oracleaddress", "0x0e42409d1a801397ce282faa828bf728e66d1c82")
	form.Add("chainID", "421614")
	form.Add("creator", "0xCf1e4c0455E69c3A82C2344648bE76C72DBcDa06")
	form.Add("signeddata", "0x1d83bf11265cd578928d6fca12ecd0ee46222321676db2670e9a79895833889d6ba5f1587b0ad22909371f9bb22a6826dec19e90f96e2836f0330fa9f5931e231b")
	form.Add("symbols", "Ethereum-0x84cA8bc7997272c7CfB4D0Cd3D55cd942B3c9419-DIA,Ethereum-0x4e3FBD56CD56c3e72c1403e103b45Db9da5B9D2B-CVX")
	form.Add("sleepseconds", "")
	form.Add("feederID", "")
	form.Add("mandatoryfrequency", "0")
	form.Add("frequency", "7200")
	form.Add("feedselection", `[{"Methodology":"MAIR","Symbol":"DIA-0x84cA8bc7997272c7CfB4D0Cd3D55cd942B3c9419","FeedSelection":[{"Address":"0x84cA8bc7997272c7CfB4D0Cd3D55cd942B3c9419","Blockchain":"Ethereum","Exchangepairs":[{"Exchange":"Binance","Pairs":["DIA-USDT","DIA-BTC"]},{"Exchange":"Bitmax","Pairs":["DIA-USDT"]},{"Exchange":"CoinBase","Pairs":["DIA-USD"]},{"Exchange":"UniswapV3","Pairs":["0xA14aFC841a2742Cbd52587b705F00f322309580e"]},{"Exchange":"OKEx","Pairs":["DIA-USDT"]},{"Exchange":"KuCoin","Pairs":["DIA-BTC","DIA-USDT"]},{"Exchange":"SushiSwap","Pairs":["0xC965BC2cf23D54ccF41a7d3aF10572580E9B8078"]}]}]},{"Methodology":"MAIR","Symbol":"CVX-0x4e3FBD56CD56c3e72c1403e103b45Db9da5B9D2B","FeedSelection":[{"Address":"0x4e3FBD56CD56c3e72c1403e103b45Db9da5B9D2B","Blockchain":"Ethereum","Exchangepairs":[{"Exchange":"Binance","Pairs":["CVX-USDT"]},{"Exchange":"UniswapV3","Pairs":["0x3Df6Bb62348D786A50b635b638fF99D850f98211"]},{"Exchange":"Curvefi","Pairs":["0xB576491F1E6e5E62f1d8F26062Ee822B40B0E0d4","0x47D5E1679Fe5f0D9f0A657c6715924e33Ce05093"]},{"Exchange":"Huobi","Pairs":["CVX-USDT"]},{"Exchange":"Bitmax","Pairs":["CVX-USDT"]},{"Exchange":"MEXC","Pairs":["CVX-USDT"]},{"Exchange":"SushiSwap","Pairs":["0x05767d9EF41dC40689678fFca0608878fb3dE906"]},{"Exchange":"BitMart","Pairs":["CVX-USDT"]}]}]}]`)
	form.Add("deviationpermille", "")
	c.Request = httptest.NewRequest("POST", "/create", strings.NewReader(form.Encode()))
	c.Request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	ob.Create(c)

	fmt.Println("-------------")

	assert.Equal(t, http.StatusCreated, w.Code)
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "0x1234", response)
	// assert.Equal(t, "1", response["chainId"])
	// assert.Equal(t, "0x5678", response["creator"])
	// assert.Equal(t, "sym1,sym2", response["symbols"])
	// assert.Equal(t, "mockPublicKey", response["publicKey"])

	// mockKeyring.AssertExpectations(t)
	// mockRelDB.AssertExpectations(t)
	// mockK8sbridgeClient.AssertExpectations(t)
}
