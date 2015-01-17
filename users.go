package zendesk

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/context"
)

type User struct {
	Id                  int               `json:"id,omitempty"`
	Url                 string            `json:"url,omitempty"`
	Name                string            `json:"name,omitempty"`
	ExternalId          string            `json:"external_id,omitempty"`
	Alias               string            `json:"alias,omitempty"`
	CreatedAt           string            `json:"created_at,omitempty"`
	UpdatedAt           string            `json:"updated_at,omitempty"`
	Active              bool              `json:"active,omitempty"`
	Verified            bool              `json:"verified,omitempty"`
	Shared              bool              `json:"shared,omitempty"`
	SharedAgent         bool              `json:"shared_agent,omitempty"`
	Locale              string            `json:"locale,omitempty"`
	LocaleId            int               `json:"locale_id,omitempty"`
	TimeZone            string            `json:"time_zone,omitempty"`
	LastLoginAt         string            `json:"last_login_at,omitempty"`
	Email               string            `json:"email,omitempty"`
	Phone               string            `json:"phone,omitempty"`
	Signature           string            `json:"signature,omitempty"`
	Details             string            `json:"details,omitempty"`
	Notes               string            `json:"notes,omitempty"`
	OrganizationId      int               `json:"organization_id,omitempty"`
	Role                string            `json:"role,omitempty"`
	CustomRoleId        string            `json:"custom_role_id,omitempty"`
	Moderator           bool              `json:"moderator,omitempty"`
	TicketRestriction   string            `json:"ticket_restriction,omitempty"`
	OnlyPrivateComments bool              `json:"only_private_comments,omitempty"`
	Tags                []string          `json:"tags,omitempty"`
	Suspended           bool              `json:"suspended,omitempty"`
	RestrictedAgent     bool              `json:"restricted_agent,omitempty"`
	Photo               *Attachment       `json:"photo,omitempty"`
	UserFields          map[string]string `json:"user_fields,omitempty"`
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

func (api *UserApi) postUser(path string, payload interface{}) (User, error) {
	response := struct {
		User User `json:"user"`
	}{}

	err := api.client.post(api.context, path, payload, &response)
	if err != nil {
		return User{}, err
	}
	return response.User, nil
}

func (api *UserApi) deleteUser(path string) (User, error) {
	response := struct {
		User User `json:"user"`
	}{}

	err := api.client.delete(api.context, path, &response)
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
	return api.postUser("/api/v2/users.json", map[string]User{"user": user})
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

func (api *UserApi) UpdateUser(id int) (User, error) {
	return User{}, NotImplementedErr
}

func (api *UserApi) Suspend(id int) (User, error) {
	return User{}, NotImplementedErr
}

func (api *UserApi) Delete(id int) (User, error) {
	path := fmt.Sprintf("/api/v2/users/%d.json", id)
	return api.deleteUser(path)
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
