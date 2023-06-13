// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/ci-tools/pkg/jobrunaggregator/jobrunaggregatorlib (interfaces: CIDataClient)

// Package jobruntestcaseanalyzer is a generated GoMock package.
package jobruntestcaseanalyzer

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	jobrunaggregatorapi "github.com/openshift/ci-tools/pkg/jobrunaggregator/jobrunaggregatorapi"
	jobrunaggregatorlib "github.com/openshift/ci-tools/pkg/jobrunaggregator/jobrunaggregatorlib"
	sets "k8s.io/apimachinery/pkg/util/sets"
)

// MockCIDataClient is a mock of CIDataClient interface.
type MockCIDataClient struct {
	ctrl     *gomock.Controller
	recorder *MockCIDataClientMockRecorder
}

// MockCIDataClientMockRecorder is the mock recorder for MockCIDataClient.
type MockCIDataClientMockRecorder struct {
	mock *MockCIDataClient
}

// NewMockCIDataClient creates a new mock instance.
func NewMockCIDataClient(ctrl *gomock.Controller) *MockCIDataClient {
	mock := &MockCIDataClient{ctrl: ctrl}
	mock.recorder = &MockCIDataClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCIDataClient) EXPECT() *MockCIDataClientMockRecorder {
	return m.recorder
}

// GetBackendDisruptionRowCountByJob mocks base method.
func (m *MockCIDataClient) GetBackendDisruptionRowCountByJob(arg0 context.Context, arg1, arg2 string) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBackendDisruptionRowCountByJob", arg0, arg1, arg2)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBackendDisruptionRowCountByJob indicates an expected call of GetBackendDisruptionRowCountByJob.
func (mr *MockCIDataClientMockRecorder) GetBackendDisruptionRowCountByJob(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackendDisruptionRowCountByJob", reflect.TypeOf((*MockCIDataClient)(nil).GetBackendDisruptionRowCountByJob), arg0, arg1, arg2)
}

// GetBackendDisruptionStatisticsByJob mocks base method.
func (m *MockCIDataClient) GetBackendDisruptionStatisticsByJob(arg0 context.Context, arg1, arg2 string) ([]jobrunaggregatorapi.BackendDisruptionStatisticsRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBackendDisruptionStatisticsByJob", arg0, arg1, arg2)
	ret0, _ := ret[0].([]jobrunaggregatorapi.BackendDisruptionStatisticsRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBackendDisruptionStatisticsByJob indicates an expected call of GetBackendDisruptionStatisticsByJob.
func (mr *MockCIDataClientMockRecorder) GetBackendDisruptionStatisticsByJob(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackendDisruptionStatisticsByJob", reflect.TypeOf((*MockCIDataClient)(nil).GetBackendDisruptionStatisticsByJob), arg0, arg1, arg2)
}

// GetJobRunForJobNameAfterTime mocks base method.
func (m *MockCIDataClient) GetJobRunForJobNameAfterTime(arg0 context.Context, arg1 string, arg2 time.Time) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJobRunForJobNameAfterTime", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobRunForJobNameAfterTime indicates an expected call of GetJobRunForJobNameAfterTime.
func (mr *MockCIDataClientMockRecorder) GetJobRunForJobNameAfterTime(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobRunForJobNameAfterTime", reflect.TypeOf((*MockCIDataClient)(nil).GetJobRunForJobNameAfterTime), arg0, arg1, arg2)
}

// GetJobRunForJobNameBeforeTime mocks base method.
func (m *MockCIDataClient) GetJobRunForJobNameBeforeTime(arg0 context.Context, arg1 string, arg2 time.Time) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJobRunForJobNameBeforeTime", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobRunForJobNameBeforeTime indicates an expected call of GetJobRunForJobNameBeforeTime.
func (mr *MockCIDataClientMockRecorder) GetJobRunForJobNameBeforeTime(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobRunForJobNameBeforeTime", reflect.TypeOf((*MockCIDataClient)(nil).GetJobRunForJobNameBeforeTime), arg0, arg1, arg2)
}

// GetLastAggregationForJob mocks base method.
func (m *MockCIDataClient) GetLastAggregationForJob(arg0 context.Context, arg1, arg2 string) (*jobrunaggregatorapi.AggregatedTestRunRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastAggregationForJob", arg0, arg1, arg2)
	ret0, _ := ret[0].(*jobrunaggregatorapi.AggregatedTestRunRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLastAggregationForJob indicates an expected call of GetLastAggregationForJob.
func (mr *MockCIDataClientMockRecorder) GetLastAggregationForJob(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastAggregationForJob", reflect.TypeOf((*MockCIDataClient)(nil).GetLastAggregationForJob), arg0, arg1, arg2)
}

// GetLastJobRunEndTimeFromTable mocks base method.
func (m *MockCIDataClient) GetLastJobRunEndTimeFromTable(arg0 context.Context, arg1 string) (*time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastJobRunEndTimeFromTable", arg0, arg1)
	ret0, _ := ret[0].(*time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLastJobRunEndTimeFromTable indicates an expected call of GetLastJobRunEndTimeFromTable.
func (mr *MockCIDataClientMockRecorder) GetLastJobRunEndTimeFromTable(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastJobRunEndTimeFromTable", reflect.TypeOf((*MockCIDataClient)(nil).GetLastJobRunEndTimeFromTable), arg0, arg1)
}

// ListAggregatedTestRunsForJob mocks base method.
func (m *MockCIDataClient) ListAggregatedTestRunsForJob(arg0 context.Context, arg1, arg2 string, arg3 time.Time) ([]jobrunaggregatorapi.AggregatedTestRunRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAggregatedTestRunsForJob", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]jobrunaggregatorapi.AggregatedTestRunRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAggregatedTestRunsForJob indicates an expected call of ListAggregatedTestRunsForJob.
func (mr *MockCIDataClientMockRecorder) ListAggregatedTestRunsForJob(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAggregatedTestRunsForJob", reflect.TypeOf((*MockCIDataClient)(nil).ListAggregatedTestRunsForJob), arg0, arg1, arg2, arg3)
}

// ListAlertHistoricalData mocks base method.
func (m *MockCIDataClient) ListAlertHistoricalData(arg0 context.Context) ([]jobrunaggregatorapi.HistoricalData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAlertHistoricalData", arg0)
	ret0, _ := ret[0].([]jobrunaggregatorapi.HistoricalData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAlertHistoricalData indicates an expected call of ListAlertHistoricalData.
func (mr *MockCIDataClientMockRecorder) ListAlertHistoricalData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAlertHistoricalData", reflect.TypeOf((*MockCIDataClient)(nil).ListAlertHistoricalData), arg0)
}

// ListAllJobs mocks base method.
func (m *MockCIDataClient) ListAllJobs(arg0 context.Context) ([]jobrunaggregatorapi.JobRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllJobs", arg0)
	ret0, _ := ret[0].([]jobrunaggregatorapi.JobRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllJobs indicates an expected call of ListAllJobs.
func (mr *MockCIDataClientMockRecorder) ListAllJobs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllJobs", reflect.TypeOf((*MockCIDataClient)(nil).ListAllJobs), arg0)
}

// ListAllKnownAlerts mocks base method.
func (m *MockCIDataClient) ListAllKnownAlerts(arg0 context.Context) ([]*jobrunaggregatorapi.KnownAlertRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllKnownAlerts", arg0)
	ret0, _ := ret[0].([]*jobrunaggregatorapi.KnownAlertRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllKnownAlerts indicates an expected call of ListAllKnownAlerts.
func (mr *MockCIDataClientMockRecorder) ListAllKnownAlerts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllKnownAlerts", reflect.TypeOf((*MockCIDataClient)(nil).ListAllKnownAlerts), arg0)
}

// ListDisruptionHistoricalData mocks base method.
func (m *MockCIDataClient) ListDisruptionHistoricalData(arg0 context.Context) ([]jobrunaggregatorapi.HistoricalData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDisruptionHistoricalData", arg0)
	ret0, _ := ret[0].([]jobrunaggregatorapi.HistoricalData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDisruptionHistoricalData indicates an expected call of ListDisruptionHistoricalData.
func (mr *MockCIDataClientMockRecorder) ListDisruptionHistoricalData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDisruptionHistoricalData", reflect.TypeOf((*MockCIDataClient)(nil).ListDisruptionHistoricalData), arg0)
}

// ListProwJobRunsSince mocks base method.
func (m *MockCIDataClient) ListProwJobRunsSince(arg0 context.Context, arg1 *time.Time) ([]*jobrunaggregatorapi.TestPlatformProwJobRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProwJobRunsSince", arg0, arg1)
	ret0, _ := ret[0].([]*jobrunaggregatorapi.TestPlatformProwJobRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProwJobRunsSince indicates an expected call of ListProwJobRunsSince.
func (mr *MockCIDataClientMockRecorder) ListProwJobRunsSince(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProwJobRunsSince", reflect.TypeOf((*MockCIDataClient)(nil).ListProwJobRunsSince), arg0, arg1)
}

// ListReleaseTags mocks base method.
func (m *MockCIDataClient) ListReleaseTags(arg0 context.Context) (sets.String, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListReleaseTags", arg0)
	ret0, _ := ret[0].(sets.String)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListReleaseTags indicates an expected call of ListReleaseTags.
func (mr *MockCIDataClientMockRecorder) ListReleaseTags(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListReleaseTags", reflect.TypeOf((*MockCIDataClient)(nil).ListReleaseTags), arg0)
}

// ListUnifiedTestRunsForJobAfterDay mocks base method.
func (m *MockCIDataClient) ListUnifiedTestRunsForJobAfterDay(arg0 context.Context, arg1 string, arg2 time.Time) (*jobrunaggregatorlib.UnifiedTestRunRowIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUnifiedTestRunsForJobAfterDay", arg0, arg1, arg2)
	ret0, _ := ret[0].(*jobrunaggregatorlib.UnifiedTestRunRowIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUnifiedTestRunsForJobAfterDay indicates an expected call of ListUnifiedTestRunsForJobAfterDay.
func (mr *MockCIDataClientMockRecorder) ListUnifiedTestRunsForJobAfterDay(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUnifiedTestRunsForJobAfterDay", reflect.TypeOf((*MockCIDataClient)(nil).ListUnifiedTestRunsForJobAfterDay), arg0, arg1, arg2)
}

// ListUploadedJobRunIDsSinceFromTable mocks base method.
func (m *MockCIDataClient) ListUploadedJobRunIDsSinceFromTable(arg0 context.Context, arg1 string, arg2 *time.Time) (map[string]bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUploadedJobRunIDsSinceFromTable", arg0, arg1, arg2)
	ret0, _ := ret[0].(map[string]bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUploadedJobRunIDsSinceFromTable indicates an expected call of ListUploadedJobRunIDsSinceFromTable.
func (mr *MockCIDataClientMockRecorder) ListUploadedJobRunIDsSinceFromTable(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUploadedJobRunIDsSinceFromTable", reflect.TypeOf((*MockCIDataClient)(nil).ListUploadedJobRunIDsSinceFromTable), arg0, arg1, arg2)
}
