package main

// Posts struct for reddit post
type Posts struct {
	Kind string `json:"kind"`
	Data struct {
		Modhash  string `json:"modhash"`
		Dist     int    `json:"dist"`
		Children []struct {
			Kind string `json:"kind"`
			Data struct {
				ApprovedAtUtc     interface{} `json:"approved_at_utc"`
				Subreddit         string      `json:"subreddit"`
				Selftext          string      `json:"selftext"`
				AuthorFullname    string      `json:"author_fullname"`
				Saved             bool        `json:"saved"`
				ModReasonTitle    interface{} `json:"mod_reason_title"`
				Gilded            int         `json:"gilded"`
				Clicked           bool        `json:"clicked"`
				Title             string      `json:"title"`
				LinkFlairRichtext []struct {
					E string `json:"e"`
					T string `json:"t"`
				} `json:"link_flair_richtext"`
				SubredditNamePrefixed      string      `json:"subreddit_name_prefixed"`
				Hidden                     bool        `json:"hidden"`
				Pwls                       int         `json:"pwls"`
				LinkFlairCSSClass          string      `json:"link_flair_css_class"`
				Downs                      int         `json:"downs"`
				ThumbnailHeight            interface{} `json:"thumbnail_height"`
				TopAwardedType             interface{} `json:"top_awarded_type"`
				HideScore                  bool        `json:"hide_score"`
				Name                       string      `json:"name"`
				Quarantine                 bool        `json:"quarantine"`
				LinkFlairTextColor         string      `json:"link_flair_text_color"`
				UpvoteRatio                float64     `json:"upvote_ratio"`
				AuthorFlairBackgroundColor string      `json:"author_flair_background_color"`
				SubredditType              string      `json:"subreddit_type"`
				Ups                        int         `json:"ups"`
				TotalAwardsReceived        int         `json:"total_awards_received"`
				MediaEmbed                 struct {
				} `json:"media_embed"`
				ThumbnailWidth        interface{}   `json:"thumbnail_width"`
				AuthorFlairTemplateID string        `json:"author_flair_template_id"`
				IsOriginalContent     bool          `json:"is_original_content"`
				UserReports           []interface{} `json:"user_reports"`
				SecureMedia           interface{}   `json:"secure_media"`
				IsRedditMediaDomain   bool          `json:"is_reddit_media_domain"`
				IsMeta                bool          `json:"is_meta"`
				Category              interface{}   `json:"category"`
				SecureMediaEmbed      struct {
				} `json:"secure_media_embed"`
				LinkFlairText       string      `json:"link_flair_text"`
				CanModPost          bool        `json:"can_mod_post"`
				Score               int         `json:"score"`
				ApprovedBy          interface{} `json:"approved_by"`
				AuthorPremium       bool        `json:"author_premium"`
				Thumbnail           string      `json:"thumbnail"`
				Edited              bool        `json:"edited"`
				AuthorFlairCSSClass string      `json:"author_flair_css_class"`
				AuthorFlairRichtext []struct {
					A string `json:"a,omitempty"`
					E string `json:"e"`
					U string `json:"u,omitempty"`
					T string `json:"t,omitempty"`
				} `json:"author_flair_richtext"`
				Gildings struct {
				} `json:"gildings"`
				PostHint          string      `json:"post_hint"`
				ContentCategories interface{} `json:"content_categories"`
				IsSelf            bool        `json:"is_self"`
				ModNote           interface{} `json:"mod_note"`
				Created           float64     `json:"created"`
				LinkFlairType     string      `json:"link_flair_type"`
				Wls               int         `json:"wls"`
				RemovedByCategory interface{} `json:"removed_by_category"`
				BannedBy          interface{} `json:"banned_by"`
				AuthorFlairType   string      `json:"author_flair_type"`
				Domain            string      `json:"domain"`
				AllowLiveComments bool        `json:"allow_live_comments"`
				SelftextHTML      string      `json:"selftext_html"`
				Likes             interface{} `json:"likes"`
				SuggestedSort     interface{} `json:"suggested_sort"`
				BannedAtUtc       interface{} `json:"banned_at_utc"`
				ViewCount         interface{} `json:"view_count"`
				Archived          bool        `json:"archived"`
				NoFollow          bool        `json:"no_follow"`
				IsCrosspostable   bool        `json:"is_crosspostable"`
				Pinned            bool        `json:"pinned"`
				Over18            bool        `json:"over_18"`
				Preview           struct {
					Images []struct {
						Source struct {
							URL    string `json:"url"`
							Width  int    `json:"width"`
							Height int    `json:"height"`
						} `json:"source"`
						Resolutions []struct {
							URL    string `json:"url"`
							Width  int    `json:"width"`
							Height int    `json:"height"`
						} `json:"resolutions"`
						Variants struct {
						} `json:"variants"`
						ID string `json:"id"`
					} `json:"images"`
					Enabled bool `json:"enabled"`
				} `json:"preview"`
				AllAwardings             []interface{} `json:"all_awardings"`
				Awarders                 []interface{} `json:"awarders"`
				MediaOnly                bool          `json:"media_only"`
				LinkFlairTemplateID      string        `json:"link_flair_template_id"`
				CanGild                  bool          `json:"can_gild"`
				Spoiler                  bool          `json:"spoiler"`
				Locked                   bool          `json:"locked"`
				AuthorFlairText          string        `json:"author_flair_text"`
				TreatmentTags            []interface{} `json:"treatment_tags"`
				Visited                  bool          `json:"visited"`
				RemovedBy                interface{}   `json:"removed_by"`
				NumReports               interface{}   `json:"num_reports"`
				Distinguished            interface{}   `json:"distinguished"`
				SubredditID              string        `json:"subreddit_id"`
				ModReasonBy              interface{}   `json:"mod_reason_by"`
				RemovalReason            interface{}   `json:"removal_reason"`
				LinkFlairBackgroundColor string        `json:"link_flair_background_color"`
				ID                       string        `json:"id"`
				IsRobotIndexable         bool          `json:"is_robot_indexable"`
				ReportReasons            interface{}   `json:"report_reasons"`
				Author                   string        `json:"author"`
				DiscussionType           interface{}   `json:"discussion_type"`
				NumComments              int           `json:"num_comments"`
				SendReplies              bool          `json:"send_replies"`
				WhitelistStatus          string        `json:"whitelist_status"`
				ContestMode              bool          `json:"contest_mode"`
				ModReports               []interface{} `json:"mod_reports"`
				AuthorPatreonFlair       bool          `json:"author_patreon_flair"`
				AuthorFlairTextColor     string        `json:"author_flair_text_color"`
				Permalink                string        `json:"permalink"`
				ParentWhitelistStatus    string        `json:"parent_whitelist_status"`
				Stickied                 bool          `json:"stickied"`
				URL                      string        `json:"url"`
				SubredditSubscribers     int           `json:"subreddit_subscribers"`
				CreatedUtc               float64       `json:"created_utc"`
				NumCrossposts            int           `json:"num_crossposts"`
				Media                    interface{}   `json:"media"`
				IsVideo                  bool          `json:"is_video"`
			} `json:"data,omitempty"`
		} `json:"children"`
		After  string      `json:"after"`
		Before interface{} `json:"before"`
	} `json:"data"`
}

// Comments struct of comment
type Comments []struct {
	Kind string `json:"kind"`
	Data struct {
		Modhash  string `json:"modhash"`
		Dist     int    `json:"dist"`
		Children []struct {
			Kind string `json:"kind"`
			Data struct {
				ApprovedAtUtc     interface{}   `json:"approved_at_utc"`
				Subreddit         string        `json:"subreddit"`
				Selftext          string        `json:"selftext"`
				UserReports       []interface{} `json:"user_reports"`
				Saved             bool          `json:"saved"`
				ModReasonTitle    interface{}   `json:"mod_reason_title"`
				Gilded            int           `json:"gilded"`
				Clicked           bool          `json:"clicked"`
				Title             string        `json:"title"`
				LinkFlairRichtext []struct {
					E string `json:"e"`
					T string `json:"t"`
				} `json:"link_flair_richtext"`
				SubredditNamePrefixed      string      `json:"subreddit_name_prefixed"`
				Hidden                     bool        `json:"hidden"`
				Pwls                       int         `json:"pwls"`
				LinkFlairCSSClass          string      `json:"link_flair_css_class"`
				Downs                      int         `json:"downs"`
				ThumbnailHeight            interface{} `json:"thumbnail_height"`
				TopAwardedType             interface{} `json:"top_awarded_type"`
				ParentWhitelistStatus      string      `json:"parent_whitelist_status"`
				HideScore                  bool        `json:"hide_score"`
				Body                       string      `json:"body"`
				Name                       string      `json:"name"`
				Quarantine                 bool        `json:"quarantine"`
				LinkFlairTextColor         string      `json:"link_flair_text_color"`
				UpvoteRatio                float64     `json:"upvote_ratio"`
				AuthorFlairBackgroundColor interface{} `json:"author_flair_background_color"`
				SubredditType              string      `json:"subreddit_type"`
				Ups                        int         `json:"ups"`
				TotalAwardsReceived        int         `json:"total_awards_received"`
				MediaEmbed                 struct {
				} `json:"media_embed"`
				ThumbnailWidth        interface{} `json:"thumbnail_width"`
				AuthorFlairTemplateID interface{} `json:"author_flair_template_id"`
				IsOriginalContent     bool        `json:"is_original_content"`
				AuthorFullname        string      `json:"author_fullname"`
				SecureMedia           interface{} `json:"secure_media"`
				IsRedditMediaDomain   bool        `json:"is_reddit_media_domain"`
				IsMeta                bool        `json:"is_meta"`
				Category              interface{} `json:"category"`
				SecureMediaEmbed      struct {
				} `json:"secure_media_embed"`
				LinkFlairText       string        `json:"link_flair_text"`
				CanModPost          bool          `json:"can_mod_post"`
				Score               int           `json:"score"`
				ApprovedBy          interface{}   `json:"approved_by"`
				AuthorPremium       bool          `json:"author_premium"`
				Thumbnail           string        `json:"thumbnail"`
				Edited              interface{}   `json:"edited"`
				AuthorFlairCSSClass interface{}   `json:"author_flair_css_class"`
				AuthorFlairRichtext []interface{} `json:"author_flair_richtext"`
				Gildings            struct {
					Gid2 int `json:"gid_2"`
				} `json:"gildings"`
				ContentCategories interface{} `json:"content_categories"`
				IsSelf            bool        `json:"is_self"`
				ModNote           interface{} `json:"mod_note"`
				Created           float64     `json:"created"`
				LinkFlairType     string      `json:"link_flair_type"`
				Wls               int         `json:"wls"`
				RemovedByCategory interface{} `json:"removed_by_category"`
				BannedBy          interface{} `json:"banned_by"`
				AuthorFlairType   string      `json:"author_flair_type"`
				Domain            string      `json:"domain"`
				AllowLiveComments bool        `json:"allow_live_comments"`
				SelftextHTML      interface{} `json:"selftext_html"`
				Likes             interface{} `json:"likes"`
				SuggestedSort     interface{} `json:"suggested_sort"`
				BannedAtUtc       interface{} `json:"banned_at_utc"`
				ViewCount         interface{} `json:"view_count"`
				Archived          bool        `json:"archived"`
				NoFollow          bool        `json:"no_follow"`
				IsCrosspostable   bool        `json:"is_crosspostable"`
				Pinned            bool        `json:"pinned"`
				Over18            bool        `json:"over_18"`
				AllAwardings      []struct {
					GiverCoinReward     interface{} `json:"giver_coin_reward"`
					SubredditID         interface{} `json:"subreddit_id"`
					IsNew               bool        `json:"is_new"`
					DaysOfDripExtension int         `json:"days_of_drip_extension"`
					CoinPrice           int         `json:"coin_price"`
					ID                  string      `json:"id"`
					PennyDonate         interface{} `json:"penny_donate"`
					CoinReward          int         `json:"coin_reward"`
					IconURL             string      `json:"icon_url"`
					DaysOfPremium       int         `json:"days_of_premium"`
					IconHeight          int         `json:"icon_height"`
					ResizedIcons        []struct {
						URL    string `json:"url"`
						Width  int    `json:"width"`
						Height int    `json:"height"`
					} `json:"resized_icons"`
					IconWidth           int         `json:"icon_width"`
					StaticIconWidth     int         `json:"static_icon_width"`
					StartDate           interface{} `json:"start_date"`
					IsEnabled           bool        `json:"is_enabled"`
					Description         string      `json:"description"`
					EndDate             interface{} `json:"end_date"`
					SubredditCoinReward int         `json:"subreddit_coin_reward"`
					Count               int         `json:"count"`
					StaticIconHeight    int         `json:"static_icon_height"`
					Name                string      `json:"name"`
					ResizedStaticIcons  []struct {
						URL    string `json:"url"`
						Width  int    `json:"width"`
						Height int    `json:"height"`
					} `json:"resized_static_icons"`
					IconFormat    interface{} `json:"icon_format"`
					AwardSubType  string      `json:"award_sub_type"`
					PennyPrice    interface{} `json:"penny_price"`
					AwardType     string      `json:"award_type"`
					StaticIconURL string      `json:"static_icon_url"`
				} `json:"all_awardings"`
				Awarders                 []interface{} `json:"awarders"`
				MediaOnly                bool          `json:"media_only"`
				LinkFlairTemplateID      string        `json:"link_flair_template_id"`
				CanGild                  bool          `json:"can_gild"`
				Spoiler                  bool          `json:"spoiler"`
				Locked                   bool          `json:"locked"`
				AuthorFlairText          interface{}   `json:"author_flair_text"`
				TreatmentTags            []interface{} `json:"treatment_tags"`
				Visited                  bool          `json:"visited"`
				RemovedBy                interface{}   `json:"removed_by"`
				NumReports               interface{}   `json:"num_reports"`
				Distinguished            interface{}   `json:"distinguished"`
				SubredditID              string        `json:"subreddit_id"`
				ModReasonBy              interface{}   `json:"mod_reason_by"`
				RemovalReason            interface{}   `json:"removal_reason"`
				LinkFlairBackgroundColor string        `json:"link_flair_background_color"`
				ID                       string        `json:"id"`
				IsRobotIndexable         bool          `json:"is_robot_indexable"`
				NumDuplicates            int           `json:"num_duplicates"`
				ReportReasons            interface{}   `json:"report_reasons"`
				Author                   string        `json:"author"`
				DiscussionType           interface{}   `json:"discussion_type"`
				NumComments              int           `json:"num_comments"`
				SendReplies              bool          `json:"send_replies"`
				Media                    interface{}   `json:"media"`
				ContestMode              bool          `json:"contest_mode"`
				AuthorPatreonFlair       bool          `json:"author_patreon_flair"`
				AuthorFlairTextColor     interface{}   `json:"author_flair_text_color"`
				Permalink                string        `json:"permalink"`
				WhitelistStatus          string        `json:"whitelist_status"`
				Stickied                 bool          `json:"stickied"`
				URL                      string        `json:"url"`
				SubredditSubscribers     int           `json:"subreddit_subscribers"`
				CreatedUtc               float64       `json:"created_utc"`
				NumCrossposts            int           `json:"num_crossposts"`
				ModReports               []interface{} `json:"mod_reports"`
				IsVideo                  bool          `json:"is_video"`
			} `json:"data"`
		} `json:"children"`
		After  interface{} `json:"after"`
		Before interface{} `json:"before"`
	} `json:"data"`
}

// Definition struct of json response
type Definition []struct {
	Meta struct {
		ID        string   `json:"id"`
		UUID      string   `json:"uuid"`
		Sort      string   `json:"sort"`
		Src       string   `json:"src"`
		Section   string   `json:"section"`
		Stems     []string `json:"stems"`
		Offensive bool     `json:"offensive"`
	} `json:"meta"`
	Hwi struct {
		Hw  string `json:"hw"`
		Prs []struct {
			Mw    string `json:"mw"`
			Sound struct {
				Audio string `json:"audio"`
				Ref   string `json:"ref"`
				Stat  string `json:"stat"`
			} `json:"sound,omitempty"`
		} `json:"prs"`
	} `json:"hwi"`
	Fl  string `json:"fl"`
	Ins []struct {
		Il  string `json:"il"`
		Ifc string `json:"ifc"`
		If  string `json:"if"`
	} `json:"ins"`
	Def []struct {
		Sseq [][][]interface{} `json:"sseq"`
	} `json:"def"`
	Date     string   `json:"date"`
	Shortdef []string `json:"shortdef"`
}
