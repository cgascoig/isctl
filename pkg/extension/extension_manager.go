package extension

import (
	"io/fs"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type ExtensionManager struct {
	extensions []*Extension
}

func NewExtensionManager() *ExtensionManager {
	return &ExtensionManager{
		extensions: []*Extension{},
	}
}

func (em *ExtensionManager) AddFS(fileSystem fs.FS, eoFn ExecuteOperationFunction, outputFn OutputFunction) {
	if fileSystem == nil {
		return
	}

	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil || !d.Type().IsRegular() {
			return nil
		}
		log.Debugf("Loading extension in %s\n", path)

		code, err := fs.ReadFile(fileSystem, path) //fileSystem.ReadFile(path)
		if err != nil {
			log.Debugf("Error loading extension: %v", err)
			return nil
		}

		e := NewExtension()
		e.ExecuteOperation = eoFn
		e.Output = outputFn
		err = e.SetCode(code)
		if err != nil {
			log.Debugf("Error loading extension: %v", err)
			return nil
		}

		em.extensions = append(em.extensions, e)

		return nil
	})
}

func (em *ExtensionManager) GetCommands() []*cobra.Command {
	cmds := []*cobra.Command{}
	for _, e := range em.extensions {
		cmd, err := e.GetCommand()
		if err != nil {
			log.Debugf("Error loading extension: %v", err)
		} else {
			cmds = append(cmds, cmd)
		}
	}

	return cmds
}
