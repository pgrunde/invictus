package create

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pgrunde/invictus/create/templates"
)

var badWindowsNames = []string{"com1", "com2", "com3", "com4", "com5", "com6", "com7", "com8", "com9", "lpt1", "lpt2", "lpt3", "lpt4", "lpt5", "lpt6", "lpt7", "lpt8", "lpt9", "con", "nul", "prn"}

var badCharacters = []rune{'^', '/', '?', '<', '>', '\\', ':', '*', '|', '"', '.'}

type CreateSettings struct {
	ProjectName string
	GoPath      string
	FullPath    string
	DbName      string
	DbUser      string
	DbPassword  string
}

func NewProject(s, dbname, dbuser, dbpassword string) (err error) {
	s = strings.ToLower(s)
	if err = hasIllegalFilename(s); err != nil {
		return fmt.Errorf("Given project name %s has an illegal filename: %s", s, err)
	}
	err = os.Mkdir("."+string(filepath.Separator)+s, 0744)
	if os.IsExist(err) {
		return fmt.Errorf("A directory of that name already exists")
	}
	settings := NewCreateSettings(s, dbname, dbuser, dbpassword)
	GenerateNew(settings)
	return nil
}

func NewCreateSettings(s, dbname, dbuser, dbpassword string) (settings CreateSettings) {
	currentFullPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	gopath := buildGoPath(s, currentFullPath, os.Getenv("GOPATH"))
	settings.ProjectName = s
	settings.GoPath = gopath
	settings.FullPath = currentFullPath
	settings.DbName = dbname
	settings.DbUser = dbuser
	settings.DbPassword = dbpassword
	return
}

// GenerateNew takes in a project and creates a folder
// with enclosing files
func GenerateNew(s CreateSettings) {
	templates.CreateMain(s.ProjectName, s.GoPath)
	templates.CreateSettings(s.ProjectName, s.DbName)

	createServerFolder(s.ProjectName, s.FullPath)
	templates.CreateServer(s.ProjectName, s.FullPath, s.GoPath)

	createParamsFolder(s.ProjectName, s.FullPath)
	templates.CreateParams(s.ProjectName, s.FullPath)

	createDbFolder(s.ProjectName, s.FullPath)
	createMigrationsFolder(s.ProjectName, s.FullPath)
	templates.CreateDbConf(s.ProjectName, s.FullPath, s.DbName, s.DbUser, s.DbPassword)

	createApiFolder(s.ProjectName, s.FullPath)
	templates.CreateAPI(s.ProjectName, s.FullPath)
	templates.CreateEndpoint(s.ProjectName, s.FullPath)
	templates.CreateErrors(s.ProjectName, s.FullPath)
	templates.CreateInit(s.ProjectName, s.FullPath)
	templates.CreateInterface(s.ProjectName, s.FullPath)
	templates.CreateOption(s.ProjectName, s.FullPath)
	templates.CreateOptionMethod(s.ProjectName, s.FullPath)
	templates.CreateRequest(s.ProjectName, s.FullPath, s.GoPath)
	templates.CreateResponse(s.ProjectName, s.FullPath)
	templates.CreateTree(s.ProjectName, s.FullPath, s.GoPath)
	templates.CreateWrite(s.ProjectName, s.FullPath)

	createV1Folder(s.ProjectName, s.FullPath)
	templates.CreateV1(s.ProjectName, s.FullPath)
}

// bulidGoPath assumes that imports follow GOPATH + "/src"
func buildGoPath(s, w, g string) string {
	return w[len(g)+5:] + "/" + s
}

func hasIllegalFilename(s string) error {
	if s == "" {
		return fmt.Errorf("No project name given. Please specify a project name")
	}
	if len(s) > 255 {
		return fmt.Errorf("Project name cannot exceed 255 characters")
	}
	for _, bwn := range badWindowsNames {
		if s == bwn {
			return fmt.Errorf("Given name contains a reserved Windows operating system file name")
		}
	}
	for _, r := range badCharacters {
		if strings.ContainsRune(s, r) {
			return fmt.Errorf("Found illegal character %s in filename", string(r))
		}
	}
	return nil
}
