package july

import "fmt"

// Contact struct defines the structure of a contact
type Contact struct {
	Name  string
	Phone string
	Email string
}

func Week3Result() string {
	var result string
	// Create an empty map to store contact details
	// Key: string (name or identifier), Value: Contact struct
	contactDetails := make(map[string]Contact)

	// Add a contact with the key "John"
	contactDetails["John"] = Contact{
		Name:  "John Flores",
		Phone: "09123456789",
		Email: "sample1@mail.com",
	}

	contactDetails["Frey"] = Contact{
		Name:  "Freya Muji",
		Phone: "09123456789",
		Email: "sample2@mail.com",
	}

	contactDetails["Freya"] = Contact{
		Name:  "Freya Muji",
		Phone: "09123456789",
		Email: "sample2@mail.com",
	}

	// Search for a contact using a key
	searchContact := "Freya"

	// Check if the contact exists in the map
	_, found := contactDetails[searchContact]
	if found == true { // Or contact, found := contacts[nameToSearch]; found
		contact := contactDetails[searchContact]
		result = fmt.Sprintf("Contact Found: Name %s Phone %s Email %s", contact.Name, contact.Phone, contact.Email)
	} else {
		result = fmt.Sprintln("Not Found.")
	}
	return result
}
