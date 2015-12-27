package localconf

import (
	"fmt"
	"os"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewConfig(t *testing.T) {
	Convey("Given filename", t, func() {
		appName := "test_app"
		fileName := "app.json"

		Convey("When call", func() {
			c := NewConfig(appName, fileName)

			Convey("Then ok", func() {
				So(c, ShouldNotBeNil)
				h := os.Getenv("HOME")
				t.Logf("%s", h)
				So(c.Path(), ShouldEqual, fmt.Sprintf("%s/.test_app/app.json", h))
			})

			Reset(func() {
				c.Cleanup()
			})
		})
	})
}

func TestSaveAndLoad(t *testing.T) {
	Convey("Given filename", t, func() {
		type Hoge struct {
			Token string `json:"token"`
			Date  int64  `json:"date"`
		}

		appName := "test_app"
		fileName := "app.json"

		Convey("When call", func() {
			c := NewConfig(appName, fileName)

			src := &Hoge{
				Token: "hogehoge",
				Date:  time.Now().Unix(),
			}
			err := c.Save(src)

			Convey("Then save ok", func() {
				So(err, ShouldBeNil)

				tt := new(Hoge)
				err := c.Load(tt)

				Convey("Then load ok", func() {
					So(err, ShouldBeNil)
					So(tt.Token, ShouldEqual, src.Token)
					So(tt.Date, ShouldEqual, src.Date)
				})

				Reset(func() {
					c.Cleanup()
				})
			})
		})
	})
}
