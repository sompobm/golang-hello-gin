package db

import (
	"hello-world/models"
	"time"

	"gopkg.in/mgo.v2/bson"
	// "time"
)

///////  is func adapter save record under database
func CreateCompany(in *models.Company) (*models.Company, error) {

	sessionCopy := pullSession()
	defer sessionCopy.Close()

	conn := sessionCopy.DB(databaseName)
	collection := conn.C("master_company")

	in.ID.Hex()
	in.CreatedAt = time.Now()
	in.UpdatedAt = in.CreatedAt

	err := collection.Insert(in)
	if err != nil {
		return nil, err
	}
	return in, nil
}

// GetAllComplaints is func get all Complaint
func GetAllCompany() ([]*models.Company, error) {
	sessionCopy := pullSession()
	defer sessionCopy.Close()
	var data []*models.Company

	conn := sessionCopy.DB(databaseName)
	collection := conn.C("master_company")

	err := collection.Find(nil).All(&data)

	if err != nil {

		return nil, err
	}
	return data, nil
}

func ValidateCompany(companyCode string) (*models.Company, error) {
	sessionCopy := pullSession()
	defer sessionCopy.Close()
	var data *models.Company

	conn := sessionCopy.DB(databaseName)
	collection := conn.C("master_company")

	err := collection.Find(bson.M{"company_code": companyCode}).One(&data)

	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetComplaint is func get one Complaint
func GetCompanyById(id string) (*models.Company, error) {

	sessionCopy := pullSession()
	defer sessionCopy.Close()
	var data *models.Company

	conn := sessionCopy.DB(databaseName)
	collection := conn.C("master_company")

	err := collection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&data)

	if err != nil {
		return nil, err
	}
	return data, nil
}

// UpdateComplaint is func update one Complaint
func UpdateCompany(id string, in *models.Company) (*models.Company, error) {
	sessionCopy := pullSession()
	defer sessionCopy.Close()
	_, err := GetCompanyById(id)
	if err != nil {
		return nil, err
	}

	query := bson.M{
		"company_name": in.Company_Name,
	}

	condition := bson.M{"_id": bson.ObjectIdHex(id)}

	conn := sessionCopy.DB(databaseName)
	collection := conn.C("master_company")

	err = collection.Update(condition, bson.M{"$set": query})

	if err != nil {
		return nil, err
	}
	data, err := GetCompanyById(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}
