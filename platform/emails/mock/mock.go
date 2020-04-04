package mock

//go:generate mockgen -package=mock -destination=ses.go github.com/aws/aws-sdk-go/service/ses/sesiface SESAPI
