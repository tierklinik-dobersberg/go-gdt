package gdt

import "errors"

// Common error definitions
var (
	ErrInvalidLength = errors.New("field has invalid length")
)

var (
	FieldSoftwareResponsible = FieldDesc{
		ID:        102,
		MaxLength: 60,
		Type:      FieldTypeText,
	}

	FieldSoftware = FieldDesc{
		ID:        103,
		MaxLength: 60,
		Type:      FieldTypeText,
	}

	FieldSoftwareRelease = FieldDesc{
		ID:        132,
		MaxLength: 60,
		Type:      FieldTypeText,
	}

	FieldPatientID = FieldDesc{
		ID:        3000,
		MaxLength: 10,
		Type:      FieldTypeText,
	}

	FieldPatientNameSuffix = FieldDesc{
		ID:        3100,
		MaxLength: 15,
		Type:      FieldTypeText,
	}

	FieldPatientName = FieldDesc{
		ID:        3101,
		MaxLength: 28,
		Type:      FieldTypeText,
	}

	FieldPatientGivenName = FieldDesc{
		ID:        3102,
		MaxLength: 28,
		Type:      FieldTypeText,
	}

	FieldPatientBrithday = FieldDesc{
		ID:     3103, // TTMMJJJJ
		Length: 8,
		Type:   FieldTypeText,
	}

	FieldPatientTitel = FieldDesc{
		ID:        3104,
		MaxLength: 15,
		Type:      FieldTypeText,
	}

	FieldPatientInsuranceNumber = FieldDesc{
		ID:        3105,
		MaxLength: 12,
		Type:      FieldTypeText,
	}

	FieldPatientCity = FieldDesc{
		ID:        3106,
		MaxLength: 30,
		Type:      FieldTypeText,
	}

	FieldPatientStreet = FieldDesc{
		ID:        3107,
		MaxLength: 28,
		Type:      FieldTypeText,
	}

	FieldPatientInsuranceType = FieldDesc{
		ID:     3108, // 1 = Mitglied, 3 = Famielnversicherter, 5 = Rentner
		Length: 1,
		Type:   FieldTypeNumber,
	}

	FieldPatientGender = FieldDesc{
		ID:     3110, // 1 = male / 2 = female
		Length: 1,
		Type:   FieldTypeNumber,
	}

	FieldPatientSize = FieldDesc{
		ID:   3622, // in cm; variable length; float
		Type: FieldTypeFloat,
	}

	FieldPatientWeight = FieldDesc{
		ID:   3623, // in kg; variable length; float
		Type: FieldTypeFloat,
	}

	FieldPatientNativeLanguage = FieldDesc{
		ID:        3628,
		MaxLength: 60,
		Type:      FieldTypeText,
	}

	FieldStudyDate = FieldDesc{
		ID:     6200, // TTMMJJJJ
		Length: 8,
		Type:   FieldTypeText,
	}

	FieldStudyTime = FieldDesc{
		ID:     6201, // HHMMSS
		Length: 6,
		Type:   FieldTypeText,
	}

	FieldDiagnose = FieldDesc{
		ID:        6205,
		MaxLength: 60,
		Type:      FieldTypeText,
	}

	FieldFindings = FieldDesc{
		ID:        6220,
		MaxLength: 60,
		Type:      FieldTypeText,
	}

	FieldExternalFindings = FieldDesc{
		ID:        6221,
		MaxLength: 60,
		Type:      FieldTypeText,
	}

	FieldResultTextLength = FieldDesc{
		ID:        6226,
		MaxLength: 60,
		Type:      FieldTypeText,
	}

	FieldComment = FieldDesc{
		ID:        6227,
		MaxLength: 60,
		Type:      FieldTypeText,
	}

	FieldResultTableText = FieldDesc{
		ID:   6228, // multiple times
		Type: FieldTypeText,
	}

	FieldArchiveID = FieldDesc{
		ID:        6302,
		MaxLength: 60,
		Type:      FieldTypeText,
	}

	FieldFileFormat = FieldDesc{
		ID:        6303,
		MaxLength: 60,
		Type:      FieldTypeText,
	}

	FieldFileContent = FieldDesc{
		ID:        6304,
		MaxLength: 60,
		Type:      FieldTypeText,
	}

	FieldFileURL = FieldDesc{
		ID:        6305,
		MaxLength: 60,
		Type:      FieldTypeText,
	}

	FieldFreeCategoryName = FieldDesc{
		ID:        6330,
		MaxLength: 60,
		Type:      FieldTypeText,
	}

	FieldFreeCategoryContent = FieldDesc{
		ID:        6331,
		MaxLength: 60,
		Type:      FieldTypeText,
	}

	FieldRecordType = FieldDesc{
		ID:     8000,
		Type:   FieldTypeText,
		Length: 4,
	}

	FieldRecordLength = FieldDesc{
		ID:     8100,
		Length: 5,
		Type:   FieldTypeNumber,
	}

	FieldReceiverID = FieldDesc{
		ID:     8315,
		Type:   FieldTypeText,
		Length: 8,
	}

	FieldSenderID = FieldDesc{
		ID:     8316,
		Type:   FieldTypeText,
		Length: 8,
	}

	FieldProcedureID = FieldDesc{
		ID:        8402,
		MaxLength: 6,
		Type:      FieldTypeText,
	}

	FieldTestIdent = FieldDesc{
		ID:        8410,
		MaxLength: 20,
		Type:      FieldTypeText,
	}

	FieldTestDescription = FieldDesc{
		ID:        8411,
		MaxLength: 60,
		Type:      FieldTypeText,
	}

	FieldTestStatus = FieldDesc{
		ID:        8418,
		MaxLength: 1,
		Type:      FieldTypeText,
	}

	FieldResultValue = FieldDesc{
		ID:   8420, // variable length; float
		Type: FieldTypeFloat,
	}

	FieldUnit = FieldDesc{
		ID:        8421,
		MaxLength: 60,
		Type:      FieldTypeText,
	}

	FieldSampleMaterialIdent = FieldDesc{
		ID:        8428,
		MaxLength: 8,
		Type:      FieldTypeText,
	}

	FieldSampleMaterialIndex = FieldDesc{
		ID:     8429,
		Length: 2,
		Type:   FieldTypeNumber,
	}

	FieldSampleMaterialDescription = FieldDesc{
		ID:        8430,
		MaxLength: 60,
		Type:      FieldTypeText,
	}

	FieldSampleMaterialSpecificatoin = FieldDesc{
		ID:        8431,
		MaxLength: 60,
		Type:      FieldTypeText,
	}

	FieldAcceptanceDate = FieldDesc{
		ID:     8432, // TTMMJJJJ
		Type:   FieldTypeText,
		Length: 8,
	}

	FieldDataUnit = FieldDesc{
		ID:        8437,
		MaxLength: 60,
		Type:      FieldTypeText,
	}

	FieldData = FieldDesc{
		ID:        8438,
		MaxLength: 60,
		Type:      FieldTypeText,
	}

	FieldAcceptanceTime = FieldDesc{
		ID:     8439, // HHMMSS
		Type:   FieldTypeText,
		Length: 6,
	}

	FieldNormalValueText = FieldDesc{
		ID:        8460,
		MaxLength: 60,
		Type:      FieldTypeText,
	}

	FieldNormalValueLowerBoundary = FieldDesc{
		ID:   8461, // variable length; float
		Type: FieldTypeFloat,
	}

	FieldNormalValueUpperBoundary = FieldDesc{
		ID:   8462, // variable length; float
		Type: FieldTypeFloat,
	}

	FieldTestAnnotation = FieldDesc{
		ID:        8470,
		MaxLength: 60,
		Type:      FieldTypeText,
	}

	FieldResultText = FieldDesc{
		ID:        8480,
		MaxLength: 60,
		Type:      FieldTypeText,
	}

	FieldSignature = FieldDesc{
		ID:        8990,
		MaxLength: 60,
		Type:      FieldTypeText,
	}

	FieldCharacterSet = FieldDesc{
		ID:     9206, // 1 byte; number
		Length: 1,
		Type:   FieldTypeNumber,
	}

	FieldVersion = FieldDesc{
		ID:   9218, // 02.10
		Type: FieldTypeText,
	}
)
