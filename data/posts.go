package data

import postpb "go-grpc-playground/protos/v1/posts"

type PostData struct {
	UserId string
	Posts  []*postpb.PostMessage
}

var UserPosts = []*PostData{
	{
		UserId: "1",
		Posts: []*postpb.PostMessage{
			{
				PostId: "1",
				Author: "",
				Title:  "title1",
				Body:   "body1",
				Tags:   []string{"Tag1, Tag2"},
			},
			{
				PostId: "2",
				Author: "",
				Title:  "title2",
				Body:   "body2",
				Tags:   []string{"Tag1, Tag3"},
			},
		},
	},
}
