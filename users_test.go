package zendesk

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestList(t *testing.T) {
	if !*runIntegrationTests {
		t.Skip("To run this test, use: go test -integration")
		return
	}

	Convey("Given the user api", t, func() {
		client, err := NewFromEnv()
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
		client, err := NewFromEnv()
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
		client, err := NewFromEnv()
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
		client, err := NewFromEnv()
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
		client, err := NewFromEnv()
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
		client, err := NewFromEnv()
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
