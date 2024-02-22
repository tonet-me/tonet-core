package entity

type User struct {
	ID              string     `bson:"_id,omitempty" json:"id"`
	FirstName       string     `bson:"first_name,omitempty" json:"first_name"`
	LastName        string     `bson:"last_name,omitempty" json:"last_mame"`
	Email           string     `bson:"email" json:"email"`
	EmailVerified   bool       `bson:"email_verified" json:"email_verified"`
	PhoneNumber     string     `bson:"phone_number" json:"phone_number"`
	ProfilePhotoURL string     `bson:"profile_photo_url,omitempty" json:"profile_photo_url"`
	Status          UserStatus `bson:"status" json:"status"`
}

type UserStatus int

const (
	UserStatusActive UserStatus = iota + 1
	UserStatusDeActive
	UserStatusSuspend
)

func (u UserStatus) IsValid() bool {
	return u >= UserStatusActive && int(u) <= len(UserStatusStrings)
}

var UserStatusStrings = map[UserStatus]string{
	UserStatusActive:   "active",
	UserStatusDeActive: "deActive",
	UserStatusSuspend:  "suspend",
}

func (u UserStatus) String() string {
	return UserStatusStrings[u]
}
