package lqfiles

import "go.uber.org/fx"

var FilesProvider = fx.Provide(NewFiles)
