package entities

type User struct {
	ID          string                 `jsonapi:"primary,users" map:"id"`
	Name        string                 `jsonapi:"attr,userName" map:"name"`
	Password    string                 `jsonapi:"attr,firstName" map:"password"`
	FirstName   string                 `jsonapi:"attr,userName" map:"firstname"`
	LastName    string                 `jsonapi:"attr,lastName" map:"lastname"`
	PhoneNumber string                 `jsonapi:"attr,phoneNumber" map:"phone_num"`
	Email       string                 `jsonapi:"attr,email" map:"email"`
	Image       string                 `jsonapi:"attr,image" map:"image"`
	Settings    map[string]interface{} `jsonapi:"attr,settings" map:"settings"`

	//Company       *CompanyData        `jsonapi:"relation,company" map:"company"`
	//Location      *LocationData       `jsonapi:"relation,location" map:"location"`
	//Roles         []*RoleData         `jsonapi:"relation,roles" map:"roles"`
	//Subscriptions []*SubscriptionData `jsonapi:"relation,subscriptions" map:"subs"`
}
