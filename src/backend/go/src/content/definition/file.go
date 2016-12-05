package definition

const (
    C_FT_FILE = "file"
    C_FT_DIR = "directory"
)

type FileNode struct {
    ID string `json:"id"`
    Name string `json:"name"`
    Type string `json:"type"`
    Owner string `json:"owner"`
    CreatedTime int64 `json:"created_time"`
    UpdatedTime int64 `json:"updated_time"`
}

type Directory struct {
    FileNode
    Children []string   `json:"children"`
}

type File struct {
    FileNode
    Ref string `json:"reference"`
}

type FileTree struct {
    Current interface{} `json:"properties"`
    Children []*FileTree `json:"children,omitempty"`
}
