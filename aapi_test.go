package pixiv

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	tokenSet bool
	testUID  uint64
)

func initTest() (err error) {
	if tokenSet {
		return
	}
	SetAuth(os.Getenv("TOKEN"), os.Getenv("REFRESH_TOKEN"))
	testUID, err = strconv.ParseUint(os.Getenv("TEST_UID"), 10, 0)
	tokenSet = true
	return err
}

func TestUserDetail(t *testing.T) {
	r := require.New(t)
	r.Nil(initTest())
	app := NewApp()
	detail, err := app.UserDetail(testUID)
	r.Nil(err)
	r.Equal(testUID, detail.User.ID)
}

func TestUserIllusts(t *testing.T) {
	r := require.New(t)
	r.Nil(initTest())
	app := NewApp()
	illusts, err := app.UserIllusts(490219, "illust", 0)
	r.Nil(err)
	r.Len(illusts, 30)
}

func TestUserBookmarksIllust(t *testing.T) {
	r := require.New(t)
	r.Nil(initTest())
	app := NewApp()
	illusts, err := app.UserBookmarksIllust(testUID, "public", 0, "")
	r.Nil(err)
	r.Equal(uint64(70095856), illusts[0].ID)
	size, err := download(illusts[0].MetaSinglePage.OriginalImageURL, ".", ".test.png", true)
	r.Nil(err)
	r.Equal(int64(993598), size)
}

func TestIllustFollow(t *testing.T) {
	r := require.New(t)
	r.Nil(initTest())
	app := NewApp()
	illusts, err := app.IllustFollow("public", 0)
	r.Nil(err)
	r.Len(illusts, 30)
}

// func TestTest(t *testing.T) {
// 	// r := require.New(t)
// 	initTest()
// 	app := NewApp()
// 	app.testResponse()
// 	panic("test")
// }
