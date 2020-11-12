// Code generated by mockery v2.4.0-beta. DO NOT EDIT.

package mocks

import (
	context "context"
	pb "in-backend/services/profile/pb"

	mock "github.com/stretchr/testify/mock"
)

// ProfileServiceServer is an autogenerated mock type for the ProfileServiceServer type
type ProfileServiceServer struct {
	mock.Mock
}

// CreateAcademicHistory provides a mock function with given fields: _a0, _a1
func (_m *ProfileServiceServer) CreateAcademicHistory(_a0 context.Context, _a1 *pb.CreateAcademicHistoryRequest) (*pb.AcademicHistory, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.AcademicHistory
	if rf, ok := ret.Get(0).(func(context.Context, *pb.CreateAcademicHistoryRequest) *pb.AcademicHistory); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.AcademicHistory)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.CreateAcademicHistoryRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateCandidate provides a mock function with given fields: _a0, _a1
func (_m *ProfileServiceServer) CreateCandidate(_a0 context.Context, _a1 *pb.CreateCandidateRequest) (*pb.Candidate, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.Candidate
	if rf, ok := ret.Get(0).(func(context.Context, *pb.CreateCandidateRequest) *pb.Candidate); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.Candidate)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.CreateCandidateRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateCompany provides a mock function with given fields: _a0, _a1
func (_m *ProfileServiceServer) CreateCompany(_a0 context.Context, _a1 *pb.CreateCompanyRequest) (*pb.Company, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.Company
	if rf, ok := ret.Get(0).(func(context.Context, *pb.CreateCompanyRequest) *pb.Company); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.Company)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.CreateCompanyRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateCourse provides a mock function with given fields: _a0, _a1
func (_m *ProfileServiceServer) CreateCourse(_a0 context.Context, _a1 *pb.CreateCourseRequest) (*pb.Course, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.Course
	if rf, ok := ret.Get(0).(func(context.Context, *pb.CreateCourseRequest) *pb.Course); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.Course)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.CreateCourseRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateDepartment provides a mock function with given fields: _a0, _a1
func (_m *ProfileServiceServer) CreateDepartment(_a0 context.Context, _a1 *pb.CreateDepartmentRequest) (*pb.Department, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.Department
	if rf, ok := ret.Get(0).(func(context.Context, *pb.CreateDepartmentRequest) *pb.Department); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.Department)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.CreateDepartmentRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateInstitution provides a mock function with given fields: _a0, _a1
func (_m *ProfileServiceServer) CreateInstitution(_a0 context.Context, _a1 *pb.CreateInstitutionRequest) (*pb.Institution, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.Institution
	if rf, ok := ret.Get(0).(func(context.Context, *pb.CreateInstitutionRequest) *pb.Institution); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.Institution)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.CreateInstitutionRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateJobHistory provides a mock function with given fields: _a0, _a1
func (_m *ProfileServiceServer) CreateJobHistory(_a0 context.Context, _a1 *pb.CreateJobHistoryRequest) (*pb.JobHistory, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.JobHistory
	if rf, ok := ret.Get(0).(func(context.Context, *pb.CreateJobHistoryRequest) *pb.JobHistory); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.JobHistory)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.CreateJobHistoryRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSkill provides a mock function with given fields: _a0, _a1
func (_m *ProfileServiceServer) CreateSkill(_a0 context.Context, _a1 *pb.CreateSkillRequest) (*pb.Skill, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.Skill
	if rf, ok := ret.Get(0).(func(context.Context, *pb.CreateSkillRequest) *pb.Skill); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.Skill)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.CreateSkillRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAcademicHistory provides a mock function with given fields: _a0, _a1
func (_m *ProfileServiceServer) DeleteAcademicHistory(_a0 context.Context, _a1 *pb.DeleteAcademicHistoryRequest) (*pb.DeleteAcademicHistoryResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.DeleteAcademicHistoryResponse
	if rf, ok := ret.Get(0).(func(context.Context, *pb.DeleteAcademicHistoryRequest) *pb.DeleteAcademicHistoryResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.DeleteAcademicHistoryResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.DeleteAcademicHistoryRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCandidate provides a mock function with given fields: _a0, _a1
func (_m *ProfileServiceServer) DeleteCandidate(_a0 context.Context, _a1 *pb.DeleteCandidateRequest) (*pb.DeleteCandidateResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.DeleteCandidateResponse
	if rf, ok := ret.Get(0).(func(context.Context, *pb.DeleteCandidateRequest) *pb.DeleteCandidateResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.DeleteCandidateResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.DeleteCandidateRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteJobHistory provides a mock function with given fields: _a0, _a1
func (_m *ProfileServiceServer) DeleteJobHistory(_a0 context.Context, _a1 *pb.DeleteJobHistoryRequest) (*pb.DeleteJobHistoryResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.DeleteJobHistoryResponse
	if rf, ok := ret.Get(0).(func(context.Context, *pb.DeleteJobHistoryRequest) *pb.DeleteJobHistoryResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.DeleteJobHistoryResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.DeleteJobHistoryRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAcademicHistory provides a mock function with given fields: _a0, _a1
func (_m *ProfileServiceServer) GetAcademicHistory(_a0 context.Context, _a1 *pb.GetAcademicHistoryRequest) (*pb.AcademicHistory, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.AcademicHistory
	if rf, ok := ret.Get(0).(func(context.Context, *pb.GetAcademicHistoryRequest) *pb.AcademicHistory); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.AcademicHistory)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.GetAcademicHistoryRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllCandidates provides a mock function with given fields: _a0, _a1
func (_m *ProfileServiceServer) GetAllCandidates(_a0 context.Context, _a1 *pb.GetAllCandidatesRequest) (*pb.GetAllCandidatesResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.GetAllCandidatesResponse
	if rf, ok := ret.Get(0).(func(context.Context, *pb.GetAllCandidatesRequest) *pb.GetAllCandidatesResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.GetAllCandidatesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.GetAllCandidatesRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllCompanies provides a mock function with given fields: _a0, _a1
func (_m *ProfileServiceServer) GetAllCompanies(_a0 context.Context, _a1 *pb.GetAllCompaniesRequest) (*pb.GetAllCompaniesResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.GetAllCompaniesResponse
	if rf, ok := ret.Get(0).(func(context.Context, *pb.GetAllCompaniesRequest) *pb.GetAllCompaniesResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.GetAllCompaniesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.GetAllCompaniesRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllCourses provides a mock function with given fields: _a0, _a1
func (_m *ProfileServiceServer) GetAllCourses(_a0 context.Context, _a1 *pb.GetAllCoursesRequest) (*pb.GetAllCoursesResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.GetAllCoursesResponse
	if rf, ok := ret.Get(0).(func(context.Context, *pb.GetAllCoursesRequest) *pb.GetAllCoursesResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.GetAllCoursesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.GetAllCoursesRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllDepartments provides a mock function with given fields: _a0, _a1
func (_m *ProfileServiceServer) GetAllDepartments(_a0 context.Context, _a1 *pb.GetAllDepartmentsRequest) (*pb.GetAllDepartmentsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.GetAllDepartmentsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *pb.GetAllDepartmentsRequest) *pb.GetAllDepartmentsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.GetAllDepartmentsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.GetAllDepartmentsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllInstitutions provides a mock function with given fields: _a0, _a1
func (_m *ProfileServiceServer) GetAllInstitutions(_a0 context.Context, _a1 *pb.GetAllInstitutionsRequest) (*pb.GetAllInstitutionsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.GetAllInstitutionsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *pb.GetAllInstitutionsRequest) *pb.GetAllInstitutionsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.GetAllInstitutionsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.GetAllInstitutionsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllSkills provides a mock function with given fields: _a0, _a1
func (_m *ProfileServiceServer) GetAllSkills(_a0 context.Context, _a1 *pb.GetAllSkillsRequest) (*pb.GetAllSkillsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.GetAllSkillsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *pb.GetAllSkillsRequest) *pb.GetAllSkillsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.GetAllSkillsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.GetAllSkillsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCandidateByID provides a mock function with given fields: _a0, _a1
func (_m *ProfileServiceServer) GetCandidateByID(_a0 context.Context, _a1 *pb.GetCandidateByIDRequest) (*pb.Candidate, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.Candidate
	if rf, ok := ret.Get(0).(func(context.Context, *pb.GetCandidateByIDRequest) *pb.Candidate); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.Candidate)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.GetCandidateByIDRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCompany provides a mock function with given fields: _a0, _a1
func (_m *ProfileServiceServer) GetCompany(_a0 context.Context, _a1 *pb.GetCompanyRequest) (*pb.Company, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.Company
	if rf, ok := ret.Get(0).(func(context.Context, *pb.GetCompanyRequest) *pb.Company); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.Company)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.GetCompanyRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCourse provides a mock function with given fields: _a0, _a1
func (_m *ProfileServiceServer) GetCourse(_a0 context.Context, _a1 *pb.GetCourseRequest) (*pb.Course, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.Course
	if rf, ok := ret.Get(0).(func(context.Context, *pb.GetCourseRequest) *pb.Course); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.Course)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.GetCourseRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDepartment provides a mock function with given fields: _a0, _a1
func (_m *ProfileServiceServer) GetDepartment(_a0 context.Context, _a1 *pb.GetDepartmentRequest) (*pb.Department, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.Department
	if rf, ok := ret.Get(0).(func(context.Context, *pb.GetDepartmentRequest) *pb.Department); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.Department)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.GetDepartmentRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInstitution provides a mock function with given fields: _a0, _a1
func (_m *ProfileServiceServer) GetInstitution(_a0 context.Context, _a1 *pb.GetInstitutionRequest) (*pb.Institution, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.Institution
	if rf, ok := ret.Get(0).(func(context.Context, *pb.GetInstitutionRequest) *pb.Institution); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.Institution)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.GetInstitutionRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetJobHistory provides a mock function with given fields: _a0, _a1
func (_m *ProfileServiceServer) GetJobHistory(_a0 context.Context, _a1 *pb.GetJobHistoryRequest) (*pb.JobHistory, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.JobHistory
	if rf, ok := ret.Get(0).(func(context.Context, *pb.GetJobHistoryRequest) *pb.JobHistory); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.JobHistory)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.GetJobHistoryRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSkill provides a mock function with given fields: _a0, _a1
func (_m *ProfileServiceServer) GetSkill(_a0 context.Context, _a1 *pb.GetSkillRequest) (*pb.Skill, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.Skill
	if rf, ok := ret.Get(0).(func(context.Context, *pb.GetSkillRequest) *pb.Skill); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.Skill)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.GetSkillRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAcademicHistory provides a mock function with given fields: _a0, _a1
func (_m *ProfileServiceServer) UpdateAcademicHistory(_a0 context.Context, _a1 *pb.UpdateAcademicHistoryRequest) (*pb.AcademicHistory, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.AcademicHistory
	if rf, ok := ret.Get(0).(func(context.Context, *pb.UpdateAcademicHistoryRequest) *pb.AcademicHistory); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.AcademicHistory)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.UpdateAcademicHistoryRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCandidate provides a mock function with given fields: _a0, _a1
func (_m *ProfileServiceServer) UpdateCandidate(_a0 context.Context, _a1 *pb.UpdateCandidateRequest) (*pb.Candidate, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.Candidate
	if rf, ok := ret.Get(0).(func(context.Context, *pb.UpdateCandidateRequest) *pb.Candidate); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.Candidate)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.UpdateCandidateRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateJobHistory provides a mock function with given fields: _a0, _a1
func (_m *ProfileServiceServer) UpdateJobHistory(_a0 context.Context, _a1 *pb.UpdateJobHistoryRequest) (*pb.JobHistory, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.JobHistory
	if rf, ok := ret.Get(0).(func(context.Context, *pb.UpdateJobHistoryRequest) *pb.JobHistory); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.JobHistory)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.UpdateJobHistoryRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
