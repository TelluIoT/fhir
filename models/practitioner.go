// Copyright (c) 2011-2017, HL7, Inc & The MITRE Corporation
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without modification,
// are permitted provided that the following conditions are met:
//
//     * Redistributions of source code must retain the above copyright notice, this
//       list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above copyright notice,
//       this list of conditions and the following disclaimer in the documentation
//       and/or other materials provided with the distribution.
//     * Neither the name of HL7 nor the names of its contributors may be used to
//       endorse or promote products derived from this software without specific
//       prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
// IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT,
// INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT
// NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
// PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
// WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Practitioner struct {
	ID             string `json:"id,omitempty"`
	DomainResource `bson:",inline"`
	Identifier     []Identifier                         `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Active         *bool                                `bson:"active,omitempty" json:"active,omitempty"`
	Name           []HumanName                          `bson:"name,omitempty" json:"name,omitempty"`
	Telecom        []ContactPoint                       `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Address        []Address                            `bson:"address,omitempty" json:"address,omitempty"`
	Gender         string                               `bson:"gender,omitempty" json:"gender,omitempty"`
	BirthDate      *FHIRDateTime                        `bson:"birthDate,omitempty" json:"birthDate,omitempty"`
	Photo          []Attachment                         `bson:"photo,omitempty" json:"photo,omitempty"`
	Qualification  []PractitionerQualificationComponent `bson:"qualification,omitempty" json:"qualification,omitempty"`
	Communication  []CodeableConcept                    `bson:"communication,omitempty" json:"communication,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Practitioner) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Practitioner"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Practitioner), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Practitioner) GetBSON() (interface{}, error) {
	x.ResourceType = "Practitioner"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "practitioner" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type practitioner Practitioner

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Practitioner) UnmarshalJSON(data []byte) (err error) {
	x2 := practitioner{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i], err = MapToResource(x2.Contained[i], true)
				if err != nil {
					return err
				}
			}
		}
		*x = Practitioner(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Practitioner) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Practitioner"
	} else if x.ResourceType != "Practitioner" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Practitioner, instead received %s", x.ResourceType))
	}
	return nil
}

type PractitionerQualificationComponent struct {
	BackboneElement `bson:",inline"`
	Identifier      []Identifier     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Code            *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Period          *Period          `bson:"period,omitempty" json:"period,omitempty"`
	Issuer          *Reference       `bson:"issuer,omitempty" json:"issuer,omitempty"`
}
