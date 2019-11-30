package twitter

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestDecodeTwitterAccount(t *testing.T) {
	// try decode a "kitchen sink", so that we can test that most features get decoded correctly
	twitterAccountApiResponseJson := strings.NewReader(`
{
  "id": 9876543,
  "id_str": "9876543",
  "name": "barry",
  "screen_name": "testingperson",
  "location": "In your hearts and minds",
  "profile_location": null,
  "description": "i am the what i was of what i will be.",
  "url": "https:\/\/t.co\/abcdefghij",
  "entities": {
    "url": {
      "urls": [
        {
          "url": "https:\/\/t.co\/abcdefghij",
          "expanded_url": "http:\/\/Instagram.com\/demotestingIGperson",
          "display_url": "Instagram.com\/demotestingIGperson",
          "indices": [
            0,
            23
          ]
        }
      ]
    },
    "description": {
      "urls": [
        
      ]
    }
  },
  "protected": false,
  "followers_count": 161,
  "friends_count": 280,
  "listed_count": 8,
  "created_at": "Wed Mar 21 18:13:14 +0000 2007",
  "favourites_count": 2279,
  "utc_offset": null,
  "time_zone": null,
  "geo_enabled": true,
  "verified": false,
  "statuses_count": 1729,
  "lang": null,
  "status": {
    "created_at": "Wed Nov 27 18:54:49 +0000 2019",
    "id": 1234567890123456789,
    "id_str": "1234567890123456789",
    "text": "Demo tweet #testing https:\/\/t.co\/abcdefgijk",
    "truncated": false,
    "entities": {
      "hashtags": [
        {
          "text": "testing",
          "indices": [
            0,
            8
          ]
        }
      ],
      "symbols": [
        
      ],
      "user_mentions": [
        
      ],
      "urls": [
        {
          "url": "https:\/\/t.co\/abcdefgijk",
          "expanded_url": "https:\/\/www.instagram.com\/p\/BAABAABAABA\/?igshid=xyxyxyxyxyxyx",
          "display_url": "instagram.com\/p\/BAABAABAABA\/\u2026",
          "indices": [
            52,
            75
          ]
        }
      ]
    },
    "source": "\u003ca href=\"http:\/\/instagram.com\" rel=\"nofollow\"\u003eInstagram\u003c\/a\u003e",
    "in_reply_to_status_id": null,
    "in_reply_to_status_id_str": null,
    "in_reply_to_user_id": null,
    "in_reply_to_user_id_str": null,
    "in_reply_to_screen_name": null,
    "geo": {
      "type": "Point",
      "coordinates": [
        34.0522,
        -118.243
      ]
    },
    "coordinates": {
      "type": "Point",
      "coordinates": [
        -118.243,
        34.0522
      ]
    },
    "place": {
      "id": "3b77caf94bfc81fe",
      "url": "https:\/\/api.twitter.com\/1.1\/geo\/id\/3b77caf94bfc81fe.json",
      "place_type": "city",
      "name": "Los Angeles",
      "full_name": "Los Angeles, CA",
      "country_code": "US",
      "country": "USA",
      "contained_within": [
        
      ],
      "bounding_box": {
        "type": "Polygon",
        "coordinates": [
          [
            [
              -118.668404,
              33.704538
            ],
            [
              -118.155409,
              33.704538
            ],
            [
              -118.155409,
              34.337041
            ],
            [
              -118.668404,
              34.337041
            ]
          ]
        ]
      },
      "attributes": {
        
      }
    },
    "contributors": null,
    "is_quote_status": false,
    "retweet_count": 0,
    "favorite_count": 0,
    "favorited": false,
    "retweeted": false,
    "possibly_sensitive": false,
    "lang": "en"
  },
  "contributors_enabled": false,
  "is_translator": false,
  "is_translation_enabled": false,
  "profile_background_color": "FFFFFF",
  "profile_background_image_url": "http:\/\/abs.twimg.com\/images\/themes\/theme1\/bg.png",
  "profile_background_image_url_https": "https:\/\/abs.twimg.com\/images\/themes\/theme1\/bg.png",
  "profile_background_tile": true,
  "profile_image_url": "http:\/\/pbs.twimg.com\/profile_images\/923335960007340032\/pIbUjNkC_normal.jpg",
  "profile_image_url_https": "https:\/\/pbs.twimg.com\/profile_images\/923335960007340032\/pIbUjNkC_normal.jpg",
  "profile_banner_url": "https:\/\/pbs.twimg.com\/profile_banners\/9876543\/1508975481",
  "profile_link_color": "0012BB",
  "profile_sidebar_border_color": "AAAAAA",
  "profile_sidebar_fill_color": "FFFFFF",
  "profile_text_color": "000000",
  "profile_use_background_image": false,
  "has_extended_profile": false,
  "default_profile": false,
  "default_profile_image": false,
  "can_media_tag": null,
  "followed_by": null,
  "following": null,
  "follow_request_sent": null,
  "notifications": null,
  "translator_type": "none"
}
`)

	var acc twitterAccount
	assertTrue(t, json.NewDecoder(twitterAccountApiResponseJson).Decode(&acc) == nil)

	// NOTE: assertions skipped for fields typed interface{}

	assertTrue(t, acc.ID == 9876543)
	assertEqualString(t, acc.IDStr, "9876543")
	assertEqualString(t, acc.ScreenName, "testingperson")
	assertEqualString(t, acc.Name, "barry")
	assertEqualString(t, acc.Location, "In your hearts and minds")
	assertEqualString(t, acc.Description, "i am the what i was of what i will be.")
	assertEqualString(t, acc.URL, "https://t.co/abcdefghij")

	assertTrue(t, !acc.Protected)
	assertTrue(t, acc.GeoEnabled)
	assertTrue(t, !acc.Verified)
	assertTrue(t, !acc.ContributorsEnabled)
	assertTrue(t, !acc.HasExtendedProfile)

	assertTrue(t, acc.FollowersCount == 161)
	assertTrue(t, acc.ListedCount == 8)
	assertTrue(t, acc.FavouritesCount == 2279)
	assertTrue(t, acc.StatusesCount == 1729)

	assertEqualString(t, acc.Lang, "")
	assertTrue(t, !acc.IsTranslator)
	assertTrue(t, !acc.IsTranslationEnabled)
	assertEqualString(t, acc.TranslatorType, "none")

	assertTrue(t, !acc.ProfileUseBackgroundImage)
	assertTrue(t, !acc.DefaultProfile)
	assertTrue(t, !acc.DefaultProfileImage)
	assertTrue(t, acc.ProfileBackgroundTile)
	assertEqualString(t, acc.ProfileBackgroundColor, "FFFFFF")
	assertEqualString(t, acc.ProfileBackgroundImageURL, "http://abs.twimg.com/images/themes/theme1/bg.png")
	assertEqualString(t, acc.ProfileBackgroundImageURLHTTPS, "https://abs.twimg.com/images/themes/theme1/bg.png")
	assertEqualString(t, acc.ProfileImageURL, "http://pbs.twimg.com/profile_images/923335960007340032/pIbUjNkC_normal.jpg")
	assertEqualString(t, acc.ProfileImageURLHTTPS, "https://pbs.twimg.com/profile_images/923335960007340032/pIbUjNkC_normal.jpg")
	assertEqualString(t, acc.ProfileBannerURL, "https://pbs.twimg.com/profile_banners/9876543/1508975481")
	assertEqualString(t, acc.ProfileLinkColor, "0012BB")
	assertEqualString(t, acc.ProfileSidebarBorderColor, "AAAAAA")
	assertEqualString(t, acc.ProfileSidebarFillColor, "FFFFFF")
	assertEqualString(t, acc.ProfileTextColor, "000000")

	latestTweet := acc.Status // shorthand

	assertEqualString(t, latestTweet.TweetIDStr, "1234567890123456789")
	assertTrue(t, latestTweet.TweetID == 1234567890123456789)
	assertTrue(t, latestTweet.User == nil)
	assertEqualString(t, latestTweet.CreatedAt, "Wed Nov 27 18:54:49 +0000 2019")
	assertEqualString(t, latestTweet.Text, "Demo tweet #testing https://t.co/abcdefgijk")
	assertEqualString(t, latestTweet.FullText, "")
	assertEqualString(t, latestTweet.Lang, "en")
	assertEqualString(t, latestTweet.Source, `<a href="http://instagram.com" rel="nofollow">Instagram</a>`)
	assertTrue(t, !latestTweet.Truncated)
	assertTrue(t, !latestTweet.PossiblySensitive)
	assertTrue(t, !latestTweet.IsQuoteStatus)

	assertEqualString(t, latestTweet.InReplyToScreenName, "")
	assertTrue(t, latestTweet.InReplyToStatusID == 0)
	assertEqualString(t, latestTweet.InReplyToStatusIDStr, "")
	assertTrue(t, latestTweet.InReplyToUserID == 0)
	assertEqualString(t, latestTweet.InReplyToUserIDStr, "")

	assertTrue(t, !latestTweet.WithheldCopyright)
	assertTrue(t, len(latestTweet.WithheldInCountries) == 0)
	assertEqualString(t, latestTweet.WithheldScope, "")

	assertTrue(t, !latestTweet.Favorited)
	assertTrue(t, latestTweet.FavoriteCount == 0)

	assertTrue(t, !latestTweet.Retweeted)
	assertTrue(t, latestTweet.RetweetedStatus == nil)
	assertTrue(t, latestTweet.RetweetCount == 0)

	assertTrue(t, len(latestTweet.DisplayTextRange) == 0)

	assertTrue(t, latestTweet.Coordinates.Latitude() == 34.0522)
	assertTrue(t, latestTweet.Coordinates.Longitude() == -118.243)

	assertTrue(t, latestTweet.ExtendedEntities == nil)
	// I was too lazy to type assertions for the "entities" hierarchy, so we're just comparing
	// re-serialized versions. this would catch if we would have had typos in JSON field
	// names (they would not get decoded, and hence would not get re-serialized)
	entitiesJson, err := json.MarshalIndent(latestTweet.Entities, "", "  ")
	assertTrue(t, err == nil)
	assertEqualString(t, string(entitiesJson), `{
  "hashtags": [
    {
      "indices": [
        0,
        8
      ],
      "text": "testing"
    }
  ],
  "symbols": [],
  "user_mentions": [],
  "urls": [
    {
      "url": "https://t.co/abcdefgijk",
      "expanded_url": "https://www.instagram.com/p/BAABAABAABA/?igshid=xyxyxyxyxyxyx",
      "display_url": "instagram.com/p/BAABAABAABA/…",
      "indices": [
        52,
        75
      ]
    }
  ],
  "polls": null
}`)
}

func assertEqualString(t *testing.T, actual string, expected string) {
	t.Helper()

	if actual != expected {
		t.Fatalf("exp=%v; got=%v", expected, actual)
	}
}

func assertTrue(t *testing.T, val bool) {
	t.Helper()

	if !val {
		t.Fatal("expected true; got false")
	}
}
