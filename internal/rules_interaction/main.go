package rules_interaction

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/sirrend/terrap-cli/internal/commons"
	"github.com/sirrend/terrap-cli/internal/utils"
)

// ruleBookObjectNameBuilder
/*
@brief:
	ruleBookObjectNameBuilder builds the name of the remote object to retrieve
@params:
	provider - string - the provider name
	sourceVersion - string - the source version
	targetVersion - string - the target version
*/
func ruleBookObjectNameBuilder(provider, sourceVersion, targetVersion string) string {
	return "rules/" + provider + "/" + sourceVersion + "_to_" + targetVersion + ".json"
}

// GetRulebook
/*
@brief:
	GetRulebook retrieves a rulebook from the remote bucket
@params:
	provider - string - the provider name
	sourceVersion - string - the source version
	targetVersion - string - the target version
@returns:
	Rulebook - the downloaded rulebook
	error - if exists, else nil
*/
func GetRulebook(provider, sourceVersion, targetVersion string) (Rulebook, error) {
	var rulebook Rulebook
	rulebookName := ruleBookObjectNameBuilder(provider, sourceVersion, targetVersion)
	clientSession := session.Must(session.NewSession())

	// Create a S3 client from just a session.
	client := s3.New(clientSession, &aws.Config{
		Region: &commons.RULES_BUCKET_REGION,
	})

	rulebookObject, err := client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(commons.RULES_BUCKET),
		Key:    aws.String(rulebookName),
	})

	if err != nil {
		return Rulebook{}, err
	}

	rulebook.Bytes = utils.StreamToByte(rulebookObject.Body)
	rulebook.SourceVersion = sourceVersion
	rulebook.TargetVersion = targetVersion

	return rulebook, nil
}
