package ntopng

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func (T User) GetName() string {
	return T.Username.String()
}

func (c *Client) GetUser(user User) (User, error) {
	return User{}, errors.New("not implemented")
}

func (c *Client) DeleteUser(user User) (User, error) {
	return user, errors.New("not implemented")
}

func (c *Client) CreateUser(user User) (User, error) {
	err := validate.Struct(user)
	if err != nil {
		return user, err
	}

	apiObjectAsJSON, err := json.Marshal(user)
	// set a new var with apiObjectAsJSON as a string, not []byte
	if err != nil {
		return user, err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", c.HostURL, "create/ntopng/user.lua"),
		strings.NewReader(string(apiObjectAsJSON)),
	)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return user, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return user, err
	}

	newAPIObject := make([]User, 1)
	err = json.Unmarshal(body, &newAPIObject[0])
	if err != nil {
		return user, err
	}
	return newAPIObject[0], nil
}

func (c *Client) UpdateUser(user User) (User, error) {
	return user, errors.New("not implemented")
}
