// Code generated by entc, DO NOT EDIT.

package drug

const (
	// Label holds the string label denoting the drug type in the database.
	Label = "drug"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldHowto holds the string denoting the howto field in the database.
	FieldHowto = "howto"
	// FieldProperty holds the string denoting the property field in the database.
	FieldProperty = "property"

	// EdgeEmployee holds the string denoting the employee edge name in mutations.
	EdgeEmployee = "employee"
	// EdgeDrugtype holds the string denoting the drugtype edge name in mutations.
	EdgeDrugtype = "drugtype"
	// EdgeDisease holds the string denoting the disease edge name in mutations.
	EdgeDisease = "disease"

	// Table holds the table name of the drug in the database.
	Table = "drugs"
	// EmployeeTable is the table the holds the employee relation/edge.
	EmployeeTable = "drugs"
	// EmployeeInverseTable is the table name for the Employee entity.
	// It exists in this package in order to avoid circular dependency with the "employee" package.
	EmployeeInverseTable = "employees"
	// EmployeeColumn is the table column denoting the employee relation/edge.
	EmployeeColumn = "employee_drug"
	// DrugtypeTable is the table the holds the drugtype relation/edge.
	DrugtypeTable = "drugs"
	// DrugtypeInverseTable is the table name for the DrugType entity.
	// It exists in this package in order to avoid circular dependency with the "drugtype" package.
	DrugtypeInverseTable = "drug_types"
	// DrugtypeColumn is the table column denoting the drugtype relation/edge.
	DrugtypeColumn = "drug_type_drug"
	// DiseaseTable is the table the holds the disease relation/edge.
	DiseaseTable = "drugs"
	// DiseaseInverseTable is the table name for the Disease entity.
	// It exists in this package in order to avoid circular dependency with the "disease" package.
	DiseaseInverseTable = "diseases"
	// DiseaseColumn is the table column denoting the disease relation/edge.
	DiseaseColumn = "disease_drug"
)

// Columns holds all SQL columns for drug fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldHowto,
	FieldProperty,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Drug type.
var ForeignKeys = []string{
	"disease_drug",
	"drug_type_drug",
	"employee_drug",
}