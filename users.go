package zendesk

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/context"
)

type User struct {
	Id                  int               `json:"id"`
	Url                 string            `json:"url"`
	Name                string            `json:"name"`
	ExternalId          string            `json:"external_id"`
	Alias               string            `json:"alias"`
	CreatedAt           string            `json:"created_at"`
	UpdatedAt           string            `json:"updated_at"`
	Active              bool              `json:"active"`
	Verified            bool              `json:"verified"`
	Shared              bool              `json:"shared"`
	SharedAgent         bool              `json:"shared_agent"`
	Locale              string            `json:"locale"`
	LocaleId            int               `json:"locale_id"`
	TimeZone            string            `json:"time_zone"`
	LastLoginAt         string            `json:"last_login_at"`
	Email               string            `json:"email"`
	Phone               string            `json:"phone"`
	Signature           string            `json:"signature"`
	Details             string            `json:"details"`
	Notes               string            `json:"notes"`
	OrganizationId      int               `json:"organization_id"`
	Role                string            `json:"role"`
	CustomRoleId        string            `json:"custom_role_id"`
	Moderator           bool              `json:"moderator"`
	TicketRestriction   string            `json:"ticket_restriction"`
	OnlyPrivateComments bool              `json:"only_private_comments"`
	Tags                []string          `json:"tags"`
	Suspended           bool              `json:"suspended"`
	RestrictedAgent     bool              `json:"restricted_agent"`
	Photo               Attachment        `json:"photo"`
	UserFields          map[string]string `json:"user_fields"`
}

type UserApi struct {
	client  *Client
	context context.Context
}

func (api *UserApi) WithContext(ctx context.Context) *UserApi {
	return &UserApi{
		client:  api.client,
		context: ctx,
	}
}

func (api *UserApi) getUsers(path string, params *url.Values) ([]User, error) {
	response := struct {
		Users []User `json:"users"`
	}{}

	err := api.client.get(api.context, path, params, &response)
	if err != nil {
		return nil, err
	}
	return response.Users, nil
}

func (api *UserApi) getUser(path string, params *url.Values) (User, error) {
	response := struct {
		User User `json:"user"`
	}{}

	err := api.client.get(api.context, path, params, &response)
	if err != nil {
		return User{}, err
	}
	return response.User, nil
}

func (api *UserApi) List() ([]User, error) {
	return api.getUsers("/api/v2/users.json", nil)
}

func (api *UserApi) ShowMany(ids ...int) ([]User, error) {
	params := url.Values{}
	params.Set("ids", strings.Join(toStringArray(ids), ","))

	return api.getUsers("/api/v2/show_many.json", &params)
}

func (api *UserApi) Show(id int) (User, error) {
	path := fmt.Sprintf("/api/v2/users/%d.json", id)
	return api.getUser(path, nil)
}

func (api *UserApi) Related(id int) (map[string]int, error) {
	path := fmt.Sprintf("/api/v2/users/%d/related.json", id)
	response := map[string]map[string]int{}

	err := api.client.get(api.context, path, nil, &response)
	if err != nil {
		return nil, err
	}

	return response["user_related"], nil
}

func (api *UserApi) Create(user User) (User, error) {
	return User{}, NotImplementedErr
}

func (api *UserApi) Merge(username, password string) (User, error) {
	return User{}, NotImplementedErr
}

func (api *UserApi) MergeByAdmin(baseUserId, toBeMergedUserId int) (User, error) {
	return User{}, NotImplementedErr
}

func (api *UserApi) CreateMany(users ...User) (JobStatus, error) {
	return JobStatus{}, NotImplementedErr
}

func (api *UserApi) UpdateUser(id string) (User, error) {
	return User{}, NotImplementedErr
}

func (api *UserApi) Suspend(id string) (User, error) {
	return User{}, NotImplementedErr
}

func (api *UserApi) Delete(id string) (User, error) {
	return User{}, NotImplementedErr
}

func (api *UserApi) SearchQuery(query string) ([]User, error) {
	params := url.Values{}
	params.Set("query", query)

	return api.getUsers("/api/v2/users/autocomplete.json", &params)
}

func (api *UserApi) SearchExternalId(externalId string) ([]User, error) {
	params := url.Values{}
	params.Set("external_id", externalId)

	return api.getUsers("/api/v2/users/autocomplete.json", &params)
}

func (api *UserApi) Autocomplete(name string) ([]User, error) {
	params := url.Values{}
	params.Set("name", name)

	return api.getUsers("/api/v2/users/autocomplete.json", &params)
}

func (api *UserApi) Me() (User, error) {
	return api.getUser("/api/v2/users/me.json", nil)
}

func (api *UserApi) SetPassword(id, password string) error {
	return NotImplementedErr
}

func (api *UserApi) ChangePassword(id, password string) error {
	return NotImplementedErr
}
