package mangodb

type TimePoint struct {
	StartTime int64 `bson:"startTime"`
	EndTime   int64 `bson:"endTime"`
}

// 存储在mongodb中的内容
type UserInfo struct {
	UserId    string    `bson:"userid"`
	UserName  string    `bson:"username"`
	Password  string    `bson:"password"`
	Email     string    `bson:"email"`
	Token     string    `bson:"token"`
	Status    string    `bson:"status"`
	HeadImage string    `bson:"image"`
	Timepoint TimePoint `bson:"timepoint"`
}
