package zendesk

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestJson(t *testing.T) {
	Convey("Given an empty user instance", t, func() {
		user := User{}

		Convey("When I marshall it to json", func() {
			data, err := json.Marshal(user)

			Convey("Then I expect no errors", func() {
				So(err, ShouldBeNil)
			})

			Convey("And I expect the json to be {}", func() {
				So(string(data), ShouldEqual, "{}")
			})
		})
	})
}

func TestList(t *testing.T) {
	if !*runIntegrationTests {
		t.Skip("To run this test, use: go test -integration")
		return
	}

	Convey("Given the user api", t, func() {
		client, err := FromEnv()
		So(err, ShouldBeNil)

		userApi := client.Users()

		Convey("When I #List the users", func() {
			users, err := userApi.List()

			Convey("Then I expect no errors", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then I expect at least one user", func() {
				So(len(users), ShouldBeGreaterThan, 0)

				user := users[0]

				Convey("With name populated", func() {
					So(user.Name, ShouldNotBeBlank)
				})

				Convey("With email populated", func() {
					So(user.Email, ShouldNotBeBlank)
				})
			})
		})
	})
}

func TestMe(t *testing.T) {
	if !*runIntegrationTests {
		t.Skip("To run this test, use: go test -integration")
		return
	}

	Convey("Given the user api", t, func() {
		client, err := FromEnv()
		So(err, ShouldBeNil)

		userApi := client.Users()

		Convey("When I ask about #Me", func() {
			user, err := userApi.Me()

			Convey("Then I expect no errors", func() {
				So(err, ShouldBeNil)
			})

			Convey("And I expect id populated", func() {
				So(user.Id, ShouldBeGreaterThan, 0)
			})

			Convey("And I expect name populated", func() {
				So(user.Name, ShouldNotBeBlank)
			})

			Convey("And I expect email populated", func() {
				So(user.Email, ShouldNotBeBlank)
			})
		})
	})
}

func TestShow(t *testing.T) {
	if !*runIntegrationTests {
		t.Skip("To run this test, use: go test -integration")
		return
	}

	Convey("Given the user api", t, func() {
		client, err := FromEnv()
		So(err, ShouldBeNil)

		userApi := client.Users()
		me, err := userApi.Me()
		So(err, ShouldBeNil)

		Convey("When I ask about #Show", func() {
			user, err := userApi.Show(me.Id)

			Convey("Then I expect no errors", func() {
				So(err, ShouldBeNil)
			})

			Convey("And I expect name populated", func() {
				So(user.Name, ShouldNotBeBlank)
			})
		})
	})
}

func TestRelated(t *testing.T) {
	if !*runIntegrationTests {
		t.Skip("To run this test, use: go test -integration")
		return
	}

	Convey("Given the user api", t, func() {
		client, err := FromEnv()
		So(err, ShouldBeNil)

		userApi := client.Users()
		me, err := userApi.Me()
		So(err, ShouldBeNil)

		Convey("When I ask about #Related", func() {
			related, err := userApi.Related(me.Id)

			Convey("Then I expect no errors", func() {
				So(err, ShouldBeNil)
			})

			Convey("And I expect some values back", func() {
				So(len(related), ShouldBeGreaterThan, 0)
			})
		})
	})
}

func TestAutocomplete(t *testing.T) {
	if !*runIntegrationTests {
		t.Skip("To run this test, use: go test -integration")
		return
	}

	Convey("Given the user api", t, func() {
		client, err := FromEnv()
		So(err, ShouldBeNil)

		userApi := client.Users()
		me, err := userApi.Me()
		So(err, ShouldBeNil)

		Convey("When I call #Autocomplete", func() {
			users, err := userApi.Autocomplete(me.Name)

			Convey("Then I expect no errors", func() {
				So(err, ShouldBeNil)
			})

			Convey("And I expect at least one user back", func() {
				So(len(users), ShouldBeGreaterThan, 0)
			})
		})
	})
}

func TestSearchQuery(t *testing.T) {
	if !*runIntegrationTests {
		t.Skip("To run this test, use: go test -integration")
		return
	}

	Convey("Given the user api", t, func() {
		client, err := FromEnv()
		So(err, ShouldBeNil)

		userApi := client.Users()
		me, err := userApi.Me()
		So(err, ShouldBeNil)

		Convey("When I call #SearchQuery", func() {
			users, err := userApi.SearchQuery(me.Name)

			Convey("Then I expect no errors", func() {
				So(err, ShouldBeNil)
			})

			Convey("And I expect at least one user back", func() {
				So(len(users), ShouldBeGreaterThan, 0)
			})
		})
	})
}

func TestCreateDeleteAndManagePassword(t *testing.T) {
	if !*runIntegrationTests {
		t.Skip("To run this test, use: go test -integration")
		return
	}

	Convey("Given the user api", t, func() {
		client, err := FromEnv()
		So(err, ShouldBeNil)

		Convey("When I create a new user via #Create", func() {
			request := User{
				Name:  "Sample User",
				Email: fmt.Sprintf("matt.ho+%d@gmail.com", time.Now().Unix()),
				UserFields: map[string]interface{}{
					"text_field":    "abc",
					"decimal_field": 123,
				},
			}

			// test 1 - create the user
			user, err := client.Users().Create(request)
			So(err, ShouldBeNil)
			So(user.Id, ShouldNotEqual, 0)
			So(user.UserFields["text_field"], ShouldEqual, "abc")
			So(user.UserFields["decimal_field"], ShouldEqual, 123)

			// test 2 - set password
			password := "password"
			err = client.Users().SetPassword(user.Id, password)
			So(err, ShouldBeNil)

			userClient := FromPassword(client.domain, request.Email, password)
			me, err := userClient.Users().Me()
			So(err, ShouldBeNil)
			So(me.Id, ShouldEqual, user.Id)

			// test 3 - delete the user
			deleted, err := client.Users().Delete(user.Id)
			So(err, ShouldBeNil)
			So(deleted.Active, ShouldBeFalse)
		})
	})
}
