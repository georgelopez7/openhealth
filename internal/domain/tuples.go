package domain

import "fmt"

const (
	EntityAccount       = "account"
	EntityDoctor        = "doctor"
	EntityNurse         = "nurse"
	EntityHospital      = "hospital"
	EntityMedicalRecord = "medical_record"
)

const (
	RelationAccount  = "account"
	RelationDoctor   = "doctor"
	RelationNurse    = "nurse"
	RelationHospital = "hospital"
	RelationOwner    = "owner"
)

type Tuple struct {
	UserType     string
	UserID       string
	UserRelation string
	Relation     string
	ObjectType   string
	ObjectID     string
}

// User - returns the tuple user in the form "userType:userID"
func (t Tuple) User() string {
	user := fmt.Sprintf("%s:%s", t.UserType, t.UserID)

	if t.UserRelation != "" {
		user = fmt.Sprintf("%s#%s", user, t.UserRelation)
	}

	return user
}

// Object - returns the tuple object in the form "objectType:objectID".
func (t Tuple) Object() string {
	return fmt.Sprintf("%s:%s", t.ObjectType, t.ObjectID)
}

// ToString - returns the tuple in the form "objectType:objectID#relation@user".
func (t Tuple) ToString() string {
	return fmt.Sprintf("%s#%s@%s", t.Object(), t.Relation, t.User())
}

// NewDoctorAccountTuple - creates the tuple that links a doctor to its account.
func NewDoctorAccountTuple(doctorID, accountID string) Tuple {
	return Tuple{
		UserType:   EntityAccount,
		UserID:     accountID,
		Relation:   RelationAccount,
		ObjectType: EntityDoctor,
		ObjectID:   doctorID,
	}
}

// NewDoctorToAccountAssignmentTuple - creates the tuple that assigns a doctor to an account.
func NewDoctorToAccountAssignmentTuple(accountID, doctorID string) Tuple {
	return Tuple{
		UserType:     EntityDoctor,
		UserID:       doctorID,
		UserRelation: RelationAccount,
		Relation:     RelationDoctor,
		ObjectType:   EntityAccount,
		ObjectID:     accountID,
	}
}

// NewNurseAccountTuple - creates the tuple that links a nurse to its account.
func NewNurseAccountTuple(nurseID, accountID string) Tuple {
	return Tuple{
		UserType:   EntityAccount,
		UserID:     accountID,
		Relation:   RelationAccount,
		ObjectType: EntityNurse,
		ObjectID:   nurseID,
	}
}

// NewNurseToHospitalAssignmentTuple - creates the tuple that assigns a nurse to a hospital.
func NewNurseToHospitalAssignmentTuple(nurseID, hospitalID string) Tuple {
	return Tuple{
		UserType:     EntityNurse,
		UserID:       nurseID,
		UserRelation: RelationAccount,
		Relation:     RelationNurse,
		ObjectType:   EntityHospital,
		ObjectID:     hospitalID,
	}
}

// NewHospitalToAccountAssignmentTuple - creates the tuple that assigns a hospital to an account.
func NewHospitalToAccountAssignmentTuple(accountID, hospitalID string) Tuple {
	return Tuple{
		UserType:   EntityHospital,
		UserID:     hospitalID,
		Relation:   RelationHospital,
		ObjectType: EntityAccount,
		ObjectID:   accountID,
	}
}

// NewMedicalRecordOwnerTuple - creates the tuple that makes an account the owner of a medical record.
func NewMedicalRecordOwnerTuple(recordID, accountID string) Tuple {
	return Tuple{
		UserType:   EntityAccount,
		UserID:     accountID,
		Relation:   RelationOwner,
		ObjectType: EntityMedicalRecord,
		ObjectID:   recordID,
	}
}
