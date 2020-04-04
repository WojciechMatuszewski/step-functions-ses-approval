package mock

//go:generate mockgen -package=mock -destination=sfn.go github.com/aws/aws-sdk-go/service/sfn/sfniface SFNAPI
