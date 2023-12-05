package parser

type Version struct {
	Info  Info   `json:"info"`
	Rules []Rule `json:"rules"`
}

type Info struct {
	RootResourceType string `json:"root_resource_type"`
	RootResourceName string `json:"root_resource_name"`
	ComponentType    string `json:"component_type"`
	ComponentName    string `json:"component_name"`
	Metadata         struct {
		DocumentationURL string `json:"documentation_url"`
	} `json:"metadata"`
}

// type Rule struct {
// 	Operation         string `json:"operation"`
// 	KeyType           string `json:"key_type"`
// 	HumanReadablePath string `json:"human_readable_path"`
// 	Notification      string `json:"notification"`
// 	IsRequired        bool   `json:"is_required"`
// 	ResourceComponent string `json:"resource_component"`
// 	Docs              string `json:"docs"`
// }
