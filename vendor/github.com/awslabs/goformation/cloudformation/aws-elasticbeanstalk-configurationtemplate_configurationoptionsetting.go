package cloudformation

// AWSElasticBeanstalkConfigurationTemplate_ConfigurationOptionSetting AWS CloudFormation Resource (AWS::ElasticBeanstalk::ConfigurationTemplate.ConfigurationOptionSetting)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-option-settings.html
type AWSElasticBeanstalkConfigurationTemplate_ConfigurationOptionSetting struct {

	// Namespace AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-option-settings.html#cfn-beanstalk-optionsettings-namespace
	Namespace string `json:"Namespace,omitempty"`

	// OptionName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-option-settings.html#cfn-beanstalk-optionsettings-optionname
	OptionName string `json:"OptionName,omitempty"`

	// Value AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-option-settings.html#cfn-beanstalk-optionsettings-value
	Value string `json:"Value,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticBeanstalkConfigurationTemplate_ConfigurationOptionSetting) AWSCloudFormationType() string {
	return "AWS::ElasticBeanstalk::ConfigurationTemplate.ConfigurationOptionSetting"
}
