package rules_interaction

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/sirrend/terrap-cli/internal/commons"
	"github.com/sirrend/terrap-cli/internal/utils"
	"os"
	"strings"
)

// ruleBookObjectNamePrefixBuilder
/*
@brief:
	ruleBookObjectNamePrefixBuilder builds the name of the remote object prefix to look for
@params:
	provider - string - the provider name
	sourceVersion - string - the source version
*/
func ruleBookObjectNamePrefixBuilder(provider, sourceVersion string) string {
	return "rules/" + provider + "/" + sourceVersion + "_to_"
}

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

// FindRulebookFile
/*
@brief:
	FindRulebookFile find the rulebook file name
@params:
	provider - string - the provider name
	sourceVersion - string - the source version
*/
func FindRulebookFile(provider, sourceVersion string) string {
	provider = strings.ReplaceAll(provider, "/", "-") // edit provider to match s3 format
	clientSession := session.Must(session.NewSession())

	// Create a S3 client from just a session.
	client := s3.New(clientSession, &aws.Config{
		Region: &commons.RULES_BUCKET_REGION,
	})

	input := &s3.ListObjectsV2Input{
		Bucket: aws.String("terrap-rulebooks"),
		Prefix: aws.String(ruleBookObjectNamePrefixBuilder(provider, sourceVersion)),
	}

	output, err := client.ListObjectsV2(input)
	if err != nil {
		fmt.Println("Failed to list objects", err)
		os.Exit(1)
	}

	if len(output.Contents) > 0 {
		return *output.Contents[0].Key
	}

	return ""
}

// GetRulebook
/*
@brief:
	GetRulebook retrieves a rulebook from the remote bucket
@params:
	rulebookName - string - the rulebook name
	sourceVersion - string - the source version
	targetVersion - string - the target version
@returns:
	Rulebook - the downloaded rulebook
	error - if exists, else nil
*/
func GetRulebook(rulebookName, sourceVersion, targetVersion string) (Rulebook, error) {
	var rulebook Rulebook

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
