package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	DocumentType string             `json:"documentType,readonly,omitempty"`
	Document     string             `json:"document,readonly,omitempty"`
	Name         string             `json:"name,readonly,omitempty"`
	LastName     string             `json:"lastName,readonly,omitempty"`
	Email        string             `json:"email,readonly,omitempty"`
	MobilePhone  string             `json:"mobilePhone,readonly,omitempty"`
	Password     string             `json:"password,omitempty"`
	UserName     string             `json:"userName,readonly,omitempty"`
	UserType     string             `json:"userType,readonly,omitempty"`
}
